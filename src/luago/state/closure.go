package state
import "luago/binchunk"

type closure struct {
	proto *binchunk.Prototype
}

//创建一个新的closure
func newLuaClosure(proto *binchunk.Prototype) *closure{
	return &closure{proto: proto}
}

