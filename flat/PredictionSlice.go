// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PredictionSlice struct {
	_tab flatbuffers.Table
}

func GetRootAsPredictionSlice(buf []byte, offset flatbuffers.UOffsetT) *PredictionSlice {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PredictionSlice{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PredictionSlice) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PredictionSlice) Table() flatbuffers.Table {
	return rcv._tab
}

/// The moment in game time that this prediction corresponds to.
/// This corresponds to 'secondsElapsed' in the GameInfo table.
func (rcv *PredictionSlice) GameSeconds() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// The moment in game time that this prediction corresponds to.
/// This corresponds to 'secondsElapsed' in the GameInfo table.
func (rcv *PredictionSlice) MutateGameSeconds(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

/// The predicted location and motion of the object.
func (rcv *PredictionSlice) Physics(obj *Physics) *Physics {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Physics)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

/// The predicted location and motion of the object.
func PredictionSliceStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PredictionSliceAddGameSeconds(builder *flatbuffers.Builder, gameSeconds float32) {
	builder.PrependFloat32Slot(0, gameSeconds, 0.0)
}
func PredictionSliceAddPhysics(builder *flatbuffers.Builder, physics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(physics), 0)
}
func PredictionSliceEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}