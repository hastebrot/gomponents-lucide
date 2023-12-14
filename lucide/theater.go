package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func Theater(children ...g.Node) g.Node {
	svg := `<path d="M2 10s3-3 3-8" /><path d="M22 10s-3-3-3-8" /><path d="M10 2c0 4.4-3.6 8-8 8" /><path d="M14 2c0 4.4 3.6 8 8 8" /><path d="M2 10s2 2 2 5" /><path d="M22 10s-2 2-2 5" /><path d="M8 15h8" /><path d="M2 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1" /><path d="M14 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
