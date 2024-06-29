//go:build windows
// +build windows

package main

import (
	"github.com/getlantern/systray"
	"github.com/uif/uifd/icon"
	"github.com/uif/uifd/uif"
)

func TrayInit() {
	systray.Run(onReady, onExit)
}

var SystrayCore *systray.MenuItem
var coreOp *systray.MenuItem

func UpdateSysTrayCore() {
	if uif.GetCoreStatus() == 0 {
		SystrayCore.Check()
		SystrayCore.SetTitle("内核运行中")
		coreOp.SetTitle("关闭内核")
	} else {
		SystrayCore.Uncheck()
		SystrayCore.SetTitle("内核未运行")
		coreOp.SetTitle("运行内核")
	}
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("UIF")
		systray.SetTooltip("UI for freedom")

		SystrayCore = systray.AddMenuItemCheckbox("-", "Core Status", false)
		SystrayCore.Disable()

		coreOp = systray.AddMenuItem("关闭内核", "Close Core")
		systray.AddSeparator()

		mUrl := systray.AddMenuItem("打开面板", "Open my home")

		systray.AddSeparator()
		mQuit := systray.AddMenuItem("退出 UIF", "Quit the whole app")

		for {
			UpdateSysTrayCore()
			select {
			case <-SystrayCore.ClickedCh:
				if SystrayCore.Checked() {
					SystrayCore.Uncheck()
					SystrayCore.SetTitle("内核未运行")
				} else {
					SystrayCore.Check()
					SystrayCore.SetTitle("内核运行中")
				}
			case <-mQuit.ClickedCh:
				uif.CloseCore()
				systray.Quit()
			case <-coreOp.ClickedCh:
				if uif.GetCoreStatus() == 0 {
					uif.CloseCore()
				} else {
					uif.RunCore()
				}
			case <-mUrl.ClickedCh:
				OpenWeb()
			}
		}
	}()
}

func onExit() {
	// clean up here
}
