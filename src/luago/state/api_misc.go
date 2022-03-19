package state 


/*
len方法等价于(#a)
作用是访问a处的索引值的值(a),然后将a的长度压入栈顶
暂时只考虑string,后面再完善
*/
func (self *luaState) Len(idx int){
	val := self.stack.get(idx)
	if s, ok := val.(string); ok{
		self.stack.push(int64(len(s)))
	} else if t, ok := val.(*luaTable); ok {
		self.stack.push(int64(t.len()))	//表的大小
	} else {
		panic("length error!")
	}
}


// [-n, +1, e]从栈顶取出n个元素,然后连接到一起方到栈顶
// http://www.lua.org/manual/5.3/manual.html#lua_concat
func (self *luaState) Concat(n int) {
	if n == 0 {
		self.stack.push("")
	} else if n >= 2 {
		for i := 1; i < n; i++ {
			if self.IsString(-1) && self.IsString(-2) {
				s2 := self.ToString(-1)
				s1 := self.ToString(-2)
				self.stack.pop()
				self.stack.pop()
				self.stack.push(s1 + s2)
				continue
			}

			panic("concatenation error!")
		}
	}
	// n == 1, do nothing
}