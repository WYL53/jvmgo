package classfile

import "fmt"

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classDate []byte) (cf *ClassFile, err error) {
	defer func(){
		if r:= recover();r != nil{
			var ok bool
			err,ok = r.(error)
			if !ok{
				err = fmt.Errorf("%v",r)
			}
		}
	}()
	cr := &ClassRead{classDate}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (this *ClassFile) read(reader *ClassRead) {
	this.readAndCheckMagic(reader)
	this.readAndCheckVersion(reader)
	this.constantPool = readConstantPool(reader)
	this.accessFlags = reader.readUint16()
	this.thisClass = reader.readUint16()
	this.superClass = reader.readUint16()
	this.interfaces = reader.readUint16s()
	this.fields = readMembers(reader,this.constantPool)
	this.methods = readMebers(reader,this.constantPool)
	this.attributes = readAttributes(reader,this.constantPool)
}

func (this *ClassFile) readAndCheckVersion(reader *ClassRead) {

}

func (this *ClassFile) readAndCheckMagic(reader *ClassRead) {

}

func (this *ClassFile) MinorVersion() uint16 {
	return this.minorVersion
}

func (this *ClassFile) MajorVersion() uint16 {
	return this.majorVersion
}

func (this *ClassFile) ConstantPool() ConstantPool {

}

func (this *ClassFile)AccessFlags() uint16{
	
}

func (this *ClassFile)Fields() []*MemberInfo(){
	
}

func (this *ClassFile)Methods() []*MemberInfo(){
	
}

func (this *ClassFile)ClassName() string{
	
}

func (this *ClassFile)SuperClassName()string{
	
}

func (this *ClassFile)InterfaceNames()[]string{
	
}