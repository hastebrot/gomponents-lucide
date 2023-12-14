package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Sticker(children ...g.Node) g.Node {
	svg := `<path d="M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z" /><path d="M15 3v6h6" /><path d="M10 16s.8 1 2 1c1.3 0 2-1 2-1" /><path d="M8 13h0" /><path d="M16 13h0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
