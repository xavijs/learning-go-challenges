package ad

type AdRepository interface {
	FindBy(id Id) Ad

	Persist(ad Ad)

	FindAll() []Ad
}
