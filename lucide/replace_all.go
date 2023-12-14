package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func ReplaceAll(children ...g.Node) g.Node {
	svg := `<path d="M14 4c0-1.1.9-2 2-2" /><path d="M20 2c1.1 0 2 .9 2 2" /><path d="M22 8c0 1.1-.9 2-2 2" /><path d="M16 10c-1.1 0-2-.9-2-2" /><path d="m3 7 3 3 3-3" /><path d="M6 10V5c0-1.7 1.3-3 3-3h1" /><rect width="8" height="8" x="2" y="14" rx="2" /><path d="M14 14c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2" /><path d="M20 14c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
