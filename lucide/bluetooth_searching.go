package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BluetoothSearching(children ...g.Node) g.Node {
	svg := `<path d="m7 7 10 10-5 5V2l5 5L7 17" /><path d="M20.83 14.83a4 4 0 0 0 0-5.66" /><path d="M18 12h.01" />`
	return Icon(g.Group(children), g.Raw(svg))
}
