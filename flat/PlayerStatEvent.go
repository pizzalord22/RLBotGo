// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Notification that a player triggers some in-game event, such as:
///		Win, Loss, TimePlayed;
///		Shot, Assist, Center, Clear, PoolShot;
///		Goal, AerialGoal, BicycleGoal, BulletGoal, BackwardsGoal, LongGoal, OvertimeGoal, TurtleGoal;
///		AerialHit, BicycleHit, BulletHit, /*BackwardsHit,*/ JuggleHit, FirstTouch, BallHit;
///		Save, EpicSave, FreezeSave;
///		HatTrick, Savior, Playmaker, MVP;
///		FastestGoal, SlowestGoal, FurthestGoal, OwnGoal;
///		MostBallTouches, FewestBallTouches, MostBoostPickups, FewestBoostPickups, BoostPickups;
///		CarTouches, Demolition, Demolish;
///		LowFive, HighFive;
type PlayerStatEvent struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerStatEvent(buf []byte, offset flatbuffers.UOffsetT) *PlayerStatEvent {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerStatEvent{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PlayerStatEvent) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerStatEvent) Table() flatbuffers.Table {
	return rcv._tab
}

/// index of the player associated with the event
func (rcv *PlayerStatEvent) PlayerIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// index of the player associated with the event
func (rcv *PlayerStatEvent) MutatePlayerIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

/// Event type
func (rcv *PlayerStatEvent) StatType() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Event type
func PlayerStatEventStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PlayerStatEventAddPlayerIndex(builder *flatbuffers.Builder, playerIndex int32) {
	builder.PrependInt32Slot(0, playerIndex, 0)
}
func PlayerStatEventAddStatType(builder *flatbuffers.Builder, statType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(statType), 0)
}
func PlayerStatEventEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}