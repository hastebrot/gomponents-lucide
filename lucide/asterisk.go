package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Asterisk(children ...g.Node) g.Node {
	svg := `<path d="M12 6v12" /><path d="M17.196 9 6.804 15" /><path d="m6.804 9 10.392 6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
