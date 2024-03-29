package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpLeftFromCircle(children ...g.Node) g.Node {
	svg := `<path d="M2 8V2h6" /><path d="m2 2 10 10" /><path d="M12 2A10 10 0 1 1 2 12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
