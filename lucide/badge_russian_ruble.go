package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BadgeRussianRuble(children ...g.Node) g.Node {
	svg := `<path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z" /><path d="M9 16h5" /><path d="M9 12h5a2 2 0 1 0 0-4h-3v9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
