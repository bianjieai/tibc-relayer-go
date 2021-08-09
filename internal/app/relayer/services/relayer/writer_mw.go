package relayer

import (
	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/domain"
	log "github.com/sirupsen/logrus"
)

var _ IRelayer = new(Writer)

type Writer struct {
	next IRelayer

	logger *log.Entry

	chainName string

	cacheWriter *domain.CacheFileWriter
}

func NewWriterMW(svc IRelayer, chainName string, logger *log.Logger, dir, filename string) IRelayer {

	entry := logger.WithFields(log.Fields{
		"chain_name": chainName,
	})
	cacheWriter := domain.NewCacheFileWriter(dir, filename)
	return &Writer{
		next:        svc,
		chainName:   chainName,
		cacheWriter: cacheWriter,
		logger:      entry,
	}
}

func (w *Writer) UpdateClient() error {
	return w.next.UpdateClient()
}

type cacheData struct {
	LatestHeight uint64 `json:"latest_height"`
}

func (w *Writer) PendingDatagrams() error {
	err := w.next.PendingDatagrams()
	if err == nil {
		return nil
	}
	ctx := w.next.Context()
	defer w.cacheWriter.Write(ctx.Height())
	return err
}

func (w *Writer) IsNotRelay() bool {
	return w.next.IsNotRelay()
}

func (w *Writer) Context() *domain.Context {
	return w.next.Context()
}
