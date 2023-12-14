package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Video(children ...g.Node) g.Node {
	svg := `<path d="m22 8-6 4 6 4V8Z" /><rect width="14" height="12" x="2" y="6" rx="2" ry="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
