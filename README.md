# gomponents-lucide

[![godoc reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/hastebrot/gomponents-lucide) [![mit license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/hastebrot/gomponents-lucide/master/LICENSE) [![isc license](http://img.shields.io/badge/license-ISC-red.svg?style=flat)](https://raw.githubusercontent.com/hastebrot/gomponents-lucide/master/LICENSE)

[Lucide](https://lucide.dev/) for [gomponents](https://www.gomponents.com/).

## Usage

Add the dependency to the package.

```shell
go get github.com/hastebrot/gomponents-lucide
```

Example code. Pick the icons from https://lucide.dev/icons/.

```go
package main

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"

	"github.com/hastebrot/gomponents-lucide/lucide"
)

func Example() g.Node {
	return Section(
		Div(
			lucide.Pizza(
				g.Attr("width", "24px"),
				g.Attr("height", "24px"),
				g.Attr("stroke", "rgb(239,68,68)"),
				g.Attr("stroke-width", "1.5"),
			),
		),
		Div(
			lucide.Pizza(StyleAttr("width: 24px; height: 24px; color: rgb(239,68,68);")),
		),
		Div(
			Class("w-[24px] h-[24px] text-red-500"),
			lucide.Pizza(Class("[stroke-width:2.5]")),
		),
	)
}
```

## Build instructions

Remove existing `icons/` and `lucide/` directories before build.

```shell
❯ curl -sSL https://github.com/lucide-icons/lucide/releases/download/0.295.0/lucide-icons-0.295.0.zip | tar -xf - -C .

❯ ls -1 icons/*.svg | wc -l
1346

❯ deno run --allow-read --allow-write build.deno.ts

❯ ls -1 lucide/*.go | wc -l
1346
```
