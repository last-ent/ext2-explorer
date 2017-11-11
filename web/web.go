package web

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/last-ent/ext2-explorer/ext2"
)

func loadfs(w http.ResponseWriter, r *http.Request) {
	s := `
	<html>
	<head>
		<title>EXT2 FS Reader: Load Page</title>
	</head>
	<body>
		<form action="/show" method="POST">
		FS Path:  <input type="textarea" name="path"/>
		</form>
	</body>
	</html>
	`
	var b bytes.Buffer
	b.WriteString(s)
	w.Write(b.Bytes())
}

func showfs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := strings.TrimSpace(r.Form.Get("path"))
	fs := ext2.NewFSReader(path, ext2.DEFAULT_BLOCK_SIZE)
	// var b bytes.Buffer
	// b.WriteString(strings.Join(fs.BufferString(), "\n"))

	// b.WriteString(fmt.Sprintf("%#v", fs.ReprMap()))
	b := fs.BufferedString()

	w.Write(b.Bytes())
}

// StartServer allows us to interact with FS from Web UI.
// Fow now it takes path to FS and returns some details related to it.
func StartServer() {
	fmt.Println("Adding handler functions...")
	http.HandleFunc("/", loadfs)
	http.HandleFunc("/show", showfs)

	fmt.Println("Starting web server at port 8080...")
	http.ListenAndServe(":8080", nil)
}
