package math

import (
	"github.com/WYL53/jvmgo/rtda"
	"github.com/WYL53/jvmgo/instructions/base"
)

//异或

type IXOR struct {
	base.NoOperandsInstruction
}

func (*IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (*LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
