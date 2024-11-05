package views

import (
	"ball/src/scenes"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Vehicle struct {
	Image     *canvas.Image
	SlotIndex int // Índice del espacio de estacionamiento
}

func NewVehicleView(slotIndex int) *Vehicle {
	return &Vehicle{Image: nil, SlotIndex: slotIndex}
}

func (v *Vehicle) AddVehicle(c *scenes.Scene) {
	if v.SlotIndex < 0 || v.SlotIndex >= len(c.ParkingSlots()) {
		fmt.Println(v.SlotIndex)
		fmt.Println(c.ParkingSlots())
		fmt.Println("Índice de espacio de estacionamiento fuera de rango")
		return
	}

	entryPos := c.Entry().Position()
	slot := c.ParkingSlots()[v.SlotIndex]
	carImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/car3.png"))
	carImage.Resize(slot.Size())
	carImage.Move(entryPos)

	v.Image = carImage
	c.AddImage(carImage)

	go func() {
		stepCount := 50
		dx := (slot.Position().X - entryPos.X) / float32(stepCount)
		dy := (slot.Position().Y - entryPos.Y) / float32(stepCount)

		for i := 0; i < stepCount; i++ {
			time.Sleep(20 * time.Millisecond)
			newPos := fyne.NewPos(carImage.Position().X+dx, carImage.Position().Y+dy)
			carImage.Move(newPos)
		}

		time.Sleep(5 * time.Second)
		v.RemoveVehicle(c)
	}()
}

func (v *Vehicle) RemoveVehicle(c *scenes.Scene) {
	if v.Image == nil {
		return
	}

	entryPos := c.Entry().Position()

	go func() {
		stepCount := 50
		dx := (entryPos.X - v.Image.Position().X) / float32(stepCount) // Corrección aquí
		dy := (entryPos.Y - v.Image.Position().Y) / float32(stepCount) // Corrección aquí

		for i := 0; i < stepCount; i++ {
			time.Sleep(40 * time.Millisecond)
			newPos := fyne.NewPos(v.Image.Position().X+dx, v.Image.Position().Y+dy)
			v.Image.Move(newPos)
		}

		c.RemoveWidget(v.Image)
	}()
}
