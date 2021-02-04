package pizza

type Pizza struct {
	Basil      bool
	Ham        bool
	Mozzerella bool
	Pepperoni  bool
	Pineapple  bool
	OliveOil   bool
	Dollars    uint
	Cents      uint
}

func NewPizza() *Pizza {
	return &Pizza{}
}
