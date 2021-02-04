package pizza

type Builder interface {
	setPrice()
	GetPizza() Pizza
}

func NewBuilder(builder_type string) (bool, Builder) {
	if builder_type == Margerita {
		return true, NewMargeritaBuilder()
	} else if builder_type == Pepperoni {
		return true, NewPepperoniBuilder()
	} else if builder_type == Hawaiian {
		return true, NewHawaiianBuilder()
	} else {
		return false, nil
	}
}

type MargeritaBuilder struct {
	dollars uint
	cents   uint
}

func NewMargeritaBuilder() *MargeritaBuilder {
	return &MargeritaBuilder{}
}

func (b *MargeritaBuilder) setPrice() {
	b.dollars = 12
	b.cents = 75
}

func (b *MargeritaBuilder) GetPizza() Pizza {
	p := Pizza{Basil: true, Mozzerella: true, OliveOil: true, Dollars: b.dollars, Cents: b.cents}
	return p
}

type PepperoniBuilder struct {
	dollars uint
	cents   uint
}

func NewPepperoniBuilder() *PepperoniBuilder {
	return &PepperoniBuilder{}
}

func (b *PepperoniBuilder) setPrice() {
	b.dollars = 10
	b.cents = 10
}

func (b *PepperoniBuilder) GetPizza() Pizza {
	p := Pizza{Mozzerella: true, Pepperoni: true, Dollars: b.dollars, Cents: b.cents}
	return p
}

type HawaiianBuilder struct {
	dollars uint
	cents   uint
}

func NewHawaiianBuilder() *HawaiianBuilder {
	return &HawaiianBuilder{}
}

func (b *HawaiianBuilder) setPrice() {
	b.dollars = 15
	b.cents = 0
}

func (b *HawaiianBuilder) GetPizza() Pizza {
	p := Pizza{Pineapple: true, Ham: true, Mozzerella: true, Dollars: b.dollars, Cents: b.cents}
	return p
}
