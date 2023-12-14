package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BaggageClaim(children ...g.Node) g.Node {
	svg := `<path d="M22 18H6a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2" /><path d="M17 14V4a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v10" /><rect width="13" height="8" x="8" y="6" rx="1" /><circle cx="18" cy="20" r="2" /><circle cx="9" cy="20" r="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
