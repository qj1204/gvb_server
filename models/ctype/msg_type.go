package ctype

type MsgType int

const (
	InRoomMsg  MsgType = 1
	TextMsg    MsgType = 2
	ImageMsg   MsgType = 3
	VoiceMsg   MsgType = 4
	VideoMsg   MsgType = 5
	SystemMsg  MsgType = 6
	OutRoomMsg MsgType = 7
)
