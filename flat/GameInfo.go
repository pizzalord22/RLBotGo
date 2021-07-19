// automatically generated by the FlatBuffers compiler, do not modify

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GameInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsGameInfo(buf []byte, offset flatbuffers.UOffsetT) *GameInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GameInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GameInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GameInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GameInfo) SecondsElapsed() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GameInfo) MutateSecondsElapsed(n float32) bool {
	return rcv._tab.MutateFloat32Slot(4, n)
}

func (rcv *GameInfo) GameTimeRemaining() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GameInfo) MutateGameTimeRemaining(n float32) bool {
	return rcv._tab.MutateFloat32Slot(6, n)
}

func (rcv *GameInfo) IsOvertime() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GameInfo) MutateIsOvertime(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *GameInfo) IsUnlimitedTime() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GameInfo) MutateIsUnlimitedTime(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

/// True when cars are allowed to move, and during the pause menu. False during replays.
func (rcv *GameInfo) IsRoundActive() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

/// True when cars are allowed to move, and during the pause menu. False during replays.
func (rcv *GameInfo) MutateIsRoundActive(n byte) bool {
	return rcv._tab.MutateByteSlot(12, n)
}

/// True when the clock is paused due to kickoff, but false during kickoff countdown. In other words, it is true
/// while cars can move during kickoff. Note that if both players sit still, game clock start and this will become false.
func (rcv *GameInfo) IsKickoffPause() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

/// True when the clock is paused due to kickoff, but false during kickoff countdown. In other words, it is true
/// while cars can move during kickoff. Note that if both players sit still, game clock start and this will become false.
func (rcv *GameInfo) MutateIsKickoffPause(n byte) bool {
	return rcv._tab.MutateByteSlot(14, n)
}

/// Turns true after final replay, the moment the 'winner' screen appears. Remains true during next match
/// countdown. Turns false again the moment the 'choose team' screen appears.
func (rcv *GameInfo) IsMatchEnded() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

/// Turns true after final replay, the moment the 'winner' screen appears. Remains true during next match
/// countdown. Turns false again the moment the 'choose team' screen appears.
func (rcv *GameInfo) MutateIsMatchEnded(n byte) bool {
	return rcv._tab.MutateByteSlot(16, n)
}

func (rcv *GameInfo) WorldGravityZ() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *GameInfo) MutateWorldGravityZ(n float32) bool {
	return rcv._tab.MutateFloat32Slot(18, n)
}

/// Game speed multiplier, 1.0 is regular game speed.
func (rcv *GameInfo) GameSpeed() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

/// Game speed multiplier, 1.0 is regular game speed.
func (rcv *GameInfo) MutateGameSpeed(n float32) bool {
	return rcv._tab.MutateFloat32Slot(20, n)
}

/// Tracks the number of physics frames the game has computed.
/// May increase by more than one across consecutive packets.
/// Data type will roll over after 207 days at 120Hz.
func (rcv *GameInfo) FrameNum() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

/// Tracks the number of physics frames the game has computed.
/// May increase by more than one across consecutive packets.
/// Data type will roll over after 207 days at 120Hz.
func (rcv *GameInfo) MutateFrameNum(n int32) bool {
	return rcv._tab.MutateInt32Slot(22, n)
}

func GameInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func GameInfoAddSecondsElapsed(builder *flatbuffers.Builder, secondsElapsed float32) {
	builder.PrependFloat32Slot(0, secondsElapsed, 0.0)
}
func GameInfoAddGameTimeRemaining(builder *flatbuffers.Builder, gameTimeRemaining float32) {
	builder.PrependFloat32Slot(1, gameTimeRemaining, 0.0)
}
func GameInfoAddIsOvertime(builder *flatbuffers.Builder, isOvertime byte) {
	builder.PrependByteSlot(2, isOvertime, 0)
}
func GameInfoAddIsUnlimitedTime(builder *flatbuffers.Builder, isUnlimitedTime byte) {
	builder.PrependByteSlot(3, isUnlimitedTime, 0)
}
func GameInfoAddIsRoundActive(builder *flatbuffers.Builder, isRoundActive byte) {
	builder.PrependByteSlot(4, isRoundActive, 0)
}
func GameInfoAddIsKickoffPause(builder *flatbuffers.Builder, isKickoffPause byte) {
	builder.PrependByteSlot(5, isKickoffPause, 0)
}
func GameInfoAddIsMatchEnded(builder *flatbuffers.Builder, isMatchEnded byte) {
	builder.PrependByteSlot(6, isMatchEnded, 0)
}
func GameInfoAddWorldGravityZ(builder *flatbuffers.Builder, worldGravityZ float32) {
	builder.PrependFloat32Slot(7, worldGravityZ, 0.0)
}
func GameInfoAddGameSpeed(builder *flatbuffers.Builder, gameSpeed float32) {
	builder.PrependFloat32Slot(8, gameSpeed, 0.0)
}
func GameInfoAddFrameNum(builder *flatbuffers.Builder, frameNum int32) {
	builder.PrependInt32Slot(9, frameNum, 0)
}
func GameInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
