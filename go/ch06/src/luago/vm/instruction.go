package vm

import "luago/api"

type Instruction uint32

const MAXARG_Bx = 1<<18 - 1       // 2^18 - 1 = 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 262143 / 2 = 131071

//提取操作码
func (self Instruction)Opcode() int{
	return int(self & 0x3F)
}

func (self Instruction)ABC() (a,b,c int){
	a = int(self >> 6 & 0XFF)
	c = int(self >> 14 & 0X1FF)
	b = int(self >> 23 & 0X1FF)
	return
}

func (self Instruction)ABx() (a,bx int){
	a = int(self >> 6 & 0xFF)
	bx = int(self >> 14)
	return 
}

func (self Instruction)AsBx() (a,sbx int){
	a,bx := self.ABx()
	return a,bx - MAXARG_sBx
}

func (self Instruction)Ax() int {
	return int(self >>6)
}

//操作码名字
func (self Instruction) OpName() string {
    return opcodes[self.Opcode()].name
}
//编码模式
func (self Instruction) OpMode() byte {
    return opcodes[self.Opcode()].opMode
}
//操作数B的使用模式
func (self Instruction) BMode() byte {
    return opcodes[self.Opcode()].argBMode
}
//操作数C的使用模式
func (self Instruction) CMode() byte {
    return opcodes[self.Opcode()].argCMode
}

func (self Instruction) Execute(vm api.LuaVM) {
	action := opcodes[self.Opcode()].action
	if action != nil {
		action(self, vm)
	} else {
		panic(self.OpName())
	}
}