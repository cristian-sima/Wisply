// Package admin manages the administrator overview page
package admin

import (
	auth "github.com/cristian-sima/Wisply/models/auth"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// Dashboard It represents the administrator's dashboard
type Dashboard struct {
	Accounts     int
	Repositories int
}

// NewDashboard It creates a new Dashboard
func NewDashboard() *Dashboard {
	return &Dashboard{
		Accounts:     auth.CountAccounts(),
		Repositories: repository.CountRepositories(),
	}
}
