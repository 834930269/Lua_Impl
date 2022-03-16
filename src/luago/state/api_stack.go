package state

//实现api/lua_state的接口

//返回栈顶索引
func (self *luaState) GetTop() int{
	return self.stack.top
}


func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

func (self *luaState) CheckStack(n int) bool{
	self.stack.check(n)
	return true
}

//从栈顶推出n个元素
func (self *luaState) Pop(n int) {
	/*
	for i := 0; i<n;i++ {
		self.stack.pop()
	}
	*/
	self.SetTop(-n-1)
}

//fromIdx的值复制到toIdx上
func (self *luaState) Copy(fromIdx,toIdx int){
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx,val)
}

//指定索引处的值推到栈顶
func (self *luaState) PushValue(idx int){
	val := self.stack.get(idx)
	self.stack.push(val)
}

//Replace是pushValue的反操作
//将值从栈顶弹出,然后写入到idx处
func (self *luaState) Replace(idx int){
	val := self.stack.pop()
	self.stack.set(idx,val)
}

//Insert方法将栈顶值弹出,然后插入到指定位置
func (self *luaState) Insert(idx int){
	self.Rotate(idx,1)
}

//将idx位置的元素移出来
func (self *luaState) Remove(idx int){
	self.Rotate(idx,-1)
	self.Pop(1)
}

//将idx到栈顶旋转n个元素
//lua内置使用的是reverse来实现的
//通过三次子元素完整旋转来快速实现n各元素的旋转
//如果不使用旋转的话,需要平移 n*len(stack) 次
func (self *luaState) Rotate(idx,n int){
	t := self.stack.top -1
	p := self.stack.absIndex(idx) -1
	var m int
	if n >= 0 {
		m = t - n
	}else {
		m = p-n-1
	}
	self.stack.reverse(p,m)
	self.stack.reverse(m+1,t)
	self.stack.reverse(p,t)
}

//设置栈顶指针为新的
func (self *luaState) SetTop(idx int){
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack empty!")
	}
	n := self.stack.top - newTop
	if n > 0 {
		for i:=0;i<n;i++ {
			self.stack.pop()
		}
	}else if n < 0 {
		for i:=0;i>n;i-- {
			self.stack.push(nil)
		}
	}
}





