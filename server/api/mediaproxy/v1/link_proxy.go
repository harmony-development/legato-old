package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
	mediaproxyv1 "github.com/harmony-development/legato/gen/mediaproxy/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	lru "github.com/hashicorp/golang-lru"
)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 1 * time.Second,
			Burst:    3,
		},
	}, "/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata")
}

type linkData struct {
	linkLRU *lru.ARCCache
}

func copyOGIntoProtobuf(og *opengraph.OpenGraph, md *mediaproxyv1.SiteMetadata) {
	md.Description = og.Description
	if len(og.Images) > 0 {
		md.Image = og.Images[0].URL
	}
	md.Kind = og.Type
	md.Url = og.URL
	md.PageTitle = og.Title
	md.SiteTitle = og.SiteName
}

func (v1 *V1) obtainOG(url string, out *mediaproxyv1.SiteMetadata) error {
	if val, ok := v1.linkLRU.Get(url); ok {
		data := val.(*opengraph.OpenGraph)
		copyOGIntoProtobuf(data, out)
	}

	req, err := http.Get(url)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(req.Body)
	if err != nil {
		return err
	}

	v1.linkLRU.Add(url, og)
	copyOGIntoProtobuf(og, out)

	return nil
}

// FetchLinkMetadata implements the FetchLinkMetadata RPC
func (v1 *V1) FetchLinkMetadata(ctx context.Context, r *mediaproxyv1.FetchLinkMetadataRequest) (resp *mediaproxyv1.SiteMetadata, err error) {
	resp = &mediaproxyv1.SiteMetadata{}
	err = v1.obtainOG(r.Url, resp)
	return
}
