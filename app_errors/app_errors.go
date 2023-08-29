package app_errors

type MaketorrentError struct {}
type MediainfoError struct{}

func (e *MaketorrentError) Error() string {
	return "err: unable to create torrent file"
}

func (e *MediainfoError) Error() string {
	return "err: unable to obtain mediainfo"
}

// CUSTOM ERRORS FOR HTTP REQUESTS
type Error403 struct{}
type Error400 struct{}
type Error500 struct{}

func (e *Error403) Error() string {
	return "problems with your key or uploading privileges"
}

func (e *Error400) Error() string {
	return "missing parameters or otherwise something wrong with your upload (you did something wrong)"
}

func (e *Error500) Error() string {
	return "internal server error (our dev did something wrong - please give us a bug report!)"
}
