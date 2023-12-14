package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ParkingSquare(children ...g.Node) g.Node {
	svg := `<rect width="18" height="18" x="3" y="3" rx="2" /><path d="M9 17V7h4a3 3 0 0 1 0 6H9" />`
	return Icon(g.Group(children), g.Raw(svg))
}
