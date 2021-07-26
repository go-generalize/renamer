package renamer

// Renamer renames struct names for go2X to retrieve a unique name for structs
type Renamer struct {
	used        map[string]struct{}
	renamed     map[string]string
	defaultOpts []Option
}

// New returns a new Renamer.
func New(defaultOpts ...Option) *Renamer {
	return &Renamer{
		used:        make(map[string]struct{}),
		renamed:     map[string]string{},
		defaultOpts: defaultOpts,
	}
}

// Renamed returns a name for id.
// For idempotency, this should be called in the same order.
// Logic:
// If id is already converted, the previously converted name will be returned.
// If preferred is not taken, preferred is returned.
// If preferred is taken, another name calculated by options will be returned.
// defaultOpts follows options.
// If no opts can generate unused name, `preferred+_+hex(sha256(id))[>=3...]` string will be returned.
func (r *Renamer) Renamed(id string, preferred string, options ...Option) (name string) {
	if v, ok := r.renamed[id]; ok {
		return v
	}
	defer func() {
		r.renamed[id] = name

		_, used := r.used[name]
		if used {
			panic(name + " is already used")
		}

		r.used[name] = struct{}{}
	}()

	if _, ok := r.used[preferred]; !ok {
		return preferred
	}

	annotations := make(map[string]string)
	for _, opt := range options {
		name, ok := opt(id, preferred, annotations)

		if !ok {
			continue
		}

		if _, ok := r.used[name]; ok {
			continue
		}

		return name
	}

	for _, opt := range r.defaultOpts {
		name, ok := opt(id, preferred, annotations)

		if !ok {
			continue
		}

		if _, ok := r.used[name]; ok {
			continue
		}

		return name
	}

	hash := sha256(id)
	for i := 3; i < len(hash); i++ {
		name := preferred + "_" + hash[:i]

		if _, ok := r.used[name]; ok {
			continue
		}

		return name
	}

	return
}

func (r *Renamer) Find(id string) (string, bool) {
	renamed, ok := r.renamed[id]

	return renamed, ok
}
