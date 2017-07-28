package conversions

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (*D2F) Executor(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (*D2I) Executor(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (*D2L) Executor(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int64(d)
	stack.PushLong(i)
}
