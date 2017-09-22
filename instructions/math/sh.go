package math

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//位移 指令 可以 分为 左移 和 右移 两种，
// 右移 指令 又可 以 分为 算术 右移（ 有 符号 右移） 和 逻辑 右移（ 无符号 右移） 两种。

//算术 右移 和 逻辑 位移 的 区别 仅在 于 符号 位 的 扩展， 如下 面的 Java 代码 所示。
// int x = -1;
// println( Integer. toBinaryString( x)); //11111111111111111111111111111111
// println( Integer. toBinaryString( x >> 8)); // 11111111111111111111111111111111
// println( Integer. toBinaryString( x >>> 8)); // 00000000111111111111111111111111

//逻辑右移，移走的位填充为0；
//算术右移，移走的位填充与符号位有关，例如如果为负数，则移走的位填充为1。

//int左移
type ISHL struct {
	base.NoOperandsInstruction
}
func (*ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

//int算数右移
type ISHR struct {
	base.NoOperandsInstruction
}
func (*ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}


//int逻辑右移
type IUSHR struct {
	base.NoOperandsInstruction
}
func (*IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

//long左移
type LSHL struct {
	base.NoOperandsInstruction
}
func (*LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

//long算数右移
type LSHR struct {
	base.NoOperandsInstruction
}
func (*LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//long逻辑右移
type LUSHR struct {
	base.NoOperandsInstruction
}
func (*LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x1f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}