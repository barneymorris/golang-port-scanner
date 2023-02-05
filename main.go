package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/urfave/cli/v2"
)

func main() {
	fmt.Printf("GOOS = %s\n", runtime.GOOS)

	isAppropriateOS := runtime.GOOS == "linux" || runtime.GOOS == "darwin"
	
	if !isAppropriateOS {
		err := fmt.Errorf("Support only linux or macOS. Exiting...\n")
		panic(err)
	}

	show := &cli.Command{
		Name: "show",
		Usage: "show all opened ports",
		Action: func(ctx *cli.Context) error {
			cmd := "sudo lsof -i -P -n | grep LISTEN"

			out, err := exec.Command("bash", "-c", cmd).Output()

			if err != nil {
				return err
			}

			fmt.Printf("%s\n", out)
			return nil
		},
	}

	kill := &cli.Command{
		Name: "kill",
		Usage: "kill some process on specific port",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "port",
				Usage: "some specific port",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			port, err := strconv.Atoi(ctx.String("port"))

			if err != nil {
				return err
			}

			cmd := fmt.Sprintf("sudo kill -9 $(sudo lsof -t -i:%s)", strconv.Itoa(port))
			
			_, err = exec.Command("bash", "-c", cmd).Output()
		
			if err != nil {
				return err
			}

			fmt.Printf("process has been killed\n")
			return nil
		},
	}

	app := &cli.App{
		Name: "port-scanner",
		Usage: "port utility cli program written on Golang",
		Flags: nil,
		Commands: []*cli.Command{show, kill},
    }

	err := app.Run(os.Args)

	if err != nil {
		fmt.Printf("failed to run cli app: %s\n", err.Error())
		os.Exit(0)
	}
}