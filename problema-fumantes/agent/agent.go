package agent

import (
	"../ingredient"
	"time"
	"math/rand"
	"fmt"
)

type Agent struct {
	ingredients [3]*ingredient.Ingredient
}

func (a *Agent) Init() *Agent {
	var ingredients [3]*ingredient.Ingredient
	ingredients[0] = ingredient.New("Tabacco")
	ingredients[1] = ingredient.New("Paper")
	ingredients[2] = ingredient.New("Matches")
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
	return a.ingredients[index].Quantity() != 0
}

func (a *Agent) SellTabacco() *ingredient.Ingredient {
	return a.sellIngredient(0)
}

func (a *Agent) SellPaper() *ingredient.Ingredient {
	return a.sellIngredient(1)
}

func (a *Agent) SellMatches() *ingredient.Ingredient {
	return a.sellIngredient(2)
}

func (a *Agent) sellIngredient(index int) *ingredient.Ingredient {
	if a.thereIsIngredient(index) {
		fmt.Printf("Agent SOLD %s\n", a.ingredients[index].Name())
		return a.ingredients[index].TakeIngredient()
	}
	return nil
}

// Add more two ingedients to sell if there is no ingredient left to sell.
func (a *Agent) AddIngredients() bool {
	if !a.ThereIsTabacco() && !a.ThereIsPaper() && !a.ThereIsMatches() {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		indexA := r1.Intn(3)
		indexB := r1.Intn(3)

		if indexA == indexB {	// indexes must be different because two different ingredients need to be increased their quantity
			return a.AddIngredients()
		}

		a.ingredients[indexA].IncrementQuantity()
		fmt.Printf("Agent NOW HAS %s\n", a.ingredients[indexA].Name())
		a.ingredients[indexB].IncrementQuantity()
		fmt.Printf("Agent NOW HAS %s\n", a.ingredients[indexB].Name())
		return true
	}

	return false
}
