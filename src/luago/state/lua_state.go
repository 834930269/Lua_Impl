package state

import "fmt"
import . "luago/api"

//LuaState的实现类
//关于栈的动作全部放回到closure
type luaState struct{
	stack 	 *luaStack
	registry *luaTable	//注册表,
}

func (self *luaState) PrintStack() {
	k := self.stack
	for i := 1; i <= k.top; i++ {
		t := typeOf(k.get(i))
		switch t {
		case LUA_TBOOLEAN:
			fmt.Printf("[%t]", self.ToBoolean(i))
		case LUA_TNUMBER:
			fmt.Printf("[%g]", self.ToNumber(i))
		case LUA_TSTRING:
			fmt.Printf("[%q]", self.ToString(i))
		default: // other values
			fmt.Printf("[%s]", self.TypeName(t))
		}
	}
	fmt.Println()
}

func New() *luaState{
	registry := newLuaTable(0,0)	//注册表,注册表本身是一个表,但是全局表也是一个表,放在注册表的LUA_RIDX_GLOBALS下
	registry.put(LUA_RIDX_GLOBALS,newLuaTable(0,0))	//全局环境

	ls := &luaState{registry:registry}
	ls.pushLuaStack(newLuaStack(LUA_MINSTACK,ls))

	return ls
}

//向栈顶推入一个lua栈
func (self *luaState) pushLuaStack(stack *luaStack) {
	stack.prev = self.stack
	self.stack = stack
}

//删除栈顶的第一个栈
func (self *luaState) popLuaStack() {
	stack := self.stack
	self.stack = stack.prev
	stack.prev = nil 
}



