package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Contrast(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="12" r="10" /><path d="M12 18a6 6 0 0 0 0-12v12z" />`
	return Icon(g.Group(children), g.Raw(svg))
}
