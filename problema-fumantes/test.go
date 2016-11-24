package main

import (
	"./agent"
	"./ingredient"
	"./smoker"
	"sync"
	"fmt"
)

func main() {
	var smokers [3]*smoker.Smoker
	smokers[0] = smoker.New("smokerTabacco", 0)
	smokers[1] = smoker.New("smokerPaper", 1)
	smokers[2] = smoker.New("smokerWatches", 2)
	agent := agent.New()

	ch := make(chan *ingredient.Ingredient, 2)

	var wg sync.WaitGroup

	for time := 0; time < 5 ; time++ {
		wg.Add(1)
		go agent.AddIngredients(ch, &wg)
		wg.Wait()

		fmt.Println("==============================")

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go negociations(agent, smokers[i], ch, &wg)
		}

		wg.Wait()
	}



}

func negociations(agent *agent.Agent, smoker *smoker.Smoker, ch chan *ingredient.Ingredient, wg *sync.WaitGroup) {
	hasTabaccoAndPaper := agent.ThereIsTabacco() && agent.ThereIsPaper()
	hasTabaccoAndMatches := agent.ThereIsTabacco() && agent.ThereIsMatches()
	hasPaperAndMatches := agent.ThereIsPaper() && agent.ThereIsMatches()

	needsTabaccoAndPaper := smoker.NeedsTabacco() && smoker.NeedsPaper()
	needsTabaccoAndMatches := smoker.NeedsTabacco() && smoker.NeedsMatches()
	needsPaperAndMatches := smoker.NeedsPaper() && smoker.NeedsMatches()

	sold := false

	if hasTabaccoAndPaper && needsTabaccoAndPaper {
		agent.SellTabacco()
		agent.SellPaper()
		sold = true

	} else if hasTabaccoAndMatches && needsTabaccoAndMatches {
		agent.SellTabacco()
		agent.SellMatches()
		sold = true
	} else if hasPaperAndMatches && needsPaperAndMatches {
		agent.SellPaper()
		agent.SellMatches()
		sold = true
	}

	if sold {
		ingredientA := <- ch
		ingredientB := <- ch

		smoker.SetIngredientA(ingredientA)
		smoker.SetIngredientB(ingredientB)
		smoker.Smoke()
	}

	wg.Done()
}
