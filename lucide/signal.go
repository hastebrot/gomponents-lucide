package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Signal(children ...g.Node) g.Node {
	svg := `<path d="M2 20h.01" /><path d="M7 20v-4" /><path d="M12 20v-8" /><path d="M17 20V8" /><path d="M22 4v16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
