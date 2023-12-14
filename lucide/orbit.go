package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Orbit(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="3" /><circle cx="19" cy="5" r="2" /><circle cx="5" cy="19" r="2" /><path d="M10.4 21.9a10 10 0 0 0 9.941-15.416" /><path d="M13.5 2.1a10 10 0 0 0-9.841 15.416" />`
	return Icon(g.Group(children), g.Raw(svg))
}
