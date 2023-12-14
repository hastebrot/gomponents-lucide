package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func GitPullRequestClosed(children ...g.Node) g.Node {
	svg := `<circle cx="6" cy="6" r="3" /><path d="M6 9v12" /><path d="m21 3-6 6" /><path d="m21 9-6-6" /><path d="M18 11.5V15" /><circle cx="18" cy="18" r="3" />`
	return Icon(g.Group(children), g.Raw(svg))
}
