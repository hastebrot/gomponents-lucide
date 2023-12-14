package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArrowUpFromLine(children ...g.Node) g.Node {
	svg := `<path d="m18 9-6-6-6 6" /><path d="M12 3v14" /><path d="M5 21h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
