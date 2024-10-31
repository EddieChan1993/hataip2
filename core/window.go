package core

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cast"
)

func Window() {
	myApp := app.New()
	myWindow := myApp.NewWindow("AIP")

	input := widget.NewEntry()
	input1 := widget.NewEntry()
	input2 := widget.NewEntry()
	input.SetPlaceHolder("Holding Cost")
	input1.SetPlaceHolder("Fund Value")
	input2.SetPlaceHolder("OriInvest Amount")

	label := widget.NewLabel("Calculate Result")

	containerGrid := container.New(layout.NewGridLayout(3), input, input1, input2)
	content := container.NewVBox(containerGrid, widget.NewButton("Calculate", func() {
		res := LogicCal(&LogicInput{
			HoldingCost:     cast.ToFloat64(input.Text),
			FundValue:       cast.ToFloat64(input1.Text),
			OriInvestAmount: cast.ToFloat64(input2.Text),
		})
		str1 := fmt.Sprintf("Rate(-%.0f%% ～ +%.0f%%)：%s%%\n", fundCaLine, fundCaLine, Float642str(res.Rate, 4))
		str3 := fmt.Sprintf("AmountRate %.4f%% ~ %.4f%%：%s%%\n", minRate, maxRate, Float642str(res.TotalRate, 0))
		str4 := fmt.Sprintf("Amount：%s \n", Float642str(res.Total, 0))
		label.SetText(str1 + str3 + str4)
	}), label)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(450, 50))
	myWindow.ShowAndRun()
}
