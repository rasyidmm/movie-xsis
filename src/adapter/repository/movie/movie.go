package movie

import (
	"context"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"gorm.io/gorm"
)

type MovieDataHandler struct {
	db *gorm.DB
}

func NewMovieDataHandler(db *gorm.DB) *MovieDataHandler {
	return &MovieDataHandler{
		db: db,
	}
}

func (m *MovieDataHandler) CreateMovie(ctx context.Context, input *entity.MovieEntity) error {
	err := m.db.WithContext(ctx).Create(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MovieDataHandler) UpdateMovie(ctx context.Context, input *entity.MovieEntity) error {
	err := m.db.WithContext(ctx).Updates(&input).Where("id = ?", input.Id).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *MovieDataHandler) DeleteMovie(ctx context.Context, id int64) error {
	err := m.db.WithContext(ctx).Delete(&entity.MovieEntity{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *MovieDataHandler) GetMovie(ctx context.Context, id int64) (*entity.MovieEntity, error) {
	var result *entity.MovieEntity
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (m *MovieDataHandler) GetMovies(ctx context.Context) (*[]entity.MovieEntity, error) {
	var result *[]entity.MovieEntity
	err := m.db.WithContext(ctx).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
