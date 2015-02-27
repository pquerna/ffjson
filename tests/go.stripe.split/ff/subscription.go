package stripe

// Subscription Statuses
const (
	SubscriptionTrialing = "trialing"
	SubscriptionActive   = "active"
	SubscriptionPastDue  = "past_due"
	SubscriptionCanceled = "canceled"
	SubscriptionUnpaid   = "unpaid"
)

// Subscriptions represents a recurring charge a customer's card.
//
// see https://stripe.com/docs/api#subscription_object
type Subscription struct {
	Customer           string `json:"customer"`
	Status             string `json:"status"`
	Plan               *Plan  `json:"plan"`
	Start              int64  `json:"start"`
	EndedAt            int64  `json:"ended_at"`
	CurrentPeriodStart int64  `json:"current_period_start"`
	CurrentPeriodEnd   int64  `json:"current_period_end"`
	TrialStart         int64  `json:"trial_start"`
	TrialEnd           int64  `json:"trial_end"`
	CanceledAt         int64  `json:"canceled_at"`
	CancelAtPeriodEnd  bool   `json:"cancel_at_period_end"`
	Quantity           int64  `json"quantity"`
}
