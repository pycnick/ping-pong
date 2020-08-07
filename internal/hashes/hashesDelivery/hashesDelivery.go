package hashesDelivery

import (
	"awsproj/internal/hashes"
	"context"
	"errors"
	"log"
)

type HashesDelivery struct {
	UnimplementedHasherServer
	hUC hashes.UseCase
}

func NewHashesDelivery(hUC hashes.UseCase) *HashesDelivery {
	return &HashesDelivery{
		hUC: hUC,
	}
}

func (hD *HashesDelivery) GetHash(ctx context.Context, request *HashRequest) (*HashReply, error) {
	log.Println("[Server call...]")
	log.Println("Requested string: " + request.Str)

	newHash := hD.hUC.GenerateHash(request.Str)

	if newHash == "" {
		return &HashReply{
			Hash: newHash,
			Error: "Error while hashing string",
		}, errors.New("procedure fail")
	}

	log.Println("Result hash: " + newHash)
	log.Println("[done...]")
	return &HashReply{
		Hash: newHash,
		Error: "",
	}, nil
}
