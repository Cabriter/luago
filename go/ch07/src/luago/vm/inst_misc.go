package vm
// 其他类型指令
import . "luago/api"

// MOVE
func move(i Instruction,vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1;
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm LuaVM) {
	a,sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo !")
	}
}