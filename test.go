package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.83.150:31000/"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}
	mapping, _ := client.GetMapping().Index("news").Do(context.Background())
	fmt.Println(mapping)
}
