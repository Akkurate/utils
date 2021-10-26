# Akkurate utils

## Simple http server

```go
package main

import (
	"github.com/Akkurate/utils/apiserver"
	"github.com/Akkurate/utils/logging"
)

// Routes Routes
type Routes struct {
}

func (r *Routes) helloWorld() map[string]string {
	return map[string]string{"hello": "world"}
}
func (r *Routes) RouteIndex(c *apiserver.Context) {
	logging.Info("coming in here")
	a := r.helloWorld()
	a["id"] = c.GetParam("id").AsString()
	c.SendJSON(a)
}

func main() {
	routes := &Routes{}
	routeMapping := map[string]apiserver.RouteHandler{
		"get /:id": routes.RouteIndex,
	}
	server := apiserver.NewAPIServer(&apiserver.APIServerProps{
		Port: "7777",
	})
	server.Start(routeMapping)
}
```
