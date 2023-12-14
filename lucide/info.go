package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Info(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M12 16v-4" /><path d="M12 8h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
