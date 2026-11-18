// Package vreme provides the functionality to manage weather stations and
//
// The package implements a map of weather stations to monitor each created station. Each station has
// a name, place and an array of measurments. The stations are added to the map using oznaka.
// The package allows users to add measurements to each station.
// The user can set the minimum and maximum temperature that the sensors can track through flags, you
// can also set the minimum number of measurements for the station to be include in the statistical report.
// There are also functions that allow the user to track the monitored stations.
//
// Example usage:
//	vreme.DodajPostajo("MB01", "Maribor Tabor", "Maribor")
//
//	vreme.DodajMeritev("MB01", 12.4)
//	vreme.DodajMeritev("MB01", 15.2)
//  vreme.DodajMeritev("MB01", 21.2)
//
//	vreme.IzpisMeritev()
//	vreme.IzpisPorocila()

package vreme

import "fmt"

// Postaja represents a single weather station
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// minTemp is the minimum temperature that a sensor can read
// maxTemp is the maximum temperature that a sensor can read
// stMeritev is the minimum number of required measurments to run statistika function
var (
	minTemp   float64 = -50
	maxTemp   float64 = 50
	stMeritev int     = 3
)

// NastaviTemp sets the variables minTemp, maxTemp and stMeritev
func NastaviTemp(min float64, max float64, st int) {
	minTemp = min
	maxTemp = max
	stMeritev = st
}

// postaje is a map of postaje
var postaje = make(map[string]Postaja)

// DodajPostajo adds a Postaja to map postaje
func DodajPostajo(oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: []float64{},
	}
}

// DodajMeritev adds a measurment to the array of measurments in a particular Postaja
func DodajMeritev(oznaka string, temperatura float64) {
	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("temperatura izven meja")
		return
	}

	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		fmt.Println("postaja ne obstaja")
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// statistika returns the minimum and maximum measurments and also the avergae of all measurments
// the function also checks if there are enough measurments using stMeritev, it stores the result in boolean ok
func statistika(oznaka string) (najnizja, najvisja, povprecna float64, ok bool) {
	ok = true
	postaja, obstaja := postaje[oznaka]
	if !obstaja {
		ok = false
		return
	}

	if len(postaja.meritve) < stMeritev {
		ok = false
		return
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]

	var sum float64
	for _, temp := range postaja.meritve {
		sum += temp
		if temp < najnizja {
			najnizja = temp
		}
		if temp > najvisja {
			najvisja = temp
		}
	}

	povprecna = sum / float64(len(postaja.meritve))

	return
}

// IzpisMeritev prints out all the stations in map postaje
func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for o, p := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", o, p.ime, p.kraj, p.meritve)
	}
}

// izpisPorocila prints out the stats from statistika for each station that has enough measurements
func IzpisPorocila() {
	for oznaka, p := range postaje {
		najnizja, najvisja, povprecna, ok := statistika(oznaka)
		fmt.Printf("%s: ", p.ime)
		if !ok {
			fmt.Println("premalo meritev")
			continue
		}

		switch {
		case povprecna >= 25.0:
			fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C -> Vroče\n", najnizja, najvisja, povprecna)

		case povprecna >= 15.0:
			fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C -> Toplo\n", najnizja, najvisja, povprecna)

		case povprecna >= 5.0:
			fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C -> Hladno\n", najnizja, najvisja, povprecna)

		default:
			fmt.Printf("min %.1f °C, max %.1f °C, povprečje %.1f °C -> Mrzlo\n", najnizja, najvisja, povprecna)
		}
	}
}
