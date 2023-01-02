package models

type Habit struct {
	Id               uint64
	Name             string
	Description      string
	TotalDuration    uint64
	RepeatDuration   uint64
	RemainingRepeats uint64
	TotalRepeats     uint64
}

func NewHabit(id uint64, name string, descr string, totalDuration uint64, totalRepeats uint64) Habit {
	h := Habit{
		Id:               id,
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
