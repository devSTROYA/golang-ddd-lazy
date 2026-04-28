package user

const MIN_NAME_LENGTH = 3

type Name struct {
	value string
}

func (n Name) Value() string {
	return n.value
}

func NewName(value string) (Name, error) {
	if len(value) < MIN_NAME_LENGTH {
		return Name{}, ErrNameTooShort
	}

	return Name{value: value}, nil
}

func NameFrom(value string) Name {
	return Name{value: value}
}

func (n Name) Equals(other Name) bool {
	return n.value == other.value
}
