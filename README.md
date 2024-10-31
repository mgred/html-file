# html-file

`html-file` is a small tool that can create HTML files and reference or include
asset files.
It's inspired by [`html-insert-assets`](https://github.com/jbedard/html-insert-assets).

## Build

```sh
go build ./cmd/html-file
```

## Usage

```sh
html-file foo.mjs bar.js baz.css
```

Prints the following to `stdout`:

```html
<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title></title>
    <base href="">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href="baz.css?r=DRljqvWP" rel="stylesheet">
  </head>
  <body>
      <script src="foo.mjs?r=DRljqvWP" type="module"></script><script src="bar.js?r=DRljqvWP"></script>
  </body>
</html>
```

See the `--help` for more options.

## Acknowledgement

- [`html-insert-assets`](https://github.com/jbedard/html-insert-assets)
- [`elem-go`](https://github.com/chasefleming/elem-go)
- [`go-random`](https://github.com/mazen160/go-random)

## License

Source code in this repository is licensed under the [GNU GPL v3](https://www.gnu.org/licenses/gpl-3.0.en.html#license-text)
