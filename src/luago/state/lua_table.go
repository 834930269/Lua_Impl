package state
import "math"
import "luago/number"

type luaTable struct {
	arr 	[]luaValue
	_map 	map[luaValue]luaValue
}

func newLuaTable(nArr,nRec int)*luaTable {
	t := &luaTable{}
	if nArr > 0 {
		t.arr = make([]luaValue,0,nArr)
	}
	if nRec > 0 {
		t._map = make(map[luaValue]luaValue,nRec)
	}
	return t
}

//从表中获取一个元素
func (self *luaTable) get(key luaValue) luaValue{
	key = _floatToInteger(key)
	//如果键是整数,检查是否在数组范围内,是的话返回数组第key项
	if idx, ok := key.(int64); ok {
		if idx >= 1 && idx <= int64(len(self.arr)) {
			return self.arr[idx - 1]
		}
	}
	//否则返回map内的
	return self._map[key]
}

//尝试将float转换成int
func _floatToInteger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i,ok := number.FloatToInteger(f); ok {
			return i
		}
	}
	return key
}

func (self *luaTable) put(key,val luaValue){
	if key == nil {
		panic("table index is nil!")
	}
	if f, ok := key.(float64); ok && math.IsNaN(f){
		panic("table index is NaN!")
	}
	//添加到数组中
	key = _floatToInteger(key)
	if idx, ok := key.(int64); ok && idx >= 1 {
		arrLen := int64(len(self.arr))
		if idx <= arrLen {
			self.arr[idx - 1] = val
			if idx == arrLen && val == nil {
				self._shrinkArray()	//删除洞
			}
			return 
		}
		//将值放到尾部
		if idx == arrLen+1{
			delete(self._map,key)
			if val != nil {
				self.arr = append(self.arr,val)
				self._expandArray()//扩容后检查
			}
			return
		}
	}

	//添加到map
	if val != nil {
		if self._map == nil {
			self._map = make(map[luaValue]luaValue,8)
		}
		self._map[key] = val
	} else {
		delete(self._map,key)
	}
	
}

//删除洞
func (self *luaTable) _shrinkArray() {
	for i := len(self.arr) - 1; i >= 0; i--{
		if self.arr[i] == nil {
			self.arr = self.arr[0:i]
		}
	}
}

//在数组扩容后,将部分map中的元素
func (self *luaTable) _expandArray(){
	for idx := int64(len(self.arr)) + 1; true; idx++ {
		if val, found := self._map[idx]; found {
			delete(self._map,idx)
			self.arr = append(self.arr,val)
		} else {
			break
		}
	}
}

func (self *luaTable) len() int {
	return len(self.arr)
}
