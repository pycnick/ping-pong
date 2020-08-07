package stringsRepository

import (
	"awsproj/internal/hashes/hashesDelivery"
	"context"
	"errors"
	"log"
)

type StringsRepository struct {
	pb hashesDelivery.HasherClient
}

func NewStringsRepository(pb hashesDelivery.HasherClient) *StringsRepository {
	return &StringsRepository{
		pb: pb,
	}
}

func (sR *StringsRepository) Hash(str string) (string, error) {
	log.Println("Client call Hash function")
	rep, err := sR.pb.GetHash(context.Background(), &hashesDelivery.HashRequest{
		Str: str,
	})

	if err != nil {
		log.Println("Error while remote procedure calling: " + err.Error())
		return "", err
	}

	if rep.Error != "" {
		log.Println("Error while remote procedure executing: " + rep.Error)
		return "", errors.New(rep.Error)
	}
	return rep.Hash, nil
}
