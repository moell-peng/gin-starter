package repository

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"moell/internal/moc/util"
	"strings"
)

func Command(c *cli.Context) error {
	if err := GenRepository(c.String("name"), c.String("dir")); err != nil {
		return err
	}
	return nil
}

func GenRepository(name string, dirName string) error {

	data := make(map[string]interface{})
	data["Name"] = name
	data["AbbrName"] = strings.ToLower(name[0:1])
	data["DirName"] = dirName

	return util.CodeTemplateConversion(&util.CodeGenParams{
		Name:     name,
		Template: Template,
		SavePath: fmt.Sprintf("./internal/%s/repositories", dirName),
		Data:     data,
	})
}
