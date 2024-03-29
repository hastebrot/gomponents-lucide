package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BoomBox(children ...g.Node) g.Node {
	svg := `<path d="M4 9V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4" /><path d="M8 8v1" /><path d="M12 8v1" /><path d="M16 8v1" /><rect width="20" height="12" x="2" y="9" rx="2" /><circle cx="8" cy="15" r="2" /><circle cx="16" cy="15" r="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
