package binchunk

/*
go结构体  成员  成员类型
*/

type binaryChunk struct{
	header					//头部
	sizeUpValues byte   	//主函数upvalue数量
	mainFunc 	 *Prototype //主函数原型
}

type header struct{
	//签名,很多语言的chunk都会以一个Magic Number开始
	//主要是用来标识文件类型,如Lua就是'ESC'Lua这些字符的ASCII码
	signature			[4]byte	
	//第二个参数记录的是Lua版本号,由三部分组成
	//大版本号,小版本号,发布版本号.如5.3.4
	version				byte
	//格式号,会检查，官方的是0
	format				byte
	//起进一步校验的作用
	luacData			[6]byte
	//cint在虚拟机中的字节数
	cinSize				byte
	//size_t在虚拟机中的字节数
	sizetSize			byte
	//Lua虚拟机指令在虚拟机中的字节数
	instructionSize		byte
	//Lua整数在虚拟机中的字节数
	luaIntegerSize		byte
	//Lua整数在虚拟机中的字节数
	luaNumberSize		byte
	//保存int的Lua整数值0x5678,主要为了检测大小端D
	luacInt				int64
	//保存浮点数370.5,检查浮点数存储格式,如IEEE 754
	luacNum				float64
}


const (
	LUA_SIGNATURE 		= "\x1bLua"
	LUAC_VERSION 		= 0x53
	LUAC_FORMAT 		= 0
	LUAC_DATA			= "\x19\x93\r\n\x1a\n"
	CINT_SIZE			= 4
	CSIZET_SIZE			= 8
	INSTRUCTION_SIZE	= 4
	LUA_INTEGER_SIZE	= 8
	LUA_NUMBER_SIZE		= 8
	LUAC_INT			= 0x5678
	LUAC_NUM			= 370.5
)

const (
	TAG_NIL 			= 0x00
	TAG_BOOLEAN			= 0x01
	TAG_NUMBER			= 0x03
	TAG_INTEGER			= 0x13
	TAG_SHORT_STR		= 0x04
	TAG_LONG_STR		= 0x14
)


//函数原型
type Prototype struct{
	//源文件名,记录只有主函数原型里才会有
	Source			string
	//起止行号,普通函数,起止行号一般都大于0
	//主函数一般起止行号都是0
	LineDefined		uint32
	//止行号
	LastLineDefined	uint32
	//固定参数个数,如函数参数,main函数是0
	NumParams		byte
	//是否是可变长参数函数
	IsVararg		byte
	//寄存器数量,记录运行一个函数需要用到多少寄存器
	MaxStackSize	byte
	//指令表,每条指令占4字节
	Code			[]uint32
	//常量表,用于记录Lua中出现的字面量,如nil,布尔值,整数,浮点数,字符串
	//每个常量都以1字节tag开头,用来标识后续存放哪种类型的常量值
	//tag参数表见P25,在C中,这个是union
	Constants		[]interface{}
	//upvalue每个元素占用2个字节,主要用在闭包中
	Upvalues		[]Upvalue
	//子函数原型表
	Protos			[]*Prototype
	//行号表,和指令表一一对应
	LineInfo		[]uint32
	//局部变量表
	LocVars			[]LocVar
	//Upvalue名列表
	UpvalueNames	[]string
}

type Upvalue struct{
	Instack 		byte
	Idx 			byte
}

//局部变量
type LocVar struct{
	VarName			string
	StartPC			uint32
	EndPC			uint32
}

//对二进制chunk格式进行校验
func Undump(data []byte) *Prototype{
	reader := &reader{data}
	reader.checkHeader()		//校验头部
	reader.readByte()			//跳过Upvalue数量
	return reader.readProto("") //读取函数原型
}

