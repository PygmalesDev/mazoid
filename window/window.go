package window

import (
	obj "github.com/PygmalesDev/mazoid/objects"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	objGr obj.ObjectGroup
}

func NewWindow() (w *Window) {
	w = &Window{
		objGr: *obj.NewObjectGroup(),
	}
	w.init()
	return
}

func (w *Window) init() {
	rl.InitWindow(900, 800, "mazoid")
	rl.SetTargetFPS(60)
}

func (w *Window) Draw() {
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		for _, obj := range w.objGr {
			obj.Draw()
		}

		rl.EndDrawing()
	}
}
