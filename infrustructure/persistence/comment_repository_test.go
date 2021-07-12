package persistence

import (
	"github.com/stretchr/testify/assert"
	"go-cource-api/domain/entity"
	"testing"
)

func TestFindAllComments_Success(t *testing.T) {
	conn := DBConn()
	_ = SeedComment(conn)

	repo := NewCommentRepository(conn)

	comments, err := repo.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, comments)
	assert.EqualValues(t, len(comments), 1)
}

func TestFindByIdComment_Success(t *testing.T) {
	conn := DBConn()
	comment := SeedComment(conn)

	repo := NewCommentRepository(conn)

	founded, err := repo.FindById(comment.Id)

	assert.Nil(t, err)
	assert.NotNil(t, founded)
	assert.EqualValues(t, founded.Id, comment.Id)
}

func TestFindByPostIdComments_Success(t *testing.T) {
	conn := DBConn()
	comment := SeedComment(conn)

	repo := NewCommentRepository(conn)

	comments, err := repo.FindByPostId(comment.PostId)

	assert.Nil(t, err)
	assert.NotNil(t, comments)
	assert.EqualValues(t, len(comments), 1)
}

func TestSaveComment_Success(t *testing.T) {
	conn := DBConn()
	post := SeedPost(conn)
	repo := NewCommentRepository(conn)

	comment := &entity.Comment{
		UserId: post.UserId,
		PostId: post.Id,
		Body: "test body",
	}

	saved, err := repo.Save(comment)

	assert.Nil(t, err)
	assert.NotNil(t, saved)
}

func TestSaveComment_Failure(t *testing.T) {
	conn := DBConn()
	repo := NewCommentRepository(conn)

	comment := &entity.Comment{}

	_, err := repo.Save(comment)

	assert.NotNil(t, err)
}

func TestDeleteComment_Success(t *testing.T) {
	conn := DBConn()
	repo := NewCommentRepository(conn)

	comment := SeedComment(conn)

	 err := repo.Delete(comment.Id)

	assert.Nil(t, err)
}

func TestUpdateComment_Success(t *testing.T) {
	conn := DBConn()
	repo := NewCommentRepository(conn)

	comment := SeedComment(conn)

	updated := &entity.Comment{
		Body: "updated",
		Id: comment.Id,
		UserId: comment.UserId,
		PostId: comment.PostId,
	}

	err := repo.Update(updated)

	assert.Nil(t, err)
}
