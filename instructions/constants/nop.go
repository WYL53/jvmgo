package constants

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (this *NOP)Execute(frame *rtda.Frame)  {
	//什么也不做
}