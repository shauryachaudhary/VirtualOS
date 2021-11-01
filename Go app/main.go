package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
var myApp fyne.App = app.New()

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget


var img fyne.CanvasObject;
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {
	var myWindow fyne.Window = myApp.NewWindow("Virtual OS")
    myApp.Settings().SetTheme(theme.LightTheme())
    img = canvas.NewImageFromFile("H:\\Go app\\background.jpg") 
    
    btn1 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func(){
        showCalculator(myWindow)   
    })
    btn2 = widget.NewButtonWithIcon("BMI Calculator", theme.InfoIcon(), func(){
        showbmi(myWindow)   
    })
    btn3 = widget.NewButtonWithIcon("Audio Player", theme.RadioButtonIcon(), func(){
        showplayer(myWindow )   
    })
    
    DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(),func(){
        myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
    })
    panelContent = container.NewVBox((container.NewGridWithColumns(4,DeskBtn,btn1,btn2,btn3)))
    myWindow.Resize(fyne.NewSize(1280,720))
    myWindow.CenterOnScreen()
    
    myWindow.SetContent(
        container.NewBorder(panelContent,nil,nil,nil,img),
    )

    myWindow.ShowAndRun()

    
}