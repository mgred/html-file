package main

import "fmt"

const HELP_MESSAGE = `
SYNOPSIS:
  %[1]s [--base <base>] [--out <output>] [--title <title>] <asset>...
  %[1]s [options] --assets <asset>...
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
  -a, --assets                 General option for left over assets
  -s, --scripts                Add scripts
      -m, --module             Set type "module"
      -a, --async              Set "async" attribute
      -h, --head               Set to head of document
      -i, --insert             Copy content of the file into the tag
  -S, --styles                 Add stylesheets
      -m, --media              Set media attribute
      -i, --insert             Copy content of the file into the tag

`

func GetHelpMessage() string {
	return fmt.Sprintf(HELP_MESSAGE, "html-file")
}
