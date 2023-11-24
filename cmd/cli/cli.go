package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers!"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name: "host",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a Particular Host",
			Flags: myFlags,
			Action: func(cCtx *cli.Context) error {
				ns, err := net.LookupNS(cCtx.String("host"))
				if err != nil {
					return fmt.Errorf("не удалось выполнить запрос nserv для %s: %w", cCtx.String("host"), err)
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a Particular Host",
			Flags: myFlags,
			Action: func(cCtx *cli.Context) error {
				ip, err := net.LookupIP(cCtx.String("host"))
				if err != nil {
					return fmt.Errorf("не удалось выполнить запрос ip для %s: %w", cCtx.String("host"), err)
				}

				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}

				return nil
			},
		},
		{
			Name:  "cn",
			Usage: "Looks up the CNAME for a Particular Host",
			Flags: myFlags,
			Action: func(cCtx *cli.Context) error {
				cname, err := net.LookupCNAME(cCtx.String("host"))
				if err != nil {
					return fmt.Errorf("не удалось выполнить запрос cname для %s: %w", cCtx.String("host"), err)
				}

				fmt.Println(cname)

				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a Particular Host",
			Flags: myFlags,
			Action: func(cCtx *cli.Context) error {
				mx, err := net.LookupMX(cCtx.String("host"))
				if err != nil {
					return fmt.Errorf("не удалось выполнить запрос mx для %s: %w", cCtx.String("host"), err)
				}

				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}

				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
