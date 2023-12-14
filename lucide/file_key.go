package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileKey(children ...g.Node) g.Node {
	svg := `<path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" /><circle cx="10" cy="16" r="2" /><path d="m16 10-4.5 4.5" /><path d="m15 11 1 1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
