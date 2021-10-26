package apiserver

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"
	"time"

	"github.com/Akkurate/utils/conv"
	"github.com/Akkurate/utils/logging"
)

type Param struct {
	Raw string
}

func (p *Param) AsInt64() int64 {
	v, e := strconv.ParseInt(p.Raw, 10, 64)
	if e != nil {
		panic(`Integer is required`)
	}
	return v
}

func (p *Param) AsInt() int {
	v, e := strconv.ParseInt(p.Raw, 10, 64)
	if e != nil {
		panic(`Integer is required`)
	}
	return int(v)
}

func (p *Param) AsString() string {
	return p.Raw
}

type Context struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Safe   func()
	Params map[string]string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (c *Context) GetParam(name string) *Param {
	if v, ok := c.Params[name]; ok {
		return &Param{
			Raw: v,
		}
	}
	return nil
}

func (c *Context) GetRequiredParam(name string) *Param {
	if v, ok := c.Params[name]; ok {
		return &Param{
			Raw: v,
		}
	}
	panic(fmt.Sprintf("Param %v is required", name))
}

func (c *Context) Unauthorized() {
	c.Res.Header().Set("WWW-Authenticate", `Basic realm="Auth required"`)
	c.Res.Write([]byte("Unauthorized"))
}

func (c *Context) GetQueryInt(name string, defaultValue int) int {
	q := c.Req.URL.Query()
	data := q.Get(name)
	if data == "" {
		return defaultValue
	}

	return conv.StringToInt(data)
}

func (c *Context) getQueryTime(name string) *time.Time {
	q := c.Req.URL.Query()
	data := q.Get(name)
	if data == "" {
		return nil
	}

	tm, e := time.Parse("2006-01-02T15:04:05", data)
	if e != nil {
		logging.Info("invalid date")
		return nil
	}
	return &tm
}
func (c *Context) GetQueryString(name string, defaultValue string) string {
	q := c.Req.URL.Query()
	data := q.Get(name)
	if data == "" {
		return defaultValue
	}
	return data
}

func (c *Context) Decode(ptr interface{}) {
	err := json.NewDecoder(c.Req.Body).Decode(ptr)
	if err != nil {
		panic(err)
	}
}

func (c *Context) SendOK() {
	c.SendJSON(map[string]interface{}{
		"ok": true,
	})
}
func (c *Context) SendJSON(data interface{}) {
	c.Res.Header().Set("Content-Type", "application/json")
	//c.Res.Header().Set("Content-Encoding", "gzip")

	b, _ := json.Marshal(data)
	logging.Info("%v", string(b))
	c.Res.WriteHeader(202)
	c.Res.Write(b)
}

func (c *Context) Error(message string) {
	c.Res.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(&ErrorResponse{Error: message})
	c.Res.WriteHeader(500)
	c.Res.Write([]byte(data))
}

type RouteHandler func(context *Context)
