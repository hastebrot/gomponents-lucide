package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Waypoints(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="4.5" r="2.5" /><path d="m10.2 6.3-3.9 3.9" /><circle cx="4.5" cy="12" r="2.5" /><path d="M7 12h10" /><circle cx="19.5" cy="12" r="2.5" /><path d="m13.8 17.7 3.9-3.9" /><circle cx="12" cy="19.5" r="2.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
