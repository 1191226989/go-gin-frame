package assets

import "embed"

var (
	//go:embed static
	Static embed.FS
	// c.FileFromFS("static/avatar.png", http.FS(assets.Static))

)
