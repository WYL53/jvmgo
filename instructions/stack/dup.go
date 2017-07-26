package stack

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//dup 系列 指令 复制 栈 顶 变量


type DUP struct {
	base.NoOperandsInstruction
}

func (this *DUP) Executor(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

type DUP_X1 struct {
	base.NoOperandsInstruction
}
func (this *DUP_X1) Executor(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}
func (this *DUP_X2) Executor(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}


type DUP2 struct {
	base.NoOperandsInstruction
}
func (this *DUP2) Executor(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}
func (this *DUP2_X1) Executor(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}
func (this *DUP2_X2) Executor(frame *rtda.Frame) {
	stack:=frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
