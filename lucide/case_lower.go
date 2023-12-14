package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CaseLower(children ...g.Node) g.Node {
	svg := `<circle cx="7" cy="12" r="3" /><path d="M10 9v6" /><circle cx="17" cy="12" r="3" /><path d="M14 7v8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
