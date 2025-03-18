package main

import (
	socialize "TikTok-rpc/rpc/socialize/kitex_gen/socialize/socializeservice"
	"log"
)

func main() {
	svr := socialize.NewServer(new(SocializeServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
