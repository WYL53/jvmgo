package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassRead, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassRead, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(read, cp),
	}
}

func (this *MemberInfo) AccessFlags() uint16 {
	return this.accessFlags
}

func (this *MemberInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}

func (this *MemberInfo) Descriptor() string {
	return this.cp.getUtf8(this.descriptorIndex)
}
