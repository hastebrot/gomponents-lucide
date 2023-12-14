package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Scaling(children ...g.Node) g.Node {
	svg := `<path d="M21 3 9 15" /><path d="M12 3H3v18h18v-9" /><path d="M16 3h5v5" /><path d="M14 15H9v-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
