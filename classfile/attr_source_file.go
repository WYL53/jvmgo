package classfile

//SourceFile 是 可 选定 长 属性， 只会 出现 在 ClassFile 结构 中， 用于 指出 源 文件名。
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16 //值 必须 是 2。是 常量 池 索引，指向CONSTANT_ Utf8_ info常量。
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readUint16()
}

func (this *SourceFileAttribute) FileName() string {
	return this.cp.getUtf8(this.sourceFileIndex)
}
