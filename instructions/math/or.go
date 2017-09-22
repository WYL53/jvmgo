package math

import (
	"github.com/WYL53/jvmgo/rtda"
	"github.com/WYL53/jvmgo/instructions/base"
)


type IOR struct {
	base.NoOperandsInstruction
}

func (*IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (*LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
