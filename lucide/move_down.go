package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveDown(children ...g.Node) g.Node {
	svg := `<path d="M8 18L12 22L16 18" /><path d="M12 2V22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
