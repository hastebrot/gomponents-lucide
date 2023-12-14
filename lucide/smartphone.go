package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Smartphone(children ...g.Node) g.Node {
	svg := `<rect width="14" height="20" x="5" y="2" rx="2" ry="2" /><path d="M12 18h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
