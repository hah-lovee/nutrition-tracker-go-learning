package main

import (
	"bufio"
	"errors"
	"fmt"
	"miniProject/food"
	foodstore "miniProject/foodstore"
	"miniProject/meal"
	"os"

	"github.com/k0kubun/pp"
)

func main() {

	foodsDB := foodstore.CreateEmptyStore()
	meat, err := food.CreateFood("meat", 150)
	var ve *food.ValidationError
	if errors.As(err, &ve) {
		fmt.Println("плохое поле:", ve.Field, ve.Value)
		return
	}
	err = foodsDB.AddFood(meat)
	if errors.Is(err, foodstore.ErrFoodExists) {
		fmt.Println("такой продукт уже существует")
	}
	makaroni, _ := food.CreateFood("makaroni", 50.0)
	foodsDB.AddFood(makaroni)

	err = foodsDB.AddFood(meat)
	if errors.Is(err, foodstore.ErrFoodExists) {
		fmt.Println("такой продукт уже существует")
	}

	foodsDB.ChangeKcalByName("meat", 100)

	currentMeal := meal.NewMeal()
	fmt.Println(currentMeal.GetDay())
	err = currentMeal.AddFoodByName(&foodsDB, "sfddf")
	if errors.Is(err, foodstore.ErrFoodNotFound) {
		fmt.Println("такой продукт нет в базе")
	}
	err = currentMeal.AddFoodByName(&foodsDB, "meat")
	if errors.Is(err, foodstore.ErrFoodNotFound) {
		fmt.Println("такой продукт нет в базе")
	} else {
		fmt.Println("Успешно добавил")
	}
	pp.Println(currentMeal.Foods)
	pp.Println(foodsDB)
}

func StartLoop(currentMeal meal.Meal, foodsDB foodstore.FoodStore) {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("Введите продукт: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("error input!")
			break
		}

		text := scanner.Text()
		if text == "New meal" {
			fmt.Println("new meal create!")
			currentMeal = meal.NewMeal()
			fmt.Println(currentMeal.GetDay())
			continue

		} else if text == "break" {
			break
		}

		if err := currentMeal.AddFoodByName(&foodsDB, text); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("New total Kcal:", currentMeal.TotalKcal())
		}
	}
}
