package comparisons

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

func (*DCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (*DCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
