package repositories

type MoveRepository interface {
	GetName() (string, error)
}

type MovieManager struct {
}

func (m *MovieManager) GetName() (string, error) {
	return "movie", nil
}

func NewMovieManager() MoveRepository {
	return &MovieManager{}
}
