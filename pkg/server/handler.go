package server

import (
	"exchangeservice/pkg/server/gen"
	"exchangeservice/pkg/storage"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.elastic.co/apm/v2"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.Logger
	storage storage.Storage
}

func (h Handler) Health(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}

func (h Handler) Exchange(ctx echo.Context, from string, to string, amount float32) error {
	span, apmCtx := apm.StartSpan(ctx.Request().Context(), "Exhange", "request")
	defer span.End()

	rate := h.storage.GetByCurrency(apmCtx, from, to)
	if rate == nil {
		h.logger.Debug("currency not found", zap.String("from", from), zap.String("to", to))
		return ctx.NoContent(http.StatusNotFound)
	}

	total := rate.Total(amount)
	output := gen.ExchangeResult{
		Total: total,
	}
	return ctx.JSON(http.StatusOK, output)
}

func NewHandler(logger *zap.Logger, storage storage.Storage) Handler {
	return Handler{
		logger:  logger,
		storage: storage,
	}
}
