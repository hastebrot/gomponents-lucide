package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ThermometerSun(children ...g.Node) g.Node {
	svg := `<path d="M12 9a4 4 0 0 0-2 7.5" /><path d="M12 3v2" /><path d="m6.6 18.4-1.4 1.4" /><path d="M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z" /><path d="M4 13H2" /><path d="M6.34 7.34 4.93 5.93" />`
	return Icon(g.Group(children), g.Raw(svg))
}
