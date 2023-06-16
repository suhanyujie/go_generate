package main

import (
	"go_generate/internal/interf"
	"go_generate/internal/service"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: GenPro,
		Commands: []*cli.Command{
			{
				Name:    "gen",
				Aliases: []string{"m"},
				Usage:   "项目生成",
				Action:  GenPro,
			},
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "configPath",
			Usage:       "指定配置文件",
			Aliases:     []string{"c"},
			Required:    false,
			TakesFile:   false,
			Value:       "./config",
			Destination: nil,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// GenPro 生成项目 todo
func GenPro(*cli.Context) error {
	srcPath := ""
	dstPath := ""
	generator := service.NewGoProGen(srcPath, dstPath)
	toolObj := generator.(interf.ToolIf)
	toolObj.Prompting()
	toolObj.Writing()
	toolObj.End()
	return nil
}
