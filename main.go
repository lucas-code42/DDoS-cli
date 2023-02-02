package main

import (
	"fmt"
	"net/http"
)

func recoverConnection() {
	if r := recover(); r != nil {
		fmt.Println("Recover OK")
	}
}

func verifyError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func main() {

	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		go workerRequest()
		go workerRequest()
		go workerRequest()
		go workerRequest()
	}

	fmt.Println("fim")

}

func workerRequest() {
	url := "https://lunyscosmeticos.com.br/"
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	verifyError(err, "1***")

	req.Header.Add("cache-control", "max-age=0")

	res, err := client.Do(req)
	verifyError(err, "2***")
	if res.StatusCode != 200 {
		verifyError(fmt.Errorf("erro dif de 200"), "erro dif de 200")
	}

	defer recoverConnection()
}
