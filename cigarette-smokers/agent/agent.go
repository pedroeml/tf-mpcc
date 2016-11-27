package agent

import (
	"../ingredient"
	"math/rand"
	"fmt"
	"sync"
)

type Agent struct {
	ingredients [3]*ingredient.Ingredient
}

func (a *Agent) Init() *Agent {
	var ingredients [3]*ingredient.Ingredient
	ingredients[0] = nil
	ingredients[1] = nil
	ingredients[2] = nil
	a.ingredients = ingredients
	return a
}

func New() *Agent {
	return new(Agent).Init()
}

func (a *Agent) ThereIsTabacco() bool {
	return a.thereIsIngredient(0)
}

func (a *Agent) ThereIsPaper() bool {
	return a.thereIsIngredient(1)
}

func (a *Agent) ThereIsMatches() bool {
	return a.thereIsIngredient(2)
}

func (a *Agent) thereIsIngredient(index int) bool {
	return a.ingredients[index] != nil
}

func (a *Agent) SellTabacco() {
	a.sellIngredient(0)
}

func (a *Agent) SellPaper() {
	a.sellIngredient(1)
}

func (a *Agent) SellMatches() {
	a.sellIngredient(2)
}

func (a *Agent) sellIngredient(index int) {
	if a.thereIsIngredient(index) {
		fmt.Printf("Agent SOLD %s\n", a.ingredients[index].Name())
		a.ingredients[index] = nil
	}
}

// Add more two ingedients to sell if there is no ingredient left to sell.
func (a *Agent) AddIngredients(ch chan *ingredient.Ingredient, wg *sync.WaitGroup) {
	if !a.ThereIsTabacco() && !a.ThereIsPaper() && !a.ThereIsMatches() {
		indexA := rand.Intn(3)
		indexB := rand.Intn(3)

		if indexA == indexB {	// indexes must be different because two different ingredients need to be increased their quantity
			a.AddIngredients(ch, wg)
		} else {
			a.ingredients[indexA] = createIngredient(indexA)
			fmt.Printf("Agent NOW HAS %s\n", a.ingredients[indexA].Name())
			ch <- a.ingredients[indexA]

			a.ingredients[indexB] = createIngredient(indexB)
			fmt.Printf("Agent NOW HAS %s\n", a.ingredients[indexB].Name())
			ch <- a.ingredients[indexB]

			wg.Done()
		}
	}
}

func createIngredient(index int) *ingredient.Ingredient {
	switch index {
	case 0:
		return ingredient.New("Tabacco")
	case 1:
		return ingredient.New("Paper")
	case 2:
		return ingredient.New("Matches")
	}
	return nil
}
