package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CloudRainWind(children ...g.Node) g.Node {
	svg := `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242" /><path d="m9.2 22 3-7" /><path d="m9 13-3 7" /><path d="m17 13-3 7" />`
	return Icon(g.Group(children), g.Raw(svg))
}
