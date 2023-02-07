package http

import (
	"aresbot/pkg/logger"
	"compress/gzip"
	"crypto/tls"
	"github.com/andybalholm/brotli"
	"io"
	"net/http"
	"net/url"
)

type GzipReadCloser struct {
	*gzip.Reader
	io.Closer
}

type BrotliReadCloser struct {
	*brotli.Reader
	io.Closer
}

func (gz GzipReadCloser) Close() error {
	return gz.Closer.Close()
}

func (br BrotliReadCloser) Close() error {
	return br.Closer.Close()
}

func NewHttpClientWithTransportOptions(cipherSuites []uint16, minTLSVersion uint16, maxTLSVersion uint16) *http.Client {
	tlsConfig := &tls.Config{
		CipherSuites: cipherSuites,
	}

	tlsConfig.PreferServerCipherSuites = true
	tlsConfig.MinVersion = minTLSVersion
	tlsConfig.MaxVersion = maxTLSVersion

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	logger.InfoLogger.Print("setup tls: %v %d %d", cipherSuites, minTLSVersion, maxTLSVersion)

	return &http.Client{Transport: tr}
}

func NewHttpClientWithTransportOptionsAndProxy(cipherSuites []uint16, minTLSVersion uint16, maxTLSVersion uint16, proxyString string) *http.Client {
	tlsConfig := &tls.Config{
		CipherSuites: cipherSuites,
	}

	tlsConfig.PreferServerCipherSuites = true
	tlsConfig.MinVersion = minTLSVersion
	tlsConfig.MaxVersion = maxTLSVersion

	proxyUrl, err := url.Parse(proxyString)

	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	if err != nil {
		logger.ErrorLogger.Print(err, "can not parse proxy string to proxy url %s", proxyString)
	} else {
		tr.Proxy = http.ProxyURL(proxyUrl)
	}

	logger.InfoLogger.Print("setup http client with proxy, cipher suites and tls: %v %d %d", cipherSuites, minTLSVersion, maxTLSVersion)

	return &http.Client{Transport: tr}
}
