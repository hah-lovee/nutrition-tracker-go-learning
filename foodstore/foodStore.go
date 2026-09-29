package foodstore

import (
	"errors"
	"fmt"
	"miniProject/food"
)

var (
	ErrFoodNotFound = errors.New("food not found")
	ErrFoodExists   = errors.New("food already exists")
)

type Store interface {
	FindFood(name string) (*food.Food, bool)
	AddFood(food *food.Food) error
}

type FoodStore struct {
	store map[string]*food.Food
}

func CreateEmptyStore() FoodStore {
	return FoodStore{
		make(map[string]*food.Food),
	}
}

func (store *FoodStore) AddFood(food *food.Food) error {
	if _, ok := store.store[food.GetName()]; !ok {
		store.store[food.GetName()] = food
		return nil
	} else {
		return fmt.Errorf("add food %q: %w", food.GetName(), ErrFoodExists)
	}
}

func (store *FoodStore) ChangeKcalByName(name string, newKcal float64) {
	if food, ok := store.store[name]; ok {
		food.SetKcal(newKcal)
	}
}

func (store *FoodStore) FindFood(name string) (*food.Food, bool) {
	if f, ok := store.store[name]; ok {
		return f, ok
	}
	return &food.Food{}, false
}
