package main

import (
	"flag"
	"log"

	"github.com/zachary-walters/anthelion_api/models"
	"github.com/zachary-walters/anthelion_api/options"
	"github.com/zachary-walters/anthelion_api/services"
)

func main() {
	var filePath string
	var tmdbID string
	var announceUrl string
	var apiKey string

	flag.StringVar(&filePath, "p", "", "file path to the video file (-p)")
	flag.StringVar(&tmdbID, "t", "", "TMDB ID (-t)")
	flag.StringVar(&announceUrl, "u", "", "Announce URL (-u)")
	flag.StringVar(&apiKey, "k", "", "API Key (-k)")
	flag.Parse()

	if filePath == "" {
		log.Println("A file path (-p) flag is required.")
		return
	}

	if tmdbID == "" {
		log.Println("TMDB ID (-t) flag is required.")
		return
	}

	options.LoadEnv()

	if options.ANTHELION_API_URL == "" {
		log.Println("Anthelion's API URL must be set as en environment variable")
		return
	}

	if announceUrl != "" {
		options.ANNOUNCE_URL = announceUrl
	}

	if apiKey != "" {
		options.API_KEY = apiKey
	}

	if options.ANNOUNCE_URL == "" || options.API_KEY == "" {
		log.Println("APIKey and/or Announce URL are missing. Aborting.")
		return
	}

	shellCommandService := services.NewShellCommandService()

	// CREATE .TORRENT FILE
	torrentPath, err := shellCommandService.MakeTorrentFile(filePath)
	if err != nil {
		return
	}
	log.Println("Creating .torrent file DONE.")

	// GENERATE MEDIAINFO
	mediainfo, err := shellCommandService.GetMediaInfo(filePath)
	if err != nil {
		return
	}
	log.Println("Generating Mediainfo DONE.")

	// UPLOAD TORRENT
	log.Println("Uploading Torrent ...")
	apiParams := models.APIInputParameter{
		APIKey:    options.API_KEY,
		Action:    "upload",
		TMDBID:    tmdbID,
		Mediainfo: mediainfo,
		FileInput: torrentPath,
	}

	apiService := services.NewAPIService()
	apiUploadResponse, err := apiService.UploadTorrent(apiParams)
	if err != nil {
		return
	}
	log.Println("Uploading Torrent DONE.")

	err = apiService.DownloadTorrent(apiUploadResponse.Download, filePath)
	if err != nil {
		return
	}
}
