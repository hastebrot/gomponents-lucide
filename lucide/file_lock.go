package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func FileLock(children ...g.Node) g.Node {
	svg := `<path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" /><rect width="8" height="6" x="8" y="12" rx="1" /><path d="M15 12v-2a3 3 0 1 0-6 0v2" />`
	return Icon(g.Group(children), g.Raw(svg))
}
