package main

import (
	"allegro"
	"runtime"
	"fmt"
)

func main() {
	fmt.Println("Starting")

	allegro.Init()
	allegro.InstallKeyboard()
	allegro.InstallMouse()

	disp := allegro.CreateDisplay(600, 400, allegro.WINDOWED)

	runtime.GOMAXPROCS(10)

	color = allegro.CreateColor(100, 0, 0, 255)
	color.Clear()
	allegro.Flip()

	fmt.Println("Created window")

	handleEvents(disp)

	fmt.Println("Ended")
}

func handleEvents(disp *allegro.Display) {
	sources := [...]*allegro.EventSource{
		disp.GetEventSource(),
		allegro.GetKeyboardEventSource(),
		allegro.GetMouseEventSource()}
	ch := allegro.GetEvents(sources[:])
	for ev := range ch {
		switch ev := ev.(type) {
		case allegro.DisplayResizeEvent:
			disp.AcknowledgeResize()
		case allegro.MouseButtonDown:
			DisplayColor()
		case allegro.DisplayCloseEvent:
			return
		case allegro.KeyDownEvent:
			if ev.Keycode == allegro.KEY_ESCAPE {
				return
			}
		}
	}
}

var color allegro.Color

func DisplayColor() {
	r, g, b, _ := color.GetRGBA()
	color = allegro.CreateColor(r + 5, g + 5, b + 5, 255)
	color.Clear()
	allegro.Flip()
}