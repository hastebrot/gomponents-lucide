package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func RadioReceiver(children ...g.Node) g.Node {
	svg := `<path d="M5 16v2" /><path d="M19 16v2" /><rect width="20" height="8" x="2" y="8" rx="2" /><path d="M18 12h0" />`
	return Icon(g.Group(children), g.Raw(svg))
}
