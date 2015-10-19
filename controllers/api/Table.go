package api

import (
	"io"
	"os"
	"time"

	"github.com/cristian-sima/Wisply/models/api"
)

var validNumberOfHours = 3600 * 24

// Table holds all the methods for downloading the Wisply tables
type Table struct {
	Controller
}

// ShowList displays the list of available tables to be downloaded
func (controller *Table) ShowList() {
	controller.Data["tables"] = api.GetAllowedTables()
	controller.Layout = "site/public-layout.tpl"
	controller.TplNames = "site/api/table/list.tpl"
	// Please use http://www.timestampgenerator.com/
	controller.SetCustomTitle("API & Developers")
	controller.IndicateLastModification(1445250987)
}

// DownloadTable starts the process of downloading the table
func (controller *Table) setHeadersDownload(filename string) {

	controller.Ctx.Output.Header("Content-Description", "File Transfer")
	controller.Ctx.Output.Header("Content-Type", "application/octet-stream")

	controller.Ctx.Output.Header("Content-Disposition", "attachment; filename="+filename)

	controller.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	controller.Ctx.Output.Header("Expires", "0")
	controller.Ctx.Output.Header("Cache-Control", "must-revalidate")
	controller.Ctx.Output.Header("Pragma", "public")

}

// GenerateTable generates the table if there is no table or it is too old
func (controller *Table) GenerateTable() {
	tableName := controller.Ctx.Input.Param(":name")
	filename := tableName + "." + "csv"
	fullPath := "cache/api/tables/"
	file, err := controller.getTable(fullPath + "/" + filename)
	if err != nil {
		api.GenerateTableFile(tableName)
	} else {
		if controller.checkFileIsStillValid(fullPath + "/" + filename) {
			controller.closeFile(file)
			controller.deleteFile(fullPath + "/" + filename)
			api.GenerateTableFile(tableName)
		}
	}
	controller.ShowBlankPage()
}

// DownloadTable downloads a table
func (controller *Table) DownloadTable() {
	tableName := controller.Ctx.Input.Param(":name")
	filename := tableName + "." + "csv"
	fullPath := "cache/api/tables/"

	if !api.IsAllowedTable(tableName) {
		controller.DisplaySimpleError("This table name is restricted.")
	} else {
		var (
			file *os.File
			err  error
		)
		file, err = controller.getTable(fullPath + "/" + filename)
		if err == nil {
			controller.setHeadersDownload("Table " + tableName + ".csv")
			controller.readFile(file, fullPath)
			controller.closeFile(file)
		} else {
			controller.ShowBlankPage()
		}
	}

}

func (controller *Table) checkFileIsStillValid(path string) bool {
	info, _ := os.Stat(path)
	duration := time.Since(info.ModTime()).Seconds()
	if int(duration) >= validNumberOfHours {
		return true
	}
	return false
}

func (controller *Table) getTable(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return file, err
	}
	return file, nil
}

func (controller *Table) deleteFile(path string) {
	if err := os.Remove(path); err != nil {
		panic("Closing error: " + err.Error())
	}
}

func (controller *Table) closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic("Closing error: " + err.Error())
	}
}

func (controller *Table) readFile(file *os.File, fullPath string) {

	// make a buffer to keep chunks that are read
	buffer := make([]byte, 1024)
	body := []byte{}
	for {
		size, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if size == 0 {
			break
		}
		body = append(body, buffer[:size]...)
	}
	controller.Ctx.Output.Body(body)
}
