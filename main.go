package main

import (
	"MVC_GUI/Controller3"
	"MVC_GUI/View3"

	"fyne.io/fyne/v2/app"
)

func main() {
	application := app.New()
	window := application.NewWindow("MVC GUI Example")

	// Creăm interfața pentru citirea datelor
	inputForm := View3.CreateInputForm(func(rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx float64) {
		// Când datele sunt introduse, calculăm soluțiile
		v, gamma, Vmax, Gammamax := Controller3.DetermineSolutions(rho, S, P, G, Cx0, k, eta, Cz, Cx, Czmax)
		// Afișăm rezultatele
		View3.DisplayResults(window, v, gamma, Vmax, Gammamax)
	})

	window.SetContent(inputForm)
	window.ShowAndRun()
}
