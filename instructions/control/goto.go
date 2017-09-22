package control

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (this *GOTO) Execute(frame *rtda.Frame) {
	base.Branck(frame, this.Offset)
}
