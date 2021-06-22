package interfaces

import (
	"encoding/json"
	"fmt"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"net/http"
)

type CreatePostDto struct {
	Title string
	Body string
	UserId uint64
}

type Posts struct {
	app application.PostAppInterface
}

func NewPosts(app application.PostAppInterface) *Posts {
	return &Posts{
		app: app,
	}
}

func (p *Posts) List(w http.ResponseWriter, r *http.Request) {
	users, err := p.app.FindAll()

	if err != nil {

	}

	a, err := json.Marshal(users)
	n := len(a)
	s := string(a[:n])
	fmt.Fprint(w, s)
}

func (p *Posts) Save(w http.ResponseWriter, r *http.Request) {
	var createPostDto CreatePostDto
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createPostDto)

	if err != nil {
		panic(err)
	}

	post := &entity.Post{
		Body: createPostDto.Body,
		Title: createPostDto.Title,
		UserId: createPostDto.UserId,
	}

	savedPost, erro := p.app.Save(post)
	if erro != nil {
		println("error on save")
	}

	a, error := json.Marshal(savedPost)
	if error != nil {
		println("error on marchmal")
	}
	s := string(a)
	fmt.Fprint(w, s)
}
