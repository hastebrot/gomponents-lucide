package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Annoyed(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M8 15h8" /><path d="M8 9h2" /><path d="M14 9h2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
