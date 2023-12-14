package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Lock(children ...g.Node) g.Node {
	svg := `<rect width="18" height="11" x="3" y="11" rx="2" ry="2" /><path d="M7 11V7a5 5 0 0 1 10 0v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
