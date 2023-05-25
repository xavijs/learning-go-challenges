package ad

type AdRepository interface {
	FindBy(id Id) (*Ad, error)

	Persist(ad *Ad) error

	FindAll() (*[]Ad, error)
}
