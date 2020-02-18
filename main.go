package main

import (
	"github.com/BaronMsk/ssl-checker/config"
	"github.com/BaronMsk/ssl-checker/certificate"
)

func main() {
	config := config.NewConfig()
	certificate.NewCheckCertificate(&config)
}