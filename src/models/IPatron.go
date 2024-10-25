package models

type Pos struct{
	X int32
	Y int32
}

// Observer es la interfaz que define el comportamiento de los observadores
type Observer interface {
	Update(pos Pos)
}

// Subject es la interfaz que define el comportamiento del sujeto observado
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}