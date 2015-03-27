// Package papertrail sends logs to Papertrail (https://papertrailapp.com/).
//
// Implements io.Writer interface (http://golang.org/pkg/io/#Writer).
//
//  writer := papertrail.Writer{
//    Port: 12345,
//    Network: papertrail.TCP,
//  }
//
//  // use writer directly
//  n, err := writer.Write([]byte("writer\n"))
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
	Server  string
}

func (w *Writer) Write(p []byte) (n int, err error) {

	var conn net.Conn

	if w.Server == "" {
		w.Server = "logs"
	}

	address := fmt.Sprintf("%s.papertrailapp.com:%d", w.Server, w.Port)

	switch w.Network {
	case UDP:
		conn, err = net.Dial(UDP, address)
	case TCP:
		roots := x509.NewCertPool()
		if ok := roots.AppendCertsFromPEM([]byte(pem)); !ok {
			return 0, errors.New("Failed to parse root certificate.")
		}
		conn, err = tls.Dial(TCP, address, &tls.Config{
			RootCAs: roots,
		})
	default:
		return 0, errors.New("Invalid Network. Neither UDP nor TCP.")
	}

	if err != nil {
		return 0, err
	}
	defer conn.Close()
	return conn.Write(p)
}
