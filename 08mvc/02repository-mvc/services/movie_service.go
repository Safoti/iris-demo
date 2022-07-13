package main

import (
	"github.com/kataras/iris/v12/_examples/mvc/overview/repositories"
	"myapp/08mvc/02repository-mvc/datamodels"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description
 **/
//定义service 接口

// MovieService handles some of the CRUID operations of the movie datamodel.
// It depends on a movie repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type MovieService interface {
	GetAll() datamodels.Movie
	GetByID(id int64) (datamodels.Movie, bool)
	DeleteByID(id int64) bool
	UpdatePosterAndGenreByID(id int64, poster string, genre string) (datamodels.Movie, error)
}

//func NewMoviceService(repo repositories.MovieRepository)MovieService{
//	return &movieService{
//		repo: repo,
//	}
//}
type movieService struct {
	repo repositories.MovieRepository
}

func (s *movieService) GetAll() []datamodels.Movie {

}
