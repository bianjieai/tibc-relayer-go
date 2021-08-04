package main

import (
	"fmt"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/jasonlvhit/gocron"
)

func main() {
	client := repostitory.NewTendermintClient("iris")
	gocron.Every(5).Second().Do(task, client)
	<-gocron.Start()
}
func task(client *repostitory.TendermintClient) {
	fmt.Println(client.GetLatestHeight())
}
