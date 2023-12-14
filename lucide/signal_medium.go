package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func SignalMedium(children ...g.Node) g.Node {
	svg := `<path d="M2 20h.01" /><path d="M7 20v-4" /><path d="M12 20v-8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
