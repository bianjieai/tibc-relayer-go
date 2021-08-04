package main

import (
	"fmt"
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	"github.com/jasonlvhit/gocron"
)

func main() {
	client := repostitory.NewTendermintClient("iris")
	err := gocron.Every(5).Second().From(gocron.NextTick()).Do(task, client)
	if err != nil {
		panic(err)
	}
	<-gocron.Start()
}
func task(client *repostitory.TendermintClient) {
	fmt.Println(client.GetLatestHeight())
}
