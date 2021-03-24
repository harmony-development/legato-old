package v1

import (
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
	mediaproxyv1 "github.com/harmony-development/legato/gen/mediaproxy/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/labstack/echo/v4"
)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 1 * time.Second,
			Burst:    3,
		},
	}, "/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata")
}

// TODO use this struct somewhere
// type linkData struct {
// 	linkLRU *lru.ARCCache
// }

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
	data, err := v1.fetch(url)
	if err != nil {
		return err
	}
	copyOGIntoProtobuf((*opengraph.OpenGraph)(data.OG), out)

	return nil
}

// FetchLinkMetadata implements the FetchLinkMetadata RPC
func (v1 *V1) FetchLinkMetadata(c echo.Context, r *mediaproxyv1.FetchLinkMetadataRequest) (resp *mediaproxyv1.FetchLinkMetadataResponse, err error) {
	resp = &mediaproxyv1.FetchLinkMetadataResponse{}

	data, err := v1.fetch(r.Url)
	if err != nil {
		return nil, err
	}

	if data.MD != nil {
		return &mediaproxyv1.FetchLinkMetadataResponse{
			Data: &mediaproxyv1.FetchLinkMetadataResponse_IsMedia{
				IsMedia: &mediaproxyv1.MediaMetadata{
					Mimetype: data.MD.mimetype,
					Filename: data.MD.filename,
				},
			},
		}, nil
	}

	return &mediaproxyv1.FetchLinkMetadataResponse{
		Data: &mediaproxyv1.FetchLinkMetadataResponse_IsSite{
			IsSite: func() (r *mediaproxyv1.SiteMetadata) {
				r = new(mediaproxyv1.SiteMetadata)
				copyOGIntoProtobuf((*opengraph.OpenGraph)(data.OG), r)
				return r
			}(),
		},
	}, nil
}
