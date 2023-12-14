package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ScanEye(children ...g.Node) g.Node {
	svg := `<path d="M3 7V5a2 2 0 0 1 2-2h2" /><path d="M17 3h2a2 2 0 0 1 2 2v2" /><path d="M21 17v2a2 2 0 0 1-2 2h-2" /><path d="M7 21H5a2 2 0 0 1-2-2v-2" /><circle cx="12" cy="12" r="1" /><path d="M5 12s2.5-5 7-5 7 5 7 5-2.5 5-7 5-7-5-7-5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
