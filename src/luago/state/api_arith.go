package state

import "math"
import . "luago/api"
import "luago/number"

var (
	iadd 	= func(a,b int64) 		int64 		{ return a + b }
	fadd 	= func(a,b float64) 	float64		{ return a + b }
	isub 	= func(a,b int64)		int64		{ return a - b }
	fsub 	= func(a,b float64) 	float64		{ return a - b }
	imul 	= func(a,b int64)		int64		{ return a * b }
	fmul 	= func(a,b float64)		float64		{ return a * b }
	imod 	= number.IMod
	fmod 	= number.FMod
	pow  	= math.Pow
	div  	= func(a,b float64)		float64		{ return a / b }
	iidiv 	= number.IFloorDiv
	fidiv 	= number.FFloorDiv
	band 	= func(a,b int64)		int64		{ return a & b }
	bor		= func(a,b int64)		int64		{ return a | b }
	bxor	= func(a,b int64)		int64		{ return a ^ b }
	shl		= number.ShiftLeft
	shr		= number.ShiftRight
	iunm	= func(a, _ int64)		int64		{   return -a  }
	funm	= func(a, _ float64)	float64		{   return -a  }
	bnot	= func(a, _ int64)		int64		{   return ^a  }
)


type operator struct{
	integerFunc		func(int64, int64) 		int64		//int类型的运算方法
	floatFunc		func(float64,float64)	float64		//float类型的运算方法
}

//不同操作的int和float对应的函数
//和consts 中的 arithmetic functions 一一对应
var operators = []operator{
	operator{iadd,fadd},
	operator{isub,fsub},
	operator{imul,fmul},
	operator{imod,fmod},
	operator{nil,pow},
	operator{nil,div},
	operator{iidiv,fidiv},
	operator{band,nil},
	operator{bor,nil},
	operator{bxor,nil},
	operator{shl,nil},
	operator{shr,nil},
	operator{iunm,funm},
	operator{bnot,nil},
}

//arithmetic - 算术运算
func (self *luaState) Arith(op ArithOp) {
	var a, b luaValue	//operands(操作数)
	b = self.stack.pop()
	if op != LUA_OPUNM && op != LUA_OPBNOT{
		a = self.stack.pop()
	}else {
		a = b
	}

	operator := operators[op]
	if result := _arith(a,b,operator); result != nil{
		self.stack.push(result)
	}else{
		panic("arithmetic error!")
	}
}

//实际的调用
func _arith(a,b luaValue,op operator) luaValue {
	if op.floatFunc == nil {	//bitwise,只有整数,即bit操作
		if x, ok := convertToInteger(a); ok {
			if y,ok := convertToInteger(b); ok{
				return op.integerFunc(x, y)
			}
		}
	} else {	//arith
		if op.integerFunc != nil {	//add,sub,mul,mod,idiv,unm
			if x, ok := a.(int64); ok{
				if y, ok := b.(int64); ok{
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := convertToFloat(a); ok{
			if y, ok := convertToFloat(b); ok{
				return op.floatFunc(x,y)
			}
		}
	}
	return nil
}