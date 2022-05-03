package main

import (
	"github.com/robfig/cron"
	"log"
	"main/servise"
	"math/rand"
	"time"
)

var sendStr = []string{
	"我顶",
	"顶一顶",
	"随手顶一下",
	"顶",
	"捞一下",
	"我再捞",
	"继续顶",
	"顶就完事了",
	"坚持顶",
	"捞一捞",
	"越来越多人啦",
}

func RandomSend() {
	rand.Seed(time.Now().UnixNano())
	content := sendStr[rand.Intn(len(sendStr))]
	servise.Send(content)
}

func main() {
	log.SetFlags(log.Lshortfile)
	RandomSend()
	c := cron.New()
	_ = c.AddFunc("30 44 */4 * * *", RandomSend)
	c.Start()
	select {}
}
