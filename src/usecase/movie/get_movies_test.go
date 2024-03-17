package movie

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rasyidmm/movie-xsis/src/adapter/repository/entity"
	"github.com/rasyidmm/movie-xsis/src/shared/mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMovieUsecase_GetMovies(t *testing.T) {
	repo := &mock.MovieRepository{}
	Convey("TestMovieUsecase_GetMovies", t, func() {
		Convey("Positive Case", func() {
			repo.On("GetMovies", context.Background()).Return(&[]entity.MovieEntity{}, nil).Once()
			uc := NewMovieUsecase(repo)
			res, err := uc.GetMovies(context.Background())
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
		})
		Convey("Negative Case", func() {
			repo.On("GetMovies", context.Background()).Return(nil, errors.New("error")).Once()
			uc := NewMovieUsecase(repo)
			res, err := uc.GetMovies(context.Background())
			So(err, ShouldNotBeNil)
			So(res, ShouldBeNil)
		})
	})
}
