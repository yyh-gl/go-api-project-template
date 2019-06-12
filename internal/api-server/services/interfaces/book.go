package interfaces

type BookService interface {
	Index(string) (string, error)
}
