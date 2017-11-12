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

func readFiles(fs *ext2.FSReader) []byte {
	var buff bytes.Buffer
	for _, de := range fs.RootDir.Dentries {
		if de.FileType != "Regular File" {
			continue
		}
		file := bytes.NewBufferString("==================================================")
		file.WriteString("==================================================\n\n")
		file.WriteString(fmt.Sprintf("File Name: %s\n", de.Name))

		file.WriteString("\n--------------------------------------------------\n\n")
		if content, err := fs.ReadFile("/", de.Name); err == nil {
			file.Write(bytes.Trim(content, "\x00"))

			file.WriteString("\n==================================================")
			file.WriteString("==================================================\n")
			buff.Write(file.Bytes())
		}
	}

	return buff.Bytes()
}

func showfs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	path := strings.TrimSpace(r.Form.Get("path"))
	fs := ext2.NewFSReader(path, ext2.DEFAULT_BLOCK_SIZE)
	b := fs.BufferedString()
	b.WriteString("\n\n")

	b.Write(readFiles(fs))

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
