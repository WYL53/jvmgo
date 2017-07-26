package base

import "github.com/WYL53/jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Executor(frame *rtda.Frame)
}

type NoOperandsInstruction struct {}

func (this *NoOperandsInstruction)FetchOperands(reader *BytecodeReader)  {

}

//跳转指令
type BranchInstruction struct {
	Offset int
}

func (this *BranchInstruction)FetchOperands(reader *BytecodeReader)  {
	this.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (this *Index8Instruction)FetchOperands(reader *BytecodeReader)  {
	this.Index = int(reader.ReadInt8())
}

type Index16Instruction struct {
	Index uint
}

func (this *Index16Instruction)FetchOperands(reader *BytecodeReader)  {
	this.Index = int(reader.ReadInt16())
}