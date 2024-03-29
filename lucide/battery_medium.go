package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BatteryMedium(children ...g.Node) g.Node {
	svg := `<rect width="16" height="10" x="2" y="7" rx="2" ry="2" /><line x1="22" x2="22" y1="11" y2="13" /><line x1="6" x2="6" y1="11" y2="13" /><line x1="10" x2="10" y1="11" y2="13" />`
	return Icon(g.Group(children), g.Raw(svg))
}
