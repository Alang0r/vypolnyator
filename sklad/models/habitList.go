package models

import "fmt"

type HabitList struct {
	Id           uint64 // id of the record
	Name         string // name of habit
	Description  string // dscr of habit
	Habits       []Habit
	NotifyPeriod uint64 // Notify repit period
}

func NewHabitList(name string, dscr string) HabitList {
	l := HabitList{
		Name:        name,
		Description: dscr,
	}
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

func (l *HabitList) Tick() {
	for i, v := range l.Habits {
		if !v.Completed {
			l.Habits[i].RemainingDuration--
			if v.RemainingDuration == 0 {
				l.Habits[i].Completed = true
			}
		}

	}
}
