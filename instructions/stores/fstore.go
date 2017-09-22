package stores

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//和 加载 指令 刚好 相反， 存储 指令 把 变量 从 操 作数 栈 顶 弹出， 然后 存入 局部 变 量表。

func _fstore(frame *rtda.Frame,index uint)  {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index,val)
}

type FSTORE struct {base.Index8Instruction}

func (this *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame,uint(this.Index))
}

type FSTORE_0 struct {base.NoOperandsInstruction}
func (this *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame,0)
}
type FSTORE_1 struct {base.NoOperandsInstruction}
func (this *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame,1)
}
type FSTORE_2 struct {base.NoOperandsInstruction}
func (this *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame,2)
}
type FSTORE_3 struct {base.NoOperandsInstruction}
func (this *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame,3)
}
