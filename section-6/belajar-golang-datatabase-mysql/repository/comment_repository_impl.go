package repository

import (
	"belajar-golang-datatabase-mysql/entity"
	"context"
	"database/sql"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	panic("implement me")
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	panic("implement me")
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	panic("implement me")
}
