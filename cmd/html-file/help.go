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
      Add the following assets as stylesheets.

      The given assets will web <link>ed unless the --insert options is set.
      In this case the content of the asset will be copied into a <style> tag.

  %[1]s --favicon [options] <asset>
      Add the following asset as favicon.

OPTIONS:

  -h, --help                   Print this Help message
  -v, --version                Print version
  -o, --out                    Output file to write to
  -b, --base                   Base path to set, default "/"
  -t, --title                  Set content of <title> element
  -A, --assets                 General option for left over assets
  -s, --scripts                Add scripts
      -a, --async              Set "async" attribute
      -h, --head               Set to head of document
      -i, --insert             Copy content of the file into the tag
      -m, --module             Set type "module"
      -p, --preload            Mark this script for preload
  -S, --styles                 Add stylesheets
      -i, --insert             Copy content of the file into the tag
      -m, --media              Set media attribute
      -p, --preload            Mark this style for preload
  -f, favicon                  Link a favicon
      -z, --sizes              Set the sizes width by length, e.g 25x25

`

func GetHelpMessage() string {
	return fmt.Sprintf(HELP_MESSAGE, "html-file")
}
