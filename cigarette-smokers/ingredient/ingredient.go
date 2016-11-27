package ingredient

type Ingredient struct {
	name string
}

func (i *Ingredient) Init(name string) *Ingredient {
	i.name = name
	return i
}

func New(name string) *Ingredient {
	return new(Ingredient).Init(name)
}

func (i *Ingredient) Name() string {
	return i.name
}
