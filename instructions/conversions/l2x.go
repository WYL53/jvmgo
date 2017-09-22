package conversions

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

type L2F struct {
	base.NoOperandsInstruction
}

func (*L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

type L2I struct {
	base.NoOperandsInstruction
}

func (*L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}

type L2D struct {
	base.NoOperandsInstruction
}

func (*L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

//type I2B struct{
//	base.NoOperandsInstruction
//}
//
//func (*I2B) Execute(frame *rtda.Frame) {
//	stack := frame.OperandStack()
//	l := stack.()
//	d := float64(l)
//	stack.PushDouble(d)
//}
//
//type I2C struct{
//	base.NoOperandsInstruction
//}
//
//func (*I2C) Execute(frame *rtda.Frame) {
//	stack := frame.OperandStack()
//	l := stack.PopLong()
//	d := float64(l)
//	stack.PushDouble(d)
//}
//
//type I2S struct{
//	base.NoOperandsInstruction
//}
//
//func (*I2S) Execute(frame *rtda.Frame) {
//	stack := frame.OperandStack()
//	l := stack.PopLong()
//	d := float64(l)
//	stack.PushDouble(d)
//}
