package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileAudio2(children ...g.Node) g.Node {
	svg := `<path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v2" /><polyline points="14 2 14 8 20 8" /><path d="M2 17v-3a4 4 0 0 1 8 0v3" /><circle cx="9" cy="17" r="1" /><circle cx="3" cy="17" r="1" />`
	return Icon(g.Group(children), g.Raw(svg))
}
