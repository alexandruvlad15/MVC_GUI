package View3

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func CreateInputForm(onSubmit func(rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx float64)) fyne.CanvasObject {
	// Creăm câmpuri de input pentru fiecare parametru
	rhoEntry := widget.NewEntry()
	rhoEntry.SetPlaceHolder("rho")

	GEntry := widget.NewEntry()
	GEntry.SetPlaceHolder("G")

	SEntry := widget.NewEntry()
	SEntry.SetPlaceHolder("S")

	PEntry := widget.NewEntry()
	PEntry.SetPlaceHolder("P")

	Cx0Entry := widget.NewEntry()
	Cx0Entry.SetPlaceHolder("Cx0")

	etaEntry := widget.NewEntry()
	etaEntry.SetPlaceHolder("eta")

	kEntry := widget.NewEntry()
	kEntry.SetPlaceHolder("k")

	CzEntry := widget.NewEntry()
	CzEntry.SetPlaceHolder("Cz")

	CzmaxEntry := widget.NewEntry()
	CzmaxEntry.SetPlaceHolder("Czmax")

	CxEntry := widget.NewEntry()
	CxEntry.SetPlaceHolder("Cx")

	// Buton pentru a trimite datele
	submitButton := widget.NewButton("Calculează", func() {
		// Citirea și conversia datelor din string în float64
		rho, _ := strconv.ParseFloat(rhoEntry.Text, 64)
		G, _ := strconv.ParseFloat(GEntry.Text, 64)
		S, _ := strconv.ParseFloat(SEntry.Text, 64)
		P, _ := strconv.ParseFloat(PEntry.Text, 64)
		Cx0, _ := strconv.ParseFloat(Cx0Entry.Text, 64)
		eta, _ := strconv.ParseFloat(etaEntry.Text, 64)
		k, _ := strconv.ParseFloat(kEntry.Text, 64)
		Cz, _ := strconv.ParseFloat(CzEntry.Text, 64)
		Czmax, _ := strconv.ParseFloat(CzmaxEntry.Text, 64)
		Cx, _ := strconv.ParseFloat(CxEntry.Text, 64)

		// Apelăm funcția de calcul
		onSubmit(rho, G, S, P, Cx0, eta, k, Cz, Czmax, Cx)
	})

	// Aranjarea câmpurilor într-un layout
	form := container.NewVBox(
		rhoEntry, GEntry, SEntry, PEntry, Cx0Entry,
		etaEntry, kEntry, CzEntry, CzmaxEntry, CxEntry,
		submitButton,
	)

	return form
}

func DisplayResults(window fyne.Window, v, gamma, Vmax, Gammamax float64) {
	// Creăm etichete pentru a afișa rezultatele
	resultLabel := widget.NewLabel(
		"Rezultate:\n" +
			"v = " + strconv.FormatFloat(v, 'f', 2, 64) + "\n" +
			"gamma = " + strconv.FormatFloat(gamma, 'f', 2, 64) + "\n" +
			"Vmax = " + strconv.FormatFloat(Vmax, 'f', 2, 64) + "\n" +
			"Gammamax = " + strconv.FormatFloat(Gammamax, 'f', 2, 64))

	// Afișăm rezultatele într-o nouă fereastră
	window.SetContent(container.NewVBox(resultLabel))
}
