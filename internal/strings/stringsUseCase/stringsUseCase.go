package stringsUseCase

import (
	"awsproj/internal/strings"
	"log"
)

type StringsUsecase struct {
	sR strings.Repository
}

func NewStringsUseCase(sR strings.Repository) *StringsUsecase {
	return &StringsUsecase{
		sR: sR,
	}
}

func (sUC *StringsUsecase) HashString(str string) string {
	log.Println("[Start hash...]")
	res, err := sUC.sR.Hash(str)

	if err != nil {
		log.Println(err)
		return res
	}

	log.Println("[End hashing...]")
	return res
}
