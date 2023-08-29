package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/zachary-walters/anthelion_api/app_errors"
	"github.com/zachary-walters/anthelion_api/models"
	"github.com/zachary-walters/anthelion_api/options"
)

type APIService struct {
	client *http.Client
}

// Create new APIService that interacts with anthelion's api
func NewAPIService() *APIService {
	return &APIService{
		client: &http.Client{},
	}
}

// POST Request to upload torrent to anthelion
func (svc *APIService) UploadTorrent(input models.APIInputParameter) (models.APIUploadResponse, error) {
	apiUploadResponse := models.APIUploadResponse{}

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	apiKey, err := w.CreateFormField("api_key")
	if err != nil {
		log.Println("error creating api_key field:", err)
		return apiUploadResponse, err
	}
	apiKey.Write([]byte(input.APIKey))

	action, err := w.CreateFormField("action")
	if err != nil {
		log.Println("error creating action field:", err)
		return apiUploadResponse, err
	}
	action.Write([]byte(input.Action))

	tmdbid, err := w.CreateFormField("tmdbid")
	if err != nil {
		log.Println("error creating tmdbid field:", err)
		return apiUploadResponse, err
	}
	tmdbid.Write([]byte(input.TMDBID))

	mediainfo, err := w.CreateFormField("mediainfo")
	if err != nil {
		log.Println("error creating mediainfo field:", err)
		return apiUploadResponse, err
	}
	mediainfo.Write([]byte(input.Mediainfo))

	fw, err := w.CreateFormFile("file_input", input.FileInput)
	if err != nil {
		log.Println("error creating form file:", err)
		return apiUploadResponse, err
	}

	fd, err := os.Open(input.FileInput)
	if err != nil {
		log.Println("error opening torrent file:", err)
		return apiUploadResponse, err
	}
	defer fd.Close()

	_, err = io.Copy(fw, fd)
	if err != nil {
		log.Println("error copying file data to form field:", err)
		return apiUploadResponse, err
	}
	w.Close()

	req, err := http.NewRequest("POST", options.ANTHELION_API_URL, buf)
	if err != nil {
		log.Println("error creating the POST request for uploading:", err)
		return apiUploadResponse, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := svc.client.Do(req)
	if err != nil {
		log.Println("error sending the POST request for uploading:", err)
		return apiUploadResponse, err
	}

	switch resp.StatusCode {
	case 403:
		err = &app_errors.Error403{}
	case 400:
		err = &app_errors.Error400{}
	case 500:
		err = &app_errors.Error500{}
	}

	if err != nil {
		log.Println(err.Error())
		return apiUploadResponse, err
	}

	log.Println("torrent uploaded successfully with status: ", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body:", err)
		return apiUploadResponse, err
	}

	err = json.Unmarshal(body, &apiUploadResponse)
	if err != nil {
		log.Println("error unmarshalling response body:", err)
		return apiUploadResponse, err
	}

	return apiUploadResponse, nil
}

func (svc *APIService) DownloadTorrent(url, filePath string) error {
	lastPathIdx := strings.LastIndex(filePath, "/")
	fileName := filePath

	if lastPathIdx > 0 {
		fileName = filePath[lastPathIdx+1:]
	}

	fileName = fmt.Sprintf("%s.torrent", fileName)

	resp, err := svc.client.Get(url)
	if err != nil || resp.StatusCode != 200{
		log.Println("error downloading torrent:", err)
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading bytes from response body:", err)
		return err
	}

	fo, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = fo.Write(body)
	if err != nil {
		log.Println("error .writing torrent file:", err)
		return err
	}

	return nil
}
