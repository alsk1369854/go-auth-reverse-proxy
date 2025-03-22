package utils

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ===== Proxy 建立器 =====
func CreateProxy(target string) *httputil.ReverseProxy {
	parsedURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("解析 URL 錯誤: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = parsedURL.Host
	}

	return proxy
}

// ===== HTTP → HTTPS 重導函式 =====
func RedirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.RequestURI()
	http.Redirect(w, r, target, http.StatusMovedPermanently)
}