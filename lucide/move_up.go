package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoveUp(children ...g.Node) g.Node {
	svg := `<path d="M8 6L12 2L16 6" /><path d="M12 2V22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
