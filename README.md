# kedgecli

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
