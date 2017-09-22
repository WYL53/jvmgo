package extended

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/rtda"
)
//goto_ w 指令 和 goto 指令 的 唯一 区别 就是 索引 从 2 字节 变成 了 4 字节。

type GOTO_W struct {
	offset int
}

func (this *GOTO_W)FetchOperands(reader *base.BytecodeReader)  {
	this.offset = int(reader.ReadInt32())
}

func (this *GOTO_W)Execute(frame *rtda.Frame)  {
	base.Branck(frame,this.offset)
}

