package proxy

import (
	"io"
	"net/http"

	"github.com/paralus/relay/pkg/utils"
)

// Func is responsible for forwarding a remote connection to local server
// and writing the response.
type Func func(w io.Writer, r io.ReadCloser, msg *utils.ControlMessage, req *http.Request)
