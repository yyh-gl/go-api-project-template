package implements

type BookService interface {
	Index(string) (string, error)
}
