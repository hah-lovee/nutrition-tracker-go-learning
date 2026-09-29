package meal

import (
	"fmt"
	"miniProject/food"
	foodstore "miniProject/foodstore"
	"time"
)

type timestamp struct {
	timeCreate time.Time
	dayCreate  time.Time
}

type Meal struct {
	Foods     []*food.Food
	timestamp // embedding
}

func NewMeal() Meal {
	t := time.Now()
	return Meal{
		Foods: []*food.Food{},
		timestamp: timestamp{
			timeCreate: t,
			dayCreate:  StartOfDay(t),
		},
	}
}

func (m *Meal) GetDay() time.Time {
	return m.dayCreate
}

func (m *Meal) Age() time.Duration {
	return time.Since(m.timeCreate)
}

func (m *Meal) AddFood(f *food.Food) {
	m.Foods = append(m.Foods, f)
}

func (m *Meal) TotalKcal() float64 {
	var total float64
	for _, food := range m.Foods {
		total += food.GetKcal()
	}

	return total
}

func (m *Meal) AddFoodByName(store foodstore.Store, name string) error {
	if food, ok := store.FindFood(name); ok {
		m.AddFood(food)
		return nil
	}
	return fmt.Errorf("add to meal: %w", foodstore.ErrFoodNotFound)
}

func StartOfDay(t time.Time) time.Time {
	t = t.UTC()
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
