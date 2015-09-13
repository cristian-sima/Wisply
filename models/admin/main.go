package admin

import (
	auth "github.com/cristian-sima/Wisply/models/auth"
	sources "github.com/cristian-sima/Wisply/models/sources"
)

type Dashboard struct {
	Accounts int
	Sources  int
}

func GetDashboard() *Dashboard {

	numberOfAccounts := auth.CountAccounts()
	numberOfSources := sources.CountSources()

	return &Dashboard{
		Accounts: numberOfAccounts,
		Sources:  numberOfSources,
	}

}
