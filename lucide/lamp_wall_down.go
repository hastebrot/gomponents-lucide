package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func LampWallDown(children ...g.Node) g.Node {
	svg := `<path d="M11 13h6l3 7H8l3-7Z" /><path d="M14 13V8a2 2 0 0 0-2-2H8" /><path d="M4 9h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4v6Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
