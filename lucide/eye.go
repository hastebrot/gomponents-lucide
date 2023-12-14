package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Eye(children ...g.Node) g.Node {
	svg := `<path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z" /><circle cx="12" cy="12" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
