package main

import (
	"ball/src/scenes"
	"ball/src/views"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type ParkingLot struct {
	availableSlots   []int
	vehicleQueue     chan *views.Vehicle
	mu               sync.Mutex
	isAccessOccupied bool // Estado de la entrada
}

func main() {
	myApp := app.New()
	stage := myApp.NewWindow("App - Parking Lot")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(815, 515))
	stage.SetFixedSize(true)

	scene := scenes.NewScene(stage)
	scene.Init()

	parkingLot := &ParkingLot{
		availableSlots: createAvailableSlots(20),
		vehicleQueue:   make(chan *views.Vehicle, 100),
	}

	go parkingLot.manageVehicles(scene)

	stage.ShowAndRun()
}

func createAvailableSlots(totalSlots int) []int {
	slots := make([]int, totalSlots)
	for i := 0; i < totalSlots; i++ {
		slots[i] = i
	}
	return slots
}

func (p *ParkingLot) manageVehicles(scene *scenes.Scene) {
	for {
		p.mu.Lock()
		if len(p.availableSlots) > 0 {
			slotIndex := p.availableSlots[0]
			p.availableSlots = p.availableSlots[1:] // Reservar el slot

			vehicle := views.NewVehicleView(slotIndex)
			vehicle.AddVehicle(scene)
			scene.UpdateParkingSlot(slotIndex, true)

			p.mu.Unlock()

			// Simular el tiempo que el vehículo está estacionado
			go p.vehicleExit(slotIndex, vehicle, scene)
		} else {
			vehicle := views.NewVehicleView(-1)
			p.vehicleQueue <- vehicle
			fmt.Println("Vehículo agregado a la cola, esperando espacio")
			p.mu.Unlock()
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (p *ParkingLot) vehicleExit(slotIndex int, vehicle *views.Vehicle, scene *scenes.Scene) {
	// Simular el tiempo que el vehículo está estacionado
	time.Sleep(time.Duration(3+rand.Intn(3)) * time.Second)
	fmt.Println("[]Saliendo...")

	p.mu.Lock()
	scene.RemoveWidget(vehicle.Image)
	scene.UpdateParkingSlot(slotIndex, false)

	// Liberar el espacio
	fmt.Println("Liberando espacio:", slotIndex)
	p.availableSlots = append(p.availableSlots, slotIndex)

	// Comprobar si hay vehículos en la cola para entrar
	if len(p.vehicleQueue) > 0 {
		nextVehicle := <-p.vehicleQueue
		nextSlot := slotIndex // Asignar el mismo slot que se liberó
		nextVehicle.AddVehicle(scene)
		scene.UpdateParkingSlot(nextSlot, true)
		fmt.Println("Vehículo desbloqueado y asignado a espacio:", nextSlot)
	}

	p.mu.Unlock()
}
