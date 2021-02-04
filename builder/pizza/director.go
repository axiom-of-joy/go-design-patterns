package pizza

type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) GetPizza() Pizza {
	d.builder.setPrice()
	p := d.builder.GetPizza()
	return p
}
