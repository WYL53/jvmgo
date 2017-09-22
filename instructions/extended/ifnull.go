package extended

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)
//根据 引用 是否 是 null 进行 跳 转， ifnull 和 ifnonnull 指令 把 栈 顶的 引用 弹出。

type IFNULL struct {
	base.BranchInstruction
}

func (this *IFNULL)Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref == nil{
		base.Branck(frame,this.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (this *IFNONNULL)Execute(frame *rtda.Frame)  {
	ref := frame.OperandStack().PopRef()
	if ref != nil{
		base.Branck(frame,this.Offset)
	}
}