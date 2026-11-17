package vreme

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var (
	minTemp   float64 = -50
	maxTemp   float64 = 50
	stMeritev int     = 3
)

func NastaviTemp(min float64, max float64, st int) {
	minTemp = min
	maxTemp = max
	stMeritev = st
}

var postaje = make(map[string]Postaja)

func DodajPostajo(oznaka string, ime string, kraj string, meritve []float64) {
	postaje[oznaka] = Postaja{
		ime:     ime,
		kraj:    kraj,
		meritve: meritve,
	}
}

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

func IzpisMeritev() {
	fmt.Println("MERITVE:")
	for o, p := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", o, p.ime, p.kraj, p.meritve)
	}
}

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
