package vm

import . "luago/api"

//把a处的值移到b处
func move(i Instruction,vm LuaVM){
	a, b, _ := i.ABC()
	a += 1; b += 1
	vm.Copy(b,a)
}

//跳转指令
func jmp(i Instruction,vm LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}

/*---------加载指令集---------*/
//将a到a+b的寄存器全部设置为nil
func loadNil(i Instruction,vm LuaVM){
	a, b, _ := i.ABC()
	a += 1

	vm.PushNil()//先往栈顶push一个nil
	for i := a; i <= a+b; i++ {
		vm.Copy(-1,i)//设置其余的都是nil
	}
	vm.Pop(1)//然后推出
}

//布尔值由B指定(0表示false,非0表示true)
//寄存器C非0的话则跳过下一条指令
func loadBool(i Instruction,vm LuaVM) {
	a,b,c := i.ABC()
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}