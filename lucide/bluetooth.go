package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Bluetooth(children ...g.Node) g.Node {
	svg := `<path d="m7 7 10 10-5 5V2l5 5L7 17" />`
	return Icon(g.Group(children), g.Raw(svg))
}
