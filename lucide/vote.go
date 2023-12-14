package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Vote(children ...g.Node) g.Node {
	svg := `<path d="m9 12 2 2 4-4" /><path d="M5 7c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v12H5V7Z" /><path d="M22 19H2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
