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

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "treinalinux.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ips na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
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

	fmt.Println(" IPs")
	for _, ip := range ips {
		fmt.Println(" ", ip)

	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(" NameServers")
	for _, servidor := range servidores {
		fmt.Println(" ", servidor.Host)
	}

}
