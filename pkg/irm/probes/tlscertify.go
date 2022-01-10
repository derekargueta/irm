package probes

import (
	"crypto/tls"
	"log"
	"net"
	"strings"
	"time"
)

type Tlscertify struct {
}

/*
Doesnt work atm, supported Max CDN websites
dont return true when ran through cidr list

*/

func (h *Tlscertify) Run(domain string) *TotalTlsCertify {
	digicert := false
	comodo := false
	Encrypt := false
	Amazon := false

	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: time.Second * 15}, "tcp", domain+":443", nil)

	if err != nil {
		log.Print("No SSL certificate: "+domain, err)
	} else {
		certificate := conn.ConnectionState().PeerCertificates[0].Issuer.Organization[0]

		if strings.Contains(certificate, "DigiCert Inc") {
			digicert = true
		} else if strings.EqualFold(certificate, "COMODO CA Limited") {
			comodo = true
		} else if strings.EqualFold(certificate, "Let's Encrypt") {
			Encrypt = true
		} else if strings.EqualFold(certificate, "Amazon") {
			Amazon = true
		}

		return &TotalTlsCertify{
			Digicert: digicert,
			Comodo:   comodo,
			Encrypt:  Encrypt,
			Amazon:   Amazon,
			Err:      err,
		}
	}
	return &TotalTlsCertify{
		Digicert: false,
		Comodo:   false,
		Encrypt:  false,
		Amazon:   false,
		Err:      err,
	}

}
