package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Sheet(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" ry="2" /><line x1="3" x2="21" y1="9" y2="9" /><line x1="3" x2="21" y1="15" y2="15" /><line x1="9" x2="9" y1="9" y2="21" /><line x1="15" x2="15" y1="9" y2="21" />`
	return Icon(g.Group(children), g.Raw(svg))
}
