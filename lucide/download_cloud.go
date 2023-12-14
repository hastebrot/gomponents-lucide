package lucide

import (
	. "github.com/hastebrot/gomponents-lucide"
	g "github.com/maragudk/gomponents"
)

func DownloadCloud(children ...g.Node) g.Node {
	svg := `<path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242" /><path d="M12 12v9" /><path d="m8 17 4 4 4-4" />`
	return Icon(g.Group(children), g.Raw(svg))
}
