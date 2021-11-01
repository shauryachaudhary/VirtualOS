package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
func showbmi(w fyne.Window) {
    a := app.New()
    c := a.NewWindow("BMI Calculator")
    
    c.Resize(fyne.NewSize(400, 400))
    label := canvas.NewText("", color.Black)
    label.Alignment = fyne.TextAlignCenter
    label.TextSize = 20
    
    result := canvas.NewText("", color.Black)
    result.Alignment = fyne.TextAlignCenter
    result.TextSize = 20
    
    inputH := widget.NewEntry()
	inputH.SetPlaceHolder("Enter height in cm")

    inputW := widget.NewEntry()
    inputW.SetPlaceHolder("Enter Weight in KG")
    btn1 := widget.NewButton("Calculate", func() {
       
		h, _ := strconv.ParseFloat(inputH.Text, 64)
        w, _ := strconv.ParseFloat(inputW.Text, 64)
		if h <= 0 || w <= 0{
			result.Text = "Error"
		}else{
        result.Text = calculateBMI(h/100, w)
	    }
	    result.Refresh()
    })
    
    c.SetContent(
        container.NewVBox(
            label,
            result,
            inputH,
            inputW,
            btn1,
        ))
    c.ShowAndRun()
}

func calculateBMI(height, weight float64) string {
	var BMI float64 = weight / math.Pow(height, 2) 
    if BMI <= 18.5 {
        fmt.Println("Underweight.")
        return "Underweight."
    } else if BMI <= 24.9 { 
        fmt.Println("Healthy.")
        return "Normal Weight"
    } else if BMI <= 29.9 { 
        fmt.Println("Overweight")
        return "Overweight"
    } else if BMI <= 34.9 { 
        fmt.Println("Obesity Class 1")
        return "Obesity Class 1"
    } else if BMI <= 39.9 { 
        fmt.Println("Obesity Class 2")
        return "Obesity Class 2"
    } else {  
		fmt.Println("Obesity Class 3")
        return "Obesity Class 3"
    }
}