package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tractor(children ...g.Node) g.Node {
	svg := `<path d="M3 4h9l1 7" /><path d="M4 11V4" /><path d="M8 10V4" /><path d="M18 5c-.6 0-1 .4-1 1v5.6" /><path d="m10 11 11 .9c.6 0 .9.5.8 1.1l-.8 5h-1" /><circle cx="7" cy="15" r=".5" /><circle cx="7" cy="15" r="5" /><path d="M16 18h-5" /><circle cx="18" cy="18" r="2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
