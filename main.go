package main

import (
	"context"
	"log"
	"os"

	"github.com/puhlica/vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "beleženje meritev vremenskih postaj",
		Flags: []cli.Flag{
			&cli.FloatFlag{Name: "minTemp", Usage: "najnižja dovoljena temperatura", Value: -50},
			&cli.FloatFlag{Name: "maxTemp", Usage: "najvišja dovoljena temperatura", Value: 50},
			&cli.IntFlag{Name: "stMeritev", Usage: "najmanjše število meritev za statistiko", Value: 3},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			postaje := make(map[string]vreme.Postaja)
			vreme.DodajPostajo(postaje, "VOGL", "Vogel", "Bohinj")
			vreme.DodajPostajo(postaje, "KRED", "Kredarica", "Kranjska Gora")
			vreme.DodajPostajo(postaje, "VRSC", "Vršič", "Bovec")
			vreme.DodajPostajo(postaje, "RUDN", "Rudno polje", "Bohinj")

			vreme.DodajMeritev(postaje, "VOGL", 6, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "VOGL", 7, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KRED", 7, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KRED", 8, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KRED", 9, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "VRSC", 1.4, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "VRSC", 1.7, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "VRSC", 0.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "RUDN", 1.5, minTemp, maxTemp)

			vreme.DodajMeritev(postaje, "VOGL", 80, minTemp, maxTemp) // neveljavna meritev
			vreme.DodajMeritev(postaje, "XXXX", 10, minTemp, maxTemp) // neveljavna postaja

			vreme.IzpisMeritev(postaje)
			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
