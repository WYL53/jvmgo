package constants

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)
//bipush 指令 从 操 作数 中 获取 一个 byte 型 整数， 扩展 成 int 型， 然后 推 入栈 顶。
//sipush 指令 从 操 作数 中 获取 一个 short 型 整数， 扩展 成 int 型， 然后 推 入栈 顶。

type BIPUSH struct {
	val int8
}

func (this *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}

func (this *BIPUSH) Executor(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (this *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}

func (this *SIPUSH) Executor(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

