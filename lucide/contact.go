package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Contact(children ...g.Node) g.Node {
	svg := `<path d="M17 18a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2" /><rect width="18" height="18" x="3" y="4" rx="2" /><circle cx="12" cy="10" r="2" /><line x1="8" x2="8" y1="2" y2="4" /><line x1="16" x2="16" y1="2" y2="4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
