package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasyidmm/movie-xsis/src/shared/web"
	"github.com/rasyidmm/movie-xsis/src/usecase/movie"
	"github.com/rasyidmm/movie-xsis/src/usecase/payload"
	"strconv"
)

type MovieService struct {
	MovieUsecase movie.MoviePort
}

func NewMovieService(m movie.MoviePort) *MovieService {
	return &MovieService{
		MovieUsecase: m,
	}
}

// GetMovie: Handler function for Get Movie
// @Summary     Get Movie by ids
// @Description Endpoint for Get Movie by id
// @Tags        Movie
// @Movie     json
// @Param id path string true "id of movie"
// @Success     200     {object} payload.MoviePayload
// @Failure     400     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     500     {object} web.Response
// @Router      /v1/movie/{id} [GET]
func (s *MovieService) GetMovie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	result, err := s.MovieUsecase.GetMovie(ctx.Context(), int64(id))
	if err != nil {
		return web.ErrResponse(ctx, err)
	}

	return web.OKResponse(ctx, fiber.StatusOK, result)
}

// GetMovies: Handler function for Get Movie
// @Summary     Get Movies
// @Description Endpoint for Get Movies
// @Tags        Movie
// @Movie     json
// @Success     200     {object} []payload.MoviePayload
// @Failure     400     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     500     {object} web.Response
// @Router      /v1/movies [GET]
func (s *MovieService) GetMovies(ctx *fiber.Ctx) error {

	result, err := s.MovieUsecase.GetMovies(ctx.Context())
	if err != nil {
		return web.ErrResponse(ctx, err)
	}

	return web.OKResponse(ctx, fiber.StatusOK, result)
}

// CreateMovie: Handler function for Create Movie
// @Summary     Create Movie
// @Description Endpoint for Create Movie
// @Tags        Movie
// @Movie     json
// @Param request body payload.CreateMoviePayload true "request"
// @Success     200     {object} web.Response
// @Failure     400     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     500     {object} web.Response
// @Router      /v1/movie [POST]
func (s *MovieService) CreateMovie(ctx *fiber.Ctx) error {
	var input payload.CreateMoviePayload
	if err := ctx.BodyParser(&input); err != nil {
		return web.ErrResponse(ctx, err)
	}

	err := web.Validate().Struct(input)
	if err != nil {
		return web.ErrResponse(ctx, err)
	}

	err = s.MovieUsecase.CreateMovie(ctx.Context(), input)
	if err != nil {
		return web.ErrResponse(ctx, err)
	}

	return web.OKResponse(ctx, fiber.StatusOK, "success")
}

// UpdateMovie: Handler function for Update Movie
// @Summary     Update Movie
// @Description Endpoint for Update Movie
// @Tags        Movie
// @Movie     json
// @Param id path string true "id of movie"
// @Param request body payload.UpdateMoviePayload true "request"
// @Success     200     {object} web.Response
// @Failure     400     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     500     {object} web.Response
// @Router      /v1/movie/{id} [PATCH]
func (s *MovieService) UpdateMovie(ctx *fiber.Ctx) error {
	var input payload.UpdateMoviePayload
	id, _ := strconv.Atoi(ctx.Params("id"))
	if err := ctx.BodyParser(&input); err != nil {
		return web.ErrResponse(ctx, err)
	}
	input.Id = int64(id)

	err := web.Validate().Struct(input)
	if err != nil {
		return web.ErrResponse(ctx, err)
	}

	err = s.MovieUsecase.UpdateMovie(ctx.Context(), input)
	if err != nil {
		return web.ErrResponse(ctx, err)
	}
	return web.OKResponse(ctx, fiber.StatusOK, "success")
}

// DeleteMovie: Handler function for Delete Movie
// @Summary    Delete Movie by id
// @Description Endpoint for Delete Movie by id
// @Tags        Movie
// @Movie     json
// @Param id path string true "id of movie"
// @Success     200     {object} web.Response
// @Failure     400     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     500     {object} web.Response
// @Router      /v1/movie/{id} [DELETE]
func (s *MovieService) DeleteMovie(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))
	err := s.MovieUsecase.DeleteMovie(ctx.Context(), int64(id))
	if err != nil {
		return web.ErrResponse(ctx, err)
	}
	return web.OKResponse(ctx, fiber.StatusOK, "success")

}

// GetMovieByLike: Handler function for Get Movie By Like
// @Summary     Get Movie By Like
// @Description Endpoint for  Get Movie By Like
// @Tags        Movie
// @Movie     json
// @Query id path string true "title"
// @Success     200     {object} []payload.MoviePayload
// @Failure     400     {object} web.Response
// @Failure     401     {object} web.Response
// @Failure     500     {object} web.Response
// @Router      /v1/movie/like [GET]
func (s *MovieService) GetMovieByLike(ctx *fiber.Ctx) error {
	title := ctx.Query("title")
	result, err := s.MovieUsecase.GetMovieByLike(ctx.Context(), title)
	if err != nil {
		return web.ErrResponse(ctx, err)
	}
	return web.OKResponse(ctx, fiber.StatusOK, result)

}
