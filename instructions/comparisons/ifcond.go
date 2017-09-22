package comparisons

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type IFEQ struct{ base.BranchInstruction }

func (this *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (this *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (this *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (this *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (this *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (this *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branck(frame, this.Offset)
	}
}
