package stripe

import (
	"time"
)

// Customer encapsulates details about a Customer registered in Stripe.
//
// see https://stripe.com/docs/api#customer_object
type Customer struct {
	Id           string        `json:"id"`
	Desc         string        `json:"description,omitempty"`
	Email        string        `json:"email,omitempty"`
	Created      int64         `json:"created"`
	Balance      int64         `json:"account_balance"`
	Delinquent   bool          `json:"delinquent"`
	Cards        CardData      `json:"cards,omitempty"`
	Discount     *Discount     `json:"discount,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	Livemode     bool          `json:"livemode"`
	DefaultCard  string        `json:"default_card"`
}

func NewCustomer() *Customer {

	return &Customer{
		Id:         "hooN5ne7ug",
		Desc:       "A very nice customer.",
		Email:      "customer@example.com",
		Created:    time.Now().UnixNano(),
		Balance:    10,
		Delinquent: false,
		Cards: CardData{
			Object: "A92F4CFE-8B6B-4176-873E-887AC0D120EB",
			Count:  1,
			Url:    "https://stripe.example.com/card/A92F4CFE-8B6B-4176-873E-887AC0D120EB",
			Data: []*Card{
				&Card{
					Name:        "John Smith",
					Id:          "7526EC97-A0B6-47B2-AAE5-17443626A116",
					Fingerprint: "4242424242424242",
					ExpYear:     time.Now().Year() + 1,
					ExpMonth:    1,
				},
			},
		},
		Discount: &Discount{
			Id:       "Ee9ieZ8zie",
			Customer: "hooN5ne7ug",
			Start:    time.Now().UnixNano(),
			End:      time.Now().UnixNano(),
			Coupon: &Coupon{
				Id:               "ieQuo5Aiph",
				Duration:         "2m",
				PercentOff:       10,
				DurationInMonths: 2,
				MaxRedemptions:   1,
				RedeemBy:         time.Now().UnixNano(),
				TimesRedeemed:    1,
				Livemode:         true,
			},
		},
		Subscription: &Subscription{
			Customer: "hooN5ne7ug",
			Status:   SubscriptionActive,
			Plan: &Plan{
				Id:              "gaiyeLua5u",
				Name:            "Great Plan (TM)",
				Amount:          10,
				Interval:        "monthly",
				IntervalCount:   3,
				Currency:        "USD",
				TrialPeriodDays: 15,
				Livemode:        true,
			},
			Start:              time.Now().UnixNano(),
			EndedAt:            0,
			CurrentPeriodStart: time.Now().UnixNano(),
			CurrentPeriodEnd:   time.Now().UnixNano(),
			TrialStart:         time.Now().UnixNano(),
			TrialEnd:           time.Now().UnixNano(),
			CanceledAt:         0,
			CancelAtPeriodEnd:  false,
			Quantity:           2,
		},
		Livemode:    true,
		DefaultCard: "7526EC97-A0B6-47B2-AAE5-17443626A116",
	}
}
