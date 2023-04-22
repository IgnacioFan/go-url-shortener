package repository

type UrlRepository interface {
	Create(url string) (uint64, error)
	FindBy(url string) (uint64, error)
	Find(id uint64) (string, error)
}
