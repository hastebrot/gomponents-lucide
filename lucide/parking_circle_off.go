package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ParkingCircleOff(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="m5 5 14 14" /><path d="M13 13a3 3 0 1 0 0-6H9v2" /><path d="M9 17v-2.34" />`
	return Icon(g.Group(children), g.Raw(svg))
}
