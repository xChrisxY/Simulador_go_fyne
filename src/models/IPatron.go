package models

type Pos struct {
	X int32
	Y int32
}

type Observer interface {
	Update(pos Pos)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}
