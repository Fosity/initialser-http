package main
import (
	"gopkg.in/urfave/cli.v2"
	"os"
	"github.com/leonlau/initialser-http/cmd"
)
const version = "0.0.4 beta"
func main() {
	app := &cli.App{}
	app.Name = "initialser"
	app.Version = version
	app.Usage = ""
	app.Commands = []*cli.Command{
		cmd.CmdHttp,
	}


	app.Run(os.Args)
}