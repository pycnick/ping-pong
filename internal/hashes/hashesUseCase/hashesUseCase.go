package hashesUseCase

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashesUseCase struct {
}

func (hUC *HashesUseCase) GenerateHash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}