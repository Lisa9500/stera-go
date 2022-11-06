package dlg

import (
	// "io/ioutil"

	"io/ioutil"
	"log"

	// "strings"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

// open file function.
func Opfl() (string, string) {
	result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
		Title: "Open File Dialog",
		Role:  "OpenFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "Text Files (*.txt)",
				Pattern:     "*.txt",
			},
			{
				DisplayName: "GeoJSON Files (*.geojson)",
				Pattern:     "*.geojson",
			},
			{
				DisplayName: "All Files (*.*)",
				Pattern:     "*.*",
			},
		},
		SelectedFileFilterIndex: 1,
		FileName:                "file.geojson",
		DefaultExtension:        "geojson",
	})
	if err == cfd.ErrorCancelled {
		log.Fatal("Dialog was cancelled by the user.")
	} else if err != nil {
		log.Fatal(err)
	}
	fn := result

	ba, er := ioutil.ReadFile(fn)
	if er != nil {
		log.Fatal(er)
	} else {
		log.Print("Open from file '" + fn + "'.")
	}
	return fn, string(ba)
}
