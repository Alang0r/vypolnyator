package models

type Habit struct {
	Id                uint64 // id of the record
	ListId            uint64
	Name              string // name of habit
	Description       string // dscr of habit
	TotalDuration     uint64 // amount of repetition
	RemainingDuration uint64
	RepeatDuration    uint64 //
	RemainingRepeats  uint64 // amount of reamaining repeats to complete goal
	TotalRepeats      uint64 // amount of total repeats to complete goal
	Completed         bool
}

func NewHabit(id uint64, listId uint64, name string, descr string, totalDuration uint64, totalRepeats uint64) Habit {
	h := Habit{
		Id:               id,
		ListId: listId,
		Name:             name,
		Description:      descr,
		TotalDuration:    totalDuration,
		TotalRepeats:     totalRepeats,
		RemainingRepeats: totalRepeats,
	}

	return h
}

func (h *Habit) calculateRemaining() uint64 {
	return h.RemainingRepeats
}

func (h *Habit) markDone() {

}
