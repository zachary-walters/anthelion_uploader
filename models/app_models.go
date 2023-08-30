package models

type Config struct {
	AnnounceURL string `json:"announce_url"`
	APIKey      string `json:"api_key"`
}

type ShellCommand struct {
	Command          string            `json:"command"`
	ShellCommandArgs []ShellCommandArg `json:"shell_command_args"`
}

type ShellCommandArg struct {
	Flag string `json:"flag"`
	Arg  string `json:"arg"`
}

type MediainfoSections struct {
	General string   `json:"general"`
	Video   string   `json:"video"`
	Audio   []string `json:"audio"`
	Text    []string `json:"text"`
}
