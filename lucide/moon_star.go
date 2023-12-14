package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MoonStar(children ...g.Node) g.Node {
	svg := `<path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z" /><path d="M19 3v4" /><path d="M21 5h-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
