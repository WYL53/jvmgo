package stores

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//和 加载 指令 刚好 相反， 存储 指令 把 变量 从 操 作数 栈 顶 弹出， 然后 存入 局部 变 量表。

func _astore(frame *rtda.Frame,index uint)  {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index,val)
}

type ASTORE struct {base.Index8Instruction}

func (this *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame,uint(this.Index))
}

type ASTORE_0 struct {base.NoOperandsInstruction}
func (this *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame,0)
}
type ASTORE_1 struct {base.NoOperandsInstruction}
func (this *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame,1)
}
type ASTORE_2 struct {base.NoOperandsInstruction}
func (this *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame,2)
}
type ASTORE_3 struct {base.NoOperandsInstruction}
func (this *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame,3)
}
