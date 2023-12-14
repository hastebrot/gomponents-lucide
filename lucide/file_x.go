package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileX(children ...g.Node) g.Node {
	svg := `<path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" /><polyline points="14 2 14 8 20 8" /><line x1="9.5" x2="14.5" y1="12.5" y2="17.5" /><line x1="14.5" x2="9.5" y1="12.5" y2="17.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
