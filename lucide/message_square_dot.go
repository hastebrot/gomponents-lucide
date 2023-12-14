package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareDot(children ...g.Node) g.Node {
	svg := `<path d="M11.7 3H5a2 2 0 0 0-2 2v16l4-4h12a2 2 0 0 0 2-2v-2.7" /><circle cx="18" cy="6" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
