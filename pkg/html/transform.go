package html

import (
	"fmt"
	"os"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
)

type Asset struct {
	Props  attrs.Props
	Parent string
	Path   string
	Insert bool
}

func (a *Asset) Content() (string, error) {
	file, err := os.ReadFile(a.Path)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func TransformToScriptElements(assets *[]Asset, hash string) []elem.Node {
	return elem.TransformEach(*assets, func(s Asset) elem.Node {
		if s.Insert {
			c, _ := s.Content()
			return elem.Script(s.Props, elem.Text(c))
		}
		props := attrs.Props{
			attrs.Src: fmt.Sprintf("%s?r=%s", s.Path, hash),
		}
		return elem.Script(attrs.Merge(props, s.Props))
	})
}

func TransformToLinkElements(assets *[]Asset, hash string) []elem.Node {
	return elem.TransformEach(*assets, func(s Asset) elem.Node {
		props := attrs.Props{
			attrs.Href: fmt.Sprintf("%s?r=%s", s.Path, hash),
			attrs.Rel:  "stylesheet",
		}
		return elem.Link(attrs.Merge(props, s.Props))
	})
}

func FilterAssetsForParent(assets []Asset, parent string) (result []Asset) {
	for _, a := range assets {
		if a.Parent == parent {
			result = append(result, a)
		}
	}
	return
}
