package repository

import (
	belajar_golang_datatabase "belajar-golang-datatabase-mysql"
	"belajar-golang-datatabase-mysql/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_datatabase.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "test repository",
	}

	newComment, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(newComment)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_datatabase.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_datatabase.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
