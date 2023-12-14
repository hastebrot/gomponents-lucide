package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Feather(children ...g.Node) g.Node {
	svg := `<path d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5z" /><line x1="16" x2="2" y1="8" y2="22" /><line x1="17.5" x2="9" y1="15" y2="15" />`
	return Icon(g.Group(children), g.Raw(svg))
}
