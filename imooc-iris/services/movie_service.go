package services

import "imooc/repositories"

type MovieService interface {
	ShowName() string
}

type MovieServiceManager struct {
	repo repositories.MoveRepository
}

func (m *MovieServiceManager) ShowName() string {
	name, _ := m.repo.GetName()
	println(name)
	return name
}

func NewMovieServiceManager(repo repositories.MoveRepository) MovieService {
	return &MovieServiceManager{
		repo: repo,
	}
}
