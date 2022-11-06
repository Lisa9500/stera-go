// WARNING! All changes made in this file will be lost!
package uigen

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UISteraGoWindowMainWindow struct {
	Centralwidget *widgets.QWidget
	LabelTitle *widgets.QLabel
	LabelTitle2 *widgets.QLabel
	LabelFileName *widgets.QLabel
	LabelFileName2 *widgets.QLabel
	BuildStartButton1 *widgets.QPushButton
	BuildStartButton2 *widgets.QPushButton
	FrameBuildings *widgets.QFrame
	LoadFileName *widgets.QLineEdit
	TextBrowser *widgets.QTextBrowser
	FileOpenButton *widgets.QPushButton
	LabelTitle3 *widgets.QLabel
	FrameTerrain *widgets.QFrame
	LoadFileName2 *widgets.QLineEdit
	FileOpenButton2 *widgets.QPushButton
	TextBrowser2 *widgets.QTextBrowser
	Menubar *widgets.QMenuBar
	MenuFile *widgets.QMenu
	Statusbar *widgets.QStatusBar
}

func (this *UISteraGoWindowMainWindow) SetupUI(MainWindow *widgets.QMainWindow) {
	MainWindow.SetObjectName("MainWindow")
	MainWindow.SetGeometry(core.NewQRect4(0, 0, 800, 600))
	this.Centralwidget = widgets.NewQWidget(MainWindow, core.Qt__Widget)
	this.Centralwidget.SetObjectName("Centralwidget")
	this.LabelTitle = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelTitle.SetObjectName("LabelTitle")
	this.LabelTitle.SetGeometry(core.NewQRect4(30, 30, 50, 24))
	this.LabelTitle2 = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelTitle2.SetObjectName("LabelTitle2")
	this.LabelTitle2.SetGeometry(core.NewQRect4(430, 30, 50, 24))
	this.LabelFileName = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelFileName.SetObjectName("LabelFileName")
	this.LabelFileName.SetGeometry(core.NewQRect4(50, 60, 50, 12))
	this.LabelFileName2 = widgets.NewQLabel(this.Centralwidget, core.Qt__Widget)
	this.LabelFileName2.SetObjectName("LabelFileName2")
	this.LabelFileName2.SetGeometry(core.NewQRect4(440, 60, 50, 12))
	this.BuildStartButton1 = widgets.NewQPushButton(this.Centralwidget)
	this.BuildStartButton1.SetObjectName("BuildStartButton1")
	this.BuildStartButton1.SetGeometry(core.NewQRect4(290, 510, 75, 23))
	this.BuildStartButton2 = widgets.NewQPushButton(this.Centralwidget)
	this.BuildStartButton2.SetObjectName("BuildStartButton2")
	this.BuildStartButton2.SetGeometry(core.NewQRect4(690, 510, 75, 23))
	this.FrameBuildings = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.FrameBuildings.SetObjectName("FrameBuildings")
	this.FrameBuildings.SetGeometry(core.NewQRect4(19, 19, 361, 521))
	this.FrameBuildings.SetFrameShape(widgets.QFrame__StyledPanel)
	this.FrameBuildings.SetFrameShadow(widgets.QFrame__Raised)
	this.LoadFileName = widgets.NewQLineEdit(this.FrameBuildings)
	this.LoadFileName.SetObjectName("LoadFileName")
	this.LoadFileName.SetGeometry(core.NewQRect4(100, 40, 241, 20))
	this.TextBrowser = widgets.NewQTextBrowser(this.FrameBuildings)
	this.TextBrowser.SetObjectName("TextBrowser")
	this.TextBrowser.SetGeometry(core.NewQRect4(40, 80, 301, 141))
	this.FileOpenButton = widgets.NewQPushButton(this.FrameBuildings)
	this.FileOpenButton.SetObjectName("FileOpenButton")
	this.FileOpenButton.SetGeometry(core.NewQRect4(150, 230, 75, 23))
	this.LabelTitle3 = widgets.NewQLabel(this.FrameBuildings, core.Qt__Widget)
	this.LabelTitle3.SetObjectName("LabelTitle3")
	this.LabelTitle3.SetGeometry(core.NewQRect4(10, 260, 50, 12))
	this.FrameTerrain = widgets.NewQFrame(this.Centralwidget, core.Qt__Widget)
	this.FrameTerrain.SetObjectName("FrameTerrain")
	this.FrameTerrain.SetGeometry(core.NewQRect4(420, 20, 361, 521))
	this.FrameTerrain.SetFrameShape(widgets.QFrame__StyledPanel)
	this.FrameTerrain.SetFrameShadow(widgets.QFrame__Raised)
	this.LoadFileName2 = widgets.NewQLineEdit(this.FrameTerrain)
	this.LoadFileName2.SetObjectName("LoadFileName2")
	this.LoadFileName2.SetGeometry(core.NewQRect4(100, 40, 241, 20))
	this.FileOpenButton2 = widgets.NewQPushButton(this.FrameTerrain)
	this.FileOpenButton2.SetObjectName("FileOpenButton2")
	this.FileOpenButton2.SetGeometry(core.NewQRect4(150, 230, 75, 23))
	this.TextBrowser2 = widgets.NewQTextBrowser(this.FrameTerrain)
	this.TextBrowser2.SetObjectName("TextBrowser2")
	this.TextBrowser2.SetGeometry(core.NewQRect4(40, 80, 301, 141))
	this.FrameTerrain.Raise()
	this.FrameBuildings.Raise()
	this.LabelTitle.Raise()
	this.LabelTitle2.Raise()
	this.LabelFileName.Raise()
	this.LabelFileName2.Raise()
	this.BuildStartButton1.Raise()
	this.BuildStartButton2.Raise()
	MainWindow.SetCentralWidget(this.Centralwidget)
	this.Menubar = widgets.NewQMenuBar(MainWindow)
	this.Menubar.SetObjectName("Menubar")
	this.Menubar.SetGeometry(core.NewQRect4(0, 0, 800, 21))
	this.MenuFile = widgets.NewQMenu(this.Menubar)
	this.MenuFile.SetObjectName("MenuFile")
	MainWindow.SetMenuBar(this.Menubar)
	this.Statusbar = widgets.NewQStatusBar(MainWindow)
	this.Statusbar.SetObjectName("Statusbar")
	MainWindow.SetStatusBar(this.Statusbar)
	this.Menubar.QWidget.AddAction(this.MenuFile.MenuAction())


    this.RetranslateUi(MainWindow)

}

func (this *UISteraGoWindowMainWindow) RetranslateUi(MainWindow *widgets.QMainWindow) {
    _translate := core.QCoreApplication_Translate
	MainWindow.SetWindowTitle(_translate("MainWindow", "MainWindow", "", -1))
	this.LabelTitle.SetText(_translate("MainWindow", "Buildings", "", -1))
	this.LabelTitle2.SetText(_translate("MainWindow", "Terrain", "", -1))
	this.LabelFileName.SetText(_translate("MainWindow", "File name", "", -1))
	this.LabelFileName2.SetText(_translate("MainWindow", "File name", "", -1))
	this.BuildStartButton1.SetText(_translate("MainWindow", "Build start", "", -1))
	this.BuildStartButton2.SetText(_translate("MainWindow", "Build start", "", -1))
	this.FileOpenButton.SetText(_translate("MainWindow", "PushButton", "", -1))
	this.LabelTitle3.SetText(_translate("MainWindow", "Build", "", -1))
	this.FileOpenButton2.SetText(_translate("MainWindow", "PushButton", "", -1))
	this.MenuFile.SetTitle(_translate("MainWindow", "File", "", -1))
}
