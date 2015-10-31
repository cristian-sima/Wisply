// Package repository contains the basic remote repository.
// It must be inherited by the other remote repositories
package repository

import "github.com/cristian-sima/Wisply/models/repository"

// Repository represents a simple remote repository
type Repository struct {
	repository *repository.Repository
}

// SetLocalRepository changes the local repository
func (repository *Repository) SetLocalRepository(local *repository.Repository) {
	repository.repository = local
}

// GetLocalRepository returns the reference to the local repository
func (repository *Repository) GetLocalRepository() *repository.Repository {
	return repository.repository
}
