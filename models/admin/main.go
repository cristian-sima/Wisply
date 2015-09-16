package admin

import (
	auth "github.com/cristian-sima/Wisply/models/auth"
	sources "github.com/cristian-sima/Wisply/models/sources"
)

// Dashboard It represents the administrator's dashboard
type Dashboard struct {
	Accounts int
	Sources  int
}

// NewDashboard It creates a new Dashboard
func NewDashboard() *Dashboard {
	numberOfAccounts := auth.CountAccounts()
	numberOfSources := sources.CountSources()
	return &Dashboard{
		Accounts: numberOfAccounts,
		Sources:  numberOfSources,
	}
}
