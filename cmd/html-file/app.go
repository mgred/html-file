package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

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

func RunApp(opts Options) (err error) {
	if opts.Help {
		fmt.Fprintf(os.Stdout, `
SYNOPSIS:
  %[1]s [--base <base>] [--out <output>] [--title <title>] <asset>...
  %[1]s [options] --scripts [--module] [--async | --insert] [--head] <asset>...
  %[1]s [options] --styles [--media] [--insert] <asset>...

DESCRIPTION:

  %[1]s --scripts [options] <asset>...
      Add the following assets as scripts.

      By default the tags will be written to the bottom of the <body>.
      The --head option changes this behavior and writes everything to <head>.

  %[1]s --styles [options] <asset>...
    TBD

OPTIONS:

  -h, --help                   Print this Help message
  -v, --version                Print version
  -o, --out                    Output file to write to
  -b, --base                   Base path to set, default "/"
  -t, --title                  Set content of <title> element
  -s, --scripts                Add scripts
      -m, --module             Set type "module"
      -a, --async              Set "async" attribute
      -h, --head               Set to head of document
      -i, --insert             Copy content of the file into the tag
  -S, --styles                 Add stylesheets
      -m, --media              Set media attribute
      -i, --insert             Copy content of the file into the tag
`, "html-file")
		os.Exit(0)
	}

	if opts.Version {
		fmt.Fprint(os.Stdout, GetVersion())
		os.Exit(0)
	}

	output, err := GetOutput(opts.Out)
	if err != nil {
		return fmt.Errorf("ERROR: Could not open Output file `%s`! [%s]", opts.Out, err.Error())
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
