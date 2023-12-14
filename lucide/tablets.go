package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Tablets(children ...g.Node) g.Node {
	svg := `<circle cx="7" cy="7" r="5" /><circle cx="17" cy="17" r="5" /><path d="M12 17h10" /><path d="m3.46 10.54 7.08-7.08" />`
	return Icon(g.Group(children), g.Raw(svg))
}
