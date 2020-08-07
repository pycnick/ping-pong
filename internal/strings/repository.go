package strings

type Repository interface {
	Hash(str string) (string, error)
}
