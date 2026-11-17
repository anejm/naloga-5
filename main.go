package main

import (
	"github.com/anejm/naloga-5/vreme",
)

func main() {
	postaje := map[string]Postaja{
		"MB01": {"Maribor Tabor", "Maribor", []float64{12.4}},
		"KP01": {"Koper Luka", "Koper", []float64{24.5, 27, 26.1, 28.4}},
		"KR01": {"Kredarica", "Triglav", []float64{2, -1.5, 3.5}},
		"LJ01": {"Ljubljana Bežigrad", "Ljubljana", []float64{18.2, 21.5, 19.8, 16.5}},
	}

	izpisMeritev(postaje)
	fmt.Println()

	dodajMeritev(postaje, "FDGZS", 50.0)
	dodajMeritev(postaje, "MB01", 70.0)
	dodajMeritev(postaje, "MB01", 20.1)
	fmt.Println()

	izpisMeritev(postaje)
	fmt.Println()

	low1, high1, avg1, ok1 := statistika(postaje, "MB01")
	fmt.Printf("min: %.1f max: %.1f avg: %.1f bool: %t\n", low1, high1, avg1, ok1)
	low2, high2, avg2, ok2 := statistika(postaje, "KP01")
	fmt.Printf("min: %.1f max: %.1f avg: %.1f bool: %t\n", low2, high2, avg2, ok2)
	low3, high3, avg3, ok3 := statistika(postaje, "fewszh")
	fmt.Printf("min: %.1f max: %.1f avg: %.1f bool: %t\n", low3, high3, avg3, ok3)
	fmt.Println()

	izpisMeritev(postaje)
	fmt.Println()

	izpisPorocila(postaje)
	fmt.Println()
}
