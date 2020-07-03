package gatherer

import (
	"github.com/NhokCrazy199/AssassinGo/logger"
	"github.com/bobesa/go-domain-util/domainutil"
	whois "github.com/likexian/whois-go"
	whois_parser "github.com/likexian/whois-parser-go"
)

// Whois queries the domain information.
type Whois struct {
	target string
	raw    string
	info   map[string]interface{}
}

// NewWhois returns a new Whois.
func NewWhois() *Whois {
	return &Whois{}
}

// Set implements Gatherer interface.
// Params should be {target string}.
func (w *Whois) Set(v ...interface{}) {
	w.target = domainutil.Domain(v[0].(string))
}

// Report implements Gatherer interface.
func (w *Whois) Report() map[string]interface{} {
	return w.info
}

// Run implements Gatherer interface.
func (w *Whois) Run() {
	logger.Green.Println("Whois Information")

	whoisRaw, err := whois.Whois(w.target)
	if err != nil {
		logger.Red.Println(err)
		return
	}
	w.raw = whoisRaw
	result, _ := whois_parser.Parse(w.raw)

	w.info = map[string]interface{}{
		"domain":          w.target,
		"registrar_name":  result.Registrar.ReferralURL,
		"admin_name":      result.Administrative.Name,
		"admin_email":     result.Administrative.Email,
		"admin_phone":     result.Administrative.Phone,
		"domain_name":     result.Domain.Name,
		"created_date":    result.Domain.CreatedDate,
		"expiration_date": result.Domain.ExpirationDate,
		"ns":              result.Domain.NameServers,
		"state":           result.Domain.Status[0],
	}
	for k, v := range w.info {
		logger.Blue.Println(k + ": " + v.(string))
	}
}
