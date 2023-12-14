package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Antenna(children ...g.Node) g.Node {
	svg := `<path d="M2 12 7 2" /><path d="m7 12 5-10" /><path d="m12 12 5-10" /><path d="m17 12 5-10" /><path d="M4.5 7h15" /><path d="M12 16v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
