package html

import (
	"fmt"
	"os"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

type AssetType int

const (
	Script AssetType = iota
	Style
	Link
)

type Asset struct {
	Props  attrs.Props
	Parent string
	Path   string
	Insert bool
	Type   AssetType
}

func (a *Asset) Content() (string, error) {
	file, err := os.ReadFile(a.Path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func RenderScript(a *Asset, hash string) (string, error) {
	if a.Insert {
		c, err := a.Content()
		return elem.Script(a.Props, elem.Text(c)).Render(), err
	}
	props := attrs.Merge(attrs.Props{
		attrs.Src: hashedPath(a.Path, hash),
	}, a.Props)
	return elem.Script(props).Render(), nil
}

func RenderStyle(a *Asset, hash string) (string, error) {
	if a.Insert {
		c, err := a.Content()
		return elem.Style(a.Props, elem.Text(c)).Render(), err
	}
	a.Props[attrs.Rel] = "stylesheet"
	return RenderLink(a, hash)
}

func RenderLink(a *Asset, hash string) (string, error) {
	props := attrs.Merge(attrs.Props{
		attrs.Href: hashedPath(a.Path, hash),
	}, a.Props)
	return elem.Link(props).Render(), nil
}

func hashedPath(path string, hash string) string {
	return fmt.Sprintf("%s?r=%s", path, hash)
}
