package persistence

import (
	"github.com/khulnasoft-lab/harbor-scanner-vul/pkg/harbor"
	"github.com/khulnasoft-lab/harbor-scanner-vul/pkg/job"
)

type Store interface {
	Create(scanJob job.ScanJob) error
	Get(scanJobID string) (*job.ScanJob, error)
	UpdateStatus(scanJobID string, newStatus job.ScanJobStatus, error ...string) error
	UpdateReport(scanJobID string, report harbor.ScanReport) error
}
