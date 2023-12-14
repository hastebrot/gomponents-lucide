package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ShieldQuestion(children ...g.Node) g.Node {
	svg := `<path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10" /><path d="M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3" /><path d="M12 17h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
