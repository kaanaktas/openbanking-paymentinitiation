package consent

import "time"

type ObWriteDomesticConsent3 struct {
	Data *ObWriteDomesticConsent3Data `json:"Data"`
	Risk *ObRisk1                     `json:"Risk"`
}

type ObWriteDomesticConsent3Data struct {
	ConsentId      string                                     `json:"ConsentId,omitempty"`
	Initiation     *ObWriteDomestic2DataInitiation            `json:"Initiation"`
	Authorisation  *ObWriteDomesticConsent3DataAuthorisation  `json:"Authorisation,omitempty"`
	SCASupportData *ObWriteDomesticConsent3DataScaSupportData `json:"SCASupportData,omitempty"`
}

// The authorisation type request from the TPP.
type ObWriteDomesticConsent3DataAuthorisation struct {
	// Type of authorisation flow requested.
	AuthorisationType string `json:"AuthorisationType"`
	// Date and time at which the requested authorisation flow must be completed.All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
	CompletionDateTime time.Time `json:"CompletionDateTime,omitempty"`
}

// Supporting Data provided by TPP, when requesting SCA Exemption.
type ObWriteDomesticConsent3DataScaSupportData struct {
	// This field allows a PISP to request specific SCA Exemption for a Payment Initiation
	RequestedSCAExemptionType string `json:"RequestedSCAExemptionType,omitempty"`
	// Specifies a character string with a maximum length of 40 characters. Usage: This field indicates whether the PSU was subject to SCA performed by the TPP
	AppliedAuthenticationApproach string `json:"AppliedAuthenticationApproach,omitempty"`
	// Specifies a character string with a maximum length of 140 characters. Usage: If the payment is recurring then the transaction identifier of the previous payment occurrence so that the ASPSP can verify that the PISP, amount and the payee are the same as the previous occurrence.
	ReferencePaymentOrderId string `json:"ReferencePaymentOrderId,omitempty"`
}

// The Initiation payload is sent by the initiating party to the ASPSP. It is used to request movement of funds from the debtor account to a creditor for a single domestic payment.
type ObWriteDomestic2DataInitiation struct {
	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction. Usage: the  instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstructionIdentification string `json:"InstructionIdentification"`
	// Unique identification assigned by the initiating party to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain. Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction. OB: The Faster Payments Scheme can only access 31 characters for the EndToEndIdentification field.
	EndToEndIdentification string                                               `json:"EndToEndIdentification"`
	LocalInstrument        string                                               `json:"LocalInstrument,omitempty"`
	InstructedAmount       *ObWriteDomestic2DataInitiationInstructedAmount      `json:"InstructedAmount"`
	DebtorAccount          *ObWriteDomestic2DataInitiationDebtorAccount         `json:"DebtorAccount,omitempty"`
	CreditorAccount        *ObWriteDomestic2DataInitiationCreditorAccount       `json:"CreditorAccount"`
	CreditorPostalAddress  *ObPostalAddress6                                    `json:"CreditorPostalAddress,omitempty"`
	RemittanceInformation  *ObWriteDomestic2DataInitiationRemittanceInformation `json:"RemittanceInformation,omitempty"`
	SupplementaryData      *ObSupplementaryData1                                `json:"SupplementaryData,omitempty"`
}

// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party. Usage: This amount has to be transported unchanged through the transaction chain.
type ObWriteDomestic2DataInitiationInstructedAmount struct {
	Amount   string `json:"Amount"`
	Currency string `json:"Currency"`
}

// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
type ObWriteDomestic2DataInitiationDebtorAccount struct {
	SchemeName     string `json:"SchemeName"`
	Identification string `json:"Identification"`
	// The account name is the name or names of the account owner(s) represented at an account level, as displayed by the ASPSP's online channels. Note, the account name is not the product name or the nickname of the account.
	Name                    string `json:"Name,omitempty"`
	SecondaryIdentification string `json:"SecondaryIdentification,omitempty"`
}

// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
type ObWriteDomestic2DataInitiationCreditorAccount struct {
	SchemeName     string `json:"SchemeName"`
	Identification string `json:"Identification"`
	// The account name is the name or names of the account owner(s) represented at an account level. Note, the account name is not the product name or the nickname of the account. OB: ASPSPs may carry out name validation for Confirmation of Payee, but it is not mandatory.
	Name                    string `json:"Name"`
	SecondaryIdentification string `json:"SecondaryIdentification,omitempty"`
}

// Information that locates and identifies a specific address, as defined by postal services.
type ObPostalAddress6 struct {
	AddressType        string   `json:"AddressType,omitempty"`
	Department         string   `json:"Department,omitempty"`
	SubDepartment      string   `json:"SubDepartment,omitempty"`
	StreetName         string   `json:"StreetName,omitempty"`
	BuildingNumber     string   `json:"BuildingNumber,omitempty"`
	PostCode           string   `json:"PostCode,omitempty"`
	TownName           string   `json:"TownName,omitempty"`
	CountrySubDivision string   `json:"CountrySubDivision,omitempty"`
	Country            string   `json:"Country,omitempty"`
	AddressLine        []string `json:"AddressLine,omitempty"`
}

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
type ObWriteDomestic2DataInitiationRemittanceInformation struct {
	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured string `json:"Unstructured,omitempty"`
	// Unique reference, as assigned by the creditor, to unambiguously refer to the payment transaction. Usage: If available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the amount of money. If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification. OB: The Faster Payments Scheme can only accept 18 characters for the ReferenceInformation field - which is where this ISO field will be mapped.
	Reference string `json:"Reference,omitempty"`
}

// Additional information that can not be captured in the structured fields and/or any other specific block.
type ObSupplementaryData1 struct {
}

// A number of monetary units specified in an active currency where the unit of currency is explicit and compliant with ISO 4217.
type ObActiveCurrencyAndAmountSimpleType struct {
}

// A code allocated to a currency by a Maintenance Agency under an international identification scheme, as described in the latest edition of the international standard ISO 4217 \"Codes for the representation of currencies and funds\".
type ActiveOrHistoricCurrencyCode struct {
}

// The Risk section is sent by the initiating party to the ASPSP. It is used to specify additional details for risk scoring for Payments.
type ObRisk1 struct {
	// Specifies the payment context
	PaymentContextCode string `json:"PaymentContextCode,omitempty"`
	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode string `json:"MerchantCategoryCode,omitempty"`
	// The unique customer identifier of the PSU with the merchant.
	MerchantCustomerIdentification string                  `json:"MerchantCustomerIdentification,omitempty"`
	DeliveryAddress                *ObRisk1DeliveryAddress `json:"DeliveryAddress,omitempty"`
}

// Information that locates and identifies a specific address, as defined by postal services or in free format text.
type ObRisk1DeliveryAddress struct {
	AddressLine        []string `json:"AddressLine,omitempty"`
	StreetName         string   `json:"StreetName,omitempty"`
	BuildingNumber     string   `json:"BuildingNumber,omitempty"`
	PostCode           string   `json:"PostCode,omitempty"`
	TownName           string   `json:"TownName"`
	CountrySubDivision []string `json:"CountrySubDivision,omitempty"`
	// Nation with its own government, occupying a particular territory.
	Country string `json:"Country"`
}

type ObWriteDomesticConsentResponse3 struct {
	Data  *ObWriteDomesticConsentResponse3Data `json:"Data"`
	Risk  *ObRisk1                             `json:"Risk"`
	Links *Links                               `json:"Links,omitempty"`
	Meta  *Meta                                `json:"Meta,omitempty"`
}

type ObWriteDomesticConsentResponse3Data struct {
	// OB: Unique identification as assigned by the ASPSP to uniquely identify the consent resource.
	ConsentId string `json:"ConsentId"`
	// Date and time at which the resource was created.All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
	CreationDateTime time.Time `json:"CreationDateTime"`
	// Specifies the status of consent resource in code form.
	Status string `json:"Status"`
	// Date and time at which the resource status was updated.All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
	StatusUpdateDateTime time.Time `json:"StatusUpdateDateTime"`
	// Specified cut-off date and time for the payment consent.All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
	CutOffDateTime time.Time `json:"CutOffDateTime,omitempty"`
	// Expected execution date and time for the payment resource.All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
	ExpectedExecutionDateTime time.Time `json:"ExpectedExecutionDateTime,omitempty"`
	// Expected settlement date and time for the payment resource.All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
	ExpectedSettlementDateTime time.Time                                    `json:"ExpectedSettlementDateTime,omitempty"`
	Charges                    []ObWriteDomesticConsentResponse3DataCharges `json:"Charges,omitempty"`
	Initiation                 *ObWriteDomestic2DataInitiation              `json:"Initiation"`
	Authorisation              *ObWriteDomesticConsent3DataAuthorisation    `json:"Authorisation,omitempty"`
	SCASupportData             *ObWriteDomesticConsent3DataScaSupportData   `json:"SCASupportData,omitempty"`
}

// Set of elements used to provide details of a charge for the payment initiation.
type ObWriteDomesticConsentResponse3DataCharges struct {
	ChargeBearer *ObChargeBearerType1Code             `json:"ChargeBearer"`
	Type_        *ObExternalPaymentChargeType1Code    `json:"Type"`
	Amount       *ObActiveOrHistoricCurrencyAndAmount `json:"Amount"`
}

// ObChargeBearerType1Code : Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
type ObChargeBearerType1Code string

// Charge type, in a coded form.
type ObExternalPaymentChargeType1Code struct {
}

// Amount of money associated with the charge type.
type ObActiveOrHistoricCurrencyAndAmount struct {
	Amount   *ObActiveCurrencyAndAmountSimpleType `json:"Amount"`
	Currency *ActiveOrHistoricCurrencyCode        `json:"Currency"`
}

// Links relevant to the payload
type Links struct {
	Self  string `json:"Self"`
	First string `json:"First,omitempty"`
	Prev  string `json:"Prev,omitempty"`
	Next  string `json:"Next,omitempty"`
	Last  string `json:"Last,omitempty"`
}

// Meta Data relevant to the payload
type Meta struct {
	TotalPages             int32        `json:"TotalPages,omitempty"`
	FirstAvailableDateTime *IsoDateTime `json:"FirstAvailableDateTime,omitempty"`
	LastAvailableDateTime  *IsoDateTime `json:"LastAvailableDateTime,omitempty"`
}

// All dates in the JSON payloads are represented in ISO 8601 date-time format.  All date-time fields in responses must include the timezone. An example is below: 2017-04-05T10:43:07+00:00
type IsoDateTime struct {
}

// ******** TYPES OF AUTHORIZE CONSENT JWT *************
type Acr struct {
	Value     string `json:"value"`
	Essential bool   `json:"essential"`
}

type OpenBankingIntentId struct {
	Value     string `json:"value"`
	Essential bool   `json:"essential"`
}

type IdToken struct {
	OpenBankingIntentId OpenBankingIntentId `json:"openbanking_intent_id"`
	Acr                 Acr                 `json:"acr"`
}

type Userinfo struct {
	OpenBankingIntentId OpenBankingIntentId `json:"openbanking_intent_id"`
}

type Claims struct {
	Userinfo Userinfo `json:"userinfo"`
	//	IdToken  IdToken  `json:"id_token"`
}

type AuthorisedConsent struct {
	Iss          string `json:"iss"`
	Aud          string `json:"aud"`
	ResponseType string `json:"response_type"`
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	//Nonce        string `json:"nonce"`
	//State        string `json:"state"`
	//Exp          int64  `json:"exp"`
	//Iat          int64  `json:"iat"`
	Claims Claims `json:"claims"`
}

type Consent struct {
	Id                          int64  `db:"id"`
	SessionReferenceId          string `db:"session_reference_id"`
	TrackingId                  string `db:"tracking_id"`
	AspspId                     string `db:"aspsp_id"`
	ConsentId                   string `db:"consent_id"`
	ConsentStatusUpdateDateTime string `db:"consent_status_update_date_time"`
	CreateDateTime              string `db:"create_date_time"`
	UpdateDateTime              string `db:"update_date_time"`
	ConsentStatus               string `db:"consent_status"`
	ConsentType                 string `db:"consent_type"`
	ObjectState                 string `db:"object_state"`
	Tokens                      []Token
}

type TokensInConsent struct {
	Consent
	Token
}

type Token struct {
	Id                      *int64  `db:"token_tid"`
	AccessToken             *string `db:"access_token"`
	ResourceAccessToken     *string `db:"resource_access_token"`
	ResourceRefreshToken    *string `db:"resource_refresh_token"`
	TokenStatus             *string `db:"token_status"`
	ExpiresIn               *int    `db:"expires_in"`
	CreateDateTime          *string `db:"create_date_time"`
	UpdateDateTime          *string `db:"update_date_time"`
	TokenExpirationDateTime *string `db:"token_expiration_date_time"`
	ConsentTid              *int64  `db:"consent_tid"`
}

type ActiveConsent struct {
	ConsentTid int64  `json:"consentTid"`
	AspspId    string `json:"aspspId"`
}
