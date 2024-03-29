package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TouchpadOff(children ...g.Node) g.Node {
	svg := `<path d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16" /><path d="M2 14h12" /><path d="M22 14h-2" /><path d="M12 20v-6" /><path d="m2 2 20 20" /><path d="M22 16V6a2 2 0 0 0-2-2H10" />`
	return Icon(g.Group(children), g.Raw(svg))
}
