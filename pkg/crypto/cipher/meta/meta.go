// Package meta provides interfaces and implementations for working with
// cipher metadata.
// metadata is required to write the cipher to a file and display the name
// for the user.
package meta

// IMeta defines the interface for cipher metadata.
type IMeta interface {
	// Code returns cipher code.
	Code() int
	// Name returns cipher name.
	Name() string
}

type Meta struct {
	code int
	name string
}

func (m Meta) Code() int {
	return m.code
}

func (m Meta) Name() string {
	return m.name
}

// NewMeta creates a new Meta instance with the given code and name.
func NewMeta(code int, name string) Meta {
	return Meta{code: code, name: name}
}

// IMetaProvider defines the interface for providing cipher metadata.
type IMetaProvider interface {
	// Meta returns metadata.
	Meta() IMeta
}

type MetaProvider struct {
	meta IMeta
}

func (mp MetaProvider) Meta() IMeta {
	if mp.meta == nil {
		mp.meta = NewMeta(0, "")
	}
	return mp.meta
}

// NewMetaProvider creates a new MetaProvider instance with the given
// IMeta instance.
func NewMetaProvider(m IMeta) MetaProvider {
	return MetaProvider{meta: m}
}
