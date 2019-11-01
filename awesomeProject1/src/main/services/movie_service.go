package services

import (
	"github.com/kataras/iris/_examples/mvc/overview/datamodels"
	"github.com/kataras/iris/_examples/mvc/overview/repositories"
)

type MovieService interface {
	GetAll() []datamodels.Movie
	GetByID(id int64)(datamodels.Movie,bool)
	DeleteByID(id int64)bool
	UpdatePosterAndGenreByID(id int64,poster string,genre string)(datamodels.Movie,error)
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{
		repo: repo,
	}
}

type movieService struct {
	repo repositories.MovieRepository
}

func (s *movieService) GetAll() []datamodels.Movie {
	return s.repo.SelectMany(func(_ datamodels.Movie) bool {
		return true
	}, -1)
}

func (s *movieService) GetByID(id int64) (datamodels.Movie,bool) {
	return s.repo.Select(func(m datamodels.Movie) bool {
		return m.ID == id
	})
}

func (s *movieService) DeleteByID(id int64) bool {
	return s.repo.Delete(func(m datamodels.Movie) bool {
		return m.ID == id
	}, 1)
}

func (s *movieService) UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error) {
	return s.repo.InsertOrUpdate(datamodels.Movie{
		ID:     id,
		Poster: poster,
		Genre:  genre,
	})
}

