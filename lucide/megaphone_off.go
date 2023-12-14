package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MegaphoneOff(children ...g.Node) g.Node {
	svg := `<path d="M9.26 9.26 3 11v3l14.14 3.14" /><path d="M21 15.34V6l-7.31 2.03" /><path d="M11.6 16.8a3 3 0 1 1-5.8-1.6" /><line x1="2" x2="22" y1="2" y2="22" />`
	return Icon(g.Group(children), g.Raw(svg))
}
