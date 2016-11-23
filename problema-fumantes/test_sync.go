package main

import (
	"./agent"
	"./ingredient"
	"./smoker"
	"fmt"
)

func main() {
	smokerTabacco := smoker.New("smokerTabacco", 0)
	smokerPaper := smoker.New("smokerPaper", 1)
	smokerMatches := smoker.New("smokerWatches", 2)
	agent := agent.New()
	agent.AddIngredients()
	var smokers [3]*smoker.Smoker
	smokers[0] = smokerTabacco
	smokers[1] = smokerPaper
	smokers[2] = smokerMatches
	for i := 0; i < 10; i++ {
		fmt.Println("==============================")
		negociations(agent, smokers)
	}
}

func negociations(agent *agent.Agent, smokers [3]*smoker.Smoker) {
	hasTabaccoAndPaper := agent.ThereIsTabacco() && agent.ThereIsPaper()
	hasTabaccoAndMatches := agent.ThereIsTabacco() && agent.ThereIsMatches()
	hasPaperAndMatches := agent.ThereIsPaper() && agent.ThereIsMatches()

	for i := 0; i < 3; i++ {
		needsTabaccoAndPaper := smokers[i].NeedsTabacco() && smokers[i].NeedsPaper()
		needsTabaccoAndMatches := smokers[i].NeedsTabacco() && smokers[i].NeedsMatches()
		needsPaperAndMatches := smokers[i].NeedsPaper() && smokers[i].NeedsMatches()

		var ingredientA, ingredientB *ingredient.Ingredient
		sold := false

		if hasTabaccoAndPaper && needsTabaccoAndPaper {
			ingredientA = agent.SellTabacco()
			ingredientB = agent.SellPaper()
			sold = true

		} else if hasTabaccoAndMatches && needsTabaccoAndMatches {
			ingredientA = agent.SellTabacco()
			ingredientB = agent.SellMatches()
			sold = true
		} else if hasPaperAndMatches && needsPaperAndMatches {
			ingredientA= agent.SellPaper()
			ingredientB= agent.SellMatches()
			sold = true
		}

		if sold {
			smokers[i].SetIngredientA(ingredientA)
			smokers[i].SetIngredientB(ingredientB)
			smokers[i].Smoke()
			agent.AddIngredients()
		}

	}
}