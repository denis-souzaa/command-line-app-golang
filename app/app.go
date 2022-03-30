package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Get Name Server"
	app.Usage = "Get name or IP server by host name"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IPs servers on internet",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "domain",
			Usage:  "Get domains servers on internet",
			Flags:  flags,
			Action: getDomainName,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getDomainName(c *cli.Context) {
	host := c.String("host")
	domains, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, domain := range domains {
		fmt.Println(domain.Host)
	}
}
