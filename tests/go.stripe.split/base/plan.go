package stripe

// Plan holds details about pricing information for different products and
// feature levels on your site. For example, you might have a $10/month plan
// for basic features and a different $20/month plan for premium features.
//
// see https://stripe.com/docs/api#plan_object
type Plan struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Amount          int64  `json:"amount"`
	Interval        string `json:"interval"`
	IntervalCount   int    `json:"interval_count"`
	Currency        string `json:"currency"`
	TrialPeriodDays int    `json:"trial_period_days"`
	Livemode        bool   `json:"livemode"`
}
