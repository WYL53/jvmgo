package comparisons

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//if_ acmpeq 和 if_ acmpne 指令 把 栈 顶的 两个 引用 弹出， 根据 引用 是否 相同 进行 跳 转。

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (this *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branck(frame, this.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (this *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branck(frame, this.Offset)
	}
}
