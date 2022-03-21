package state
import "luago/binchunk"
import . "luago/api"

type closure struct {
	proto *binchunk.Prototype
	goFunc GoFunction			//go函数渗透部分
}

//创建一个新的closure
func newLuaClosure(proto *binchunk.Prototype) *closure{
	return &closure{proto: proto}
}

func newGoClosure(f GoFunction) *closure{
	return &closure{goFunc:f}
}