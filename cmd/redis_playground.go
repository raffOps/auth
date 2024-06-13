package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/redis/go-redis/v9"
	"io/ioutil"
	"log"
	"time"
)

var ctx = context.Background()

func main() {
	// Carregar o certificado da Autoridade Certificadora
	caCert, err := ioutil.ReadFile("redis/tls/ca.crt")
	if err != nil {
		log.Fatalf("Failed to load CA cert: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatalf("Failed to append CA cert to pool")
	}

	// Carregar o certificado do cliente e a chave privada
	clientCert, err := tls.LoadX509KeyPair("redis/tls/client.crt", "redis/tls/client.key")
	if err != nil {
		log.Fatalf("Failed to load client cert/key: %v", err)
	}

	// Configurar o tls.Config
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{clientCert},
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		TLSConfig:    tlsConfig,
		DialTimeout:  50 * time.Second, // Tempo limite de conexão
		ReadTimeout:  30 * time.Second, // Tempo limite de leitura
		WriteTimeout: 30 * time.Second, // Tempo limite de escrita

	})

	err = rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = rdb.Set(ctx, "teste", "teste", 0).Err()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
