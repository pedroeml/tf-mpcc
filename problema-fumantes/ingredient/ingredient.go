package ingredient

type Ingredient struct {
	name string
	quantity int
}

func (i *Ingredient) Init(name string) *Ingredient {
	i.name = name
	i.quantity = 0
	return i
}

func New(name string) *Ingredient {
	return new(Ingredient).Init(name)
}

func (i *Ingredient) Name() string {
	return i.name
}

func (i *Ingredient) Quantity() int {
	return i.quantity
}

func (i *Ingredient) IncrementQuantity() {
	i.quantity += 1
}

func (i *Ingredient) TakeIngredient() *Ingredient {
	if i.quantity == 0 {
		return nil
	}

	i.quantity -= 1
	takenIngredient := &Ingredient{name:i.name, quantity:1}
	return takenIngredient
}
