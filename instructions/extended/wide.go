package extended

import (
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/instructions/loads"
	"github.com/WYL53/jvmgo/instructions/math"
	"github.com/WYL53/jvmgo/rtda"
	"github.com/WYL53/jvmgo/instructions/stores"
)
//　 wide 指令 加载 类 指令、 存储 类 指令、 ret 指令 和 iinc 指令
// 需要 按 索引 访问 局部 变 量表， 索引 以 uint8 的 形式 存在 字节 码 中。

//wide 指令 改变 其他 指令 的 行为， modifiedInstruction 字段 存放 被 改变 的 指令。
//wide 指令 只是 增加 了 索引 宽度， 并不 改变 子 指令 操作，

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (this *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: //iload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x16:  //lload
		linst := &loads.LLOAD{}
		linst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = linst
	case 0x17: //fload
		finst := &loads.FLOAD{}
		finst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = finst
	case 0x18: //dload
		dinst := &loads.DLOAD{}
		dinst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = dinst
	case 0x19:  //aload
		ainst := &loads.ALOAD{}
		ainst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = ainst
	case 0x36: //istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0x37: //lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0x38: //fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0x39: //dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0x3a: //astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0x84: //iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0xa9:
		panic("Unsupported opcode:oxa9!")
	}
}

func (this *WIDE) Execute(frame *rtda.Frame) {
	this.modifiedInstruction.Execute(frame)
}
