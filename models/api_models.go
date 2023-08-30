package models

// flags
const (
	DIRECTORS   = "Directors"
	EXTENDED    = "Extended"
	UNCUT       = "Uncut"
	UNRATED     = "Unrated"
	HDR10       = "HDR10"
	DV          = "DV"
	_4KREMASTER = "4KRemaster"
	ATMOS       = "Atmos"
	DUALAUDIO   = "DualAudio"
	COMMENTARY  = "Commentary"
	REMUX       = "Remux"
	_3D         = "3D"
	CRITERION   = "Criterion"
)

type APIInputParameter struct {
	APIKey    string   `json:"api_key"`    // your API key
	Action    string   `json:"action"`     // "upload", "download", or "search"
	TMDBID    string   `json:"tmdbid"`     // film identifier from TMDb
	Mediainfo string   `json:"mediainfo"`  // newline-delimited mediainfo of the primary video file in the torrent
	FileInput string   `json:"file_input"` // the .torrent file
	Flags     []string `json:"flags[]"`    // extra flags
}

type APIUploadResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Download string `json:"download"` // torrent download url
}
