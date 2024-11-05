package scenes

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type Scene struct {
	scene        fyne.Window
	container    *fyne.Container
	parkingSlots []*canvas.Rectangle
	entry        *canvas.Rectangle // Entrada del estacionamiento
}

func NewScene(scene fyne.Window) *Scene {
	return &Scene{scene: scene, container: nil, parkingSlots: make([]*canvas.Rectangle, 20)}
}

func (s *Scene) Init() {

	rect := canvas.NewRectangle(color.RGBA{R: 50, G: 50, B: 50, A: 255})
	rect.Resize(fyne.NewSize(815, 515))
	rect.Move(fyne.NewPos(0, 0))

	s.container = container.NewWithoutLayout(rect)
	s.scene.SetContent(s.container)

	s.createParkingSlots()

	s.createEntry()

	s.container.Add(s.entry)
	s.container.Refresh()
}
func (s *Scene) createParkingSlots() {
	slotWidth := float32(90)
	slotHeight := float32(120)
	padding := float32(10)

	for i := 0; i < 20; i++ {
		x := float32(i%5)*(slotWidth+padding) + padding  // Colocación horizontal con margen
		y := float32(i/5)*(slotHeight+padding) + padding // Colocación vertical con margen

		slot := canvas.NewRectangle(color.RGBA{R: 70, G: 70, B: 70, A: 255})
		slot.Resize(fyne.NewSize(slotWidth, slotHeight))
		slot.Move(fyne.NewPos(x, y))

		if i%5 != 4 {
			line := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 0, A: 255}) // Línea amarilla
			line.Resize(fyne.NewSize(2, slotHeight))
			line.Move(fyne.NewPos(x+slotWidth, y))
			s.container.Add(line)
		}

		// Agregar un número de espacio para identificar cada plaza
		numberText := canvas.NewText(fmt.Sprintf("%d", i+1), color.White)
		numberText.TextSize = 14
		numberText.Move(fyne.NewPos(x+slotWidth/2-5, y+slotHeight/2-10))
		s.container.Add(numberText)

		s.parkingSlots[i] = slot
		s.container.Add(slot)
	}

	s.container.Refresh()
}

func (s *Scene) createEntry() {

	entryWidth := float32(220)
	entryHeight := float32(80)

	s.entry = canvas.NewRectangle(color.RGBA{R: 200, G: 200, B: 200, A: 255})

	// Mover la entrada a una posición más visible
	s.entry.Resize(fyne.NewSize(entryWidth, entryHeight))
	s.entry.Move(fyne.NewPos(550, 225))
}

// Método exportado para acceder a los espacios de estacionamiento
func (s *Scene) ParkingSlots() []*canvas.Rectangle {
	return s.parkingSlots
}

// Actualizar el estado del estacionamiento
func (s *Scene) UpdateParkingSlot(slotIndex int, occupied bool) {
	if slotIndex < 0 || slotIndex >= len(s.parkingSlots) {
		return // Index fuera de rango
	}

	if occupied {
		s.parkingSlots[slotIndex].FillColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	} else {
		s.parkingSlots[slotIndex].FillColor = color.RGBA{R: 200, G: 200, B: 200, A: 255}
	}
	s.parkingSlots[slotIndex].Refresh()
}

func (s *Scene) AddWidget(widget fyne.Widget) {
	s.container.Add(widget)
	s.container.Refresh()
}

func (s *Scene) AddImage(image *canvas.Image) {
	s.container.Add(image)
	s.container.Refresh()
}

func (s *Scene) RemoveWidget(widget fyne.CanvasObject) {
	s.container.Remove(widget)
	s.container.Refresh()
}

func (s *Scene) Entry() *canvas.Rectangle {
	return s.entry
}
