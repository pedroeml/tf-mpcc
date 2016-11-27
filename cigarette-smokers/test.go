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
	smokers[2] = smoker.New("smokerMatches", 2)
	agent := agent.New()

	ch := make(chan *ingredient.Ingredient, 2)

	var wg sync.WaitGroup

	for time := 0; time < 5 ; time++ {
		fmt.Println("========================================")

		wg.Add(1)
		go agent.AddIngredients(ch, &wg)	// Agent dispose two ingredients and store these to the channel and to his array of ingredients.
		wg.Wait()	// Waits Agent finishes the ingredients disposal.

		fmt.Println("========================================")

		for i := 0; i < 3; i++ {
			wg.Add(1)
			go negociations(agent, smokers[i], ch, &wg)	// Starts negociations among Agent and Smokers.
		}

		wg.Wait()	// Wait untill all negociations are done, which means at least one smoker bought the ingredients.
	}



}

func negociations(agent *agent.Agent, smoker *smoker.Smoker, ch chan *ingredient.Ingredient, wg *sync.WaitGroup) {
	hasTabaccoAndPaper := agent.ThereIsTabacco() && agent.ThereIsPaper()
	hasTabaccoAndMatches := agent.ThereIsTabacco() && agent.ThereIsMatches()
	hasPaperAndMatches := agent.ThereIsPaper() && agent.ThereIsMatches()

	needsTabaccoAndPaper := smoker.NeedsTabacco() && smoker.NeedsPaper()
	needsTabaccoAndMatches := smoker.NeedsTabacco() && smoker.NeedsMatches()
	needsPaperAndMatches := smoker.NeedsPaper() && smoker.NeedsMatches()

	sold := false	// Flag to indicate that two ingredients were sold.

	if hasTabaccoAndPaper && needsTabaccoAndPaper {		// If this smoker needs tabacco and paper and the agent has these ingredients:
		agent.SellTabacco()	// Agent sell these ingredients, which means he doesn't have these ingredients anymore.
		agent.SellPaper()	// So, only the channel has these ingredients.
		sold = true

	} else if hasTabaccoAndMatches && needsTabaccoAndMatches {	// If this smoker needs tabacco and matches and the agent has these ingredients:
		agent.SellTabacco()	// Agent sell these ingredients, which means he doesn't have these ingredients anymore.
		agent.SellMatches()	// So, only the channel has these ingredients.
		sold = true
	} else if hasPaperAndMatches && needsPaperAndMatches {		// If this smoker needs matches and paper and the agent has these ingredients:
		agent.SellPaper()	// Agent sell these ingredients, which means he doesn't have these ingredients anymore.
		agent.SellMatches()	// So, only the channel has these ingredients.
		sold = true
	}

	if sold {	// If the ingredients were sold to this smoker:
		ingredientA := <- ch	// Remove the two sold ingredients from the channel.
		ingredientB := <- ch

		smoker.SetIngredientA(ingredientA)	// The smoker who needs these ingredients now has them.
		smoker.SetIngredientB(ingredientB)
		smoker.Smoke()		// Smoke a cigarette by removing the bought ingredients that he needed.
	}

	wg.Done()	// This smoker is done: he may has smoked, because smoking requires the agent to has the ingredients that he needs.
}
