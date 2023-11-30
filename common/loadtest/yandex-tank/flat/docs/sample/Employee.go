// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Employee struct {
	_tab flatbuffers.Table
}

func GetRootAsEmployee(buf []byte, offset flatbuffers.UOffsetT) *Employee {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Employee{}
	x.Init(buf, n+offset)
	return x
}

func FinishEmployeeBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsEmployee(buf []byte, offset flatbuffers.UOffsetT) *Employee {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Employee{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedEmployeeBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Employee) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Employee) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Employee) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Employee) Surname() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Employee) Code() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func EmployeeStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func EmployeeAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func EmployeeAddSurname(builder *flatbuffers.Builder, surname flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(surname), 0)
}
func EmployeeAddCode(builder *flatbuffers.Builder, code flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(code), 0)
}
func EmployeeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}