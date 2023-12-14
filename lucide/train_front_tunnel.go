package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func TrainFrontTunnel(children ...g.Node) g.Node {
	svg := `<path d="M2 22V12a10 10 0 1 1 20 0v10" /><path d="M15 6.8v1.4a3 2.8 0 1 1-6 0V6.8" /><path d="M10 15h.01" /><path d="M14 15h.01" /><path d="M10 19a4 4 0 0 1-4-4v-3a6 6 0 1 1 12 0v3a4 4 0 0 1-4 4Z" /><path d="m9 19-2 3" /><path d="m15 19 2 3" />`
	return Icon(g.Group(children), g.Raw(svg))
}