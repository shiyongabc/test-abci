package main

import (
	"github.com/shiyongabc/test-abci/server"
)

type cliArgs struct {
	//cli.Helper
	//ListenAddress      string `cli:"*l,*listen" usage:"listen host and port" dft:"$API_HOST_LS"`

}

func main() {
	app := server.NewApplication()
	server.New(app).Start("0.0.0.0:81")
	//cli.Run(new(cliArgs), func(ctx *cli.Context) error {
	//	//argv := ctx.Argv().(*cliArgs)
	//	app := server.NewApplication()
	//	server.New(app).Start("0.0.0.0:81")
	//
	//	return nil
	//})



}
