package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Grab(children ...g.Node) g.Node {
	svg := `<path d="M18 11.5V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1.4" /><path d="M14 10V8a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v2" /><path d="M10 9.9V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5" /><path d="M6 14v0a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0" /><path d="M18 11v0a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8 2 2 0 1 1 4 0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
