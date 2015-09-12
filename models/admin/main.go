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

	numberOfAccounts := auth.Count()
	numberOfSources := sources.Count()

	return &Dashboard{
		Accounts: numberOfAccounts,
		Sources:  numberOfSources,
	}

}
