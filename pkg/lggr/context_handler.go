package lggr

import (
	"context"
	"log/slog"
	"sso/pkg/base"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	actionID := ctx.Value(base.CtxActionID)
	if actionID != nil {
		r.AddAttrs(slog.String(base.CtxActionID, actionID.(string)))
	}
	return h.Handler.Handle(ctx, r)
}
