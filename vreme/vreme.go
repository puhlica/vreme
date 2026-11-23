// Package vreme omogoča beleženje temperaturnih meritev vremenskih postaj in izračun statistike
package vreme

import "fmt"

// Postaja predstavlja vremensko postajo z imenom, krajem in meritvami
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}


// DodajMeritev postaji z oznako doda meritev, če je temperatura med minTemp in maxTemp
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura, minTemp, maxTemp float64) {
	if temperatura <= maxTemp && temperatura >= minTemp {
		p, ok := postaje[oznaka]
		if !ok {
			fmt.Println("Postaja ne obstaja")
		} else {
			p.meritve = append(p.meritve, temperatura)
			postaje[oznaka] = p
		}
	} else {
		fmt.Println("Temperatura izven dovoljenega območja")
	}
}

func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {
	p, ok := postaje[oznaka]
	if !ok {
		return
	}
	l := len(p.meritve)
	if l < stMeritev {
		ok = false
		return
	}
	najnizja = p.meritve[0]
	najvisja = p.meritve[0]
	sum := 0.0
	for _, t := range p.meritve {
		if t < najnizja {
			najnizja = t
		}
		if t > najvisja {
			najvisja = t
		}
		sum += t
	}
	povprecna = sum / float64(l)
	return
}

// IzpisMeritev izpiše meritve postaj
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")
	for oznaka, p := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

// IzpisPorocila izpise porocilo postaj za vse ki imajo vsaj stMeritev meritev, za ostale izpise da je meritev premalo
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {
	opis := ""
	for oznaka, p := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(postaje, oznaka, stMeritev)
		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}
		switch {
		case povprecna >= 25:
			opis = "Vroče"
		case povprecna >= 15 && povprecna < 25:
			opis = "Toplo"
		case povprecna >= 5 && povprecna < 15:
			opis = "Hladno"
		default:
			opis = "Mrzlo"
		}

		fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", p.ime, najnizja, najvisja, povprecna, opis)
	}
}

// DodajPostajo doda postajo v slovar postaje doda novo postajo z oznako, imenom in krajem
func DodajPostajo(postaje map[string]Postaja, oznaka, ime, kraj string) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj}
}
