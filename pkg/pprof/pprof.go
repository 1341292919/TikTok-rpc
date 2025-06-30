package pprof

import (
	"TikTok-rpc/pkg/utils"
	"log"
	"net/http"
	"os"
	"runtime"

	_ "net/http/pprof"
)

func Load(service string) {
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
	addr, err := utils.GetPortForPprof(service)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
}
