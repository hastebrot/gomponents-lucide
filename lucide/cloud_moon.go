package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CloudMoon(children ...g.Node) g.Node {
	svg := `<path d="M13 16a3 3 0 1 1 0 6H7a5 5 0 1 1 4.9-6Z" /><path d="M10.1 9A6 6 0 0 1 16 4a4.24 4.24 0 0 0 6 6 6 6 0 0 1-3 5.197" />`
	return Icon(g.Group(children), g.Raw(svg))
}
