package repository

import (
	"encoding/json"

	"github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// Repository manages the operations for an repository
type Repository struct {
	Controller
}

// Display shows the public page for an repository
func (controller *Repository) Display() {
	repo := controller.GetRepository()
	controller.SetCustomTitle(repo.Name)
	controller.Data["institution"] = repo.GetInstitution()
	controller.Data["process"] = harvest.NewProcess(repo.LastProcess)

	if repo.HasBeenProcessed() {
		process := harvest.NewProcess(repo.LastProcess)
		storage := wisply.NewStorage(repo)
		collections := storage.GetCollections()
		collectionsJSON, _ := json.Marshal(collections)
		controller.Data["collections"] = collections
		controller.Data["collectionsJSON"] = string(collectionsJSON)
		controller.Data["process"] = process
		controller.IndicateLastModification(process.Process.End)
	}

	controller.LoadTemplate("home")
}

// ShowResource shows the page for the resource
func (controller *Repository) ShowResource() {
	resourceID := controller.Ctx.Input.Param(":resource")
	repo := controller.GetRepository()
	resource, errResource := wisply.GetRecordByID(resourceID)

	if errResource != nil {
		controller.Abort("show-database-error")
	} else {
		controller.Data["repository"] = repo
		controller.Data["institution"] = repo.GetInstitution()
		controller.Data["resource"] = resource
		controller.LoadTemplate("resource")
	}
}

// GetResourceContent gets the content of the resource
func (controller *Repository) GetResourceContent() {
	// externalURL := resource.Keys.GetURL()
	// //
	// // fmt.Println("url extern")
	// // fmt.Println(externalURL)
	// //
	// // resp, _ := http.Get()
	// // bytesFromHTML, _ := ioutil.ReadAll(resp.Body)
	//
	// doc, err := goquery.NewDocument(externalURL)
	// // handle err
	// if err != nil {
	// 	panic(err)
	// }
	// value := ""
	// doc.Find(".ep_summary_content_main").Each(func(i int, s *goquery.Selection) {
	// 	value, _ = s.Html()
	// })
	//
	// //value = strings.Replace(value, "src=\"/images", "src=\"http://www.edshare.soton.ac.uk/images", -1)
	//
	// controller.Data["value"] = template.HTML([]byte(value))
	// controller.TplNames = "site/public/repository/content.tpl"
}

// ShowRepository shows the details regarding a repository
func (controller *Repository) ShowRepository() {
	repo := controller.GetRepository()
	controller.Data["repository"] = repo
	controller.SetCustomTitle(repo.Name)

	controller.Data["institution"] = repo.GetInstitution()
	controller.Data["identification"] = repo.GetIdentification()

	if repo.HasBeenProcessed() {
		process := harvest.NewProcess(repo.LastProcess)

		storage := wisply.NewStorage(repo)

		collections := storage.GetCollections()

		controller.Data["collections"] = collections

		collectionsJSON, _ := json.Marshal(collections)
		controller.Data["collectionsJSON"] = string(collectionsJSON)

		controller.Data["process"] = process
		controller.IndicateLastModification(process.Process.End)
	}
	controller.LoadTemplate("repository")
}
