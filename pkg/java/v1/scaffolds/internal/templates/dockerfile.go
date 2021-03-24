package templates

import "sigs.k8s.io/kubebuilder/v3/pkg/model/file"

var _ file.Template = &DockerFile{}

type DockerFile struct {
	file.TemplateMixin
}

func (f *DockerFile) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = "Dockerfile"
	}

	f.TemplateBody = dockerfileTemplate

	return nil
}

// TODO: pass in the name of the operator i.e. replace Memcached
const dockerfileTemplate = `#FROM quay.io/operator-framework/java-operator:{{ .JavaOperatorVersion }}

#COPY requirements.yml ${HOME}/requirements.yml
#RUN ansible-galaxy collection install -r ${HOME}/requirements.yml \
#&& chmod -R ug+rwx ${HOME}/.ansible

#COPY watches.yaml ${HOME}/watches.yaml
#COPY {{ .RolesDir }}/ ${HOME}/{{ .RolesDir }}/
#COPY {{ .PlaybooksDir }}/ ${HOME}/{{ .PlaybooksDir }}/
`

