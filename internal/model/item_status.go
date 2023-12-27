package model

import "errors"

type Status int

const (
	Unset Status = iota
	NotOk
	Meh
	Ok
)

func (s Status) String() string {
	switch s {
	case Unset:
		return "unset"
	case NotOk:
		return "not-ok"
	case Meh:
		return "meh"
	case Ok:
		return "ok"
	}
	return "unknown"
}

func ParseStatus(str string) (Status, error) {
	switch str {
	case "unset":
		return 0, nil
	case "not-ok":
		return 1, nil
	case "meh":
		return 2, nil
	case "ok":
		return 3, nil
	}

	return -1, errors.New("unknown status")
}
