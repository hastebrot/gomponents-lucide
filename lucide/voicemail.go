package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Voicemail(children ...g.Node) g.Node {
	svg := `<circle cx="6" cy="12" r="4" /><circle cx="18" cy="12" r="4" /><line x1="6" x2="18" y1="16" y2="16" />`
	return Icon(g.Group(children), g.Raw(svg))
}
