package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MonitorOff(children ...g.Node) g.Node {
	svg := `<path d="M17 17H4a2 2 0 0 1-2-2V5c0-1.5 1-2 1-2" /><path d="M22 15V5a2 2 0 0 0-2-2H9" /><path d="M8 21h8" /><path d="M12 17v4" /><path d="m2 2 20 20" />`
	return Icon(g.Group(children), g.Raw(svg))
}
