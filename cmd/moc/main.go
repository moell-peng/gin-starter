package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"moell/internal/moc/curd"
	"moell/internal/moc/handler"
	"moell/internal/moc/model"
	"moell/internal/moc/repository"
	"moell/internal/moc/service"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "moc",
		Usage: "moell cli application",
		Commands: []*cli.Command{
			{
				Name:  "curd",
				Usage: "生成CURD代码",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "生成CURD的名字选项 ",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "dir",
						Usage:       "App 目录， 如填写 --dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
					&cli.StringFlag{
						Name:        "msr-dir",
						Usage:       "Model Repository Service 目录， 如填写 --msr-dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
				},
				Action: curd.Command,
			},
			{
				Name:  "api",
				Usage: "生成API代码",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "生成API的名字选项 ",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "dir",
						Usage:       "生成API文件的目录， 如填写 --dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
					&cli.StringFlag{
						Name:        "msr-dir",
						Usage:       "Model Repository Service 目录， 如填写 --msr-dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
				},
				Action: handler.Command,
			},
			{
				Name:  "model",
				Usage: "生成Model代码",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "生成Model的名字选项 ",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "dir",
						Usage:       "生成Model文件的目录， 如填写 --dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
				},
				Action: model.Command,
			},
			{
				Name:  "repo",
				Usage: "生成Repository代码",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "生成Repository代码的名字选项 ",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "dir",
						Usage:       "生成Repository文件的目录， 如填写 --dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
				},
				Action: repository.Command,
			},
			{
				Name:  "service",
				Usage: "生成Service代码",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "生成Service代码的名字选项 ",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "dir",
						Usage:       "生成Service文件的目录， 如填写 --dir=app ,生成的代码路径为./internal/app/* ",
						Value:       "app",
						DefaultText: "app",
					},
				},
				Action: service.Command,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
