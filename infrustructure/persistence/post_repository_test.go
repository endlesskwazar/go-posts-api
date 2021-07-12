package persistence

import (
	"github.com/stretchr/testify/assert"
	"go-cource-api/domain/entity"
	"testing"
)

func TestSavePost_Success(t *testing.T) {
	conn := DBConn()
	testUser := SeedUser(conn)

	var post = entity.Post{
		Title: "test title",
		Body: "test body",
		UserId: testUser.Id,
	}

	repo := NewPostRepository(conn)

	saved, saveErr := repo.Save(&post)
	assert.Nil(t, saveErr)
	assert.EqualValues(t, saved.Title, post.Title)
	assert.EqualValues(t, saved.Body, post.Body)
	assert.EqualValues(t, saved.UserId, testUser.Id)
}

func TestSavePost_Failure(t *testing.T) {
	conn := DBConn()

	var post = entity.Post{}

	repo := NewPostRepository(conn)

	_, saveErr := repo.Save(&post)

	assert.NotNil(t, saveErr)
}
