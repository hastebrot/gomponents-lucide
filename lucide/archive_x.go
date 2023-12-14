package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ArchiveX(children ...g.Node) g.Node {
	svg := `<rect width="20" height="5" x="2" y="3" rx="1" /><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8" /><path d="m9.5 17 5-5" /><path d="m9.5 12 5 5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
