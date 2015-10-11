package remote

import "github.com/cristian-sima/Wisply/models/repository"

// Repository represents a simple remote repository
type Repository struct {
	repository *repository.Repository
}

// GetRepository returns the reference to the local repository
func (remote *Repository) GetRepository() *repository.Repository {
	return remote.repository
}
