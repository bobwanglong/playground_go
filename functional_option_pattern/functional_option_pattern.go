package main

import "go.uber.org/zap"

type options struct {
	cache  bool
	logger *zap.Logger
}

type Option interface {
	apply(*options)
}

type cacheOption bool

func (c cacheOption) apply(opts *options) {
	opts.cache = bool(c)
}
func WithCache(c bool) Option {
	return cacheOption(c)
}

type loggerOption struct {
	log *zap.Logger
}

func (l loggerOption) apply(opts *options) {
	opts.logger = l.log
}

func WithLogger(log *loggerOption) Option {
	return loggerOption{log: log.log}
}

func Open(addr string, opts ...options) (*Connection, error) {
	// options := options{
	// 	cache: ,
	// }
	options := options{
		cache:  defaultCache,
		logger: zap.NewNop(),
	}
	for _, o := range opts {
		o.apply(&options)
	}

}
func main() {

}
