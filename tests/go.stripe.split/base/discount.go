package stripe

// Discount represents the actual application of a coupon to a particular
// customer.
//
// see https://stripe.com/docs/api#discount_object
type Discount struct {
	Id       string  `json:"id"`
	Customer string  `json:"customer"`
	Start    int64   `json:"start"`
	End      int64   `json:"end"`
	Coupon   *Coupon `json:"coupon"`
}
