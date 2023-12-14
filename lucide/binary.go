package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Binary(children ...g.Node) g.Node {
	svg := `<rect x="14" y="14" width="4" height="6" rx="2" /><rect x="6" y="4" width="4" height="6" rx="2" /><path d="M6 20h4" /><path d="M14 10h4" /><path d="M6 14h2v6" /><path d="M14 4h2v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
