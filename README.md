# kedgecli

Client library of [kedge](https://github.com/liut/kedge) with Golang

## interfaces

```go

type ClientI interface {
	Add(rd io.Reader, hat string) error
	Drop(hash string) error
	DropWithData(hash string) error
	Exist(hash string) bool
	GetHashes() Hashes
	GetTorrents() ([]TorrentStatus, error)
	GetTorrent(hash string) (*TorrentStatus, error)
	Session() (*TeSession, error)
	Stats() (*TeStatistics, error)
}

```

## usage

```go

	root := "/var/lib/store/root"
	uri := "http://localhost:16180"
	c := kedgecli.New(root, uri)

	// read metainfo as mi *MetaInfo
	var buf bytes.Buffer
	if err := mi.Write(&buf); err != nil {
		return
	}
	err := c.Add(buf, "subdir")

```
