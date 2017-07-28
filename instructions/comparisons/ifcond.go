package comparisons

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type IFEQ struct{ base.BranchInstruction }

func (this *IFEQ) Executor(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (this *IFNE) Executor(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (this *IFLT) Executor(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (this *IFLE) Executor(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (this *IFGT) Executor(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branck(frame, this.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (this *IFGE) Executor(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branck(frame, this.Offset)
	}
}
