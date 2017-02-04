package classfile

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (this *ExceptionAttribute) readInfo(reader *ClassReader) {

}

func (this *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return this.exceptionIndexTable
}
