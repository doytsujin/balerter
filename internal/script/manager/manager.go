package manager

import (
	"github.com/balerter/balerter/internal/config"
	folderProvider "github.com/balerter/balerter/internal/script/provider/folder"
	"github.com/balerter/balerter/internal/script/script"
)

type Provider interface {
	Get() ([]*script.Script, error)
}

type Manager struct {
	providers map[string]Provider
}

func New() *Manager {
	m := &Manager{
		providers: make(map[string]Provider),
	}

	return m
}

func (m *Manager) Init(cfg config.ScriptsSources) error {

	for _, folderConfig := range cfg.Folder {
		m.providers[folderConfig.Name] = folderProvider.New(folderConfig)
	}

	return nil
}

func (m *Manager) Get() ([]*script.Script, error) {
	ss := make([]*script.Script, 0)

	for _, p := range m.providers {
		s, err := p.Get()
		if err != nil {
			return nil, err
		}

		ss = append(ss, s...)
	}

	return ss, nil
}
