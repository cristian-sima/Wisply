package wisply

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
)

// Captcha encapsulates the operations for Captcha
type Captcha struct {
	Adapter
}

// Serve is an adaptation for beegoo for the captcha package
// The code is from the server.go from captcha file
// Credits 	"github.com/dchest/captcha"
func (controller *Captcha) Serve() {
	r := controller.Ctx.Request
	w := controller.Ctx.ResponseWriter

	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if controller.serve(w, r, id, ext, lang, download) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
	// Ignore other errors.
}

func (controller *Captcha) serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, captcha.StdWidth, captcha.StdHeight)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
