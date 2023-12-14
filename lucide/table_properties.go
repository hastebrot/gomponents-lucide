package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TableProperties(children ...g.Node) g.Node {
	svg := `<path d="M15 3v18" /><rect width="18" height="18" x="3" y="3" rx="2" /><path d="M21 9H3" /><path d="M21 15H3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
