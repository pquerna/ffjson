package stripe

// Credit Card Types accepted by the Stripe API.
const (
	AmericanExpress = "American Express"
	DinersClub      = "Diners Club"
	Discover        = "Discover"
	JCB             = "JCB"
	MasterCard      = "MasterCard"
	Visa            = "Visa"
	UnknownCard     = "Unknown"
)

// Card represents details about a Credit Card entered into Stripe.
type Card struct {
	Id                string `json:"id"`
	Name              string `json:"name,omitempty"`
	Type              string `json:"type"`
	ExpMonth          int    `json:"exp_month"`
	ExpYear           int    `json:"exp_year"`
	Last4             string `json:"last4"`
	Fingerprint       string `json:"fingerprint"`
	Country           string `json:"country,omitempty"`
	AddrUess1         string `json:"address_line1,omitempty"`
	Address2          string `json:"address_line2,omitempty"`
	AddressCountry    string `json:"address_country,omitempty"`
	AddressState      string `json:"address_state,omitempty"`
	AddressZip        string `json:"address_zip,omitempty"`
	AddressCity       string `json:"address_city"`
	AddressLine1Check string `json:"address_line1_check,omitempty"`
	AddressZipCheck   string `json:"address_zip_check,omitempty"`
	CVCCheck          string `json:"cvc_check,omitempty"`
}
