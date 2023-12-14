package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func BluetoothOff(children ...g.Node) g.Node {
	svg := `<path d="m17 17-5 5V12l-5 5" /><path d="m2 2 20 20" /><path d="M14.5 9.5 17 7l-5-5v4.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
