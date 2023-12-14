package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CaseSensitive(children ...g.Node) g.Node {
	svg := `<path d="m3 15 4-8 4 8" /><path d="M4 13h6" /><circle cx="18" cy="12" r="3" /><path d="M21 9v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
