package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		cli.Command{
			Name: "server",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port,p",
					Usage: "port number to listen",
					Value: 8080,
				},
				cli.StringFlag{
					Name:   "user,u",
					Usage:  "user name for authentication",
					EnvVar: "AUTH_USER",
					Value:  "",
				},
				cli.StringFlag{
					Name:   "password,P",
					Usage:  "password for authentication",
					EnvVar: "AUTH_PASSWORD",
					Value:  "",
				},
			},
			Action: func(ctx *cli.Context) error {
				port := ctx.Int("port")
				user := ctx.String("user")
				password := ctx.String("password")
				e := echo.New()
				if user != "" && password != "" {
					log.Println("Authentication enabled")
					auth := middleware.BasicAuth(func(u string, p string, c echo.Context) (bool, error) {
						if u == user && p == password {
							return true, nil
						}
						return false, nil
					})
					e.Use(auth)
				} else {
					log.Println("Authentication not enabled")
				}
				e.POST("/clean", func(c echo.Context) error {
					err := run()
					result := echo.Map{}
					if err != nil {
						result["result"] = "fail"
						result["err"] = err
						return c.JSON(500, result)
					} else {
						result["result"] = "success"
						return c.JSON(200, result)
					}
				})
				return e.Start(fmt.Sprintf(":%d", port))
			},
		},
		cli.Command{
			Name: "run",
			Action: func(ctx *cli.Context) error {
				return run()
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	cmd := exec.Command("/bin/registry", "garbage-collect", "-m", "/etc/docker/registry/config.yml")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
