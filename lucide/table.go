package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Table(children ...g.Node) g.Node {
	svg := `<path d="M12 3v18" /><rect width="18" height="18" x="3" y="3" rx="2" /><path d="M3 9h18" /><path d="M3 15h18" />`
	return Icon(g.Group(children), g.Raw(svg))
}
