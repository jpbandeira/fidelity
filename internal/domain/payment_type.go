package domain

type PaymentType uint

const (
	invalidPaymentType PaymentType = iota
	credit
	debit
	pix
	money
	serviceExchange
)

// Supported payment types
const (
	creditPayment          = "Credito"
	debitPayment           = "Debito"
	pixPayment             = "Pix"
	moneyPayment           = "Dinheiro"
	serviceExchangePayment = "Troca de Serviço"
)

var supportedPaymentTypes = map[PaymentType]string{
	credit:          creditPayment,
	debit:           debitPayment,
	pix:             pixPayment,
	money:           moneyPayment,
	serviceExchange: serviceExchangePayment,
}

func (st PaymentType) String() string {
	if s, ok := supportedPaymentTypes[st]; ok {
		return s
	}

	return ""
}

func ToPaymentType(paymentType string) PaymentType {
	for k, v := range supportedPaymentTypes {
		if v == paymentType {
			return k
		}
	}

	return invalidPaymentType
}
