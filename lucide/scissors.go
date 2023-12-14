package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Scissors(children ...g.Node) g.Node {
	svg := `<circle cx="6" cy="6" r="3" /><path d="M8.12 8.12 12 12" /><path d="M20 4 8.12 15.88" /><circle cx="6" cy="18" r="3" /><path d="M14.8 14.8 20 20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
