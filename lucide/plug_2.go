package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Plug2(children ...g.Node) g.Node {
	svg := `<path d="M9 2v6" /><path d="M15 2v6" /><path d="M12 17v5" /><path d="M5 8h14" /><path d="M6 11V8h12v3a6 6 0 1 1-12 0v0Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
