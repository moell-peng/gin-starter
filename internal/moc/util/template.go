package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"moell/pkg/utils/helper"
	"os"
	"text/template"
)

type CodeGenParams struct {
	Name     string
	Template string
	SavePath string
	Data     interface{}
}

func CodeTemplateConversion(c *CodeGenParams) error {
	if err := os.MkdirAll(c.SavePath, os.ModePerm); err != nil {
		return err
	}
	
	buf := new(bytes.Buffer)
	temp, _ := template.New(c.Name).Parse(c.Template)

	if err := temp.Execute(buf, c.Data); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s.go", c.SavePath, helper.UnderscoreName(c.Name))
	if err := ioutil.WriteFile(filePath, buf.Bytes(), 0666); err != nil {
		return err
	}

	fmt.Printf("\"%s\" Generated successfully \n", filePath)

	return nil
}
