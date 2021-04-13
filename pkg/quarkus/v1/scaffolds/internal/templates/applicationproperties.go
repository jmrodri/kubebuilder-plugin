package templates

import (
	"fmt"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/java-operator-sdk/kubebuilder-plugin/pkg/quarkus/v1/scaffolds/internal/templates/util"
)

var _ machinery.Template = &ApplicationPropertiesFile{}

type ApplicationPropertiesFile struct {
	machinery.TemplateMixin
	OrgName     string
	ProjectName string
}

func (f *ApplicationPropertiesFile) SetTemplateDefaults() error {
	if f.ProjectName == "" {
		return fmt.Errorf("invalid Application Properties name")
	}

	if f.Path == "" {
		f.Path = util.PrependResourcePath("application.properties")
	}

	f.TemplateBody = ApplicationPropertiesTemplate

	return nil
}

// TODO: pass in the name of the operator i.e. replace Memcached
const ApplicationPropertiesTemplate = `quarkus.container-image.build=true
#quarkus.container-image.group=
quarkus.container-image.name={{ .ProjectName }}-service
`
