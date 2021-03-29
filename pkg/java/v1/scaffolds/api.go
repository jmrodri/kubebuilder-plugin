package scaffolds

import (
	"fmt"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/model/resource"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugins"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

)

type apiScaffolder struct {
	fs machinery.Filesystem

	config   config.Config
	resource resource.Resource

	doRole, doPlaybook bool
}

// NewCreateAPIScaffolder returns a new plugins.Scaffolder for project initialization operations
func NewCreateAPIScaffolder(cfg config.Config, res resource.Resource, doRole, doPlaybook bool) plugins.Scaffolder {
	return &apiScaffolder{
		config:     cfg,
		resource:   res,
		doRole:     doRole,
		doPlaybook: doPlaybook,
	}
}

func (s *apiScaffolder) InjectFS(fs machinery.Filesystem) {
	s.fs = fs
}

func (s *apiScaffolder) Scaffold() error {
	fmt.Println("api.Scaffold()")
	return nil
}
