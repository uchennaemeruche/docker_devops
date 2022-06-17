
package main

import (
	"fmt"
	"github.com/ichtrojan/horus"
	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

var logger, _ = thoth.Init("log")

type Error struct{
	Message string
}


type BankJSON struct{
	Name string	`json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	USSD string `json:"ussd"`

}

type Bank struct{
	Name string	`json:"name"`
	Slug string `json:"slug"`
	Code string `json:"code"`
	USSD string `json:"ussd"`
	Logo string `json:"logo"`
}

func main(){

	if err := godotenv.Load(); err != nil{
		logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")
	if !exist{
		logger.Log(errors.New("PORT not set in .env"))
		log.Fatal("PORT not set in .env")
	}

	host, exist := os.LookupEnv("HOST")
	if !exist {
		logger.Log(errors.New("HOST not set in .env"))
		log.Fatal("HOST not set in .env")
	}

	bankJson, err := ioutil.ReadFile("./banks.json")
	if err != nil {
		logger.Log(err)
	}

	horusDbUser, exist := os.
}
