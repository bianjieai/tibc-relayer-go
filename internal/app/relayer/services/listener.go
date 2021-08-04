package services

import (
	"context"
	"sync"

	"github.com/bianjieai/tibc-relayer-go/internal/app/relayer/repostitory"
	log "github.com/sirupsen/logrus"
)

type IListener interface {
	Listen() error
}

type Listener struct {
	chainMap map[string]repostitory.IChain

	ctxMap sync.Map
	logger *log.Logger
}

func NewScanner(
	chainMap map[string]repostitory.IChain,
	logger *log.Logger) IListener {
	listener := &Listener{
		chainMap: chainMap,
		ctxMap:   sync.Map{},
		logger:   logger,
	}
	return listener
}

func (listener *Listener) Listen() error {

	// 启动N个goroutine去处理
	for chainName, _ := range listener.chainMap {
		ctx, cancel := context.WithCancel(context.Background())
		listener.ctxMap.Store(chainName, cancel)
		go listener.start(ctx, chainName)
	}
	listener.ctxMap.Range(listener.walk)

	return nil
}

func (listener *Listener) start(ctx context.Context, chainName string) {
	// todo
	//
}

func (listener *Listener) cancelCtx(locality string) {
	if value, ok := listener.ctxMap.Load(locality); ok {
		cancel := value.(context.CancelFunc)
		cancel()
		listener.ctxMap.Delete(locality)
	}
}

func (listener *Listener) walk(key, value interface{}) bool {
	listener.logger.WithFields(log.Fields{"chain": key}).Info("start")
	return true
}
