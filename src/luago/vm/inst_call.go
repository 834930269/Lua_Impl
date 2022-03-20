package vm
import . "luago/api"
import "fmt"

//CLOSURE指令,把当前Lua函数的子函数原型实例化成闭包,放入A指定的寄存器中
func closure(i Instruction, vm LuaVM) {
	a, bx := i.ABx()
	a += 1

	vm.LoadProto(bx)
	vm.Replace(a)
}

//call指令
func call(i Instruction,vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	fmt.Printf("preCall ")
	vm.PrintStack()
	nArgs := _pushFuncAndArgs(a,b,vm)
	vm.Call(nArgs,c-1)
	_popResults(a,c,vm)
	fmt.Printf("postCall ")
	vm.PrintStack()
}

//将函数和参数推入栈顶
func _pushFuncAndArgs(a,b int,vm LuaVM) (nArgs int) {
	if b >= 1 {
		vm.CheckStack(b)
		for i := a; i < a + b; i++ {
			vm.PushValue(i)
		}
		return b-1
	}else {
		//未指定返回值数量时,返回全部
		_fixStack(a,vm)
		return vm.GetTop() - vm.RegisterCount() - 1
	}
}


//弹出结果,pop主要是将返回值移到对应的寄存器内
func _popResults(a,c int,vm LuaVM) {
	if c == 1 {//无结果
	}else if c > 1 {
		for i := a+c-2;i>=a;i--{
			vm.Replace(i)
		}
	}else{
		vm.CheckStack(1)
		vm.PushInteger(int64(a))
	}
}

//如果函数中有未知返回个数,就需要等待执行完毕后，检查栈顶元素(即返回个数)
func _fixStack(a int,vm LuaVM) {
	x := int(vm.ToInteger(-1))	//栈顶的
	vm.Pop(1)

	vm.CheckStack(x - a)//x-函数寄存器索引(这部分是把原来的函数前面的参数部分重新压入栈顶(冗余了))
	for i := a; i < x;i++ {
		vm.PushValue(i)
	}
	//函数需要的参数都在栈顶,但是反的
	vm.Rotate(vm.RegisterCount() + 1,x-a)
}

func _return(i Instruction,vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1

	if b == 1 {
	}else if b > 1 {
		vm.CheckStack(b - 1)
		for i := a;i<=a+b-2;i++{
			vm.PushValue(i)
		}
	}else{
		_fixStack(a,vm)	//将函数另一部分推入栈顶,然后旋转
	}
}

func vararg(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1

	if b != 1 { // b==0 or b>1
		vm.LoadVararg(b - 1)
		_popResults(a, b, vm)
	}
}

//尾递归,尾递归不需要上一个函数的环境了,直接返回主函数
// return R(A)(R(A+1), ... ,R(A+B-1))
func tailCall(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1

	// todo: optimize tail call!
	c := 0
	nArgs := _pushFuncAndArgs(a, b, vm)
	vm.Call(nArgs, c-1)
	_popResults(a, c, vm)
}

// R(A+1) := R(B); R(A) := R(B)[RK(C)]
func self(i Instruction, vm LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1

	vm.Copy(b, a+1)
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}
