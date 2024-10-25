package models

import (
	"time"
)

type Ball struct {
	posX, posY int32
	status bool	
	observers []Observer
}

func NewBall() *Ball {
	return &Ball{posX: 0, posY: 0, status: true}
}

func (b *Ball) Run() {
	var incX int32 = 30
	var incY int32 = 30
	b.status = true
	b.posX = 60
	b.posY = 60
	for b.status {
		if b.posX < 30 || b.posX >770 {
			incX *= -1	
		}
		if b.posY < 30 || b.posY > 470 {
			incY *= -1
		}
		b.posX += incX
		b.posY += incY
		b.NotifyAll()
		time.Sleep(500 * time.Millisecond)
	}	
}

func (b *Ball) SetStatus(status bool) {
	b.status = status
}

// Register añade un observador a la lista
func (b *Ball) Register(observer Observer) {
	b.observers = append(b.observers, observer)
}

// Unregister elimina un observador de la lista
func (b *Ball) Unregister(observer Observer) {
	for i, o := range b.observers {
		if o == observer {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

// NotifyAll notifica a todos los observadores sobre una actualización
func (b *Ball) NotifyAll() {
	for _, observer := range b.observers {
		observer.Update(Pos{X:b.posX, Y:b.posY})
	}
}

