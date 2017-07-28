package math

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (this *IINC) FetchOperands(reader *base.BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
	this.Const = int32(reader.ReadInt8())
}

func (this *IINC) Executor(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(this.Index)
	val += this.Const
	localVars.SetInt(this.Index, val)
}