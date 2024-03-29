package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileClock(children ...g.Node) g.Node {
	svg := `<path d="M16 22h2c.5 0 1-.2 1.4-.6.4-.4.6-.9.6-1.4V7.5L14.5 2H6c-.5 0-1 .2-1.4.6C4.2 3 4 3.5 4 4v3" /><polyline points="14 2 14 8 20 8" /><circle cx="8" cy="16" r="6" /><path d="M9.5 17.5 8 16.25V14" />`
	return Icon(g.Group(children), g.Raw(svg))
}
