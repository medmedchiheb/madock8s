package pkg

import (
	"github.com/MaibornWolff/maDocK8s/core/types/exporter/notifier"
	"github.com/MaibornWolff/maDocK8s/core/types/services/mdstorage"
	"github.com/MaibornWolff/maDocK8s/exporter/gitlab/adapter"
	"github.com/sirupsen/logrus"
)

type service struct {
	templateFile string
	storage      mdstorage.MdstorageServer
	mdFiles      func(map[string]string) ([]adapter.GitlabFile, error)
	logger       *logrus.Entry
	name         string
}

func NewService(name string, storage mdstorage.MdstorageServer, mdFiles func(map[string]string) ([]adapter.GitlabFile, error), logger *logrus.Entry) notifier.NotifierServer {
	template := "/etc/gitlab/static/template.md"

	return &service{
		templateFile: template,
		storage:      storage,
		mdFiles:      mdFiles,
		logger:       logger,
		name:         name,
	}
}
