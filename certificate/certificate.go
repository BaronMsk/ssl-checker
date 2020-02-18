package certificate

import (
	"crypto/x509"
	"time"
	"log"
	"crypto/tls"
	"github.com/BaronMsk/ssl-checker/config"
	"net"
	"github.com/BaronMsk/ssl-checker/notification"
)



func NewCheckCertificate(config *config.ConfigurationStruct) {
	tlsConf := &tls.Config{}

	for {
		for _, domain := range config.Domains {
			dialer := &net.Dialer{
				Timeout: config.Timeout,
				FallbackDelay: config.Timeout,
			}

			conn, err := tls.DialWithDialer(dialer, "tcp", domain, tlsConf)

			if err != nil {
				log.Println(err)
			}

			stateCert := conn.ConnectionState().PeerCertificates

			for _, cert := range stateCert {
				if validCert(cert, config) == false {
					log.Println(cert.DNSNames)
					notification.NewNotification(config, CertInfo(cert))
				}
			}
		}
		time.Sleep(config.Interval)
	}

}

func validCert(cert *x509.Certificate, config *config.ConfigurationStruct) bool {
	duration, err := time.ParseDuration(config.Trigger)
	if err != nil {
		log.Println(err)
	}

	if time.Now().After(cert.NotAfter.Add(duration))  {
		return false
	}
	return true
}

func CertInfo(cert *x509.Certificate) *notification.CertificateInfoStruct {
	info := &notification.CertificateInfoStruct{
		cert.DNSNames,
		cert.SerialNumber.String(),
		cert.NotAfter.String(),
	}
	return info
}
