package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Speaker(children ...g.Node) g.Node {
	svg := `<rect width="16" height="20" x="4" y="2" rx="2" /><path d="M12 6h.01" /><circle cx="12" cy="14" r="4" /><path d="M12 14h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
