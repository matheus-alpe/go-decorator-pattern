package shop

type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// first decorator
type Milk struct {
	Coffee
}

func (m *Milk) Cost() float64 {
	return m.Coffee.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.Coffee.Description() + ", Milk"
}

// second decorator
type Caramel struct {
	Coffee
}

func (c *Caramel) Cost() float64 {
	return c.Coffee.Cost() + 1.0
}

func (c *Caramel) Description() string {
	return c.Coffee.Description() + ", Caramel"
}
