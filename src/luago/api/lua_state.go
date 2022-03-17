package api

type LuaType = int
type ArithOp = int
type CompareOp = int	

type LuaState interface {
	/* 基础栈操作 */
	GetTop()								int
	AbsIndex(idx int)						int
	CheckStack(n int)   					bool
	Pop(n int)			
	Copy(fromIdx,toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx,n int)
	SetTop(idx int)
	TypeName(tp LuaType) 					string
	Type(idx int)							LuaType
	IsNone(idx int)							bool
	IsNil(idx int)							bool
	IsNoneOrNil(idx int)					bool
	IsBoolean(idx int)						bool
	IsInteger(idx int)						bool
	IsNumber(idx int)						bool
	IsString(idx int)						bool
	ToBoolean(idx int) 						bool
	ToInteger(idx int)						int64
	ToIntegerX(idx int)						(int64,bool)
	ToNumber(idx int)						float64
	ToNumberX(idx int)						(float64,bool)
	ToString(idx int)						string
	ToStringX(idx int)						(string,bool)
	/* push函数 go->stack */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)

	//执行算术和按位运算
	Arith(op ArithOp)
	//比较运算
	Compare(idx1,idx2 int, op CompareOp) 	bool
	Len(idx int)
	//string的concat方法
	Concat(n int)
}