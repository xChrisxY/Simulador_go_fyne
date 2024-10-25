package main

import (
	"ball/src/models"
	"ball/src/scenes"
	"ball/src/views"
	_ "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New();
	stage := myApp.NewWindow("App - Ball")
	stage.CenterOnScreen()
	stage.Resize(fyne.NewSize(815,515));
	stage.SetFixedSize(true);
	
	// Create scene
	scene := scenes.NewScene(stage)
	scene.Init()

	//Add a new widget
	button := widget.NewButton("Click", func() {
		// Creamos el objeto observado
		b1 := models.NewBall()
		// Add Balon (Observador)
		ball := views.NewBall()
		ball.AddBall(*scene)
		//Registramos a ball como observador de la goroutine b1
		b1.Register(ball)
		go b1.Run()
	})
	button.Move(fyne.NewPos(100,100))
	button.Resize(fyne.NewSize(100,50))
	scene.AddWidget(button)

	stage.ShowAndRun()
}