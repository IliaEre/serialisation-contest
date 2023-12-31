// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package sample

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FindRequest struct {
	_tab flatbuffers.Table
}

func GetRootAsFindRequest(buf []byte, offset flatbuffers.UOffsetT) *FindRequest {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FindRequest{}
	x.Init(buf, n+offset)
	return x
}

func FinishFindRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsFindRequest(buf []byte, offset flatbuffers.UOffsetT) *FindRequest {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FindRequest{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedFindRequestBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *FindRequest) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FindRequest) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FindRequest) Limit() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FindRequest) MutateLimit(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *FindRequest) Offset() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FindRequest) MutateOffset(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func FindRequestStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FindRequestAddLimit(builder *flatbuffers.Builder, limit int32) {
	builder.PrependInt32Slot(0, limit, 0)
}
func FindRequestAddOffset(builder *flatbuffers.Builder, offset int32) {
	builder.PrependInt32Slot(1, offset, 0)
}
func FindRequestEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
