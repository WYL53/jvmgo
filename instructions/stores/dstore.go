package stores

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//和 加载 指令 刚好 相反， 存储 指令 把 变量 从 操 作数 栈 顶 弹出， 然后 存入 局部 变 量表。

func _dstore(frame *rtda.Frame,index uint)  {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index,val)
}

type DSTORE struct {base.Index8Instruction}

func (this *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame,uint(this.Index))
}

type DSTORE_0 struct {base.NoOperandsInstruction}
func (this *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame,0)
}
type DSTORE_1 struct {base.NoOperandsInstruction}
func (this *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame,1)
}
type DSTORE_2 struct {base.NoOperandsInstruction}
func (this *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame,2)
}
type DSTORE_3 struct {base.NoOperandsInstruction}
func (this *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame,3)
}
