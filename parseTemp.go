package utils

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"
)

func sub(left, right int) int {
	return left - right
}

func addSuffix(host string, suffix string) string {
	if !strings.HasSuffix(host, suffix) {
		host += suffix
	}
	return host
}

// ParseTemp 解析templates
func ParseTemp(filePath string, data interface{}) (b *bytes.Buffer, err error) {
	var (
		t *template.Template
	)

	tempFuncMap := template.FuncMap{
		"sub":    sub,
		"suffix": addSuffix,
	}

	t = template.New(path.Base(filePath)).Funcs(tempFuncMap)
	if t, err = t.ParseFiles(filePath); err != nil {
		err = fmt.Errorf("ERROR: Parse Template File [%s] Error. %v", filePath, err)
		return
	}

	b = new(bytes.Buffer)
	if err = t.Execute(b, data); err != nil {
		err = fmt.Errorf("ERROR: Execute Template [%s] Error. %v", filePath, err)
		return
	}
	return
}
