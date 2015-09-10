package admin

import (
	auth "github.com/cristian-sima/Wisply/models/auth"
	sources "github.com/cristian-sima/Wisply/models/sources"
)

type Dashboard struct {
	Users   int
	Sources int
}

func GetDashboard() *Dashboard {

	numberOfUsers := auth.Count()
	numberOfSources := sources.Count()

	return &Dashboard{
		Users:   numberOfUsers,
		Sources: numberOfSources,
	}

}
