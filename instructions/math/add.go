package math

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (*IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

type LADD struct {
	base.NoOperandsInstruction
}

func (*LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}


type FADD struct {
	base.NoOperandsInstruction
}

func (*FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}


type DADD struct {
	base.NoOperandsInstruction
}

func (*DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}