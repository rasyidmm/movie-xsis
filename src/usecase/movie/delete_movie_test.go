package movie

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rasyidmm/movie-xsis/src/shared/mock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMovieUsecase_DeleteMovie(t *testing.T) {
	repo := &mock.MovieRepository{}
	Convey("TestMovieUsecase_DeleteMovie", t, func() {
		Convey("Positive Case", func() {
			repo.On("DeleteMovie", context.Background(), int64(1)).Return(nil).Once()
			uc := NewMovieUsecase(repo)
			err := uc.DeleteMovie(context.Background(), int64(1))
			So(err, ShouldBeNil)
		})
		Convey("Negative Case", func() {
			repo.On("DeleteMovie", context.Background(), int64(1)).Return(errors.New("error")).Once()
			uc := NewMovieUsecase(repo)
			err := uc.DeleteMovie(context.Background(), int64(1))
			So(err, ShouldNotBeNil)
		})
	})
}
