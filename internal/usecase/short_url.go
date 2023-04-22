package usecase

type ShortUrl interface {
	Create(url string) (string, error)
	Redirect(url string) (string, error)
}
