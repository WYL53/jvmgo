package control

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//tableswitch 指令 操作 码 的 后面 有 0~ 3 字节 的 padding，
// 以 保证 defaultOffset 在 字节 码 中的 地址 是 4 的 倍数。

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (this *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.low = reader.ReadInt32()
	this.high = reader.ReadInt32()
	jumpOffsetCount := this.high - this.low + 1
	this.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}

func (this *TABLE_SWITCH) Executor(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= this.low && index <= this.high {
		offset = int(this.jumpOffsets[index-this.low])
	} else {
		offset = int(this.defaultOffset)
	}
	base.Branck(frame, offset)
}
