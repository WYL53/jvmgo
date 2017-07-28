package control

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)

//switch是控制选择的一种方式,编译器生成代码时可以对这种结构进行特定的优化,从而产生效率比较高的代码。
// 在java中，编译器根据分支的情况，分别产生tableswitch,lookupswitch两中情况，
// 其中tableswitch适用于分支比较集中的情况，而lookupswitch适用与分支比较稀疏的情况。
// 不过怎么算稀疏，怎么算集中就是编译器的决策问题了，这里不做深入的分析。

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffset   []int32
}

func (this *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffset = reader.ReadInt32s(this.npairs * 2)
}

func (this *LOOKUP_SWITCH) Executor(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < this.npairs; i += 2 {
		if this.matchOffset[i] == key {
			offset := this.matchOffset[i+1]
			base.Branck(frame, int(offset))
			return
		}
	}
	base.Branck(frame, int(this.defaultOffset))
}
