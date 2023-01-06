package main

import (
	"github.com/Alang0r/vypolnyator/models"
)

func main() {

	l := models.NewHabitList("My first list", "Description of my first list")

	h := models.NewHabit(1, 1,"TestHabit", "Descr for test", 365, 40)
	h2 := models.NewHabit(2,2, "TesSecondtHabit", "ololo", 30, 4)
	l.Add(h)
	l.Add(h2)

	for i := 0; i < 10; i++ {
		l.ShowRemainigTime()
		l.MarkHabitRepeat(1, 2)
		l.Tick()

	}

}
