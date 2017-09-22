package stack

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//pop 和 pop2 指令 将 栈 顶 变量 弹出，

type POP struct {
	base.NoOperandsInstruction
}

func (this *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}
func (this *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
