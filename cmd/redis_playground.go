package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var ctx = context.Background()

func main() {
	cert, err := tls.LoadX509KeyPair("redis/tls/client.crt", "redis/tls/client.key")
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := os.ReadFile("redis/tls/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// Configurar o tls.Config
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:      "localhost:6379",
		TLSConfig: tlsConfig,
	})

	err = rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("%v", err)
	}

	//err = rdb.Set(ctx, "teste", "teste", 0).Err()
	//if err != nil {
	//	log.Fatalf("%v", err)
	//}

	//hashMap := map[string]interface{}{
	//	"field3": "value3",
	//	"field4": "value4",
	//}
	//for k, v := range hashMap {
	//	err = rdb.HSet(ctx, "sessionId:12233434", k, v).Err()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}

	rdb.SAdd(ctx, "set1", "value1", "value2", "value3")
}
