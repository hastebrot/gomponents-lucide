package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Pointer(children ...g.Node) g.Node {
	svg := `<path d="M22 14a8 8 0 0 1-8 8" /><path d="M18 11v-1a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0" /><path d="M14 10V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1" /><path d="M10 9.5V4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v10" /><path d="M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15" />`
	return Icon(g.Group(children), g.Raw(svg))
}
