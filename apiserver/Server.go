package apiserver

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/Akkurate/utils/logging"

	pathToRegexp "github.com/soongo/path-to-regexp"
)

// APIServer APIServer
type APIServer struct {
	port string
}

type APIServerProps struct {
	Port string
}

// NewAPIServer NewAPIServer
func NewAPIServer(props *APIServerProps) *APIServer {
	x := &APIServer{
		port: props.Port,
	}
	return x
}

func (s *APIServer) Start(routeMapping map[string]RouteHandler) {
	logging.Info("Starting API Server on port %v", s.port)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		var currentURL = req.URL.Path
		logging.Info("Handle %v", currentURL)
		safe := func() {
			if r := recover(); r != nil {
				errLine := fmt.Sprintf("%v", r)
				stackTrace := errLine + "\n\n" + string(debug.Stack())
				logging.Error("%v", stackTrace)
				res.Header().Set("Content-Type", "application/json")
				data, _ := json.Marshal(&ErrorResponse{Error: errLine})
				res.WriteHeader(500)
				res.Write([]byte(data))
			}
		}
		defer safe()

		context := &Context{
			Res:    res,
			Req:    req,
			Safe:   safe,
			Params: map[string]string{},
		}

		method := strings.ToLower(req.Method)
		for path, fn := range routeMapping {
			s := strings.Split(path, " ")
			if s[0] == method {
				var tokens []pathToRegexp.Token
				regexp := pathToRegexp.Must(pathToRegexp.PathToRegexp(s[1], &tokens, nil))
				match, _ := regexp.FindStringMatch(currentURL)
				if match != nil {
					for index, g := range match.Groups() {
						if index > 0 {
							if len(tokens) >= index {
								t := tokens[index-1]
								context.Params[fmt.Sprintf("%v", t.Name)] = g.String()
							}
						}
					}
					fn(context)
					return
				}
			}
		}

		res.WriteHeader(404)
		res.Write([]byte("Not found"))

	})

	http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil)
}
