package v1

import (
	"bytes"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/cixtor/readability"
	"github.com/dyatlov/go-opengraph/opengraph"
)

type sitedata opengraph.OpenGraph
type mediadata struct {
	mimetype string
	filename string
}

type data struct {
	OG *sitedata
	MD *mediadata
	RD *string
}

func (v1 *V1) fetch(reqURL string) (d *data, err error) {
	if v, ok := v1.dataLRU.Get(reqURL); ok {
		return v.(*data), nil
	}

	data, err := fetchSite(reqURL)
	if err != nil {
		return nil, err
	}

	v1.dataLRU.Add(reqURL, data)
	return data, nil
}

func fetchSite(reqURL string) (d *data, err error) {
	d = new(data)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	parsedURL, err := url.Parse(reqURL)
	if err != nil {
		return nil, err
	}

	mimetype, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		mimetype, _, err = mime.ParseMediaType(mime.TypeByExtension(path.Ext(path.Base(parsedURL.Path))))
		if err != nil {
			return nil, err
		}
	}
	if !strings.Contains(mimetype, "text/html") {
		_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
		if err != nil {
			return nil, err
		}

		filename, ok := params["filename"]
		if ok {
			return &data{
				MD: &mediadata{
					filename: filename,
					mimetype: mimetype,
				},
			}, nil
		}
		return &data{
			MD: &mediadata{
				filename: path.Base(parsedURL.Path),
				mimetype: mimetype,
			},
		}, nil
	}

	read := readability.New()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if read.IsReadable(bytes.NewReader(data)) {
		body, err := read.Parse(bytes.NewReader(data), reqURL)
		if err != nil {
			return nil, err
		}

		converted, err := converter.ConvertString(body.Content)
		if err != nil {
			return nil, err
		}

		d.RD = &converted
	}

	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	sd := sitedata(*og)
	d.OG = &sd

	return d, nil
}
