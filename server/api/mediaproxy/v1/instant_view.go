package v1

import (
	"context"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/cixtor/readability"

	mediaproxyv1 "github.com/harmony-development/legato/gen/mediaproxy/v1"
	lru "github.com/hashicorp/golang-lru"
)

type instantViewData struct {
	instantLRU *lru.ARCCache
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
		data := val.(string)
		resp.Content = data
		return
	}

	req, err := http.Get(r.Url)
	if err != nil {
		return
	}
	defer req.Body.Close()

	read := readability.New()
	body, err := read.Parse(req.Body, r.Url)
	if err != nil {
		return
	}

	converted, err := converter.ConvertString(body.Content)
	if err != nil {
		return
	}

	v1.instantLRU.Add(r.Url, converted)
	resp.Content = converted

	return
}
