package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func AtSign(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="4" /><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
