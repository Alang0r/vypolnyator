package models

import "fmt"

type HabitList struct {
	Habits []Habit
}

func NewHabitList() HabitList {
	l := HabitList{}
	l.Habits = make([]Habit, 0)
	return l
}

func (l *HabitList) Add(h Habit) {
	l.Habits = append(l.Habits, h)
}

func (l *HabitList) Show() {
	fmt.Printf("%+v", l)
}

func (l *HabitList) ShowRemainigTime() {
	for _, h := range l.Habits {
		repeatCicle := h.TotalDuration / h.RemainingRepeats
		fmt.Printf("%s:\n%s\nRemaining repeats: %d. To complete challenge you need repeat habit at least 1 time every %d day\n", h.Name, h.Description, h.RemainingRepeats, repeatCicle)
	}
}

func (l *HabitList) MarkHabitRepeat(id uint64, count uint64) {
	for i, h := range l.Habits {
		if h.Id == id {
			l.Habits[i].RemainingRepeats -= count
		}
	}
}
