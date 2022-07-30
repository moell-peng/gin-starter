package api

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"moell/internal/moc/util"
	"strings"
)

func Command(c *cli.Context) error {
	name := c.String("name")
	msrDir := c.String("msr-dir")
	appDir := c.String("dir")

	if err := GenApi(name, appDir, msrDir); err != nil {
		return err
	}

	return nil
}

func GenApi(name string, appDir string, msrDir string) error {
	data := make(map[string]interface{})
	data["Name"] = name
	data["AbbrName"] = strings.ToLower(name[0:1])
	data["AppDir"] = appDir
	data["MSRDir"] = msrDir

	if data["AppDir"] != "app" && data["MSRDir"] == "app" {
		data["MSRDir"] = data["AppDir"]
	}

	_ = util.CodeTemplateConversion(&util.CodeGenParams{
		Name:     name,
		Template: RequestTemplate,
		SavePath: fmt.Sprintf("./internal/%s/requests", data["AppDir"]),
		Data:     data,
	})

	return util.CodeTemplateConversion(&util.CodeGenParams{
		Name:     name,
		Template: Template,
		SavePath: fmt.Sprintf("./internal/%s/api", data["AppDir"]),
		Data:     data,
	})
}
