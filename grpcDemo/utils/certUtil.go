package utils

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func GetTls() credentials.TransportCredentials {
	creds, err := credentials.NewServerTLSFromFile("./identity/server.crt", "./identity/server_no_password.key")
	if err != nil {
		log.Fatal(err)
	}
	return  creds
}

func GetServerTls() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("./identity/server.pem", "./identity/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("./identity/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	return creds
}

func GetClientTls() credentials.TransportCredentials {
	cert, _ := tls.LoadX509KeyPair("./identity/client.pem", "./identity/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("./identity/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "localhost",
		RootCAs: certPool,
	})

	return creds
}
