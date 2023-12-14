package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Power(children ...g.Node) g.Node {
	svg := `<path d="M12 2v10" /><path d="M18.4 6.6a9 9 0 1 1-12.77.04" />`
	return Icon(g.Group(children), g.Raw(svg))
}
