package views

import "embed"

// 静态资源打包

//go:embed sse/testSSE.html
var SSEStaticFS embed.FS

//go:embed all:dist
var WebStaticFS embed.FS
