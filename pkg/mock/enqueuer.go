package mock

import (
	"github.com/khulnasoft-lab/harbor-scanner-vul/pkg/harbor"
	"github.com/khulnasoft-lab/harbor-scanner-vul/pkg/job"
	"github.com/stretchr/testify/mock"
)

type Enqueuer struct {
	mock.Mock
}

func NewEnqueuer() *Enqueuer {
	return &Enqueuer{}
}

func (em *Enqueuer) Enqueue(request harbor.ScanRequest) (job.ScanJob, error) {
	args := em.Called(request)
	return args.Get(0).(job.ScanJob), args.Error(1)
}
