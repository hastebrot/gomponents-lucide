package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CableCar(children ...g.Node) g.Node {
	svg := `<path d="M10 3h.01" /><path d="M14 2h.01" /><path d="m2 9 20-5" /><path d="M12 12V6.5" /><rect width="16" height="10" x="4" y="12" rx="3" /><path d="M9 12v5" /><path d="M15 12v5" /><path d="M4 17h16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
