package auth

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"github.com/PUDDLEEE/puddleee_back/ent"
	"github.com/PUDDLEEE/puddleee_back/ent/verification"
	authdto "github.com/PUDDLEEE/puddleee_back/internal/auth/dto"
	"github.com/google/uuid"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

type AuthRepository struct {
}

func (a *AuthRepository) Create(ctx context.Context, client *ent.Client, dto authdto.CreateVerificationDTO) (*ent.Verification, error) {
	uuid, err := uuid.Parse(dto.UUID)

	if err != nil {
		return nil, err
	}
	newVerification, err := client.Verification.Create().SetUUID(uuid).SetCode(dto.Code).Save(ctx)

	if err != nil {
		return nil, err
	}
	return newVerification, nil
}

func (a *AuthRepository) Update(ctx context.Context, client *ent.Client, uuid string, dto authdto.UpdateVerificationDTO) (*ent.Verification, error) {
	verification, err := a.FindOneByUUID(ctx, client, uuid)

	if err != nil {
		return nil, err
	}

	updatedVerification, err := client.Verification.UpdateOne(verification).SetCode(dto.Code).Save(ctx)
	if err != nil {
		return nil, err
	}
	return updatedVerification, nil
}

func (a *AuthRepository) FindOneByUUID(ctx context.Context, client *ent.Client, uuidStr string) (*ent.Verification, error) {
	parsedUUID, err := uuid.Parse(uuidStr)

	if err != nil {
		return nil, err
	}
	verification, err := client.Verification.Query().Where(verification.UUID(parsedUUID)).Only(ctx)

	if err != nil {
		return nil, err
	}
	return verification, nil
}

func (a *AuthRepository) Delete(ctx context.Context, client *ent.Client, uuid string) error {
	verification, err := a.FindOneByUUID(ctx, client, uuid)

	if err != nil {
		return err
	}

	err = client.Verification.DeleteOne(verification).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) GenerateUUIDFromString(email string) uuid.UUID {
	md5hash := md5.New()
	md5hash.Write([]byte(email))

	md5string := hex.EncodeToString(md5hash.Sum(nil))

	uuid, err := uuid.FromBytes([]byte(md5string[0:16]))
	if err != nil {
		log.Fatal(err)
	}
	return uuid
}

func (a *AuthRepository) GenerateCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}
