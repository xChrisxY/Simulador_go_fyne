package models

import (
	"math/rand"
	"time"
)

type Vehicle struct {
	posX, posY int32
	status     bool
	observers  []Observer
}

func NewVehicle() *Vehicle {
	return &Vehicle{posX: 0, posY: 0, status: true}
}

func (v *Vehicle) Run() {

	v.status = true
	v.posX = rand.Int31n(700) + 50
	v.posY = rand.Int31n(400) + 50
	v.NotifyAll()

	parkDuration := time.Duration(rand.Intn(3)+3) * time.Second
	time.Sleep(parkDuration)

	v.status = false
	v.NotifyAll()
}

func (v *Vehicle) SetStatus(status bool) {
	v.status = status
}

func (v *Vehicle) Register(observer Observer) {
	v.observers = append(v.observers, observer)
}

func (v *Vehicle) Unregister(observer Observer) {
	for i, o := range v.observers {
		if o == observer {
			v.observers = append(v.observers[:i], v.observers[i+1:]...)
			break
		}
	}
}

func (v *Vehicle) NotifyAll() {
	for _, observer := range v.observers {
		observer.Update(Pos{X: v.posX, Y: v.posY})
	}
}
