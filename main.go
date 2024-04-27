package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
	// tgClient = telegram.New(token)
	// fatcher = fatcher.New(tgClient)
	// processor = processor.New(tgClient)
	// consumer.Start(processor, fatcher)
}

func mustToken() string {
	token := flag.String(
		"token-tg-bot",
		"",
		"token for acces to TG bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")

	}
	return *token

}
