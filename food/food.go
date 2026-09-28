package food

import (
	"fmt"
)

type ValidationError struct {
	Field string
	Value float64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %v", e.Field, e.Value)
}

type Food struct {
	name string
	kcal float64
}

func CreateFood(name string, kcal float64) (*Food, error) {
	if kcal < 0 {
		e := &ValidationError{Field: "kcal", Value: kcal}
		return &Food{}, fmt.Errorf("create food %w", e)
	}
	return &Food{
		name: name,
		kcal: kcal,
	}, nil
}

func (food *Food) SetName(name string) {
	food.name = name
}

func (food *Food) SetKcal(kcal float64) {
	food.kcal = kcal
}

func (food Food) GetName() string {
	return food.name
}

func (food Food) GetKcal() float64 {
	return food.kcal
}
