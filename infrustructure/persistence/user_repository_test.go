package persistence

import (
	"github.com/stretchr/testify/assert"
	"go-cource-api/domain/entity"
	"gopkg.in/guregu/null.v4"
	"testing"
)

func TestFindByIdUser_Success(t *testing.T) {
	conn := DBConn()
	user := SeedUser(conn)
	repo := NewUserRepository(conn)

	founded, err := repo.FindById(user.Id)

	assert.Nil(t, err)
	assert.EqualValues(t, user.Id, founded.Id)
}

func TestFindByIdUser_Failure(t *testing.T) {
	conn := DBConn()
	repo := NewUserRepository(conn)
	notFoundId := int64(1)

	_, err := repo.FindById(notFoundId)

	assert.NotNil(t, err)
}

func TestSaveUser_Success(t *testing.T) {
	conn := DBConn()
	repo := NewUserRepository(conn)

	user := &entity.User{
		Name:     null.StringFrom("test"),
		Email:    null.StringFrom("test"),
		Password: null.StringFrom("qweqweqw098798q6475u23hrwrkl"),
	}

	saved, err := repo.Save(user)

	assert.Nil(t, err)
	assert.NotNil(t, saved)
	assert.EqualValues(t, saved.Name, user.Name)
	assert.EqualValues(t, saved.Email, user.Email)
}

func TestSaveUser_Failure(t *testing.T) {
	conn := DBConn()
	seeded := SeedUser(conn)
	repo := NewUserRepository(conn)

	user := &entity.User{
		Name:     null.StringFrom("test"),
		Email:    seeded.Email,
		Password: null.StringFrom("qweqweqw098798q6475u23hrwrkl"),
	}

	_, err := repo.Save(user)

	assert.NotNil(t, err)
}

func TestFindByEmailUser_Success(t *testing.T) {
	conn := DBConn()
	seeded := SeedUser(conn)
	repo := NewUserRepository(conn)

	found, err := repo.FindByEmail(seeded.Email.String)

	assert.Nil(t, err)
	assert.EqualValues(t, seeded.Id, found.Id)
}
