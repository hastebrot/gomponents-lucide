package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UtilityPole(children ...g.Node) g.Node {
	svg := `<path d="M12 2v20" /><path d="M2 5h20" /><path d="M3 3v2" /><path d="M7 3v2" /><path d="M17 3v2" /><path d="M21 3v2" /><path d="m19 5-7 7-7-7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
