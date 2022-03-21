package state

import . "luago/api"

//这个类中定义了对Lua栈的所有操作

//定义Lua栈
//当lua进行到函数阶段后,所有的操作都从luaState里的一个stack变成每个函数一个stack
//而函数与函数之间形成了链表
type luaStack struct{
	slots 		[]luaValue		//栈中每个元素都是luaValue(类似C中union)
	top 		int
	// 下面是新增加的字段
	prev 		*luaStack
	closure		*closure
	varargs		[]luaValue
	pc			int
	state 		*luaState
}

//创建一个新的luaStack
func newLuaStack(size int,state *luaState) *luaStack{
	return &luaStack{
		slots: make([]luaValue,size),
		top:   0,
		state: state,
	}
}

//check方法用于检查栈的空闲空间是否还可以容纳(推入)至少n个值
//不满足,则进行扩容,这里是2倍扩容
func (self *luaStack) check(n int){
	free := len(self.slots) - self.top
	newN := len(self.slots) * 2
	for i := free; i<newN; i++{
		self.slots = append(self.slots,nil)
	}
}

//push方法将值推入栈顶,如果溢出,则先panic终止程序
func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots){
		panic("stack overBow!")
	}
	self.slots[self.top] = val
	self.top++
}

//从栈顶pop出一个值
func (self *luaStack) pop() luaValue{
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

//相对索引改成绝对索引
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0{
		return idx
	}
	if idx <= LUA_REGISTRYINDEX {
		return idx	//如果比LUA_REGIS...小,代表是伪索引
	}
	return idx + self.top + 1
}

//判断索引是否有效
func (self *luaStack) isValid(idx int) bool{
	if idx == LUA_REGISTRYINDEX {//注册表伪索引是
		return true
	}
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

//根据索引从栈中取值,如果索引无效则返回nil值
func (self *luaStack) get(idx int) luaValue{
	if idx == LUA_REGISTRYINDEX {
		return self.state.registry
	}
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1]
	}
	return nil
}

//根据索引往栈里写入值
func (self *luaStack) set(idx int,val luaValue){
	if idx == LUA_REGISTRYINDEX {
		self.state.registry = val.(*luaTable)
		return 
	}
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

//旋转,这个方式是整个旋转
func (self *luaStack) reverse(from,to int){
	slots := self.slots
	for from < to {
		slots[from],slots[to] = slots[to],slots[from]
		from++
		to--
	}
}

//从栈顶pop出N个
func (self *luaStack) popN(n int) []luaValue{
	vals := make([]luaValue,n)
	for i:=n-1;i>=0;i-- {
		vals[i] = self.pop()
	}
	return vals
}

//向栈顶push n个LuaValue
func (self *luaStack) pushN(vals []luaValue,n int){
	nVals := len(vals)
	if n<0 {n = nVals}
	for i := 0; i < n; i++ {
		if i < nVals {
			self.push(vals[i])
		} else {
			self.push(nil)
		}
	}
}