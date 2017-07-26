package stores

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//和 加载 指令 刚好 相反， 存储 指令 把 变量 从 操 作数 栈 顶 弹出， 然后 存入 局部 变 量表。

func _lstore(frame *rtda.Frame,index uint)  {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index,val)
}

type LSTORE struct {base.Index8Instruction}

func (this *LSTORE) Executor(frame *rtda.Frame) {
	_lstore(frame,uint(this.Index))
}

type LSTORE_0 struct {base.NoOperandsInstruction}
func (this *LSTORE_0) Executor(frame *rtda.Frame) {
	_lstore(frame,0)
}
type LSTORE_1 struct {base.NoOperandsInstruction}
func (this *LSTORE_1) Executor(frame *rtda.Frame) {
	_lstore(frame,1)
}
type LSTORE_2 struct {base.NoOperandsInstruction}
func (this *LSTORE_2) Executor(frame *rtda.Frame) {
	_lstore(frame,2)
}
type LSTORE_3 struct {base.NoOperandsInstruction}
func (this *LSTORE_3) Executor(frame *rtda.Frame) {
	_lstore(frame,3)
}
