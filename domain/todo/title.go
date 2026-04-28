package todo

const MIN_TITLE_LENGTH = 3

type Title struct {
	value string
}

func (t Title) Value() string {
	return t.value
}

func NewTitle(value string) (Title, error) {
	if len(value) < MIN_TITLE_LENGTH {
		return Title{}, ErrTitleTooShort
	}

	return Title{value: value}, nil
}

func TitleFrom(value string) Title {
	return Title{value: value}
}

func (t Title) Equals(other Title) bool {
	return t.value == other.value
}
