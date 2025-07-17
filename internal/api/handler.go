package api

import (
	"GoAI/internal/core"
	"GoAI/internal/models"
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIHandler handles API requests.
type APIHandler struct {
	service *core.Service
	logger  *slog.Logger
}

// NewAPIHandler creates a new APIHandler.
func NewAPIHandler(service *core.Service, logger *slog.Logger) *APIHandler {
	return &APIHandler{
		service: service,
		logger:  logger,
	}
}

// GenerateHandler handles the /api/v1/generate endpoint.
func (h *APIHandler) GenerateHandler(c *gin.Context) {
	var req models.GenerateRequest
	if c.ShouldBindJSON(&req) != nil {
		// Handle error
	}

	// Force stream mode by default to handle long responses correctly
	req.Stream = true
	h.streamedGenerate(c, &req)
}

func (h *APIHandler) nonStreamedGenerate(c *gin.Context, req *models.GenerateRequest) {
	completion, err := h.service.Generate(c.Request.Context(), req, nil)
	if err != nil {
		h.logger.Error("failed to generate completion", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate completion"})
		return
	}
	c.JSON(http.StatusOK, models.GenerateResponse{Text: completion})
}

func (h *APIHandler) streamedGenerate(c *gin.Context, req *models.GenerateRequest) {
	c.Header("Content-Type", "text/event-stream; charset=utf-8")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	writer := &streamWriter{c.Writer}
	c.Stream(func(w io.Writer) bool {
		_, err := h.service.Generate(c.Request.Context(), req, writer)
		if err != nil {
			h.logger.Error("error during stream generation", "error", err)
			// Since we are streaming, we can't set a JSON error response here.
			// The error is logged, and the stream will just close.
		}
		return false // End the stream
	})
}

// streamWriter is a helper to flush data after each write.
type streamWriter struct {
	io.Writer
}

func (w *streamWriter) Write(p []byte) (n int, err error) {
	n, err = w.Writer.Write(p)
	if err != nil {
		return n, err
	}
	if f, ok := w.Writer.(http.Flusher); ok {
		f.Flush()
	}
	return
}

// HealthCheckHandler handles the /health endpoint.
func (h *APIHandler) HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
