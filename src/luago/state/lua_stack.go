package state

//这个类中定义了对Lua栈的所有操作

//定义Lua栈
type luaStack struct{
	slots []luaValue		//栈中每个元素都是luaValue(类似C中union)
	top int
}

//创建一个新的luaStack
func newLuaStack(size int) *luaStack{
	return &luaStack{
		slots: make([]luaValue,size),
		top:   0,
	}
}

//check方法用于检查栈的空闲空间是否还可以容纳(推入)至少n个值
//不满足,则进行扩容,这里是2倍扩容
func (self *luaStack) check(n int){
	free := len(self.slots) - self.top
	for i := free; i<len(self.slots)*2; i++{
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
	return idx + self.top + 1
}

//判断索引是否有效
func (self *luaStack) isValid(idx int) bool{
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

//根据索引从栈中取值,如果索引无效则返回nil值
func (self *luaStack) get(idx int) luaValue{
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1]
	}
	return nil
}

//根据索引往栈里写入值
func (self *luaStack) set(idx int,val luaValue){
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

