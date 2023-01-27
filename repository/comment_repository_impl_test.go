package repository

import (
	"belajar_golang_database"
	"belajar_golang_database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "paulus.lestyo@gmail.com",
		Comment: "Test Repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comment, err := commentRepository.FindByID(context.Background(), 27)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
