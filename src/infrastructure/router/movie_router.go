package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rasyidmm/movie-xsis/src/infrastructure/services"
)

func NewMovie(f fiber.Router, service *services.MovieService) *fiber.Router {
	r := f.Group("/movie")
	r.Get("/like", service.GetMovieByLike)
	r.Get("/:id", service.GetMovie)
	r.Get("/", service.GetMovies)
	r.Post("/", service.CreateMovie)
	r.Patch("/:id", service.UpdateMovie)
	r.Delete("/:id", service.DeleteMovie)

	return &f
}
