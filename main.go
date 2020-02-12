package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/gogf/gf/text/gstr"
)

/**
golagn版本极速扫描网段内可ping ip小工具，主要特征：轻量，简洁，极速。
*/
func main() {
	a := app.NewWithID("io.dw.ipscan")
	a.SetIcon(theme.FyneLogo())
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("goLang IpScan")
	w.Resize(fyne.Size{Width: 550, Height: 300})
	w.SetFixedSize(true)

	infProgress := widget.NewProgressBarInfinite()

	before := widget.NewEntry()
	before.SetPlaceHolder("127.0.0.")
	start := widget.NewEntry()
	start.SetPlaceHolder("1")
	stop := widget.NewEntry()
	stop.SetPlaceHolder("20")
	ipscan := widget.NewMultiLineEntry()
	form := &widget.Form{
		OnCancel: func() {
			fmt.Println("Cancelled")
		},
		OnSubmit: func() {
			b := before.Text
			last := gstr.SubStr(b, len(b)-1, len(b))
			if len(b) > 0 && last == "." {
				infProgress.Show()
				before.Disable()
				start.Disable()
				stop.Disable()

				ips, info := pingAll(before.Text, start.Text, stop.Text)
				ipscan.SetText(info + ips)
				infProgress.Hide()

				before.Enable()
				start.Enable()
				stop.Enable()
			} else {
				err := errors.New("ip field Illegal format!")
				dialog.ShowError(err, w)
			}

		},
	}
	form.Append("ip", before)
	form.Append("start", start)
	form.Append("stop", stop)
	form.Append("can ask", ipscan)
	form.Append("progrss", infProgress)
	infProgress.Hide()
	w.SetMaster()
	w.SetContent(form)
	w.ShowAndRun()
}
