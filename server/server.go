package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dshills/apix/config"
	"github.com/dshills/apix/errorx"
)

// Serve will serve http based on config and routes
func Serve(con *config.Server, routes http.Handler) error {
	if con.Port == "" {
		return fmt.Errorf("%v is required, quitting", con.Prefix+"_SERVER_PORT")
	}
	adr := fmt.Sprintf("%v:%v", con.Host, con.Port)
	log.Printf("Starting server at %v\n", adr)
	log.Fatal(http.ListenAndServe(adr, routes))
	return nil
}

// ConfigLog will setup a log if configured
func ConfigLog(con *config.Server) (*os.File, error) {
	// Verbose logging with file name and line number
	log.SetFlags(log.LstdFlags)
	if con.Log != "" {
		f, err := os.OpenFile(con.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		log.SetOutput(f)
		return f, nil
	}
	return nil, nil
}

// GetJSONBody will unmarshal a json http body
func GetJSONBody(r *http.Request, i interface{}) *errorx.Err {
	buf := new(bytes.Buffer)
	io.Copy(buf, r.Body)
	r.Body.Close()
	if err := json.Unmarshal(buf.Bytes(), i); err != nil {
		return errorx.New("Failed to decode json", err, errorx.BadRequest)
	}
	return nil
}

// SendJSON sends data in i as JSON
func SendJSON(w http.ResponseWriter, r *http.Request, i interface{}) {
	if i == nil {
		errorx.New("No data to send", fmt.Errorf("interface is nil"), errorx.InternalServerError).Write(w, r)
		return
	}
	b, err := json.Marshal(&i)
	if ierr := errorx.NewIfErr("Failed to create json", err, errorx.InternalServerError); ierr != nil {
		ierr.Write(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/javascript")
	fmt.Fprint(w, string(b))
}
