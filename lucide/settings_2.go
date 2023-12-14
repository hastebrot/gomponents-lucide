package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Settings2(children ...g.Node) g.Node {
	svg := `<path d="M20 7h-9" /><path d="M14 17H5" /><circle cx="17" cy="17" r="3" /><circle cx="7" cy="7" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
