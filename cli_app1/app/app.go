package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar : vai retornar app cli
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App cli"
	app.Usage = "Busca IPs e nomes DNS"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca Ips na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "treinalinux.com",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)

	}

}
