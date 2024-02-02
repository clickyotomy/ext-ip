// Package resolve queries Google's DNS name servers for
// the client's external IP address.
package resolve

import (
	"fmt"
	"net"
	"time"

	"github.com/miekg/dns"
)

const (
	// This is (probably) a virtual IP for `ns{1,2,..,}.google.com'.
	server = "ns.google.com:53"
	record = dns.TypeTXT
	target = "o-o.myaddr.google.com."
)

// ExtIP fetches the external IP address of the client.
func ExtIP() (string, time.Duration, error) {
	var (
		client  dns.Client
		message dns.Msg
		err     error
		rtt     time.Duration
		result  *dns.Msg
		ok      bool
		txt     *dns.TXT
		rec     dns.RR
		addr    string
	)

	// Initiaize the client.
	client = dns.Client{}
	message = dns.Msg{}

	// Set the query.
	message.SetQuestion(target, record)

	// Execute the query; check for errors.
	result, rtt, err = client.Exchange(&message, server)
	if err != nil {
		return "", time.Duration(-1), fmt.Errorf("net: %s", err)
	}
	if len(result.Answer) == 0 {
		return "", time.Duration(-1), fmt.Errorf("dns: empty answer")
	}

	// Iterate through the answers, return the first `TXT' record.
	for _, rec = range result.Answer {
		if txt, ok = rec.(*dns.TXT); ok {
			addr = txt.Txt[0]
			break
		} else {
			err = fmt.Errorf("reflect: type assertion failed: (%s)", addr)
			return "", time.Duration(-1), err
		}
	}

	// Check if the returned IP address is valid.
	if net.ParseIP(addr) == nil {
		err = fmt.Errorf("net: invalid IP address: %s", addr)
		return "", time.Duration(-1), err
	}

	return addr, rtt, nil
}
