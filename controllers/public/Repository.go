package public

import (
	"encoding/json"
	"html/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// Repository managers the operations for displaying repositories
type Repository struct {
	Controller
	model repository.Model
}

// List shows all the institutions
func (controller *Repository) List() {
	var exists bool
	list := controller.model.GetAllInstitutions()
	exists = (len(list) != 0)
	controller.Data["anything"] = exists
	controller.Data["institutions"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/public/institution/list.tpl"
}

// ShowResource shows the page for the resource
func (controller *Repository) ShowResource() {
	repositoryID := controller.Ctx.Input.Param(":repository")
	resourceID := controller.Ctx.Input.Param(":resource")
	repo, errRepository := repository.NewRepository(repositoryID)
	resource, errResource := wisply.GetRecordByID(resourceID)
	if errRepository != nil || errResource != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["repository"] = repo
		controller.Data["institution"] = repo.GetInstitution()
		controller.Data["resource"] = resource
		controller.TplNames = "site/public/repository/resource.tpl"
	}
}

// GetResourceContent gets the content of the resource
func (controller *Repository) GetResourceContent() {
	resourceID := controller.Ctx.Input.Param(":resource")
	resource, errResource := wisply.GetRecordByID(resourceID)
	if errResource != nil {
		controller.Abort("databaseError")
	} else {
		externalURL := resource.Keys.GetURL()
		//
		// fmt.Println("url extern")
		// fmt.Println(externalURL)
		//
		// resp, _ := http.Get()
		// bytesFromHTML, _ := ioutil.ReadAll(resp.Body)

		doc, err := goquery.NewDocument(externalURL)
		// handle err
		if err != nil {
			panic(err)
		}
		value := ""
		doc.Find(".ep_summary_content_main").Each(func(i int, s *goquery.Selection) {
			value, _ = s.Html()
		})

		//value = strings.Replace(value, "src=\"/images", "src=\"http://www.edshare.soton.ac.uk/images", -1)

		controller.Data["value"] = template.HTML([]byte(value))
		controller.TplNames = "site/public/repository/content.tpl"

	}
}

// ShowRepository shows the details regarding a repository
func (controller *Repository) ShowRepository() {
	ID := controller.Ctx.Input.Param(":repository")
	repo, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
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

		controller.TplNames = "site/public/repository/repository.tpl"
	}
}
