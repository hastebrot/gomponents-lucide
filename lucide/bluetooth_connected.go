package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BluetoothConnected(children ...g.Node) g.Node {
	svg := `<path d="m7 7 10 10-5 5V2l5 5L7 17" /><line x1="18" x2="21" y1="12" y2="12" /><line x1="3" x2="6" y1="12" y2="12" />`
	return Icon(g.Group(children), g.Raw(svg))
}
