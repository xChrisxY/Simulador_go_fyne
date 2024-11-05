package models

import (
	"ball/src/scenes"
	"ball/src/views"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ParkingLot struct {
	AvailableSlots   []int
	VehicleQueue     chan *views.Vehicle
	Mu               sync.Mutex
	IsAccessOccupied bool
}

func (p *ParkingLot) ManageVehicles(scene *scenes.Scene) {
	for {
		p.Mu.Lock()
		if len(p.AvailableSlots) > 0 {
			slotIndex := p.AvailableSlots[0]
			p.AvailableSlots = p.AvailableSlots[1:]

			vehicle := views.NewVehicleView(slotIndex)
			vehicle.AddVehicle(scene)
			scene.UpdateParkingSlot(slotIndex, true)

			p.Mu.Unlock()

			go p.VehicleExit(slotIndex, vehicle, scene)
		} else {
			vehicle := views.NewVehicleView(-1)
			p.VehicleQueue <- vehicle
			fmt.Println("Vehículo agregado a la cola, esperando espacio")
			p.Mu.Unlock()
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func (p *ParkingLot) VehicleExit(slotIndex int, vehicle *views.Vehicle, scene *scenes.Scene) {

	time.Sleep(time.Duration(3+rand.Intn(3)) * time.Second)
	fmt.Println("[!] Saliendo...")

	p.Mu.Lock()
	scene.RemoveWidget(vehicle.Image)
	scene.UpdateParkingSlot(slotIndex, false)

	fmt.Println("Liberando espacio:", slotIndex)
	p.AvailableSlots = append(p.AvailableSlots, slotIndex)

	if len(p.VehicleQueue) > 0 {
		nextVehicle := <-p.VehicleQueue
		nextSlot := slotIndex
		nextVehicle.AddVehicle(scene)
		scene.UpdateParkingSlot(nextSlot, true)
		fmt.Println("Vehículo desbloqueado y asignado a espacio:", nextSlot)
	}

	p.Mu.Unlock()

}
