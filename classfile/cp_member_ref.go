package classfile
//
//CONSTANT_ Fieldref_ info 表示 字段 符号 引用，
//CONSTANT_ Methodref_ info 表示 普通（ 非 接口）方法 符号 引用，
//CONSTANT_ InterfaceMethodref_ info 表示 接口 方法 符号 引用。
// 这 三种 常量 结构 一模一样

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

func (this *ConstantMemberrefInfo) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}

func (this *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return this.cp.getNameAndType(this.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
