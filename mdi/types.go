package mdi

const (
	Version = "v7.4.47"
)

type Icon struct {
	id         string
	baseIconID string
	name       string
	tags       []string
	author     string
	version    string
	deprecated bool
	data       []byte
}

func (i Icon) ID() string {
	return i.id
}

func (i Icon) BaseIconID() string {
	return i.baseIconID
}

func (i Icon) Name() string {
	return i.name
}

func (i Icon) Tags() []string {
	cp := make([]string, len(i.tags))
	copy(cp, i.tags)
	return cp
}

func (i Icon) Author() string {
	return i.author
}

func (i Icon) Version() string {
	return i.version
}

func (i Icon) Deprecated() bool {
	return i.deprecated
}

func (i Icon) Data() []byte {
	cp := make([]byte, len(i.data))
	copy(cp, i.data)
	return cp
}
