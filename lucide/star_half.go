package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func StarHalf(children ...g.Node) g.Node {
	svg := `<path d="M12 17.8 5.8 21 7 14.1 2 9.3l7-1L12 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
