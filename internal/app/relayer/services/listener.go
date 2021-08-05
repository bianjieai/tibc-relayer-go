package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const DefaultTimeout = 10

type IListener interface {
	Listen() error
}

type Listener struct {
	relayerMap map[string]IRelayer

	ctxMap sync.Map
	logger *log.Logger
}

func NewScanner(
	relayerMap map[string]IRelayer,
	logger *log.Logger) IListener {
	listener := &Listener{
		relayerMap: relayerMap,
		ctxMap:     sync.Map{},
		logger:     logger,
	}
	return listener
}

func (listener *Listener) Listen() error {

	// 启动N个goroutine去处理
	for chainName, _ := range listener.relayerMap {
		ctx, cancel := context.WithCancel(context.Background())
		listener.ctxMap.Store(chainName, cancel)
		go listener.start(ctx, chainName)
	}
	listener.ctxMap.Range(listener.walk)
	select {}
}

func (listener *Listener) start(ctx context.Context, chainName string) {

	for {
		select {
		case <-ctx.Done():
			listener.logger.WithFields(log.Fields{
				"chain_name": chainName,
			}).Info("canceled")
			return
		default:
			if !listener.relayerMap[chainName].IsNotRelay() {
				time.Sleep(DefaultTimeout * time.Second)
			} else {
				listener.relayerMap[chainName].PendingDatagrams()
			}

		}
	}
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
