package constants

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)
//const 系列 指令
// 这一 系列 指令 把 隐含 在 操作 码 中的 常量 值 推入 操 作数 栈 顶。

type ACONST_NULL struct {base.NoOperandsInstruction}

func (this *ACONST_NULL) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct {base.NoOperandsInstruction}
func (this *DCONST_0) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {base.NoOperandsInstruction}
func (this *DCONST_1) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct {base.NoOperandsInstruction}
func (this *FCONST_0) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct {base.NoOperandsInstruction}
func (this *FCONST_1) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct {base.NoOperandsInstruction}
func (this *FCONST_2) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

type ICONST_M1 struct {base.NoOperandsInstruction}
func (this *ICONST_M1) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {base.NoOperandsInstruction}
func (this *ICONST_0) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {base.NoOperandsInstruction}
func (this *ICONST_1) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {base.NoOperandsInstruction}
func (this *ICONST_2) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {base.NoOperandsInstruction}
func (this *ICONST_3) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {base.NoOperandsInstruction}
func (this *ICONST_4) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {base.NoOperandsInstruction}
func (this *ICONST_5) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct {base.NoOperandsInstruction}
func (this *LCONST_0) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {base.NoOperandsInstruction}
func (this *LCONST_1) Executor(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}