package models

type Habit struct {
	TotalDuration    uint64
	RepeatDuration   uint64
	RamainingRepeats uint64
	TotalRepeats     uint64
}

func NewHabitat() Habit {
	h := Habit{}

	return h
}

func (obj Habit) calculateRemaining() {

}

func (obj Habit) markDone() {

}
