package types

type AnimalRepo interface {
	GetAll() ([]Animal, error)
	Get(id string) (*Animal, error)
	Create(a *Animal) (*Animal, error)
	Update(a *Animal) (*Animal, error)
	Delete(id string) error
}
type Animal struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
