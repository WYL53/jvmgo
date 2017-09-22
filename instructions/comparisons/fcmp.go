package comparisons

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//这 两条 指令 和 lcmp 指令 很像， 但是 除了 比较 的 变量 类型 不同 以外，
// 还有 一个 重要的 区别。 由于 浮点 数 计算 有可能 产生 NaN（ Not a Number） 值，
// 所以 比较 两个 浮 点数 时， 除了 大于、 等于、 小于 之外， 还有 第 4 种 结果： 无法 比较。
type FCMPG struct {
	base.NoOperandsInstruction
}

func (*FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (*FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
