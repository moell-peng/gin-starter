package curd

import (
	"github.com/urfave/cli/v2"
	"moell/internal/moc/api"
	"moell/internal/moc/model"
	"moell/internal/moc/repository"
	"moell/internal/moc/service"
)

func Command(c *cli.Context) error {

	name := c.String("name")
	msrDir := c.String("msr-dir")
	appDir := c.String("dir")
	if appDir != "app" && msrDir == "app" {
		msrDir = appDir
	}

	if err := api.GenApi(name, msrDir, appDir); err != nil {
		return err
	}

	if err := repository.GenRepository(name, msrDir); err != nil {
		return err
	}

	if err := model.GenModel(name, msrDir); err != nil {
		return err
	}

	if err := service.GenService(name, msrDir); err != nil {
		return err
	}

	return nil
}
