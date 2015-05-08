package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type FieldInfo struct {
	MemberInfo
}

/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MethodInfo struct {
	MemberInfo
}

type MemberInfo struct {
	cp              *ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	AttributeTable
}

func (self *MemberInfo) read(reader *ClassReader) {
	self.accessFlags = reader.readUint16()
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
	self.attributes = readAttributes(reader, self.cp)
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
func (self *MemberInfo) Signature() string {
	signatureAttr := self.SignatureAttribute()
	if signatureAttr != nil {
		return signatureAttr.Signature()
	}
	return ""
}