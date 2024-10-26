package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chasefleming/elem-go"
	"github.com/mgred/html-filer/pkg/html"
	"github.com/mgred/html-filer/pkg/utils"
)

func GetOutput(path string) (*os.File, error) {
	if path != "" {
		if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
			return nil, err
		}
		return os.Create(path)
	}
	return os.Stdout, nil
}

func main() {
	opts, err := ProcessArgs()

	if err != nil {
		log.Fatal(err.Error())
	}

	if opts.Help {
		fmt.Fprintf(os.Stdout, "%s", "Help")
		os.Exit(0)
	}

	if opts.Version {
		fmt.Fprint(os.Stdout, GetVersion())
		os.Exit(0)
	}

	output, err := GetOutput(opts.Out)
	if err != nil {
		log.Fatalf("ERROR: Could not open Output file `%s`! [%s]", opts.Out, err.Error())
	}
	defer output.Close()

	hash := utils.GenerateHash()
	var head bytes.Buffer
	var body bytes.Buffer

	if len(opts.Styles) > 0 {
		elements := html.TransformToLinkElements(&opts.Styles, hash)
		head.WriteString(elem.Fragment(elements...).Render())
	}

	if len(opts.Scripts) > 0 {
		m := map[string]*bytes.Buffer{
			html.BODY_TAG: &body,
			html.HEAD_TAG: &head,
		}
		for tag, parent := range m {
			assets := html.FilterAssetsForParent(opts.Scripts, tag)
			elements := html.TransformToScriptElements(&assets, hash)
			parent.WriteString(elem.Fragment(elements...).Render())
		}
	}

	if err = html.RenderDefaultHtml(output, html.DefaultData{
		Base:  opts.Base,
		Title: opts.Title,
		Body:  body.String(),
		Head:  head.String(),
	}); err != nil {
		log.Fatalf("ERROR: Could not write to Output! [%s]", err.Error())
	}
}
