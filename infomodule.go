package m1

type InfoModule struct {
	*Module
}

func newInfoModule(client *client, info ModuleInfo, msysVersion Version) (*InfoModule, error) {
	i, err := newModule(client, "INFO", info, msysVersion)
	if err != nil {
		return nil, err
	}

	return &InfoModule{i}, nil
}
