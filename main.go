package main

import (
	"ball/src/scenes"
	"ball/src/views"

	"ball/src/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	stage := myApp.NewWindow("App - Parking Lot")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(815, 515))
	stage.SetFixedSize(true)

	scene := scenes.NewScene(stage)
	scene.Init()

	parkingLot := models.ParkingLot{
		AvailableSlots: createAvailableSlots(20),
		VehicleQueue:   make(chan *views.Vehicle, 100),
	}

	go parkingLot.ManageVehicles(scene)

	stage.ShowAndRun()
}

func createAvailableSlots(totalSlots int) []int {
	slots := make([]int, totalSlots)
	for i := 0; i < totalSlots; i++ {
		slots[i] = i
	}
	return slots
}
