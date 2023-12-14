package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Subtitles(children ...g.Node) g.Node {
	svg := `<path d="M7 13h4" /><path d="M15 13h2" /><path d="M7 9h2" /><path d="M13 9h4" /><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2Z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
