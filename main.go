package main

import (
	"UrlScan/internal"
	"net/http"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	atomicLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	var encoder zapcore.Encoder
	encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		atomicLevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()

	mux := http.NewServeMux()
	handler := internal.NewHandler(logger)
	mux.Handle("/scan", handler)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	logger.Info("Starting server", zap.String("port", ":8080"))
	server.ListenAndServe()
}

func loggingMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handling request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
		)

		_, pattern := next.(*http.ServeMux).Handler(r)
		if pattern == "" {
			logger.Warn("Route not found",
				zap.String("path", r.URL.Path),
			)
		}

		next.ServeHTTP(w, r)
	})
}
