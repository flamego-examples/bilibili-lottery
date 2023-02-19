package context

import (
	"encoding/json"
	"net/http"

	"github.com/flamego/flamego"
)

// Context represents context of a request.
type Context struct {
	flamego.Context
}

// Contexter initializes a classic context for a request.
func Contexter() flamego.Handler {
	return func(ctx flamego.Context) {
		c := Context{
			Context: ctx,
		}
		c.Map(c)
	}
}

func (c *Context) Success(data ...interface{}) error {
	c.ResponseWriter().Header().Set("Content-Type", "application/json; charset=utf-8")
	c.ResponseWriter().WriteHeader(http.StatusOK)

	var d interface{}
	if len(data) == 1 {
		d = data[0]
	} else {
		d = ""
	}

	return json.NewEncoder(c.ResponseWriter()).Encode(
		map[string]interface{}{
			"data": d,
		},
	)
}

func (c *Context) Error(err error) error {
	c.ResponseWriter().Header().Set("Content-Type", "application/json; charset=utf-8")
	c.ResponseWriter().WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(c.ResponseWriter()).Encode(
		map[string]interface{}{
			"error": err.Error(),
		},
	)
}

func (c *Context) ServerError() error {
	c.ResponseWriter().Header().Set("Content-Type", "application/json; charset=utf-8")
	c.ResponseWriter().WriteHeader(http.StatusInternalServerError)

	return json.NewEncoder(c.ResponseWriter()).Encode(
		map[string]interface{}{
			"error": "internal server error",
		},
	)
}
