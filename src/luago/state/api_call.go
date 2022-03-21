package state 

import "luago/binchunk"
import "luago/vm"

//Load方法既可以加载主函数实例化成闭包推入栈顶,也可以直接加载Lua脚本,
//mode可以是b、t或者bt
//b代表二进制chunk
//t代表文本chunk
//bt代表文本或二进制
//如果加载错误会在栈顶留下一个错误信息
// [-0, +1, –]
// http://www.lua.org/manual/5.3/manual.html#lua_load
func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk) // todo
	c := newLuaClosure(proto)
	self.stack.push(c)
	return 0
}

// [-(nargs+1), +nresults, e]
// http://www.lua.org/manual/5.3/manual.html#lua_call
func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	if c, ok := val.(*closure); ok {
		if c.proto != nil {	//如果proto是空的话,代表是原生语言的函数,可以直接进入下个else
			self.callLuaClosure(nArgs, nResults, c)
		} else {	//调用go函数,如果失败,则没有函数
			self.callGoClosure(nArgs,nResults,c)
		}
	} else {
		panic("not function!")
	}
}

//函数调用
func (self *luaState) callLuaClosure(nArgs,nResults int,c *closure) {
	nRegs := int(c.proto.MaxStackSize)//先找参数
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1	//是否是可变长

	newStack := newLuaStack(nRegs + 20,self)	//进行适当扩大(为了预留少量栈空间)
	newStack.closure = c

	//调用栈创建好了后,将参数一次性从当前栈pop出来,然后压入新的调用栈中
	funcAndArgs := self.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:],nParams)
	newStack.top = nRegs
	if nArgs > nParams && isVararg {
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	//然后将新调用栈作为栈顶,一个调用栈类似于C++中一个栈帧,然后执行closure
	if nArgs > nParams && isVararg {
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	self.pushLuaStack(newStack)
	self.runLuaClosure()
	self.popLuaStack()

	//执行结束
	if nResults != 0 {
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results,nResults)
	}
}

//执行函数closure
func (self *luaState) runLuaClosure(){
	for{
		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)
		if inst.OpCode() == vm.OP_RETURN {
			break
		}
	}
}

//调用Go函数
func (self *luaState) callGoClosure(nArgs,nResults int,c *closure) {
	newStack := newLuaStack(nArgs + 20,self)	//预留大小
	newStack.closure = c

	args := self.stack.popN(nArgs)
	newStack.pushN(args,nArgs)
	self.stack.pop()
	self.pushLuaStack(newStack)
	r := c.goFunc(self)
	self.popLuaStack()

	if nResults != 0 {
		results := newStack.popN(r)
		self.stack.check(len(results))
		self.stack.pushN(results,nResults)
	}
}