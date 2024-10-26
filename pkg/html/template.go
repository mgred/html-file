package html

import (
	_ "embed"
	"io"
	"text/template"
)

//go:embed default.html
var DefaultHtml string

type DefaultData struct {
	Base  string
	Title string
	Body  string
	Head  string
}

func RenderDefaultHtml(w io.Writer, data any) error {
	t, err := template.New("default").Parse(DefaultHtml)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

const (
	BODY_TAG = "body"
	HEAD_TAG = "head"
)
