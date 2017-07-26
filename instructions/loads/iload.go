package loads

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//加载 指令 从 局部 变量 表 获取 变量， 然后 推入 操 作数 栈 顶。
// 加载 指令 共 33 条， 按照 所 操作 变量 的 类型 可以 分为 6 类：
// aload 系列 指令 操作 引用 类型 变量、 dload 系列 操作 double 类型 变量、
// fload 系列 操作 float 变量、 iload 系列 操作 int 变量、 lload 系列 操作 long 变量、
// xaload 操作 数组。

func _iload(frame *rtda.Frame,index uint)  {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

type ILOAD struct {base.Index8Instruction}

func (this *ILOAD) Executor(frame *rtda.Frame) {
	_iload(frame,uint(this.Index))
}

type ILOAD_0 struct {base.NoOperandsInstruction}
func (this *ILOAD_0) Executor(frame *rtda.Frame) {
	_iload(frame,0)
}
type ILOAD_1 struct {base.NoOperandsInstruction}
func (this *ILOAD_1) Executor(frame *rtda.Frame) {
	_iload(frame,1)
}
type ILOAD_2 struct {base.NoOperandsInstruction}
func (this *ILOAD_2) Executor(frame *rtda.Frame) {
	_iload(frame,2)
}
type ILOAD_3 struct {base.NoOperandsInstruction}
func (this *ILOAD_3) Executor(frame *rtda.Frame) {
	_iload(frame,3)
}
