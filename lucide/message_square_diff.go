package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func MessageSquareDiff(children ...g.Node) g.Node {
	svg := `<path d="m5 19-2 2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2" /><path d="M9 10h6" /><path d="M12 7v6" /><path d="M9 17h6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
