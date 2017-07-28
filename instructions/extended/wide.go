package extended

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/instructions/loads"
	"github.com/WYL53/jvmgo/instructions/math"
	"github.com/WYL53/jvmgo/rtda"
)

//wide 指令 改变 其他 指令 的 行为， modifiedInstruction 字段 存放 被 改变 的 指令。

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (this *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x16:
	case 0x17:
	case 0x18:
	case 0x19:
	case 0x36:
	case 0x37:
	case 0x38:
	case 0x39:
	case 0x3a:
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0xa9:
		panic("Unsupported opcode:oxa9!")
	}
}

func (this *WIDE) Executor(frame *rtda.Frame) {
	this.modifiedInstruction.Executor(frame)
}
