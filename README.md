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
