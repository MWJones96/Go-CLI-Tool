package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website Lookup"
	app.Usage = "Query IP, CNAME, MX, Name Servers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "mwj96.dev",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Look up Name Server",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}

				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up host IP address",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
