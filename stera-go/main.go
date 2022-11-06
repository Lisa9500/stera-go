package main

import (
	"os"
	"stera/dlg"
	"stera/pkg"
	myui "stera/uigen"

	"log"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var (
	filename string
)

type Window struct {
	myui.UISteraGoWindowMainWindow
	Widget *widgets.QMainWindow
}

func NewWidget(parent widgets.QWidget_ITF) *Window {
	window := &Window{
		Widget: widgets.NewQMainWindow(parent, core.Qt__Window),
	}

	window.SetupUI(window.Widget)
	return window
}

// TODO: Add UI logic here
func (w *Window) OpenButtonClicked(checked bool) {
	log.Print("OpenButtonClicked")
	// fn, ba := openfile()
	fn, ba := dlg.Opfl()
	filename = fn
	// ファイル名
	w.LoadFileName.SetText(string(fn))
	// ファイルの内容
	w.TextBrowser.SetText(string(ba))
}
func (w *Window) BuildButtonClicked(checked bool) {
	log.Print("BuildButtonClicked")
	pkg.DivideLine(filename)
	// 普通建物の処理，四角形分割して屋根モデルを作成する
	// pkg.ParseJson("C:/data/other_list.txt")
	// pkg.ParseJson("C:/data/hutsu_list.txt")
	// 堅ろう建物の場合，三角メッシュ分割する
	// 無壁舎の場合，三角メッシュ分割する
}

// func perseLine(fn string) error {
//
// }

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	w := NewWidget(nil)
	w.Widget.Show()

	w.FileOpenButton.ConnectClicked(w.OpenButtonClicked)
	w.BuildStartButton1.ConnectClicked(w.BuildButtonClicked)

	os.Exit(app.Exec())

}
