package mock

import (
	"github.com/khulnasoft-lab/harbor-scanner-vul/pkg/harbor"
	"github.com/khulnasoft-lab/harbor-scanner-vul/pkg/vul"
	"github.com/stretchr/testify/mock"
)

type Transformer struct {
	mock.Mock
}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) Transform(artifact harbor.Artifact, source []vul.Vulnerability) harbor.ScanReport {
	args := t.Called(artifact, source)
	return args.Get(0).(harbor.ScanReport)
}
