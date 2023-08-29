package models

type APIInputParameter struct {
	APIKey    string `json:"api_key"`    // your API key
	Action    string `json:"action"`     // "upload", "download", or "search"
	TMDBID    string `json:"tmdbid"`     // film identifier from TMDb
	Mediainfo string `json:"mediainfo"`  // newline-delimited mediainfo of the primary video file in the torrent
	FileInput string `json:"file_input"` // the .torrent file
}

type APIUploadResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Download string `json:"download"` // torrent download url
}
