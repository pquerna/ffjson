package stripe

// Coupon represents percent-off discount you might want to apply to a customer.
//
// see https://stripe.com/docs/api#coupon_object
type Coupon struct {
	Id               string `json:"id"`
	Duration         string `json:"duration"`
	PercentOff       int    `json:"percent_off"`
	DurationInMonths int    `json:"duration_in_months,omitempty"`
	MaxRedemptions   int    `json:"max_redemptions,omitempty"`
	RedeemBy         int64  `json:"redeem_by,omitempty"`
	TimesRedeemed    int    `json:"times_redeemed,omitempty"`
	Livemode         bool   `json:"livemode"`
}
