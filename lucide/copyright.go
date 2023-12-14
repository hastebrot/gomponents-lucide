package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Copyright(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M14.83 14.83a4 4 0 1 1 0-5.66" />`
	return Icon(g.Group(children), g.Raw(svg))
}
