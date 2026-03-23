package middleware

import "net/http"

// CORS adds cross-origin resource sharing headers.
// In development (appURL = "http://localhost"), it allows all origins.
// In production, only the specified appURL is allowed.
func CORS(appURL string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// Allow the configured app URL origin, or localhost for development
			allowed := origin == appURL ||
				origin == "http://localhost" ||
				origin == "http://localhost:3000" ||
				origin == "http://localhost:5173"

			if allowed && origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
