package math

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}
func (*DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1,v2)
	stack.PushDouble(result)
}

type FREM struct {
	base.NoOperandsInstruction
}
func (*FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := math.Mod(float64(v1),float64(v2))
	stack.PushFloat(float32(result))
}

type IREM struct {
	base.NoOperandsInstruction
}
func (*IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0{
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct {
	base.NoOperandsInstruction
}
func (*LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0{
		panic("java.lang.ArithmeticException:/ by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}