package main

import (
	"fmt"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false

func showplayer(w fyne.Window) {
	go func(msg string){
		fmt.Println(msg)

		if streamer == nil{

		}else{
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("going")

	time.Sleep(time.Second)

	a := app.New()
	c := a.NewWindow("Music player")
	c.Resize(fyne.NewSize(800,600))
	

	logo := canvas.NewImageFromFile("logo.png")
	logo.FillMode = canvas.ImageFillOriginal

	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			
			speaker.Init(format.SampleRate,format.SampleRate.N(time.Second/10))
		    speaker.Play(streamer)
		}),

		
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			}else if pause{
				pause = false 
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
		}),
		widget.NewToolbarSpacer(),
	)
		label := widget.NewLabel("Audio MP3")
		label.Alignment = fyne.TextAlignCenter

		label2 := widget.NewLabel("play MP3")
		label2.Alignment = fyne.TextAlignCenter

		browse_files := widget.NewButton("Browse", func(){
            fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error){
				streamer,format, _ = mp3.Decode(uc)
 				label2.Refresh()
			}, c)
			fd.Show()
			fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))

		})
		b := container.NewVBox(label,browse_files,label2,toolbar)
		c.SetContent(
			container.NewBorder(logo,nil,nil,nil,b),
		)
		c.ShowAndRun()

}
