package main

import (
	"toolapi/apiserver"
	"github.com/droundy/goopt"
	"toolapi/dao"
)

var optBundle = &dao.OptBundle{
	ReserveP2pCandidate: goopt.String([]string{"-p", "--p2p"}, "", "true"),
}

func main() {
	goopt.Parse(nil)

	apiServer := apiserver.InitGinServer(optBundle)
	apiServer.ListenAndServe()
}