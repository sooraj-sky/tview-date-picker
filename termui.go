package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	header := widgets.NewParagraph()
	header.Text = "Press q to quit, Press h or l to switch tabs"
	header.SetRect(0, 0, 50, 1)
	header.Border = false
	header.TextStyle.Bg = ui.ColorBlue


	monthsList := widgets.NewList()
	monthsList.Title = "Select A Month To Search"
	monthsList.Rows = []string{
		"[0] January",
		"[1] February",
		"[2] March",
		"[3] April",
		"[4] May",
		"[5] June",
		"[6] July",
		"[7] August",
		"[8] September",
		"[9] October",
		"[10] November",
		"[11] December",
	}
	monthsList.TextStyle = ui.NewStyle(ui.ColorYellow)
	monthsList.WrapText = false
	monthsList.SetRect(5, 5, 35, 20)


	tabpane := widgets.NewTabPane("2018", "2019", "2020", "2021", "2022")
	tabpane.SetRect(0, 1, 50, 4)
	tabpane.Border = true

	renderTab := func() {
		switch tabpane.ActiveTabIndex {
		case 0:
			ui.Render(monthsList)
		case 1:
			ui.Render(monthsList)
		case 2:
			ui.Render(monthsList)
		case 3:
			ui.Render(monthsList)
		case 4:
			ui.Render(monthsList)
		case 5:
			ui.Render(monthsList)
		}
	}

	ui.Render(header, tabpane, monthsList)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(header, tabpane)
			renderTab()
		}
	}
}