package repository

import (
	"github.com/SawitProRecruitment/UserService/internal/user/port/driven"
	"golang.org/x/crypto/bcrypt"
)

var _ driven.Encyptor = new(BcyrpEncryption)

type BcyrpEncryption struct{}

// Encrypt implements driven.Encyptor.
func (*BcyrpEncryption) Encrypt(data []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, cost)
}

func (*BcyrpEncryption) CompareEncryptedAndData(encrypted, data []byte) error {
	return bcrypt.CompareHashAndPassword(encrypted, data)
}
