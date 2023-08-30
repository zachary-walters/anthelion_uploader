package services

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/zachary-walters/anthelion_api/app_errors"
	"github.com/zachary-walters/anthelion_api/models"
	"github.com/zachary-walters/anthelion_api/options"
)

type ShellCommandService struct{}

// Create New ShellCommandService
func NewShellCommandService() *ShellCommandService {
	return &ShellCommandService{}
}

// Creates the torrent file and puts it into /tmp/
func (svc *ShellCommandService) MakeTorrentFile(filePath string) (string, error) {
	path, err := svc.findBinaryPath("mktorrent")
	if err != nil {
		log.Println("error finding mktorrent binary:", err)
		return "", err
	}

	filePath = strings.ReplaceAll(filePath, " ", "\\ ")

	lastPathIdx := strings.LastIndex(filePath, "/")
	fileName := filePath

	if lastPathIdx > 0 {
		fileName = filePath[lastPathIdx:]
	}

	torrentPath := strings.Trim(fmt.Sprintf("/tmp%s.torrent", strings.ReplaceAll(strings.ReplaceAll(fileName, "\\", ""), " ", ".")), " ")

	cmd := models.ShellCommand{
		Command: strings.Trim(path, "\n"),
		ShellCommandArgs: []models.ShellCommandArg{
			{
				Flag: "-p",
			},
			{
				Flag: "-l",
				Arg:  "24",
			},
			{
				Flag: "-a",
				Arg:  options.ANNOUNCE_URL,
			},
			{
				Arg: filePath,
			},
			{
				Flag: "-o",
				Arg:  torrentPath,
			},
		},
	}

	makeTorrentOut, err := svc.execCommand(cmd)
	if err != nil {
		log.Println("error executing mktorrent:", err)
	}

	// This check actually confirms if the torrent file was written properly
	err = func(out string) error {
		if !strings.Contains(out, "Writing metainfo file... done.") {
			return &app_errors.MaketorrentError{}
		}

		return nil
	}(makeTorrentOut)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return torrentPath, nil
}

func (svc *ShellCommandService) GetMediaInfo(filePath string) (string, error) {
	path, err := svc.findBinaryPath("mediainfo")
	if err != nil {
		log.Println("error finding mediainfo binary:", err)
		return "", err
	}

	filePath = strings.ReplaceAll(filePath, " ", "\\ ")

	cmd := models.ShellCommand{
		Command: strings.Trim(path, "\n"),
		ShellCommandArgs: []models.ShellCommandArg{
			{
				Arg: filePath,
			},
		},
	}

	mediainfo, err := svc.execCommand(cmd)
	if err != nil {
		log.Println("error executing mediainfo:", err)
		return "", err
	}

	err = func(m string) error {
		if m == "" {
			return &app_errors.MediainfoError{}
		}
		return nil
	}(mediainfo)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return mediainfo, nil
}

func (svc *ShellCommandService) findBinaryPath(binaryName string) (string, error) {
	path, err := svc.execCommand(models.ShellCommand{
		Command: "which",
		ShellCommandArgs: []models.ShellCommandArg{
			{
				Arg: binaryName,
			},
		},
	})
	if err != nil {
		log.Printf("error finding binary path for %s", binaryName)
		log.Println(err.Error())
	}

	return path, err
}

func (svc *ShellCommandService) execCommand(command models.ShellCommand) (string, error) {
	args := []string{
		command.Command,
	}
	for _, a := range command.ShellCommandArgs {
		arg := ""
		if a.Flag != "" && a.Arg != "" {
			arg = fmt.Sprintf("%s %s", a.Flag, a.Arg)
		} else if a.Flag != "" {
			arg = a.Flag
		} else if a.Arg != "" {
			arg = a.Arg
		}

		args = append(args, arg)
	}

	cmd := exec.Command("bash", "-c", strings.Join(args, " "))

	stdout, err := cmd.Output()
	if err != nil {
		log.Println("error executing command:", err)
		return "", err
	}

	return string(stdout), nil
}
