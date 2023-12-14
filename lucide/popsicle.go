package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Popsicle(children ...g.Node) g.Node {
	svg := `<path d="M18.6 14.4c.8-.8.8-2 0-2.8l-8.1-8.1a4.95 4.95 0 1 0-7.1 7.1l8.1 8.1c.9.7 2.1.7 2.9-.1Z" /><path d="m22 22-5.5-5.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
