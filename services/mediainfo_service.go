package services

import (
	"regexp"
	"strings"

	"github.com/zachary-walters/anthelion_api/models"
)

type MediainfoParsingService struct{}

func NewMediainfoParsingService() *MediainfoParsingService {
	return &MediainfoParsingService{}
}

func (svc *MediainfoParsingService) GetFlags(mediainfo string) []string {
	flagMap := map[string]bool{}
	flags := []string{}
	mediainfoSections := models.MediainfoSections{}

	mediainfo, mediainfoSections.General = svc.parseSection(mediainfo, "General")
	mediainfo, mediainfoSections.Video = svc.parseSection(mediainfo, "Video")

	mediainfoAudio := mediainfo
	for {
		mediainfo, audioTrack := svc.parseSection(mediainfoAudio, "Audio")
		if audioTrack == "" {
			break
		}

		mediainfoSections.Audio = append(mediainfoSections.Audio, audioTrack)
		mediainfoAudio = mediainfo
	}

	mediainfoText := mediainfoAudio
	for {
		mediainfo, textChannel := svc.parseSection(mediainfoText, "Text")
		if textChannel == "" {
			break
		}

		mediainfoSections.Text = append(mediainfoSections.Text, textChannel)
		mediainfoText = mediainfo
	}

	remuxRegex := regexp.MustCompile(`^Complete name.*:.*(?i)remux`)
	unratedRegex := regexp.MustCompile(`^Complete name.*:.*(?i)unrated`)
	uncutRegex := regexp.MustCompile(`^Complete name.*:.*(?i)uncut`)
	extendedRegex := regexp.MustCompile(`^Complete name.*:.*(?i)extended`)
	criterionRegex := regexp.MustCompile(`^Complete name.*:.*(?i)criterion`)
	numTracks := len(mediainfoSections.Audio)
	commentaryRegex := regexp.MustCompile(`^Title.*:.*Commentary`)
	atmosRegex := regexp.MustCompile(`^Commercial name.*:.*Atmos`)
	hdr10Regex := regexp.MustCompile(`^HDR.*:.*HDR10`)
	for _, line := range strings.Split(mediainfoSections.General, "\n") {
		if remuxRegex.MatchString(line) {
			flagMap[models.REMUX] = true
		}

		if unratedRegex.MatchString(line) {
			flagMap[models.UNRATED] = true
		}

		if uncutRegex.MatchString(line) {
			flagMap[models.UNCUT] = true
		}

		if extendedRegex.MatchString(line) {
			flagMap[models.EXTENDED] = true
		}

		if criterionRegex.MatchString(line) {
			flagMap[models.CRITERION] = true
		}
	}

	for _, audioChannel := range mediainfoSections.Audio {
		for _, line := range strings.Split(audioChannel, "\n") {
			if commentaryRegex.MatchString(line) {
				flagMap[models.COMMENTARY] = true
				numTracks = numTracks - 1
			}

			if atmosRegex.MatchString(line) {
				flagMap[models.ATMOS] = true
			}
		}
	}

	if numTracks > 1 {
		flagMap[models.DUALAUDIO] = true
	}

	for _, line := range strings.Split(mediainfoSections.Video, "\n") {
		if hdr10Regex.MatchString(line) {
			flagMap[models.HDR10] = true
		}
	}

	for k := range flagMap {
		flags = append(flags, k)
	}

	return flags
}

func (svc *MediainfoParsingService) parseSection(mediainfo, section string) (string, string) {
	start := strings.Index(mediainfo, section)
	end := strings.Index(mediainfo, "\n\n")

	if start > end || start == -1 {
		return mediainfo[end+1 : len(mediainfo)-1], ""
	}

	return mediainfo[end+1 : len(mediainfo)-1], mediainfo[start:end]
}
