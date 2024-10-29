package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mgred/html-filer/pkg/cli"
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

func RunApp(opts cli.Options) (err error) {
	if opts.Help {
		fmt.Fprint(os.Stdout, GetHelpMessage())
		os.Exit(0)
	}

	if opts.Version {
		fmt.Fprint(os.Stdout, GetVersion())
		os.Exit(0)
	}

	output, err := GetOutput(opts.Out)
	if err != nil {
		return fmt.Errorf("ERROR: could not open output file `%s`! [%s]", opts.Out, err.Error())
	}
	defer func() {
		e := output.Close()
		if e != nil {
			err = e
		}
	}()

	hash := utils.GenerateHash()
	var head bytes.Buffer
	var body bytes.Buffer

	for _, asset := range opts.Assets {
		var content string
		switch asset.Type {
		case html.Script:
			content, err = html.RenderScript(&asset, hash)
		case html.Style:
			content, err = html.RenderStyle(&asset, hash)
		}

		if err != nil {
			return fmt.Errorf("ERROR: Cannot render [%s]", err.Error())
		}

		if asset.Parent == html.HEAD_TAG {
			head.WriteString(content)
		} else {
			body.WriteString(content)
		}
	}

	if err = html.RenderDefaultHtml(output, html.DefaultData{
		Base:  opts.Base,
		Title: opts.Title,
		Body:  body.String(),
		Head:  head.String(),
	}); err != nil {
		return fmt.Errorf("ERROR: Could not write to Output! [%s]", err.Error())
	}

	return
}
