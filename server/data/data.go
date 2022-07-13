package data

type POI struct {
	ID          UUID
	Name        string
	Description string
	Thumbnail   UUID
}

type Route []RoutePoint

type RoutePoint struct {
	Fixed    bool
	Location struct {
		X float64
		Y float64
	}
	Show struct {
		POI         UUID
		Instruction InstructionDirection
	}
	Play UUID
}

type InstructionDirection string

const (
	UTurn           InstructionDirection = "u_turn"
	LeftUTurn       InstructionDirection = "left_u_turn"
	KeepLeft        InstructionDirection = "keep_left"
	LeaveRoundabout InstructionDirection = "leave_roundabout"
	SharpLeft       InstructionDirection = "sharp_left"
	Left            InstructionDirection = "left"
	SlightLeft      InstructionDirection = "slight_left"
	Continue        InstructionDirection = "continue"
	SlightRight     InstructionDirection = "slight_right"
	Right           InstructionDirection = "right"
	SharpRight      InstructionDirection = "sharp_right"
	EnterRoundabout InstructionDirection = "enter_roundabout"
	KeepRight       InstructionDirection = "keep_right"
	RightUTurn      InstructionDirection = "right_u_turn"
)
