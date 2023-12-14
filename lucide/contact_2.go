package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Contact2(children ...g.Node) g.Node {
	svg := `<path d="M16 18a4 4 0 0 0-8 0" /><circle cx="12" cy="11" r="3" /><rect width="18" height="18" x="3" y="4" rx="2" /><line x1="8" x2="8" y1="2" y2="4" /><line x1="16" x2="16" y1="2" y2="4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
