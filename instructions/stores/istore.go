package stores

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//和 加载 指令 刚好 相反， 存储 指令 把 变量 从 操 作数 栈 顶 弹出， 然后 存入 局部 变 量表。

func _istore(frame *rtda.Frame,index uint)  {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index,val)
}

type ISTORE struct {base.Index8Instruction}

func (this *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame,uint(this.Index))
}

type ISTORE_0 struct {base.NoOperandsInstruction}
func (this *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame,0)
}
type ISTORE_1 struct {base.NoOperandsInstruction}
func (this *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame,1)
}
type ISTORE_2 struct {base.NoOperandsInstruction}
func (this *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame,2)
}
type ISTORE_3 struct {base.NoOperandsInstruction}
func (this *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame,3)
}
