package classfile

//Deprecated和Synthetic是最简单的两种属性，仅起标记作用，不包含任何数据。

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

func (this *MarkerAttribute) readInfo(reader *ClassReader) {}
