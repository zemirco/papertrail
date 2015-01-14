// Package papertrail sends logs to Papertrail (https://papertrailapp.com/).
//
// Implements io.Writer interface (http://golang.org/pkg/io/#Writer).
//
//  writer := papertrail.Writer{
//    Port: 12345,
//    Network: papertrail.UDP,
//  }
//
//  // use writer directly
//  n, err := writer.Write([]byte("writer"))
//  if err != nil {
//    panic(err)
//  }
//  fmt.Printf("number of bytes written: %d\n", n)
//
//  // or create a new logger
//  logger := log.New(&writer, "", log.LstdFlags)
//  logger.Print("logger")
package papertrail

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
)

const (
	TCP = "tcp"
	UDP = "udp"
)

type Writer struct {
	Port    int
	Network string
}

func (w *Writer) Write(p []byte) (n int, err error) {

	address := fmt.Sprintf("logs.papertrailapp.com:%d", w.Port)

	if w.Network == UDP {
		conn, err := net.Dial(UDP, address)
		if err != nil {
			return 0, err
		}
		defer conn.Close()
		return conn.Write(p)
	}

	if w.Network == TCP {
		roots := x509.NewCertPool()
		ok := roots.AppendCertsFromPEM([]byte(pem))
		if !ok {
			return 0, errors.New("Failed to parse root certificate.")
		}
		conn, err := tls.Dial(TCP, address, &tls.Config{
			RootCAs: roots,
		})
		if err != nil {
			return 0, err
		}
		defer conn.Close()
		return conn.Write(p)
	}

	return 0, errors.New("Invalid Protocol. Neither UDP nor TCP.")
}
