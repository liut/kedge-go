package kedge

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"time"

	"gopkg.in/resty.v1"
)

// vars
var (
	Timeout          = 12 * time.Second
	KeepAliveTimeout = 60 * time.Second
	IdleConnTimeout  = 30 * time.Second

	afterResponseHookFunc = func(c *resty.Client, r *resty.Response) error {

		if !r.IsSuccess() {
			return fmt.Errorf("%s %q bad status:%s", r.Request.Method, r.Request.URL, r.Status())
		}
		return nil
	}
)

type clientImpl struct {
	uri  string
	root string
	rc   *resty.Client
}

// New return a client instance with args: [root[, uri]]
func New(args ...string) ClientI {
	root := "."
	uri := "http://localhost:16180/api"
	if s, ok := os.LookupEnv("KC_KEDGE_URI"); ok && len(s) > 0 {
		uri = s
	}

	argc := len(args)
	if argc > 0 {
		root = args[0]
		if argc > 1 {
			uri = args[1]
		}
	}

	rclient := resty.New().
		SetTimeout(Timeout).
		SetTransport(&http.Transport{
			Dial:            (&net.Dialer{KeepAlive: KeepAliveTimeout}).Dial,
			IdleConnTimeout: IdleConnTimeout,
		}).OnAfterResponse(afterResponseHookFunc)

	client := &clientImpl{
		root: root,
		uri:  uri,
		rc:   rclient.SetRESTMode(),
	}

	return client
}

// Add a torrent (metainfo) from io reader and sub directory
func (c *clientImpl) Add(rd io.Reader, hat string) (err error) {
	uri := c.uri + "/torrents"
	err = doRequest(c.rc.R().
		SetHeader("x-save-path", path.Join(c.root, hat)).
		SetBody(&rd).Post, uri)
	return
}

func (c *clientImpl) Drop(hash string) (err error) {
	uri := c.uri + "/torrent/" + hash
	err = doRequest(c.rc.R().Delete, uri)
	return
}

func (c *clientImpl) DropWithData(hash string) (err error) {
	uri := c.uri + "/torrent/" + hash + "/with_data"
	err = doRequest(c.rc.R().Delete, uri)
	return
}

// Exist check a hash exist
func (c *clientImpl) Exist(hash string) bool {
	uri := c.uri + "/torrent/" + hash
	err := doRequest(c.rc.R().Head, uri)
	return err == nil
}

// GetHashes return all hashes
func (c *clientImpl) GetHashes() (res Hashes) {
	uri := c.uri + "/hashes"
	doRequest(c.rc.R().SetResult(&res).Get, uri)
	return
}

// GetTorrents get all torrents
func (c *clientImpl) GetTorrents() (data []TorrentStatus, err error) {
	uri := c.uri + "/torrents"
	err = doRequest(c.rc.R().SetResult(&data).Get, uri)
	return
}

// GetTorrent get a torrent with hash string
func (c *clientImpl) GetTorrent(hash string) (*TorrentStatus, error) {
	uri := c.uri + "/torrent/" + hash
	var ts TorrentStatus
	if err := doRequest(c.rc.R().SetResult(&ts).Get, uri); err != nil {
		return nil, err
	}
	return &ts, nil
}

func (c *clientImpl) Session() (res *TeSession, err error) {
	uri := c.uri + "/session/info"
	res = new(TeSession)
	err = doRequest(c.rc.R().SetResult(res).Get, uri)
	return
}

func (c *clientImpl) Stats() (res *TeStatistics, err error) {
	uri := c.uri + "/session/stats"
	res = new(TeStatistics)
	err = doRequest(c.rc.R().SetResult(res).Get, uri)
	return
}

type methodFunc func(uri string) (*resty.Response, error)

func doRequest(mf methodFunc, uri string) (err error) {
	var resp *resty.Response
	resp, err = mf(uri)
	if err != nil {
		log.Printf("request fail: uri %q, err %s", uri, err)
		return
	}
	if resp.StatusCode() >= 400 {
		err = fmt.Errorf("request fail: %s", resp.Status())
		return
	}
	return
}
