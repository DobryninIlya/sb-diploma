package billing

type BillingData struct {
	CreateCustomer bool `json:"create_customer,omitempty"`
	Purchase       bool `json:"purchase,omitempty"`
	Payout         bool `json:"payout,omitempty"`
	Recurring      bool `json:"recurring,omitempty"`
	FraudControl   bool `json:"fraud_control,omitempty"`
	CheckoutPage   bool `json:"checkout_page,omitempty"`
}
