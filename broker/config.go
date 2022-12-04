package broker

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/fhmq/hmq/logger"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"io/ioutil"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Worker  int     `json:"workerNum"`
	Host    string  `json:"host"`
	Port    string  `json:"port"`
	Router  string  `json:"router"`
	TlsHost string  `json:"tlsHost"`
	TlsPort string  `json:"tlsPort"`
	WsPath  string  `json:"wsPath"`
	WsPort  string  `json:"wsPort"`
	WsTLS   bool    `json:"wsTLS"`
	TlsInfo TLSInfo `json:"tlsInfo"`
	Debug   bool    `json:"debug"`
}

type TLSInfo struct {
	Verify   bool   `json:"verify"`
	CaFile   string `json:"caFile"`
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

var DefaultConfig *Config = &Config{
	Worker: 4096,
	Host:   "0.0.0.0",
	Port:   "1883",
}

var (
	log = logger.Prod().Named("broker")
)

func NewTLSConfig(tlsInfo TLSInfo) (*tls.Config, error) {

	cert, err := tls.LoadX509KeyPair(tlsInfo.CertFile, tlsInfo.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing X509 certificate/key pair: %v", zap.Error(err))
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing certificate: %v", zap.Error(err))
	}

	// Create TLSConfig
	// We will determine the cipher suites that we prefer.
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	// Require client certificates as needed
	if tlsInfo.Verify {
		config.ClientAuth = tls.RequireAndVerifyClientCert
	}
	// Add in CAs if applicable.
	if tlsInfo.CaFile != "" {
		rootPEM, err := ioutil.ReadFile(tlsInfo.CaFile)
		if err != nil || rootPEM == nil {
			return nil, err
		}
		pool := x509.NewCertPool()
		ok := pool.AppendCertsFromPEM(rootPEM)
		if !ok {
			return nil, fmt.Errorf("failed to parse root ca certificate")
		}
		config.ClientCAs = pool
	}

	return &config, nil
}
