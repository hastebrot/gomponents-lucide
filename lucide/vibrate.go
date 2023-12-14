package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Vibrate(children ...g.Node) g.Node {
	svg := `<path d="m2 8 2 2-2 2 2 2-2 2" /><path d="m22 8-2 2 2 2-2 2 2 2" /><rect width="8" height="14" x="8" y="5" rx="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
