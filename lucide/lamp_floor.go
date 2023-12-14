package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func LampFloor(children ...g.Node) g.Node {
	svg := `<path d="M9 2h6l3 7H6l3-7Z" /><path d="M12 9v13" /><path d="M9 22h6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
