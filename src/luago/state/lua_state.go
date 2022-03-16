package state

//LuaState的实现类

type luaState struct{
	stack *luaStack
}

func New() *luaState{
	return &luaState{
		stack: newLuaStack(20),
	}
}