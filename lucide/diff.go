package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Diff(children ...g.Node) g.Node {
	svg := `<path d="M12 3v14" /><path d="M5 10h14" /><path d="M5 21h14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
