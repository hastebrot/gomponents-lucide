package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ParkingMeter(children ...g.Node) g.Node {
	svg := `<path d="M9 9a3 3 0 1 1 6 0" /><path d="M12 12v3" /><path d="M11 15h2" /><path d="M19 9a7 7 0 1 0-13.6 2.3C6.4 14.4 8 19 8 19h8s1.6-4.6 2.6-7.7c.3-.8.4-1.5.4-2.3" /><path d="M12 19v3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
