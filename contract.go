package gopagseguro

// Charge represents the PagSeguro v4 API charge payload.
type Charge struct {
	ID               string           `json:"id,omitempty"`
	ReferenceID      string           `json:"reference_id,omitempty"`
	Status           string           `json:"status,omitempty"`
	Description      string           `json:"description,omitempty"`
	Amount           Amount           `json:"amount"`
	PaymentMethod    *PaymentMethod   `json:"payment_method,omitempty"`
	PaymentResponse  *PaymentResponse `json:"payment_response,omitempty"`
	Links            []Links          `json:"links,omitempty"`
	NotificationUrls []string         `json:"notification_urls,omitempty"`
	CreatedAt        string           `json:"created_at,omitempty"`
	PaidAt           string           `json:"paid_at,omitempty"`
}

// Summary represents the charge amount summary.
type Summary struct {
	Total    int `json:"total"`
	Paid     int `json:"paid"`
	Refunded int `json:"refunded"`
}

// Amount represents the charge amount.
type Amount struct {
	Value    int      `json:"value"`
	Currency string   `json:"currency,omitempty"`
	Summary  *Summary `json:"summary,omitempty"`
}

// Holder represents the charge card payment method holder.
type Holder struct {
	Name    string   `json:"name"`
	TaxID   string   `json:"tax_id"`
	Email   string   `json:"email"`
	Address *Address `json:"address"`
}

// Card represents the charge payment method card.
type Card struct {
	Number       string  `json:"number,omitempty"`
	SecurityCode string  `json:"security_code,omitempty"`
	Encrypted    string  `json:"encrypted,omitempty"`
	Brand        string  `json:"brand,omitempty"`
	FirstDigits  string  `json:"first_digits,omitempty"`
	LastDigits   string  `json:"last_digits,omitempty"`
	ExpMonth     string  `json:"exp_month,omitempty"`
	ExpYear      string  `json:"exp_year,omitempty"`
	Holder       *Holder `json:"holder,omitempty"`
}

// PaymentMethod represents the charge payment method.
type PaymentMethod struct {
	Type           string  `json:"type"`
	Installments   int     `json:"installments,omitempty"`
	Capture        bool    `json:"capture"`
	Card           *Card   `json:"card,omitempty"`
	CaptureBefore  string  `json:"capture_before,omitempty"`
	SoftDescriptor string  `json:"soft_descriptor,omitempty"`
	Boleto         *Boleto `json:"boleto,omitempty"`
}

// PaymentResponse represents the charge payment method response.
type PaymentResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
}

// Links represents the charge links response.
type Links struct {
	Rel   string `json:"rel"`
	Href  string `json:"href"`
	Media string `json:"media"`
	Type  string `json:"type"`
}

// InstructionLines represents the charge payment method boleto instruction lines.
type InstructionLines struct {
	Line1 string `json:"line_1"`
	Line2 string `json:"line_2"`
}

// Address represents the charge payment method holder address.
type Address struct {
	Country    string `json:"country"`
	Region     string `json:"region"`
	RegionCode string `json:"region_code"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Locality   string `json:"locality"`
}

// Address represents the charge payment method boleto.
type Boleto struct {
	ID               string            `json:"id"`
	Barcode          string            `json:"barcode"`
	FormattedBarcode string            `json:"formatted_barcode"`
	DueDate          string            `json:"due_date"`
	InstructionLines *InstructionLines `json:"instruction_lines"`
	Holder           Holder            `json:"holder"`
}
