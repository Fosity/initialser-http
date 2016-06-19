package main
import (
	"gopkg.in/urfave/cli.v2"
	"os"
)
const version = "0.0.1 beta"
func main() {
	app := cli.NewApp()
	app.Name = "initialser"
	app.Version = version
	app.Usage = ""
	app.Commands = []*cli.Command{
		cmdHttp,
	}
	app.Run(os.Args)
}