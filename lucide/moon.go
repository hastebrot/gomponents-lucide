package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Moon(children ...g.Node) g.Node {
	svg := `<path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
