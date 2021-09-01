package kedge

// InfoHash hex string
type InfoHash string

// Hashes slice of InfoHash
type Hashes []InfoHash

// State value of torrentStatus.state
type State int16

// const of state, value from libtorrent/torrent_status.hpp
const (
	QueuedForChecking   State = iota
	CheckingFiles             // 1
	DownloadingMetadata       // 2
	Downloading               // 3
	Finished                  // 4
	Seeding                   // 5
	Allocating                // 6
	CheckingResumeData        // 7
)

// TorrentStatus ... Ratio = allTimeUploaded / TotalDone
type TorrentStatus struct {
	ActiveDuration      int64    `json:"active_duration,omitempty"`
	FinishedDuration    int64    `json:"finished_duration,omitempty"`
	SeedingDuration     int64    `json:"seeding_duration,omitempty"`
	AddedTime           int64    `json:"added_time,omitempty"`     // unix timestamp
	CompletedTime       int64    `json:"completed_time,omitempty"` // unix timestamp
	TotalDone           uint64   `json:"total_done,omitempty"`
	TotalWanted         uint64   `json:"total_wanted,omitempty"`
	TotalWantedDone     uint64   `json:"total_wanted_done,omitempty"`
	TotalDownloaded     uint64   `json:"total_downloaded,omitempty"`         // this session
	TotalUploaded       uint64   `json:"total_uploaded,omitempty"`           // this session
	TotalDataDownloaded uint64   `json:"total_payload_downloaded,omitempty"` // this session
	TotalDataUploaded   uint64   `json:"total_payload_uploaded,omitempty"`   // this session
	AllTimeDownloaded   uint64   `json:"all_time_download,omitempty"`
	AllTimeUploaded     uint64   `json:"all_time_upload,omitempty"`
	DownloadPayloadRate int      `json:"download_payload_rate,omitempty"`
	UploadPayloadRate   int      `json:"upload_payload_rate,omitempty"`
	ConnectCandidates   int      `json:"connect_candidates,omitempty"`
	NumConnections      int      `json:"num_connections,omitempty"`
	ListSeeds           int      `json:"list_seeds,omitempty"`
	ListPeers           int      `json:"list_peers,omitempty"`
	NumPieces           int      `json:"num_pieces,omitempty"`
	Infohash            InfoHash `json:"info_hash"`
	State               State    `json:"state"`
	Name                string   `json:"name"`
	SavePath            string   `json:"save_path"`
	CurrentTracker      string   `json:"current_tracker,omitempty"`
	Progress            float32  `json:"progress,omitempty"`
	ProgressPpm         int      `json:"progress_ppm,omitempty"`
	IsFinished          bool     `json:"is_finished,omitempty"`
	IsSeeding           bool     `json:"is_seeding,omitempty"`
}

// GetRatio return a ratio of shared bytes
func (ts *TorrentStatus) GetRatio() float64 {
	if ts.TotalDone > 0 {
		return float64(ts.AllTimeUploaded) / float64(ts.TotalDone)
	}
	return -1.0
}

// GetETA return seconds of remaining
func (ts *TorrentStatus) GetETA() int64 {
	if ts.DownloadPayloadRate > 0 {
		left := ts.TotalWanted - ts.TotalWantedDone
		if left > 0 {
			return int64(left) / int64(ts.DownloadPayloadRate)
		}
	}
	return 0
}

// TeSession represent session basic infomation
type TeSession struct {
	PeerPort uint16 `json:"peerPort,omitempty"`
	PeerID   string `json:"peerID,omitempty"`
	UptimeMs int64  `json:"uptimeMs,omitempty"` // Milliseconds
	Version  string `json:"version,omitempty"`

	// TODO: more configurable fields
}

// TeStatistics represent session statistics
type TeStatistics struct {
	BytesRecv       uint64 `json:"bytesRecv,omitempty"`
	BytesSent       uint64 `json:"bytesSent,omitempty"`
	BytesDataRecv   uint64 `json:"bytesDataRecv,omitempty"`
	BytesDataSent   uint64 `json:"bytesDataSent,omitempty"`
	RateRecv        uint64 `json:"rateRecv,omitempty"`
	RateSent        uint64 `json:"rateSent,omitempty"`
	BytesFailed     uint64 `json:"bytesFailed,omitempty"`
	BytesQueued     uint64 `json:"bytesQueued,omitempty"`
	BytesWasted     uint64 `json:"bytesWasted,omitempty"`
	UptimeMs        uint64 `json:"uptimeMs,omitempty"`
	Uptime          uint32 `json:"uptime,omitempty"`
	NumConnected    uint16 `json:"numPeersConnected,omitempty"`
	NumHalfOpen     uint16 `json:"numPeersHalfOpen,omitempty"`
	TaskCount       uint16 `json:"taskCount,omitempty"`
	ActiveTaskCount uint16 `json:"activeCount,omitempty"`
	PausedTaskCount uint16 `json:"pausedCount,omitempty"`
	Version         string `json:"version,omitempty"`
}
