package native

import (
	"github.com/machbase/jsh/engine"
	"github.com/machbase/jsh/native/http"
	"github.com/machbase/jsh/native/mqtt"
	"github.com/machbase/jsh/native/net"
	"github.com/machbase/jsh/native/parser"
	"github.com/machbase/jsh/native/readline"
	"github.com/machbase/jsh/native/shell"
	"github.com/machbase/jsh/native/stream"
	"github.com/machbase/jsh/native/ws"
	"github.com/machbase/jsh/native/zlib"
)

func Enable(n *engine.JSRuntime) {
	n.RegisterNativeModule("@jsh/process", n.Process)
	n.RegisterNativeModule("@jsh/fs", n.Filesystem)
	n.RegisterNativeModule("@jsh/shell", shell.Module)
	n.RegisterNativeModule("@jsh/readline", readline.Module)
	n.RegisterNativeModule("@jsh/http", http.Module)
	n.RegisterNativeModule("@jsh/ws", ws.Module)
	n.RegisterNativeModule("@jsh/mqtt", mqtt.Module)
	n.RegisterNativeModule("@jsh/stream", stream.Module)
	n.RegisterNativeModule("@jsh/zlib", zlib.Module)
	n.RegisterNativeModule("@jsh/net", net.Module)
	n.RegisterNativeModule("@jsh/parser", parser.Module)
}
