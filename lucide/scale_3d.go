package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Scale3d(children ...g.Node) g.Node {
	svg := `<circle cx="19" cy="19" r="2" /><circle cx="5" cy="5" r="2" /><path d="M5 7v12h12" /><path d="m5 19 6-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
