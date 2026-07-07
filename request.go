package resty

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// Request struct is used to build a single request.
type Request struct {
	URL            string
	Method         string
	QueryParam     url.Values
	FormData       url.Values
	Header         http.Header
	Time           time.Time
	Body           interface{}
	Result         interface{}
	Error          interface{}
	RawRequest     *http.Request
	client         *Client
	bodyBuf        *bytes.Buffer
	isMultiPart    bool
	ctx            context.Context
	fallbackResult interface{}
	fallbackError  interface{}
}

// SetBody method sets the request body. It accepts string, []byte, struct, map, slice and io.Reader.
func (r *Request) SetBody(body interface{}) *Request {
	r.Body = body
	return r
}

// Context returns the request context.
func (r *Request) Context() context.Context {
	if r.ctx == nil {
		return context.Background()
	}
	return r.ctx
}

// SetContext method sets the request context.
func (r *Request) SetContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}
