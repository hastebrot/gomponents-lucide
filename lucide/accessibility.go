package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Accessibility(children ...g.Node) g.Node {
	svg := `<circle cx="16" cy="4" r="1" /><path d="m18 19 1-7-6 1" /><path d="m5 8 3-3 5.5 3-2.36 3.5" /><path d="M4.24 14.5a5 5 0 0 0 6.88 6" /><path d="M13.76 17.5a5 5 0 0 0-6.88-6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
