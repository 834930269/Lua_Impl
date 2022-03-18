package vm

import . "luago/api"

//将常量表索引bx处的值加入到a寄存器中
func loadK(i Instruction,vm LuaVM) {
	a,bx := i.ABx()
	a += 1

	vm.GetConst(bx)	//GetConst会从常量表中获取常量然后推入栈顶
	vm.Replace(a)	//再将栈顶元素移到a位置
}

//如果常量表超过2^18调用这个
func loadKx(i Instruction,vm LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()

	vm.GetConst(ax)
	vm.Replace(a)
}

