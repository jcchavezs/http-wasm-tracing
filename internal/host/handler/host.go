package handler

import (
	"context"

	apihandler "github.com/jcchavezs/http-wasm-tracing/internal/host/api/handler"
	"github.com/jcchavezs/http-wasm-tracing/trace"
)

var Extractor = func(ctx context.Context) (trace.Span, bool) {
	return nil, false
}

type host struct{}

var _ apihandler.Host = host{}

func (host) SetSpanAttribute(ctx context.Context, key, value string) {
	if s, ok := Extractor(ctx); ok {
		s.SetAttribute(key, value)
	}
}
