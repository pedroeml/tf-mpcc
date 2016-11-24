package smoker

import (
	"fmt"
	"../ingredient"
)

type Smoker struct {
	name string
	kind int
	ingredients [2]*ingredient.Ingredient
}

func (s *Smoker) Init(name string, kind int) *Smoker {
	s.name = name
	s.kind = kind
	var ingredients [2]*ingredient.Ingredient
	ingredients[0] = nil
	ingredients[1] = nil
	s.ingredients = ingredients
	return s
}

func New(name string, kind int) *Smoker {
	return new(Smoker).Init(name, kind)
}

func (s *Smoker) Name() string {
	return s.name
}

/**
 * Number which indefies this object by his unlimited ingredient:
 * 	0 for Tabacco
 * 	1 for Paper
 * 	2 for Matches
 */
func (s *Smoker) Kind() int {
	return s.kind
}

func (s *Smoker) NeedsTabacco() bool {
	return s.needsIngredient(0)
}

func (s *Smoker) NeedsPaper() bool {
	return s.needsIngredient(1)
}

func (s *Smoker) NeedsMatches() bool {
	return s.needsIngredient(2)
}

func (s *Smoker) needsIngredient(kind int) bool {
	return s.kind != kind
}

// Sets the first ingredient this object needs.
func (s *Smoker) SetIngredientA(ing *ingredient.Ingredient) {
	s.setIngredient(ing, 0)
}

// Sets the second ingredient this object needs.
func (s *Smoker) SetIngredientB(ing *ingredient.Ingredient) {
	s.setIngredient(ing, 1)
}

func (s *Smoker) setIngredient(ing *ingredient.Ingredient, index int) {
	s.ingredients[index] = ing
	fmt.Printf("%s NOW HAS INGREDIENT %s\n", s.name, ing.Name())
}

// Performs the removal of two different ingredients this object has.
func (s *Smoker) Smoke() bool {
	if s.ingredients[0] != nil && s.ingredients[1] != nil {
		s.ingredients[0] = nil
		s.ingredients[1] = nil
		fmt.Printf("%s SMOKED\n", s.name)
		return true
	}

	return false
}
