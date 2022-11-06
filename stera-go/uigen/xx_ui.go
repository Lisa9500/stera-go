// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UIXxMainWindow struct {
	Centralwidget *widgets.QWidget
	PushButton *widgets.QPushButton
	ListView *widgets.QListView
	Label *widgets.QLabel
	Label2 *widgets.QLabel
	Menubar *widgets.QMenuBar
	Statusbar *widgets.QStatusBar
}

func (this *UIXxMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 640, 480))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.PushButton = widgets.NewQPushButton(this.Centralwidget)
	this.PushButton.SetObjectName("PushButton")
	this.PushButton.SetGeometry(core.NewQRect4(60, 70, 75, 23))
	this.ListView = widgets.NewQListView(this.Centralwidget)
	this.ListView.SetObjectName("ListView")
	this.ListView.SetGeometry(core.NewQRect4(210, 70, 256, 192))
	this.Label = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.Label.SetObjectName("Label")
	this.Label.SetGeometry(core.NewQRect4(60, 50, 50, 12))
	this.Label2 = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.Label2.SetObjectName("Label2")
	this.Label2.SetGeometry(core.NewQRect4(210, 50, 50, 12))
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 640, 21))
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)


    this.RetranslateUi(MainWindow)

}

func (this *UIXxMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.PushButton.SetText(_translate("MainWindow", "PushButton", "", -1))
	this.Label.SetText(_translate("MainWindow", "TextLabel", "", -1))
	this.Label2.SetText(_translate("MainWindow", "TextLabel", "", -1))
}
