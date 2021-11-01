package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalculator(w fyne.Window) {
	var historyArr []string;
	ishistory := false
	historystr := ""
	history := widget.NewLabel(historystr)
	output := ""
	input := widget.NewLabel(output)
	a := app.New()
	c := a.NewWindow("Calculator")
	c.Resize(fyne.NewSize(268,360))
    historybtn := widget.NewButton("History", func(){
		if ishistory{
             historystr = ""
		}else{
			for i:=len(historyArr) -1; i>=0;i--{
				historystr = historystr+historyArr[i];
				historystr = "\n";
			}
		}
		ishistory = !ishistory;
        history.SetText(historystr);
	})
	backbtn := widget.NewButton("Back", func(){
		if len(output) > 0{
		output :=output[:len(output)-1]
		input.SetText(output)}

	})
	clearbtn := widget.NewButton("Clear", func(){
        output = "";
		input.SetText(output);

	})

	openbtn := widget.NewButton("(", func(){
        output = output + "(";
		input.SetText(output);
	})
	closebtn := widget.NewButton(")", func(){
		output = output + ")";
		input.SetText(output);
	})
	dividebtn := widget.NewButton("/", func(){
		output = output + "/";
		input.SetText(output);
	})
	ninebtn := widget.NewButton("9", func(){
		output = output + "9";
		input.SetText(output);
	})
	eightbtn := widget.NewButton("8", func(){
		output = output + "8";
		input.SetText(output);
	})
	sevenbtn := widget.NewButton("7", func(){
		output = output + "7";
		input.SetText(output);
	})
	multiplybtn := widget.NewButton("*", func(){
		output = output + "*";
		input.SetText(output);
	})
	

    fourbtn := widget.NewButton("4", func(){
		output = output + "4";
		input.SetText(output);
	})
	fivebtn := widget.NewButton("5", func(){
		output = output + "5";
		input.SetText(output);
	})
	sixbtn := widget.NewButton("6", func(){
		output = output + "6";
		input.SetText(output);
	})
	
	minusbtn := widget.NewButton("-", func(){
		output = output + "-";
		input.SetText(output);
	})
	
	onebtn := widget.NewButton("1", func(){
		output = output + "1";
		input.SetText(output);
	})
	
	twobtn := widget.NewButton("2", func(){
		output = output + "2";
		input.SetText(output);
	})
	
	threebtn := widget.NewButton("3", func(){
		output = output + "3";
		input.SetText(output);
	})
	
	plusbtn := widget.NewButton("+", func(){
		output = output + "+";
		input.SetText(output);
	})
	
	zerobtn := widget.NewButton("0", func(){
		output = output + "0";
		input.SetText(output);
	})
	dotbtn := widget.NewButton(".", func(){
		output = output + ".";
		input.SetText(output);
	})
	
	equalsbtn := widget.NewButton("=", func(){
        expression, err := govaluate.NewEvaluableExpression(output);
		if err == nil{
            result, err := expression.Evaluate(nil);
	        if err ==nil{
				answer := strconv.FormatFloat(result.(float64), 'f', -1, 64);
			    strToAppend := output + " = " + answer;
				historyArr = append(historyArr, strToAppend)
				output = answer;   
			}else{
				output = "Error";
			}
		}else{
			output = "Error";			
		}
	
		input.SetText(output)

		
		
	})
	



    	


	
	c.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
		container.NewGridWithColumns(2,historybtn,
		backbtn,),
		container.NewGridWithColumns(4,
			clearbtn,
			openbtn,
			closebtn,
			dividebtn),
	    container.NewGridWithColumns(4,
		ninebtn,
		eightbtn,
		sevenbtn,
		multiplybtn),
		container.NewGridWithColumns(4,
        fourbtn,fivebtn,sixbtn,minusbtn),
		container.NewGridWithColumns(4,
		onebtn,twobtn,threebtn,plusbtn,),
        container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				zerobtn,
				dotbtn,
	    	),
	    equalsbtn,),
		),
	),
	)	
	c.ShowAndRun()
}