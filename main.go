package main

import (
	"fmt"

	"github.com/WYL53/jvmgo/classpath"
	"github.com/WYL53/jvmgo/classfile"
	"github.com/WYL53/jvmgo/rtda"
	"github.com/WYL53/jvmgo/instructions/base"
	"github.com/WYL53/jvmgo/instructions"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	className := strings.Replace(cmd.class,".","/",-1)
	cf := loadClass(className,cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil{
		interpret(mainMethod)
	}else {
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}
}

func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile {
	classData,_,err := cp.ReadClass(className)
	if err != nil{
		panic(err)
	}
	cf ,err := classfile.Parse(classData)
	if err != nil{
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo{
	for _,m := range cf.Methods(){
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V"{
			return m
		}
	}
	return nil
}

func interpret(methodInfo *classfile.MemberInfo)  {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	byteCode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread,byteCode)
}

func catchErr(frame *rtda.Frame)  {
	if r := recover();r != nil{
		fmt.Printf("LocalVars:%v\n",frame.LocalVars())
		fmt.Printf("OperandStack:%v\n",frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread,byteCode []byte)  {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for  {
		pc := frame.NextPC()
		thread.SetPC(pc)
		//解码
		reader.Reset(byteCode,pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		//执行
		fmt.Printf("pc:%2d inst:%T %v\n",pc,inst,inst)
		inst.Execute(frame)
	}
}