package services

type BookService interface {
	Uppercase(string) (string, error)
	Count(string) int
}
