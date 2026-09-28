package diary

import (
	"fmt"
	"miniProject/meal"
	"time"
)

type Diary struct {
	days map[time.Time][]*meal.Meal
}

func NewDiary() Diary {
	return Diary{
		days: map[time.Time][]*meal.Meal{},
	}
}

func (d Diary) AddMeal(m *meal.Meal) {
	currentDay := m.GetDay()
	d.days[currentDay] = append(d.days[currentDay], m)
}

func (d Diary) DiaryTotal(day time.Time) (float64, error) {
	if meals, ok := d.days[day]; ok {
		total := 0.0
		for _, m := range meals {
			total += m.TotalKcal()
		}

		return total, nil
	} else {
		return 0, fmt.Errorf("haven't info about %s", day)
	}
}
