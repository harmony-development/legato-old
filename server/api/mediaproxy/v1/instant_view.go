package v1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/cixtor/readability"

	mediaproxyv1 "github.com/harmony-development/legato/gen/mediaproxy/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	lru "github.com/hashicorp/golang-lru"
)

type instantViewData struct {
	instantLRU *lru.ARCCache
}

type instantViewResult struct {
	ok   bool
	data string
}

var converter = md.NewConverter("", true, nil)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    1,
		},
		Auth: true,
	}, "/protocol.mediaproxy.v1.MediaProxyService/InstantView")
}

func (v1 *V1) CanInstantView(ctx context.Context, r *mediaproxyv1.InstantViewRequest) (resp *mediaproxyv1.CanInstantViewResponse, err error) {
	resp = &mediaproxyv1.CanInstantViewResponse{
		CanInstantView: true,
	}

	if val, ok := v1.instantLRU.Get(r.Url); ok {
		data := val.(instantViewResult)
		resp.CanInstantView = data.ok
		return
	}

	req, err := http.Get(r.Url)
	if err != nil {
		return
	}
	defer req.Body.Close()

	read := readability.New()
	if !read.IsReadable(req.Body) {
		v1.instantLRU.Add(r.Url, instantViewResult{
			ok: false,
		})
		resp.CanInstantView = false
	}

	return
}

// InstantView implements the InstantView RPC
func (v1 *V1) InstantView(ctx context.Context, r *mediaproxyv1.InstantViewRequest) (resp *mediaproxyv1.InstantViewResponse, err error) {
	resp = &mediaproxyv1.InstantViewResponse{
		Metadata: &mediaproxyv1.SiteMetadata{},
	}
	err = v1.obtainOG(r.Url, resp.Metadata)
	if err != nil {
		return
	}

	if val, ok := v1.instantLRU.Get(r.Url); ok {
		data := val.(instantViewResult)
		resp.IsValid = data.ok
		resp.Content = data.data
		return
	}

	req, err := http.Get(r.Url)
	if err != nil {
		return
	}
	defer req.Body.Close()

	buffer := new(bytes.Buffer)
	tee := io.TeeReader(req.Body, buffer)

	read := readability.New()

	if !read.IsReadable(tee) {
		v1.instantLRU.Add(r.Url, instantViewResult{
			ok: false,
		})
		resp.IsValid = false
		return
	}

	body, err := read.Parse(io.MultiReader(buffer, tee), r.Url)
	if err != nil {
		return
	}

	converted, err := converter.ConvertString(body.Content)
	if err != nil {
		return
	}

	v1.instantLRU.Add(r.Url, instantViewResult{
		ok:   true,
		data: converted,
	})
	resp.IsValid = true
	resp.Content = converted

	return
}
