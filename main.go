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

    termWidth, _ := ui.TerminalDimensions()
	tabpane := widgets.NewTabPane("SPY", "MSFT", "AAPL")
	tabpane.SetRect(0, 0, termWidth, 4)
	tabpane.Border = true

	ui.Render(tabpane)

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h":
			tabpane.FocusLeft()
			ui.Clear()
			ui.Render(tabpane)
		case "l":
			tabpane.FocusRight()
			ui.Clear()
			ui.Render(tabpane)
		}
	}
}
