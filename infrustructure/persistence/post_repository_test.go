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

func TestDeletePost_Success(t *testing.T) {
	conn := DBConn()
	post := SeedPost(conn)
	repo := NewPostRepository(conn)

	err := repo.Delete(post.Id)

	assert.Nil(t, err)
}

func TestUpdatePost_Success(t *testing.T) {
	conn := DBConn()
	post := SeedPost(conn)
	repo := NewPostRepository(conn)

	updatedPost := &entity.Post{
		Id: post.Id,
		Title: "Updated",
		Body: "Updated",
	}

	err := repo.Update(updatedPost)

	assert.Nil(t, err)
}

func TestFindAll_Success(t *testing.T) {
	conn := DBConn()
	_ = SeedPost(conn)
	repo := NewPostRepository(conn)

	posts, err := repo.FindAll()
	assert.Nil(t, err)
	assert.EqualValues(t, len(posts), 1)
}

func TestFindById_Success(t *testing.T) {
	conn := DBConn()
	post := SeedPost(conn)
	repo := NewPostRepository(conn)

	foundPost, err := repo.FindById(post.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, post.Id, foundPost.Id)
}

func TestFindById_Failure(t *testing.T) {
	conn := DBConn()
	repo := NewPostRepository(conn)

	notFoundId := uint64(1)

	_, err := repo.FindById(notFoundId)
	assert.NotNil(t, err)
}

func TestFindByIdAndUserId_Success(t *testing.T) {
	conn := DBConn()
	post := SeedPost(conn)
	repo := NewPostRepository(conn)

	foundedPost, err := repo.FindByIdAndUserId(post.Id, post.UserId)
	assert.Nil(t, err)
	assert.NotNil(t, foundedPost)
	assert.EqualValues(t, post.Id, foundedPost.Id)
}


