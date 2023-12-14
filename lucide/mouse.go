package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Mouse(children ...g.Node) g.Node {
	svg := `<rect x="5" y="2" width="14" height="20" rx="7" /><path d="M12 6v4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
