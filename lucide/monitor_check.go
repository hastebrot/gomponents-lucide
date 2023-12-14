package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MonitorCheck(children ...g.Node) g.Node {
	svg := `<path d="m9 10 2 2 4-4" /><rect width="20" height="14" x="2" y="3" rx="2" /><path d="M12 17v4" /><path d="M8 21h8" />`
	return Icon(g.Group(children), g.Raw(svg))
}
