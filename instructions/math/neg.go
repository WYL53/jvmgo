package math

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type INEG struct {
	base.NoOperandsInstruction
}

func (*INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	stack.PushInt(-v)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (*LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopLong()
	stack.PushLong(-v)
}


type FNEG struct {
	base.NoOperandsInstruction
}

func (*FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopFloat()
	stack.PushFloat(-v)
}


type DNEG struct {
	base.NoOperandsInstruction
}

func (*DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	stack.PushDouble(-v)
}

