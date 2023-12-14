package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CheckCheck(children ...g.Node) g.Node {
	svg := `<path d="M18 6 7 17l-5-5" /><path d="m22 10-7.5 7.5L13 16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
