package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GitCommitVertical(children ...g.Node) g.Node {
	svg := `<path d="M12 3v6" /><circle cx="12" cy="12" r="3" /><path d="M12 15v6" />`
	return Icon(g.Group(children), g.Raw(svg))
}
