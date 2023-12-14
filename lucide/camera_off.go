package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func CameraOff(children ...g.Node) g.Node {
	svg := `<line x1="2" x2="22" y1="2" y2="22" /><path d="M7 7H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16" /><path d="M9.5 4h5L17 7h3a2 2 0 0 1 2 2v7.5" /><path d="M14.121 15.121A3 3 0 1 1 9.88 10.88" />`
	return Icon(g.Group(children), g.Raw(svg))
}
