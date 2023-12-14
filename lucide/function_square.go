package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FunctionSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><path d="M9 17c2 0 2.8-1 2.8-2.8V10c0-2 1-3.3 3.2-3" /><path d="M9 11.2h5.7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
