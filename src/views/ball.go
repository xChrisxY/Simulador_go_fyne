package views

import (
	"ball/src/models"
	"ball/src/scenes"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Ball struct {
	Ball *canvas.Image
}

func NewBall() *Ball{
	return &Ball{Ball: nil}
}

func (b *Ball) AddBall(c scenes.Scene) {
	pelota := canvas.NewImageFromURI(storage.NewFileURI("./assets/pelota.png"))
	pelota.Resize(fyne.NewSize(30,30))
	pelota.Move(fyne.NewPos(200,100))
	b.Ball = pelota
	c.AddImage(pelota)
}

// Update define lo que el observador hará cuando reciba una notificación
func (b *Ball) Update(pos models.Pos) {
	fmt.Printf("%d : %d\n", pos.X, pos.Y)
	b.Ball.Move(fyne.NewPos(float32(pos.X),float32(pos.Y)))
}