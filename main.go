package main

import (
	"context"
	"fmt"
	"log"

	"github.com/anejm/naloga-5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50,
				Usage: "Najnižja dovoljena temperatura",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "Najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Najmanjše število meritev za izračun statistike",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			vreme.NastaviTemp(cmd.Float("minTemp"), cmd.Float("maxTemp"), cmd.Int("stMeritev"))

			vreme.DodajPostajo("MB01", "Maribor Tabor", "Maribor")
			vreme.DodajPostajo("KP01", "Koper Luka", "Koper")
			vreme.DodajPostajo("KR01", "Kredarica", "Triglav")
			vreme.DodajPostajo("LJ01", "Ljubljana Bežigrad", "Ljubljana")

			vreme.DodajMeritev("MB01", 12.4)

			vreme.DodajMeritev("KP01", 24.5)
			vreme.DodajMeritev("KP01", 27.0)
			vreme.DodajMeritev("KP01", 26.1)
			vreme.DodajMeritev("KP01", 28.4)

			vreme.DodajMeritev("KR01", 2.0)
			vreme.DodajMeritev("KR01", -1.5)
			vreme.DodajMeritev("KR01", 3.5)

			vreme.DodajMeritev("LJ01", 18.2)
			vreme.DodajMeritev("LJ01", 21.5)
			vreme.DodajMeritev("LJ01", 19.8)
			vreme.DodajMeritev("LJ01", 16.5)

			vreme.IzpisMeritev()
			fmt.Println()

			vreme.DodajMeritev("FDGZS", 50.0)
			vreme.DodajMeritev("MB01", 70.0)
			vreme.DodajMeritev("MB01", 20.1)
			fmt.Println()

			vreme.IzpisMeritev()
			fmt.Println()

			vreme.IzpisMeritev()
			fmt.Println()

			vreme.IzpisPorocila()
			fmt.Println()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
}
