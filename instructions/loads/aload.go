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

//iload 指令 的 索引 来自 操作 数，
//iload_ 1指令索引 隐含 在 操作 码 中

func _aload(frame *rtda.Frame,index uint)  {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

type ALOAD struct {base.Index8Instruction}

func (this *ALOAD) Execute(frame *rtda.Frame) {
	_dload(frame,uint(this.Index))
}

type ALOAD_0 struct {base.NoOperandsInstruction}
func (this *ALOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame,0)
}
type ALOAD_1 struct {base.NoOperandsInstruction}
func (this *ALOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame,1)
}
type ALOAD_2 struct {base.NoOperandsInstruction}
func (this *ALOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame,2)
}
type ALOAD_3 struct {base.NoOperandsInstruction}
func (this *ALOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame,3)
}
