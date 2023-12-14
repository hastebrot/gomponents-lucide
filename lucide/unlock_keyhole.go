package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func UnlockKeyhole(children ...g.Node) g.Node {
	svg := `<circle cx="12" cy="16" r="1" /><rect x="3" y="10" width="18" height="12" rx="2" /><path d="M7 10V7a5 5 0 0 1 9.33-2.5" />`
	return Icon(g.Group(children), g.Raw(svg))
}
