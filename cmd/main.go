package main

import (
	"github.com/Alang0r/vypolnyator/models"
)

func main() {
	h := models.NewHabit(1, "TestHabit", "Descr for test", 365, 40)
	l := models.NewHabitList()
	l.Add(h)

	for i := 0; i < 10; i++ {
		l.ShowRemainigTime()
		l.MarkHabitRepeat(1, 2)

	}

}
