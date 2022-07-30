package model

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"moell/internal/moc/util"
)

func Command(c *cli.Context) error {
	return GenModel(c.String("name"), c.String("dir"))
}

func GenModel(name string, dirName string) error {
	data := make(map[string]interface{})
	data["Name"] = name
	data["DirName"] = dirName

	return util.CodeTemplateConversion(&util.CodeGenParams{
		Name:     name,
		Template: Template,
		SavePath: fmt.Sprintf("./internal/%s/models", dirName),
		Data:     data,
	})
}
