package service

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"moell/internal/moc/util"
	"strings"
)

func Command(c *cli.Context) error {
	return GenService(c.String("name"), c.String("dir"))
}

func GenService(name string, dirName string) error {
	data := make(map[string]interface{})
	data["Name"] = name
	data["AbbrName"] = strings.ToLower(name[0:1])
	data["DirName"] = dirName

	return util.CodeTemplateConversion(&util.CodeGenParams{
		Name:     name,
		Template: Template,
		SavePath: fmt.Sprintf("./internal/%s/services", dirName),
		Data:     data,
	})
}
