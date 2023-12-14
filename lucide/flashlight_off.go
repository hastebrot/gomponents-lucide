package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FlashlightOff(children ...g.Node) g.Node {
	svg := `<path d="M16 16v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4" /><path d="M7 2h11v4c0 2-2 2-2 4v1" /><line x1="11" x2="18" y1="6" y2="6" /><line x1="2" x2="22" y1="2" y2="22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
