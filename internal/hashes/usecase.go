package hashes

type UseCase interface {
	GenerateHash(str string) string
}
