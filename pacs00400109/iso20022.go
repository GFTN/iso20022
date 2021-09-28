package pacs00400109

import "encoding/xml"

type AMLIndicator string

// Specifies the unique identification of an account as assigned by the account servicer.
type AccountIdentification4Choice struct {
	XMLName xml.Name

	// International Bank Account Number (IBAN) - identifier used internationally by financial institutions to uniquely identify the account of a customer. Further specifications of the format and content of the IBAN can be found in the standard ISO 13616 "Banking and related financial services - International Bank Account Number (IBAN)" version 1997-10-01, or later revisions.
	IBAN *IBAN2007Identifier `xml:"IBAN" json:"iban" `

	// Unique identification of an account, as assigned by the account servicer, using an identification scheme.
	Othr *GenericAccountIdentification1 `xml:"Othr" json:"othr" `
}

func (a *AccountIdentification4Choice) SetIBAN(value string) {
	a.IBAN = (*IBAN2007Identifier)(&value)
}

func (a *AccountIdentification4Choice) AddOthr() *GenericAccountIdentification1 {
	a.Othr = new(GenericAccountIdentification1)
	return a.Othr
}

// Sets of elements to identify a name of the identification scheme.
type AccountSchemeName1Choice struct {
	XMLName xml.Name

	// Name of the identification scheme, in a coded form as published in an external list.
	Cd *ExternalAccountIdentification1Code `xml:"Cd" json:"cd" `

	// Name of the identification scheme, in a free text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (a *AccountSchemeName1Choice) SetCd(value string) {
	a.Cd = (*ExternalAccountIdentification1Code)(&value)
}

func (a *AccountSchemeName1Choice) SetPrtry(value string) {
	a.Prtry = (*Max35Text)(&value)
}

// A number of monetary units specified in an active currency where the unit of currency is explicit and compliant with ISO 4217.
type ActiveCurrencyAndAmount struct {
	XMLName  xml.Name
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewActiveCurrencyAndAmount(value, currency string) *ActiveCurrencyAndAmount {
	return &ActiveCurrencyAndAmount{Value: value, Currency: currency}
}

// A number of monetary units specified in an active or a historic currency where the unit of currency is explicit and compliant with ISO 4217.
type ActiveOrHistoricCurrencyAndAmount struct {
	XMLName  xml.Name
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewActiveOrHistoricCurrencyAndAmount(value, currency string) *ActiveOrHistoricCurrencyAndAmount {
	return &ActiveOrHistoricCurrencyAndAmount{Value: value, Currency: currency}
}

type ActiveOrHistoricCurrencyCode string

type AddressType2Code string

// Choice of formats for the type of address.
type AddressType3Choice struct {
	XMLName xml.Name

	// Type of address expressed as a code.
	Cd *AddressType2Code `xml:"Cd" json:"cd" `

	// Type of address expressed as a proprietary code.
	Prtry *GenericIdentification30 `xml:"Prtry" json:"prtry" `
}

func (a *AddressType3Choice) SetCd(value string) {
	a.Cd = (*AddressType2Code)(&value)
}

func (a *AddressType3Choice) AddPrtry() *GenericIdentification30 {
	a.Prtry = new(GenericIdentification30)
	return a.Prtry
}

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails12 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OrgnlMndtId *Max35Text `xml:"OrgnlMndtId,omitempty" json:"orgnl_mndt_id" `

	// Original creditor scheme identification that has been modified.
	OrgnlCdtrSchmeId *PartyIdentification125 `xml:"OrgnlCdtrSchmeId,omitempty" json:"orgnl_cdtr_schme_id" `

	// Original creditor agent that has been modified.
	OrgnlCdtrAgt *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlCdtrAgt,omitempty" json:"orgnl_cdtr_agt" `

	// Original creditor agent account that has been modified.
	OrgnlCdtrAgtAcct *CashAccount24 `xml:"OrgnlCdtrAgtAcct,omitempty" json:"orgnl_cdtr_agt_acct" `

	// Original debtor that has been modified.
	OrgnlDbtr *PartyIdentification125 `xml:"OrgnlDbtr,omitempty" json:"orgnl_dbtr" `

	// Original debtor account that has been modified.
	OrgnlDbtrAcct *CashAccount24 `xml:"OrgnlDbtrAcct,omitempty" json:"orgnl_dbtr_acct" `

	// Original debtor agent that has been modified.
	OrgnlDbtrAgt *BranchAndFinancialInstitutionIdentification5 `xml:"OrgnlDbtrAgt,omitempty" json:"orgnl_dbtr_agt" `

	// Original debtor agent account that has been modified.
	OrgnlDbtrAgtAcct *CashAccount24 `xml:"OrgnlDbtrAgtAcct,omitempty" json:"orgnl_dbtr_agt_acct" `

	// Original final collection date that has been modified.
	OrgnlFnlColltnDt *ISODate `xml:"OrgnlFnlColltnDt,omitempty" json:"orgnl_fnl_colltn_dt" `

	// Original frequency that has been modified.
	OrgnlFrqcy *Frequency36Choice `xml:"OrgnlFrqcy,omitempty" json:"orgnl_frqcy" `

	// Original reason for the mandate to allow the user to distinguish between different mandates for the same creditor.
	OrgnlRsn *MandateSetupReason1Choice `xml:"OrgnlRsn,omitempty" json:"orgnl_rsn" `

	// Original number of tracking days that has been modified.
	OrgnlTrckgDays *Exact2NumericText `xml:"OrgnlTrckgDays,omitempty" json:"orgnl_trckg_days" `
}

func (a *AmendmentInformationDetails12) SetOrgnlMndtId(value string) {
	a.OrgnlMndtId = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails12) AddOrgnlCdtrSchmeId() *PartyIdentification125 {
	a.OrgnlCdtrSchmeId = new(PartyIdentification125)
	return a.OrgnlCdtrSchmeId
}

func (a *AmendmentInformationDetails12) AddOrgnlCdtrAgt() *BranchAndFinancialInstitutionIdentification5 {
	a.OrgnlCdtrAgt = new(BranchAndFinancialInstitutionIdentification5)
	return a.OrgnlCdtrAgt
}

func (a *AmendmentInformationDetails12) AddOrgnlCdtrAgtAcct() *CashAccount24 {
	a.OrgnlCdtrAgtAcct = new(CashAccount24)
	return a.OrgnlCdtrAgtAcct
}

func (a *AmendmentInformationDetails12) AddOrgnlDbtr() *PartyIdentification125 {
	a.OrgnlDbtr = new(PartyIdentification125)
	return a.OrgnlDbtr
}

func (a *AmendmentInformationDetails12) AddOrgnlDbtrAcct() *CashAccount24 {
	a.OrgnlDbtrAcct = new(CashAccount24)
	return a.OrgnlDbtrAcct
}

func (a *AmendmentInformationDetails12) AddOrgnlDbtrAgt() *BranchAndFinancialInstitutionIdentification5 {
	a.OrgnlDbtrAgt = new(BranchAndFinancialInstitutionIdentification5)
	return a.OrgnlDbtrAgt
}

func (a *AmendmentInformationDetails12) AddOrgnlDbtrAgtAcct() *CashAccount24 {
	a.OrgnlDbtrAgtAcct = new(CashAccount24)
	return a.OrgnlDbtrAgtAcct
}

func (a *AmendmentInformationDetails12) SetOrgnlFnlColltnDt(value string) {
	a.OrgnlFnlColltnDt = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails12) AddOrgnlFrqcy() *Frequency36Choice {
	a.OrgnlFrqcy = new(Frequency36Choice)
	return a.OrgnlFrqcy
}

func (a *AmendmentInformationDetails12) AddOrgnlRsn() *MandateSetupReason1Choice {
	a.OrgnlRsn = new(MandateSetupReason1Choice)
	return a.OrgnlRsn
}

func (a *AmendmentInformationDetails12) SetOrgnlTrckgDays(value string) {
	a.OrgnlTrckgDays = (*Exact2NumericText)(&value)
}

// Provides further details on the list of direct debit mandate elements that have been modified when the amendment indicator has been set.
type AmendmentInformationDetails13 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the creditor, to unambiguously identify the original mandate.
	OrgnlMndtId *Max35Text `xml:"OrgnlMndtId,omitempty" json:"orgnl_mndt_id" `

	// Original creditor scheme identification that has been modified.
	OrgnlCdtrSchmeId *PartyIdentification135 `xml:"OrgnlCdtrSchmeId,omitempty" json:"orgnl_cdtr_schme_id" `

	// Original creditor agent that has been modified.
	OrgnlCdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlCdtrAgt,omitempty" json:"orgnl_cdtr_agt" `

	// Original creditor agent account that has been modified.
	OrgnlCdtrAgtAcct *CashAccount38 `xml:"OrgnlCdtrAgtAcct,omitempty" json:"orgnl_cdtr_agt_acct" `

	// Original debtor that has been modified.
	OrgnlDbtr *PartyIdentification135 `xml:"OrgnlDbtr,omitempty" json:"orgnl_dbtr" `

	// Original debtor account that has been modified.
	OrgnlDbtrAcct *CashAccount38 `xml:"OrgnlDbtrAcct,omitempty" json:"orgnl_dbtr_acct" `

	// Original debtor agent that has been modified.
	OrgnlDbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlDbtrAgt,omitempty" json:"orgnl_dbtr_agt" `

	// Original debtor agent account that has been modified.
	OrgnlDbtrAgtAcct *CashAccount38 `xml:"OrgnlDbtrAgtAcct,omitempty" json:"orgnl_dbtr_agt_acct" `

	// Original final collection date that has been modified.
	OrgnlFnlColltnDt *ISODate `xml:"OrgnlFnlColltnDt,omitempty" json:"orgnl_fnl_colltn_dt" `

	// Original frequency that has been modified.
	OrgnlFrqcy *Frequency36Choice `xml:"OrgnlFrqcy,omitempty" json:"orgnl_frqcy" `

	// Original reason for the mandate to allow the user to distinguish between different mandates for the same creditor.
	OrgnlRsn *MandateSetupReason1Choice `xml:"OrgnlRsn,omitempty" json:"orgnl_rsn" `

	// Original number of tracking days that has been modified.
	OrgnlTrckgDays *Exact2NumericText `xml:"OrgnlTrckgDays,omitempty" json:"orgnl_trckg_days" `
}

func (a *AmendmentInformationDetails13) SetOrgnlMndtId(value string) {
	a.OrgnlMndtId = (*Max35Text)(&value)
}

func (a *AmendmentInformationDetails13) AddOrgnlCdtrSchmeId() *PartyIdentification135 {
	a.OrgnlCdtrSchmeId = new(PartyIdentification135)
	return a.OrgnlCdtrSchmeId
}

func (a *AmendmentInformationDetails13) AddOrgnlCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	a.OrgnlCdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return a.OrgnlCdtrAgt
}

func (a *AmendmentInformationDetails13) AddOrgnlCdtrAgtAcct() *CashAccount38 {
	a.OrgnlCdtrAgtAcct = new(CashAccount38)
	return a.OrgnlCdtrAgtAcct
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtr() *PartyIdentification135 {
	a.OrgnlDbtr = new(PartyIdentification135)
	return a.OrgnlDbtr
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtrAcct() *CashAccount38 {
	a.OrgnlDbtrAcct = new(CashAccount38)
	return a.OrgnlDbtrAcct
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	a.OrgnlDbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return a.OrgnlDbtrAgt
}

func (a *AmendmentInformationDetails13) AddOrgnlDbtrAgtAcct() *CashAccount38 {
	a.OrgnlDbtrAgtAcct = new(CashAccount38)
	return a.OrgnlDbtrAgtAcct
}

func (a *AmendmentInformationDetails13) SetOrgnlFnlColltnDt(value string) {
	a.OrgnlFnlColltnDt = (*ISODate)(&value)
}

func (a *AmendmentInformationDetails13) AddOrgnlFrqcy() *Frequency36Choice {
	a.OrgnlFrqcy = new(Frequency36Choice)
	return a.OrgnlFrqcy
}

func (a *AmendmentInformationDetails13) AddOrgnlRsn() *MandateSetupReason1Choice {
	a.OrgnlRsn = new(MandateSetupReason1Choice)
	return a.OrgnlRsn
}

func (a *AmendmentInformationDetails13) SetOrgnlTrckgDays(value string) {
	a.OrgnlTrckgDays = (*Exact2NumericText)(&value)
}

// Specifies the amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
type AmountType4Choice struct {
	XMLName xml.Name

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	//
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt" json:"instd_amt" `

	// Amount of money to be moved between the debtor and creditor, expressed in the currency of the debtor's account, and the currency in which the amount is to be moved.
	EqvtAmt *EquivalentAmount2 `xml:"EqvtAmt" json:"eqvt_amt" `
}

func (a *AmountType4Choice) SetInstdAmt(value, currency string) {
	a.InstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountType4Choice) AddEqvtAmt() *EquivalentAmount2 {
	a.EqvtAmt = new(EquivalentAmount2)
	return a.EqvtAmt
}

type AnyBICDec2014Identifier string

type AnyBICIdentifier string

// Provides the details on the user identification or any user key that allows to check if the initiating party is allowed to issue the transaction.
type Authorisation1Choice struct {
	XMLName xml.Name

	// Specifies the authorisation, in a coded form.
	Cd *Authorisation1Code `xml:"Cd" json:"cd" `

	// Specifies the authorisation, in a free text form.
	Prtry *Max128Text `xml:"Prtry" json:"prtry" `
}

func (a *Authorisation1Choice) SetCd(value string) {
	a.Cd = (*Authorisation1Code)(&value)
}

func (a *Authorisation1Choice) SetPrtry(value string) {
	a.Prtry = (*Max128Text)(&value)
}

type Authorisation1Code string

type BICFIDec2014Identifier string

type BICFIIdentifier string

type BaseOneRate string

type BatchBookingIndicator string

// Set of elements used to uniquely and unambiguously identify a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification5 struct {
	XMLName xml.Name

	// Unique and unambiguous identification of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinInstnId *FinancialInstitutionIdentification8 `xml:"FinInstnId" json:"fin_instn_id" `

	// Identifies a specific branch of a financial institution.
	//
	// Usage: This component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BrnchId *BranchData2 `xml:"BrnchId,omitempty" json:"brnch_id" `
}

func (b *BranchAndFinancialInstitutionIdentification5) AddFinInstnId() *FinancialInstitutionIdentification8 {
	b.FinInstnId = new(FinancialInstitutionIdentification8)
	return b.FinInstnId
}

func (b *BranchAndFinancialInstitutionIdentification5) AddBrnchId() *BranchData2 {
	b.BrnchId = new(BranchData2)
	return b.BrnchId
}

// Unique and unambiguous identification of a financial institution or a branch of a financial institution.
type BranchAndFinancialInstitutionIdentification6 struct {
	XMLName xml.Name

	// Unique and unambiguous identification of a financial institution, as assigned under an internationally recognised or proprietary identification scheme.
	FinInstnId *FinancialInstitutionIdentification18 `xml:"FinInstnId" json:"fin_instn_id" `

	// Identifies a specific branch of a financial institution.
	//
	// Usage: This component should be used in case the identification information in the financial institution component does not provide identification up to branch level.
	BrnchId *BranchData3 `xml:"BrnchId,omitempty" json:"brnch_id" `
}

func (b *BranchAndFinancialInstitutionIdentification6) AddFinInstnId() *FinancialInstitutionIdentification18 {
	b.FinInstnId = new(FinancialInstitutionIdentification18)
	return b.FinInstnId
}

func (b *BranchAndFinancialInstitutionIdentification6) AddBrnchId() *BranchData3 {
	b.BrnchId = new(BranchData3)
	return b.BrnchId
}

// Information that locates and identifies a specific branch of a financial institution.
type BranchData2 struct {
	XMLName xml.Name

	// Unique and unambiguous identification of a branch of a financial institution.
	Id *Max35Text `xml:"Id,omitempty" json:"id" `

	// Name by which an agent is known and which is usually used to identify that agent.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress6 `xml:"PstlAdr,omitempty" json:"pstl_adr" `
}

func (b *BranchData2) SetId(value string) {
	b.Id = (*Max35Text)(&value)
}

func (b *BranchData2) SetNm(value string) {
	b.Nm = (*Max140Text)(&value)
}

func (b *BranchData2) AddPstlAdr() *PostalAddress6 {
	b.PstlAdr = new(PostalAddress6)
	return b.PstlAdr
}

// Information that locates and identifies a specific branch of a financial institution.
type BranchData3 struct {
	XMLName xml.Name

	// Unique and unambiguous identification of a branch of a financial institution.
	Id *Max35Text `xml:"Id,omitempty" json:"id" `

	// Legal entity identification for the branch of the financial institution.
	LEI *LEIIdentifier `xml:"LEI,omitempty" json:"lei" `

	// Name by which an agent is known and which is usually used to identify that agent.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress24 `xml:"PstlAdr,omitempty" json:"pstl_adr" `
}

func (b *BranchData3) SetId(value string) {
	b.Id = (*Max35Text)(&value)
}

func (b *BranchData3) SetLEI(value string) {
	b.LEI = (*LEIIdentifier)(&value)
}

func (b *BranchData3) SetNm(value string) {
	b.Nm = (*Max140Text)(&value)
}

func (b *BranchData3) AddPstlAdr() *PostalAddress24 {
	b.PstlAdr = new(PostalAddress24)
	return b.PstlAdr
}

// Specifies the Business Application Header of the Business Message.
// Can be used when replying to a query; can also be used when canceling or amending.
type BusinessApplicationHeader1 struct {
	XMLName xml.Name

	// Contains the character set of the text-based elements used in the Business Message.
	CharSet *UnicodeChartsCode `xml:"CharSet,omitempty" json:"char_set" `

	// The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.
	//
	// Note	the sending MessagingEndpoint might be different from the sending address potentially contained in the transport header (as defined in the transport layer).
	Fr *Party9Choice `xml:"Fr" json:"fr" `

	// The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.
	//
	// Note the receiving MessagingEndpoint might be different from the receiving address potentially contained in the transport header (as defined in the transport layer).
	To *Party9Choice `xml:"To" json:"to" `

	// Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.
	BizMsgIdr *Max35Text `xml:"BizMsgIdr" json:"biz_msg_idr" `

	// Contains the MessageIdentifier that defines the BusinessMessage.
	// It must contain a MessageIdentifier published on the ISO 20022 website.
	//
	// example	camt.001.001.03.
	MsgDefIdr *Max35Text `xml:"MsgDefIdr" json:"msg_def_idr" `

	// Specifies the business service agreed between the two MessagingEndpoints under which rules this Business Message is exchanged.
	//  To be used when there is a choice of processing services or processing service levels.
	// Example: E&I.
	BizSvc *Max35Text `xml:"BizSvc,omitempty" json:"biz_svc" `

	// Date and time when this Business Message (header) was created.
	// Note Times must be normalized, using the "Z" annotation.
	CreDt *ISONormalisedDateTime `xml:"CreDt" json:"cre_dt" `

	// Indicates whether the message is a Copy, a Duplicate or a copy of a duplicate of a previously sent ISO 20022 Message.
	CpyDplct *CopyDuplicate1Code `xml:"CpyDplct,omitempty" json:"cpy_dplct" `

	// Flag indicating if the Business Message exchanged between the MessagingEndpoints is possibly a duplicate.
	// If the receiving MessagingEndpoint did not receive the original, then this Business Message should be processed as if it were the original.
	//
	// If the receiving MessagingEndpoint did receive the original, then it should perform necessary actions to avoid processing this Business Message again.
	//
	// This will guarantee business idempotent behaviour.
	//
	// NOTE: this is named "PossResend" in FIX - this is an application level resend not a network level retransmission.
	PssblDplct *YesNoIndicator `xml:"PssblDplct,omitempty" json:"pssbl_dplct" `

	// Relative indication of the processing precedence of the message over a (set of) Business Messages with assigned priorities.
	Prty *BusinessMessagePriorityCode `xml:"Prty,omitempty" json:"prty" `

	// Contains the digital signature of the Business Entity authorised to sign this Business Message.
	Sgntr *SignatureEnvelope `xml:"Sgntr,omitempty" json:"sgntr" `
}

func (b *BusinessApplicationHeader1) SetCharSet(value string) {
	b.CharSet = (*UnicodeChartsCode)(&value)
}

func (b *BusinessApplicationHeader1) AddFr() *Party9Choice {
	b.Fr = new(Party9Choice)
	return b.Fr
}

func (b *BusinessApplicationHeader1) AddTo() *Party9Choice {
	b.To = new(Party9Choice)
	return b.To
}

func (b *BusinessApplicationHeader1) SetBizMsgIdr(value string) {
	b.BizMsgIdr = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeader1) SetMsgDefIdr(value string) {
	b.MsgDefIdr = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeader1) SetBizSvc(value string) {
	b.BizSvc = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeader1) SetCreDt(value string) {
	b.CreDt = (*ISONormalisedDateTime)(&value)
}

func (b *BusinessApplicationHeader1) SetCpyDplct(value string) {
	b.CpyDplct = (*CopyDuplicate1Code)(&value)
}

func (b *BusinessApplicationHeader1) SetPssblDplct(value string) {
	b.PssblDplct = (*YesNoIndicator)(&value)
}

func (b *BusinessApplicationHeader1) SetPrty(value string) {
	b.Prty = (*BusinessMessagePriorityCode)(&value)
}

func (b *BusinessApplicationHeader1) AddSgntr() *SignatureEnvelope {
	b.Sgntr = new(SignatureEnvelope)
	return b.Sgntr
}

type BusinessMessagePriorityCode string

type CancellationIndividualStatus1Code string

// Specifies the reason for the cancellation request.
type CancellationReason33Choice struct {
	XMLName xml.Name

	// Reason for the cancellation request, in a coded form.
	Cd *ExternalCancellationReason1Code `xml:"Cd" json:"cd" `

	// Reason for the cancellation request, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *CancellationReason33Choice) SetCd(value string) {
	c.Cd = (*ExternalCancellationReason1Code)(&value)
}

func (c *CancellationReason33Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Specifies the reason for the transaction cancellation status.
type CancellationStatusReason3Choice struct {
	XMLName xml.Name

	// Reason for the cancellation status, in a coded form.
	Cd *ExternalPaymentCancellationRejection1Code `xml:"Cd" json:"cd" `

	// Reason for the status, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *CancellationStatusReason3Choice) SetCd(value string) {
	c.Cd = (*ExternalPaymentCancellationRejection1Code)(&value)
}

func (c *CancellationStatusReason3Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Provides further details on the status of the cancellation request.
type CancellationStatusReason4 struct {
	XMLName xml.Name

	// Party that issues the cancellation status.
	Orgtr *PartyIdentification135 `xml:"Orgtr,omitempty" json:"orgtr" `

	// Specifies the reason for the status report.
	Rsn *CancellationStatusReason3Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Further details on the cancellation status reason.
	AddtlInf []*Max105Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (c *CancellationStatusReason4) AddOrgtr() *PartyIdentification135 {
	c.Orgtr = new(PartyIdentification135)
	return c.Orgtr
}

func (c *CancellationStatusReason4) AddRsn() *CancellationStatusReason3Choice {
	c.Rsn = new(CancellationStatusReason3Choice)
	return c.Rsn
}

func (c *CancellationStatusReason4) AddAddtlInf(value string) {
	c.AddtlInf = append(c.AddtlInf, (*Max105Text)(&value))
}

// Provides further details to identify an investigation case.
type Case5 struct {
	XMLName xml.Name

	// Uniquely identifies the case.
	Id *Max35Text `xml:"Id" json:"id" `

	// Party that created the investigation case.
	Cretr *Party40Choice `xml:"Cretr" json:"cretr" `

	// Indicates whether or not the case was previously closed and is now re-opened.
	ReopCaseIndctn *YesNoIndicator `xml:"ReopCaseIndctn,omitempty" json:"reop_case_indctn" `
}

func (c *Case5) SetId(value string) {
	c.Id = (*Max35Text)(&value)
}

func (c *Case5) AddCretr() *Party40Choice {
	c.Cretr = new(Party40Choice)
	return c.Cretr
}

func (c *Case5) SetReopCaseIndctn(value string) {
	c.ReopCaseIndctn = (*YesNoIndicator)(&value)
}

// Represents the assignment of a case to a party.
type CaseAssignment5 struct {
	XMLName xml.Name

	// Uniquely identifies the case assignment.
	Id *Max35Text `xml:"Id" json:"id" `

	// Party who assigns the case.
	// Usage: This is also the sender of the message.
	Assgnr *Party40Choice `xml:"Assgnr" json:"assgnr" `

	// Party to which the case is assigned.
	// Usage: This is also the receiver of the message.
	Assgne *Party40Choice `xml:"Assgne" json:"assgne" `

	// Date and time at which the assignment was created.
	CreDtTm *ISODateTime `xml:"CreDtTm" json:"cre_dt_tm" `
}

func (c *CaseAssignment5) SetId(value string) {
	c.Id = (*Max35Text)(&value)
}

func (c *CaseAssignment5) AddAssgnr() *Party40Choice {
	c.Assgnr = new(Party40Choice)
	return c.Assgnr
}

func (c *CaseAssignment5) AddAssgne() *Party40Choice {
	c.Assgne = new(Party40Choice)
	return c.Assgne
}

func (c *CaseAssignment5) SetCreDtTm(value string) {
	c.CreDtTm = (*ISODateTime)(&value)
}

// Status of a case resulting from a case assignment.
type CaseForwardingNotification3 struct {
	XMLName xml.Name

	// Justification for the forward action.
	Justfn *CaseForwardingNotification3Code `xml:"Justfn" json:"justfn" `
}

func (c *CaseForwardingNotification3) SetJustfn(value string) {
	c.Justfn = (*CaseForwardingNotification3Code)(&value)
}

type CaseForwardingNotification3Code string

// Provides the details to identify an account.
type CashAccount24 struct {
	XMLName xml.Name

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Id *AccountIdentification4Choice `xml:"Id" json:"id" `

	// Specifies the nature, or use of the account.
	Tp *CashAccountType2Choice `xml:"Tp,omitempty" json:"tp" `

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies
	// and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Ccy *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:"ccy" `

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Nm *Max70Text `xml:"Nm,omitempty" json:"nm" `
}

func (c *CashAccount24) AddId() *AccountIdentification4Choice {
	c.Id = new(AccountIdentification4Choice)
	return c.Id
}

func (c *CashAccount24) AddTp() *CashAccountType2Choice {
	c.Tp = new(CashAccountType2Choice)
	return c.Tp
}

func (c *CashAccount24) SetCcy(value string) {
	c.Ccy = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount24) SetNm(value string) {
	c.Nm = (*Max70Text)(&value)
}

// Provides the details to identify an account.
type CashAccount38 struct {
	XMLName xml.Name

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Id *AccountIdentification4Choice `xml:"Id" json:"id" `

	// Specifies the nature, or use of the account.
	Tp *CashAccountType2Choice `xml:"Tp,omitempty" json:"tp" `

	// Identification of the currency in which the account is held.
	//
	// Usage: Currency should only be used in case one and the same account number covers several currencies
	// and the initiating party needs to identify which currency needs to be used for settlement on the account.
	Ccy *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty" json:"ccy" `

	// Name of the account, as assigned by the account servicing institution, in agreement with the account owner in order to provide an additional means of identification of the account.
	//
	// Usage: The account name is different from the account owner name. The account name is used in certain user communities to provide a means of identifying the account, in addition to the account owner's identity and the account number.
	Nm *Max70Text `xml:"Nm,omitempty" json:"nm" `

	// Specifies an alternate assumed name for the identification of the account.
	Prxy *ProxyAccountIdentification1 `xml:"Prxy,omitempty" json:"prxy" `
}

func (c *CashAccount38) AddId() *AccountIdentification4Choice {
	c.Id = new(AccountIdentification4Choice)
	return c.Id
}

func (c *CashAccount38) AddTp() *CashAccountType2Choice {
	c.Tp = new(CashAccountType2Choice)
	return c.Tp
}

func (c *CashAccount38) SetCcy(value string) {
	c.Ccy = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (c *CashAccount38) SetNm(value string) {
	c.Nm = (*Max70Text)(&value)
}

func (c *CashAccount38) AddPrxy() *ProxyAccountIdentification1 {
	c.Prxy = new(ProxyAccountIdentification1)
	return c.Prxy
}

// Nature or use of the account.
type CashAccountType2Choice struct {
	XMLName xml.Name

	// Account type, in a coded form.
	Cd *ExternalCashAccountType1Code `xml:"Cd" json:"cd" `

	// Nature or use of the account in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *CashAccountType2Choice) SetCd(value string) {
	c.Cd = (*ExternalCashAccountType1Code)(&value)
}

func (c *CashAccountType2Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
type CategoryPurpose1Choice struct {
	XMLName xml.Name

	// Category purpose, as published in an external category purpose code list.
	Cd *ExternalCategoryPurpose1Code `xml:"Cd" json:"cd" `

	// Category purpose, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *CategoryPurpose1Choice) SetCd(value string) {
	c.Cd = (*ExternalCategoryPurpose1Code)(&value)
}

func (c *CategoryPurpose1Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

type ChargeBearerType1Code string

type ChargeIncludedIndicator string

// Specifies the charge type.
type ChargeType3Choice struct {
	XMLName xml.Name

	// Charge type, in a coded form.
	Cd *ExternalChargeType1Code `xml:"Cd" json:"cd" `

	// Type of charge in a proprietary form, as defined by the issuer.
	Prtry *GenericIdentification3 `xml:"Prtry" json:"prtry" `
}

func (c *ChargeType3Choice) SetCd(value string) {
	c.Cd = (*ExternalChargeType1Code)(&value)
}

func (c *ChargeType3Choice) AddPrtry() *GenericIdentification3 {
	c.Prtry = new(GenericIdentification3)
	return c.Prtry
}

// Set of elements used to provide information on the charges related to the payment transaction.
type Charges2 struct {
	XMLName xml.Name

	// Transaction charges to be paid by the charge bearer.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"agt" `
}

func (c *Charges2) SetAmt(value, currency string) {
	c.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges2) AddAgt() *BranchAndFinancialInstitutionIdentification5 {
	c.Agt = new(BranchAndFinancialInstitutionIdentification5)
	return c.Agt
}

// Provides further details on the charges related to the payment transaction.
type Charges6 struct {
	XMLName xml.Name

	// Total of all charges and taxes applied to the entry.
	TtlChrgsAndTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlChrgsAndTaxAmt,omitempty" json:"ttl_chrgs_and_tax_amt" `

	// Provides details of the individual charges record.
	Rcrd []*ChargesRecord3 `xml:"Rcrd,omitempty" json:"rcrd" `
}

func (c *Charges6) SetTtlChrgsAndTaxAmt(value, currency string) {
	c.TtlChrgsAndTaxAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges6) AddRcrd() *ChargesRecord3 {
	newValue := new(ChargesRecord3)
	c.Rcrd = append(c.Rcrd, newValue)
	return newValue
}

// Provides information on the charges related to the payment transaction.
type Charges7 struct {
	XMLName xml.Name

	// Transaction charges to be paid by the charge bearer.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt" json:"agt" `
}

func (c *Charges7) SetAmt(value, currency string) {
	c.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *Charges7) AddAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.Agt = new(BranchAndFinancialInstitutionIdentification6)
	return c.Agt
}

// Provides further individual record details on the charges related to the payment transaction.
type ChargesRecord3 struct {
	XMLName xml.Name

	// Transaction charges to be paid by the charge bearer.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `

	// Indicates whether the charges amount is a credit or a debit amount.
	// Usage: A zero amount is considered to be a credit.
	CdtDbtInd *CreditDebitCode `xml:"CdtDbtInd,omitempty" json:"cdt_dbt_ind" `

	// Indicates whether the charge should be included in the amount or is added as pre-advice.
	ChrgInclInd *ChargeIncludedIndicator `xml:"ChrgInclInd,omitempty" json:"chrg_incl_ind" `

	// Specifies the type of charge.
	Tp *ChargeType3Choice `xml:"Tp,omitempty" json:"tp" `

	// Rate used to calculate the amount of the charge or fee.
	Rate *PercentageRate `xml:"Rate,omitempty" json:"rate" `

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	Br *ChargeBearerType1Code `xml:"Br,omitempty" json:"br" `

	// Agent that takes the transaction charges or to which the transaction charges are due.
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt,omitempty" json:"agt" `

	// Provides details on the tax applied to charges.
	Tax *TaxCharges2 `xml:"Tax,omitempty" json:"tax" `
}

func (c *ChargesRecord3) SetAmt(value, currency string) {
	c.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *ChargesRecord3) SetCdtDbtInd(value string) {
	c.CdtDbtInd = (*CreditDebitCode)(&value)
}

func (c *ChargesRecord3) SetChrgInclInd(value string) {
	c.ChrgInclInd = (*ChargeIncludedIndicator)(&value)
}

func (c *ChargesRecord3) AddTp() *ChargeType3Choice {
	c.Tp = new(ChargeType3Choice)
	return c.Tp
}

func (c *ChargesRecord3) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargesRecord3) SetBr(value string) {
	c.Br = (*ChargeBearerType1Code)(&value)
}

func (c *ChargesRecord3) AddAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.Agt = new(BranchAndFinancialInstitutionIdentification6)
	return c.Agt
}

func (c *ChargesRecord3) AddTax() *TaxCharges2 {
	c.Tax = new(TaxCharges2)
	return c.Tax
}

// Specifies details of a claim non receipt status advice.
type ClaimNonReceipt2 struct {
	XMLName xml.Name

	// Specifies the date the original payment instruction was processed.
	DtPrcd *ISODate `xml:"DtPrcd" json:"dt_prcd" `

	// Specifies the next party the original payment instruction was sent to.
	OrgnlNxtAgt *BranchAndFinancialInstitutionIdentification6 `xml:"OrgnlNxtAgt,omitempty" json:"orgnl_nxt_agt" `
}

func (c *ClaimNonReceipt2) SetDtPrcd(value string) {
	c.DtPrcd = (*ISODate)(&value)
}

func (c *ClaimNonReceipt2) AddOrgnlNxtAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.OrgnlNxtAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.OrgnlNxtAgt
}

// Specifies whether the claim non receipt is accepted or rejected.
type ClaimNonReceipt2Choice struct {
	XMLName xml.Name

	// Claim non-receipt is accepted and processed by the agent.
	Accptd *ClaimNonReceipt2 `xml:"Accptd" json:"accptd" `

	// Specifies that the claim non receipt has been rejected and provides the reason for the rejection.
	Rjctd *ClaimNonReceiptRejectReason1Choice `xml:"Rjctd" json:"rjctd" `
}

func (c *ClaimNonReceipt2Choice) AddAccptd() *ClaimNonReceipt2 {
	c.Accptd = new(ClaimNonReceipt2)
	return c.Accptd
}

func (c *ClaimNonReceipt2Choice) AddRjctd() *ClaimNonReceiptRejectReason1Choice {
	c.Rjctd = new(ClaimNonReceiptRejectReason1Choice)
	return c.Rjctd
}

// Specifies the rejection reason of a claim non receipt.
type ClaimNonReceiptRejectReason1Choice struct {
	XMLName xml.Name

	// Reason for the rejection, in a coded form.
	Cd *ExternalClaimNonReceiptRejection1Code `xml:"Cd" json:"cd" `

	// Reason for the rejection, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *ClaimNonReceiptRejectReason1Choice) SetCd(value string) {
	c.Cd = (*ExternalClaimNonReceiptRejection1Code)(&value)
}

func (c *ClaimNonReceiptRejectReason1Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

type ClearingChannel2Code string

// Choice of a clearing system identifier.
type ClearingSystemIdentification2Choice struct {
	XMLName xml.Name

	// Identification of a clearing system, in a coded form as published in an external list.
	Cd *ExternalClearingSystemIdentification1Code `xml:"Cd" json:"cd" `

	// Identification code for a clearing system, that has not yet been identified in the list of clearing systems.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *ClearingSystemIdentification2Choice) SetCd(value string) {
	c.Cd = (*ExternalClearingSystemIdentification1Code)(&value)
}

func (c *ClearingSystemIdentification2Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Specifies the clearing system identification.
type ClearingSystemIdentification3Choice struct {
	XMLName xml.Name

	// Infrastructure through which the payment instruction is processed, as published in an external clearing system identification code list.
	Cd *ExternalCashClearingSystem1Code `xml:"Cd" json:"cd" `

	// Clearing system identification in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *ClearingSystemIdentification3Choice) SetCd(value string) {
	c.Cd = (*ExternalCashClearingSystem1Code)(&value)
}

func (c *ClearingSystemIdentification3Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Unique identification, as assigned by a clearing system, to unambiguously identify a member of the clearing system.
type ClearingSystemMemberIdentification2 struct {
	XMLName xml.Name

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClrSysId *ClearingSystemIdentification2Choice `xml:"ClrSysId,omitempty" json:"clr_sys_id" `

	// Identification of a member of a clearing system.
	MmbId *Max35Text `xml:"MmbId" json:"mmb_id" `
}

func (c *ClearingSystemMemberIdentification2) AddClrSysId() *ClearingSystemIdentification2Choice {
	c.ClrSysId = new(ClearingSystemIdentification2Choice)
	return c.ClrSysId
}

func (c *ClearingSystemMemberIdentification2) SetMmbId(value string) {
	c.MmbId = (*Max35Text)(&value)
}

// Specifies the details of a payment compensation.
type Compensation2 struct {
	XMLName xml.Name

	// Amount of money to be paid in compensation.
	Amt *ActiveCurrencyAndAmount `xml:"Amt" json:"amt" `

	// Financial institution servicing an account for the debtor.
	//
	// Usage: The debtor agent is the payer of the compensation amount.
	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"dbtr_agt" `

	// Financial institution servicing an account for the creditor.
	//
	// Usage: The creditor agent is the payee of the compensation amount.
	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"cdtr_agt" `

	// Reason for the payment compensation.
	Rsn *CompensationReason1Choice `xml:"Rsn" json:"rsn" `
}

func (c *Compensation2) SetAmt(value, currency string) {
	c.Amt = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Compensation2) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.DbtrAgt
}

func (c *Compensation2) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.CdtrAgt
}

func (c *Compensation2) AddRsn() *CompensationReason1Choice {
	c.Rsn = new(CompensationReason1Choice)
	return c.Rsn
}

// Specifies the reason for the payment compensation.
type CompensationReason1Choice struct {
	XMLName xml.Name

	// Reason for the payment compensation, in a coded form.
	Cd *ExternalPaymentCompensationReason1Code `xml:"Cd" json:"cd" `

	// Reason for the payment compensation, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *CompensationReason1Choice) SetCd(value string) {
	c.Cd = (*ExternalPaymentCompensationReason1Code)(&value)
}

func (c *CompensationReason1Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Specifies the details of the contact person.
type Contact4 struct {
	XMLName xml.Name

	// Specifies the terms used to formally address a person.
	NmPrfx *NamePrefix2Code `xml:"NmPrfx,omitempty" json:"nm_prfx" `

	// Name by which a party is known and which is usually used to identify that party.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhneNb *PhoneNumber `xml:"PhneNb,omitempty" json:"phne_nb" `

	// Collection of information that identifies a mobile phone number, as defined by telecom services.
	MobNb *PhoneNumber `xml:"MobNb,omitempty" json:"mob_nb" `

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNb *PhoneNumber `xml:"FaxNb,omitempty" json:"fax_nb" `

	// Address for electronic mail (e-mail).
	EmailAdr *Max2048Text `xml:"EmailAdr,omitempty" json:"email_adr" `

	// Purpose for which an email address may be used.
	EmailPurp *Max35Text `xml:"EmailPurp,omitempty" json:"email_purp" `

	// Title of the function.
	JobTitl *Max35Text `xml:"JobTitl,omitempty" json:"job_titl" `

	// Role of a person in an organisation.
	Rspnsblty *Max35Text `xml:"Rspnsblty,omitempty" json:"rspnsblty" `

	// Identification of a division of a large organisation or building.
	Dept *Max70Text `xml:"Dept,omitempty" json:"dept" `

	// Contact details in another form.
	Othr []*OtherContact1 `xml:"Othr,omitempty" json:"othr" `

	// Preferred method used to reach the contact.
	PrefrdMtd *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty" json:"prefrd_mtd" `
}

func (c *Contact4) SetNmPrfx(value string) {
	c.NmPrfx = (*NamePrefix2Code)(&value)
}

func (c *Contact4) SetNm(value string) {
	c.Nm = (*Max140Text)(&value)
}

func (c *Contact4) SetPhneNb(value string) {
	c.PhneNb = (*PhoneNumber)(&value)
}

func (c *Contact4) SetMobNb(value string) {
	c.MobNb = (*PhoneNumber)(&value)
}

func (c *Contact4) SetFaxNb(value string) {
	c.FaxNb = (*PhoneNumber)(&value)
}

func (c *Contact4) SetEmailAdr(value string) {
	c.EmailAdr = (*Max2048Text)(&value)
}

func (c *Contact4) SetEmailPurp(value string) {
	c.EmailPurp = (*Max35Text)(&value)
}

func (c *Contact4) SetJobTitl(value string) {
	c.JobTitl = (*Max35Text)(&value)
}

func (c *Contact4) SetRspnsblty(value string) {
	c.Rspnsblty = (*Max35Text)(&value)
}

func (c *Contact4) SetDept(value string) {
	c.Dept = (*Max70Text)(&value)
}

func (c *Contact4) AddOthr() *OtherContact1 {
	newValue := new(OtherContact1)
	c.Othr = append(c.Othr, newValue)
	return newValue
}

func (c *Contact4) SetPrefrdMtd(value string) {
	c.PrefrdMtd = (*PreferredContactMethod1Code)(&value)
}

// Communication device number or electronic address used for communication.
type ContactDetails2 struct {
	XMLName xml.Name

	// Specifies the terms used to formally address a person.
	NmPrfx *NamePrefix1Code `xml:"NmPrfx,omitempty" json:"nm_prfx" `

	// Name by which a party is known and which is usually used to identify that party.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhneNb *PhoneNumber `xml:"PhneNb,omitempty" json:"phne_nb" `

	// Collection of information that identifies a mobile phone number, as defined by telecom services.
	MobNb *PhoneNumber `xml:"MobNb,omitempty" json:"mob_nb" `

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNb *PhoneNumber `xml:"FaxNb,omitempty" json:"fax_nb" `

	// Address for electronic mail (e-mail).
	EmailAdr *Max2048Text `xml:"EmailAdr,omitempty" json:"email_adr" `

	// Contact details in another form.
	Othr *Max35Text `xml:"Othr,omitempty" json:"othr" `
}

func (c *ContactDetails2) SetNmPrfx(value string) {
	c.NmPrfx = (*NamePrefix1Code)(&value)
}

func (c *ContactDetails2) SetNm(value string) {
	c.Nm = (*Max140Text)(&value)
}

func (c *ContactDetails2) SetPhneNb(value string) {
	c.PhneNb = (*PhoneNumber)(&value)
}

func (c *ContactDetails2) SetMobNb(value string) {
	c.MobNb = (*PhoneNumber)(&value)
}

func (c *ContactDetails2) SetFaxNb(value string) {
	c.FaxNb = (*PhoneNumber)(&value)
}

func (c *ContactDetails2) SetEmailAdr(value string) {
	c.EmailAdr = (*Max2048Text)(&value)
}

func (c *ContactDetails2) SetOthr(value string) {
	c.Othr = (*Max35Text)(&value)
}

// Provides details of the number of transactions and the control sum of the message.
type ControlData1 struct {
	XMLName xml.Name

	// Number of individual transactions contained in the message.
	NbOfTxs *Max15NumericText `xml:"NbOfTxs" json:"nb_of_txs" `

	// Total of all individual amounts included in the message, irrespective of currencies.
	CtrlSum *DecimalNumber `xml:"CtrlSum,omitempty" json:"ctrl_sum" `
}

func (c *ControlData1) SetNbOfTxs(value string) {
	c.NbOfTxs = (*Max15NumericText)(&value)
}

func (c *ControlData1) SetCtrlSum(value string) {
	c.CtrlSum = (*DecimalNumber)(&value)
}

type CopyDuplicate1Code string

// Set of elements used to provide information on the group of the corrective transaction, to which the resolution message refers.
type CorrectiveGroupInformation1 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MsgId *Max35Text `xml:"MsgId" json:"msg_id" `

	// Specifies the message name identifier to which the message refers.
	MsgNmId *Max35Text `xml:"MsgNmId" json:"msg_nm_id" `

	// Date and time at which the message was created.
	CreDtTm *ISODateTime `xml:"CreDtTm,omitempty" json:"cre_dt_tm" `
}

func (c *CorrectiveGroupInformation1) SetMsgId(value string) {
	c.MsgId = (*Max35Text)(&value)
}

func (c *CorrectiveGroupInformation1) SetMsgNmId(value string) {
	c.MsgNmId = (*Max35Text)(&value)
}

func (c *CorrectiveGroupInformation1) SetCreDtTm(value string) {
	c.CreDtTm = (*ISODateTime)(&value)
}

// Provides information on the corrective interbank transaction, to which the resolution message refers.
type CorrectiveInterbankTransaction2 struct {
	XMLName xml.Name

	// Set of elements used to provide corrective information for the group header of the message under investigation.
	GrpHdr *CorrectiveGroupInformation1 `xml:"GrpHdr,omitempty" json:"grp_hdr" `

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstrId *Max35Text `xml:"InstrId,omitempty" json:"instr_id" `

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndId *Max35Text `xml:"EndToEndId,omitempty" json:"end_to_end_id" `

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TxId *Max35Text `xml:"TxId,omitempty" json:"tx_id" `

	// Universally unique identifier to provide an end-to-end reference of a payment transaction.
	UETR *UUIDv4Identifier `xml:"UETR,omitempty" json:"uetr" `

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt" json:"intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt" json:"intr_bk_sttlm_dt" `
}

func (c *CorrectiveInterbankTransaction2) AddGrpHdr() *CorrectiveGroupInformation1 {
	c.GrpHdr = new(CorrectiveGroupInformation1)
	return c.GrpHdr
}

func (c *CorrectiveInterbankTransaction2) SetInstrId(value string) {
	c.InstrId = (*Max35Text)(&value)
}

func (c *CorrectiveInterbankTransaction2) SetEndToEndId(value string) {
	c.EndToEndId = (*Max35Text)(&value)
}

func (c *CorrectiveInterbankTransaction2) SetTxId(value string) {
	c.TxId = (*Max35Text)(&value)
}

func (c *CorrectiveInterbankTransaction2) SetUETR(value string) {
	c.UETR = (*UUIDv4Identifier)(&value)
}

func (c *CorrectiveInterbankTransaction2) SetIntrBkSttlmAmt(value, currency string) {
	c.IntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CorrectiveInterbankTransaction2) SetIntrBkSttlmDt(value string) {
	c.IntrBkSttlmDt = (*ISODate)(&value)
}

// Provides information on the corrective payment initiation transaction, to which the resolution message refers.
type CorrectivePaymentInitiation4 struct {
	XMLName xml.Name

	// Set of elements used to provide corrective information for the group header of the message under investigation.
	GrpHdr *CorrectiveGroupInformation1 `xml:"GrpHdr,omitempty" json:"grp_hdr" `

	// Unique identification, as assigned by a sending party, to unambiguously identify the payment information group within the message.
	PmtInfId *Max35Text `xml:"PmtInfId,omitempty" json:"pmt_inf_id" `

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstrId *Max35Text `xml:"InstrId,omitempty" json:"instr_id" `

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndId *Max35Text `xml:"EndToEndId,omitempty" json:"end_to_end_id" `

	// Universally unique identifier to provide an end-to-end reference of a payment transaction.
	UETR *UUIDv4Identifier `xml:"UETR,omitempty" json:"uetr" `

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt" json:"instd_amt" `

	// Date or date time at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date or date time on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty" json:"reqd_exctn_dt" `

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	ReqdColltnDt *ISODate `xml:"ReqdColltnDt,omitempty" json:"reqd_colltn_dt" `
}

func (c *CorrectivePaymentInitiation4) AddGrpHdr() *CorrectiveGroupInformation1 {
	c.GrpHdr = new(CorrectiveGroupInformation1)
	return c.GrpHdr
}

func (c *CorrectivePaymentInitiation4) SetPmtInfId(value string) {
	c.PmtInfId = (*Max35Text)(&value)
}

func (c *CorrectivePaymentInitiation4) SetInstrId(value string) {
	c.InstrId = (*Max35Text)(&value)
}

func (c *CorrectivePaymentInitiation4) SetEndToEndId(value string) {
	c.EndToEndId = (*Max35Text)(&value)
}

func (c *CorrectivePaymentInitiation4) SetUETR(value string) {
	c.UETR = (*UUIDv4Identifier)(&value)
}

func (c *CorrectivePaymentInitiation4) SetInstdAmt(value, currency string) {
	c.InstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CorrectivePaymentInitiation4) AddReqdExctnDt() *DateAndDateTime2Choice {
	c.ReqdExctnDt = new(DateAndDateTime2Choice)
	return c.ReqdExctnDt
}

func (c *CorrectivePaymentInitiation4) SetReqdColltnDt(value string) {
	c.ReqdColltnDt = (*ISODate)(&value)
}

// Specifies the corrective transaction on which the investigation is processed.
type CorrectiveTransaction4Choice struct {
	XMLName xml.Name

	// Set of elements used to reference the details of the corrective payment initiation.
	Initn *CorrectivePaymentInitiation4 `xml:"Initn" json:"initn" `

	// Set of elements used to reference the details of the corrective interbank payment transaction.
	IntrBk *CorrectiveInterbankTransaction2 `xml:"IntrBk" json:"intr_bk" `
}

func (c *CorrectiveTransaction4Choice) AddInitn() *CorrectivePaymentInitiation4 {
	c.Initn = new(CorrectivePaymentInitiation4)
	return c.Initn
}

func (c *CorrectiveTransaction4Choice) AddIntrBk() *CorrectiveInterbankTransaction2 {
	c.IntrBk = new(CorrectiveInterbankTransaction2)
	return c.IntrBk
}

type CountryCode string

type CreditDebitCode string

// Provides further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction30 struct {
	XMLName xml.Name

	// Set of elements used to reference a payment instruction.
	PmtId *PaymentIdentification3 `xml:"PmtId" json:"pmt_id" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt" json:"intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SttlmPrty *Priority3Code `xml:"SttlmPrty,omitempty" json:"sttlm_prty" `

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SttlmTmIndctn *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty" json:"sttlm_tm_indctn" `

	// Provides information on the requested settlement time(s) of the payment instruction.
	SttlmTmReq *SettlementTimeRequest2 `xml:"SttlmTmReq,omitempty" json:"sttlm_tm_req" `

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AccptncDtTm *ISODateTime `xml:"AccptncDtTm,omitempty" json:"accptnc_dt_tm" `

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolgAdjstmntDt *ISODate `xml:"PoolgAdjstmntDt,omitempty" json:"poolg_adjstmnt_dt" `

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty" json:"instd_amt" `

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	XchgRate *BaseOneRate `xml:"XchgRate,omitempty" json:"xchg_rate" `

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChrgBr *ChargeBearerType1Code `xml:"ChrgBr" json:"chrg_br" `

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	ChrgsInf []*Charges2 `xml:"ChrgsInf,omitempty" json:"chrgs_inf" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt1 *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt1,omitempty" json:"prvs_instg_agt_1" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt1Acct *CashAccount24 `xml:"PrvsInstgAgt1Acct,omitempty" json:"prvs_instg_agt_1_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt2 *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt2,omitempty" json:"prvs_instg_agt_2" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt2Acct *CashAccount24 `xml:"PrvsInstgAgt2Acct,omitempty" json:"prvs_instg_agt_2_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt3 *BranchAndFinancialInstitutionIdentification5 `xml:"PrvsInstgAgt3,omitempty" json:"prvs_instg_agt_3" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt3Acct *CashAccount24 `xml:"PrvsInstgAgt3Acct,omitempty" json:"prvs_instg_agt_3_acct" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstgAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstdAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty" json:"instd_agt" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt1,omitempty" json:"intrmy_agt_1" `

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntrmyAgt1Acct *CashAccount24 `xml:"IntrmyAgt1Acct,omitempty" json:"intrmy_agt_1_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt2,omitempty" json:"intrmy_agt_2" `

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntrmyAgt2Acct *CashAccount24 `xml:"IntrmyAgt2Acct,omitempty" json:"intrmy_agt_2_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification5 `xml:"IntrmyAgt3,omitempty" json:"intrmy_agt_3" `

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntrmyAgt3Acct *CashAccount24 `xml:"IntrmyAgt3Acct,omitempty" json:"intrmy_agt_3_acct" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltmtDbtr *PartyIdentification125 `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Party that initiates the payment.
	// Usage: This can be either the debtor or a party that initiates the credit transfer on behalf of the debtor.
	InitgPty *PartyIdentification125 `xml:"InitgPty,omitempty" json:"initg_pty" `

	// Party that owes an amount of money to the (ultimate) creditor.
	Dbtr *PartyIdentification125 `xml:"Dbtr" json:"dbtr" `

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DbtrAcct *CashAccount24 `xml:"DbtrAcct,omitempty" json:"dbtr_acct" `

	// Financial institution servicing an account for the debtor.
	DbtrAgt *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt" json:"dbtr_agt" `

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DbtrAgtAcct *CashAccount24 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct" `

	// Financial institution servicing an account for the creditor.
	CdtrAgt *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt" json:"cdtr_agt" `

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CdtrAgtAcct *CashAccount24 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct" `

	// Party to which an amount of money is due.
	Cdtr *PartyIdentification125 `xml:"Cdtr" json:"cdtr" `

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CdtrAcct *CashAccount24 `xml:"CdtrAcct,omitempty" json:"cdtr_acct" `

	// Ultimate party to which an amount of money is due.
	UltmtCdtr *PartyIdentification125 `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstrForCdtrAgt []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty" json:"instr_for_cdtr_agt" `

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstrForNxtAgt []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty" json:"instr_for_nxt_agt" `

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purp *Purpose2Choice `xml:"Purp,omitempty" json:"purp" `

	// Information needed due to regulatory and statutory requirements.
	RgltryRptg []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty" json:"rgltry_rptg" `

	// Provides details on the tax.
	Tax *TaxInformation6 `xml:"Tax,omitempty" json:"tax" `

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RltdRmtInf []*RemittanceLocation4 `xml:"RltdRmtInf,omitempty" json:"rltd_rmt_inf" `

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RmtInf *RemittanceInformation15 `xml:"RmtInf,omitempty" json:"rmt_inf" `

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"SplmtryData,omitempty" json:"splmtry_data" `
}

func (c *CreditTransferTransaction30) AddPmtId() *PaymentIdentification3 {
	c.PmtId = new(PaymentIdentification3)
	return c.PmtId
}

func (c *CreditTransferTransaction30) AddPmtTpInf() *PaymentTypeInformation21 {
	c.PmtTpInf = new(PaymentTypeInformation21)
	return c.PmtTpInf
}

func (c *CreditTransferTransaction30) SetIntrBkSttlmAmt(value, currency string) {
	c.IntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction30) SetIntrBkSttlmDt(value string) {
	c.IntrBkSttlmDt = (*ISODate)(&value)
}

func (c *CreditTransferTransaction30) SetSttlmPrty(value string) {
	c.SttlmPrty = (*Priority3Code)(&value)
}

func (c *CreditTransferTransaction30) AddSttlmTmIndctn() *SettlementDateTimeIndication1 {
	c.SttlmTmIndctn = new(SettlementDateTimeIndication1)
	return c.SttlmTmIndctn
}

func (c *CreditTransferTransaction30) AddSttlmTmReq() *SettlementTimeRequest2 {
	c.SttlmTmReq = new(SettlementTimeRequest2)
	return c.SttlmTmReq
}

func (c *CreditTransferTransaction30) SetAccptncDtTm(value string) {
	c.AccptncDtTm = (*ISODateTime)(&value)
}

func (c *CreditTransferTransaction30) SetPoolgAdjstmntDt(value string) {
	c.PoolgAdjstmntDt = (*ISODate)(&value)
}

func (c *CreditTransferTransaction30) SetInstdAmt(value, currency string) {
	c.InstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction30) SetXchgRate(value string) {
	c.XchgRate = (*BaseOneRate)(&value)
}

func (c *CreditTransferTransaction30) SetChrgBr(value string) {
	c.ChrgBr = (*ChargeBearerType1Code)(&value)
}

func (c *CreditTransferTransaction30) AddChrgsInf() *Charges2 {
	newValue := new(Charges2)
	c.ChrgsInf = append(c.ChrgsInf, newValue)
	return newValue
}

func (c *CreditTransferTransaction30) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification5 {
	c.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification5)
	return c.PrvsInstgAgt1
}

func (c *CreditTransferTransaction30) AddPrvsInstgAgt1Acct() *CashAccount24 {
	c.PrvsInstgAgt1Acct = new(CashAccount24)
	return c.PrvsInstgAgt1Acct
}

func (c *CreditTransferTransaction30) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification5 {
	c.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification5)
	return c.PrvsInstgAgt2
}

func (c *CreditTransferTransaction30) AddPrvsInstgAgt2Acct() *CashAccount24 {
	c.PrvsInstgAgt2Acct = new(CashAccount24)
	return c.PrvsInstgAgt2Acct
}

func (c *CreditTransferTransaction30) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification5 {
	c.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification5)
	return c.PrvsInstgAgt3
}

func (c *CreditTransferTransaction30) AddPrvsInstgAgt3Acct() *CashAccount24 {
	c.PrvsInstgAgt3Acct = new(CashAccount24)
	return c.PrvsInstgAgt3Acct
}

func (c *CreditTransferTransaction30) AddInstgAgt() *BranchAndFinancialInstitutionIdentification5 {
	c.InstgAgt = new(BranchAndFinancialInstitutionIdentification5)
	return c.InstgAgt
}

func (c *CreditTransferTransaction30) AddInstdAgt() *BranchAndFinancialInstitutionIdentification5 {
	c.InstdAgt = new(BranchAndFinancialInstitutionIdentification5)
	return c.InstdAgt
}

func (c *CreditTransferTransaction30) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification5 {
	c.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntrmyAgt1
}

func (c *CreditTransferTransaction30) AddIntrmyAgt1Acct() *CashAccount24 {
	c.IntrmyAgt1Acct = new(CashAccount24)
	return c.IntrmyAgt1Acct
}

func (c *CreditTransferTransaction30) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification5 {
	c.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntrmyAgt2
}

func (c *CreditTransferTransaction30) AddIntrmyAgt2Acct() *CashAccount24 {
	c.IntrmyAgt2Acct = new(CashAccount24)
	return c.IntrmyAgt2Acct
}

func (c *CreditTransferTransaction30) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification5 {
	c.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification5)
	return c.IntrmyAgt3
}

func (c *CreditTransferTransaction30) AddIntrmyAgt3Acct() *CashAccount24 {
	c.IntrmyAgt3Acct = new(CashAccount24)
	return c.IntrmyAgt3Acct
}

func (c *CreditTransferTransaction30) AddUltmtDbtr() *PartyIdentification125 {
	c.UltmtDbtr = new(PartyIdentification125)
	return c.UltmtDbtr
}

func (c *CreditTransferTransaction30) AddInitgPty() *PartyIdentification125 {
	c.InitgPty = new(PartyIdentification125)
	return c.InitgPty
}

func (c *CreditTransferTransaction30) AddDbtr() *PartyIdentification125 {
	c.Dbtr = new(PartyIdentification125)
	return c.Dbtr
}

func (c *CreditTransferTransaction30) AddDbtrAcct() *CashAccount24 {
	c.DbtrAcct = new(CashAccount24)
	return c.DbtrAcct
}

func (c *CreditTransferTransaction30) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification5 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification5)
	return c.DbtrAgt
}

func (c *CreditTransferTransaction30) AddDbtrAgtAcct() *CashAccount24 {
	c.DbtrAgtAcct = new(CashAccount24)
	return c.DbtrAgtAcct
}

func (c *CreditTransferTransaction30) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification5 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification5)
	return c.CdtrAgt
}

func (c *CreditTransferTransaction30) AddCdtrAgtAcct() *CashAccount24 {
	c.CdtrAgtAcct = new(CashAccount24)
	return c.CdtrAgtAcct
}

func (c *CreditTransferTransaction30) AddCdtr() *PartyIdentification125 {
	c.Cdtr = new(PartyIdentification125)
	return c.Cdtr
}

func (c *CreditTransferTransaction30) AddCdtrAcct() *CashAccount24 {
	c.CdtrAcct = new(CashAccount24)
	return c.CdtrAcct
}

func (c *CreditTransferTransaction30) AddUltmtCdtr() *PartyIdentification125 {
	c.UltmtCdtr = new(PartyIdentification125)
	return c.UltmtCdtr
}

func (c *CreditTransferTransaction30) AddInstrForCdtrAgt() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstrForCdtrAgt = append(c.InstrForCdtrAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction30) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstrForNxtAgt = append(c.InstrForNxtAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction30) AddPurp() *Purpose2Choice {
	c.Purp = new(Purpose2Choice)
	return c.Purp
}

func (c *CreditTransferTransaction30) AddRgltryRptg() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	c.RgltryRptg = append(c.RgltryRptg, newValue)
	return newValue
}

func (c *CreditTransferTransaction30) AddTax() *TaxInformation6 {
	c.Tax = new(TaxInformation6)
	return c.Tax
}

func (c *CreditTransferTransaction30) AddRltdRmtInf() *RemittanceLocation4 {
	newValue := new(RemittanceLocation4)
	c.RltdRmtInf = append(c.RltdRmtInf, newValue)
	return newValue
}

func (c *CreditTransferTransaction30) AddRmtInf() *RemittanceInformation15 {
	c.RmtInf = new(RemittanceInformation15)
	return c.RmtInf
}

func (c *CreditTransferTransaction30) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SplmtryData = append(c.SplmtryData, newValue)
	return newValue
}

// Provide further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction36 struct {
	XMLName xml.Name

	// Set of elements used to reference a payment instruction.
	PmtId *PaymentIdentification7 `xml:"PmtId" json:"pmt_id" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation28 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"IntrBkSttlmAmt" json:"intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	SttlmPrty *Priority3Code `xml:"SttlmPrty,omitempty" json:"sttlm_prty" `

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SttlmTmIndctn *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty" json:"sttlm_tm_indctn" `

	// Provides information on the requested settlement time(s) of the payment instruction.
	SttlmTmReq *SettlementTimeRequest2 `xml:"SttlmTmReq,omitempty" json:"sttlm_tm_req" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"prvs_instg_agt_1" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt1Acct *CashAccount38 `xml:"PrvsInstgAgt1Acct,omitempty" json:"prvs_instg_agt_1_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"prvs_instg_agt_2" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt2Acct *CashAccount38 `xml:"PrvsInstgAgt2Acct,omitempty" json:"prvs_instg_agt_2_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"prvs_instg_agt_3" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt3Acct *CashAccount38 `xml:"PrvsInstgAgt3Acct,omitempty" json:"prvs_instg_agt_3_acct" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"instd_agt" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"intrmy_agt_1" `

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntrmyAgt1Acct *CashAccount38 `xml:"IntrmyAgt1Acct,omitempty" json:"intrmy_agt_1_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"intrmy_agt_2" `

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntrmyAgt2Acct *CashAccount38 `xml:"IntrmyAgt2Acct,omitempty" json:"intrmy_agt_2_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"intrmy_agt_3" `

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntrmyAgt3Acct *CashAccount38 `xml:"IntrmyAgt3Acct,omitempty" json:"intrmy_agt_3_acct" `

	// Ultimate financial institution that owes an amount of money to the (ultimate) institutional creditor.
	UltmtDbtr *BranchAndFinancialInstitutionIdentification6 `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Financial institution that owes an amount of money to the (ultimate) financial institutional creditor.
	Dbtr *BranchAndFinancialInstitutionIdentification6 `xml:"Dbtr" json:"dbtr" `

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DbtrAcct *CashAccount38 `xml:"DbtrAcct,omitempty" json:"dbtr_acct" `

	// Financial institution servicing an account for the debtor.
	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"dbtr_agt" `

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DbtrAgtAcct *CashAccount38 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct" `

	// Financial institution servicing an account for the creditor.
	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"cdtr_agt" `

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CdtrAgtAcct *CashAccount38 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct" `

	// Financial institution that receives an amount of money from the financial institutional debtor.
	Cdtr *BranchAndFinancialInstitutionIdentification6 `xml:"Cdtr" json:"cdtr" `

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CdtrAcct *CashAccount38 `xml:"CdtrAcct,omitempty" json:"cdtr_acct" `

	// Ultimate financial institution to which an amount of money is due.
	UltmtCdtr *BranchAndFinancialInstitutionIdentification6 `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstrForCdtrAgt []*InstructionForCreditorAgent2 `xml:"InstrForCdtrAgt,omitempty" json:"instr_for_cdtr_agt" `

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstrForNxtAgt []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty" json:"instr_for_nxt_agt" `

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purp *Purpose2Choice `xml:"Purp,omitempty" json:"purp" `

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RmtInf *RemittanceInformation2 `xml:"RmtInf,omitempty" json:"rmt_inf" `

	// Provides information on the underlying customer credit transfer for which cover is provided.
	UndrlygCstmrCdtTrf *CreditTransferTransaction37 `xml:"UndrlygCstmrCdtTrf,omitempty" json:"undrlyg_cstmr_cdt_trf" `

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"SplmtryData,omitempty" json:"splmtry_data" `
}

func (c *CreditTransferTransaction36) AddPmtId() *PaymentIdentification7 {
	c.PmtId = new(PaymentIdentification7)
	return c.PmtId
}

func (c *CreditTransferTransaction36) AddPmtTpInf() *PaymentTypeInformation28 {
	c.PmtTpInf = new(PaymentTypeInformation28)
	return c.PmtTpInf
}

func (c *CreditTransferTransaction36) SetIntrBkSttlmAmt(value, currency string) {
	c.IntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CreditTransferTransaction36) SetIntrBkSttlmDt(value string) {
	c.IntrBkSttlmDt = (*ISODate)(&value)
}

func (c *CreditTransferTransaction36) SetSttlmPrty(value string) {
	c.SttlmPrty = (*Priority3Code)(&value)
}

func (c *CreditTransferTransaction36) AddSttlmTmIndctn() *SettlementDateTimeIndication1 {
	c.SttlmTmIndctn = new(SettlementDateTimeIndication1)
	return c.SttlmTmIndctn
}

func (c *CreditTransferTransaction36) AddSttlmTmReq() *SettlementTimeRequest2 {
	c.SttlmTmReq = new(SettlementTimeRequest2)
	return c.SttlmTmReq
}

func (c *CreditTransferTransaction36) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt1
}

func (c *CreditTransferTransaction36) AddPrvsInstgAgt1Acct() *CashAccount38 {
	c.PrvsInstgAgt1Acct = new(CashAccount38)
	return c.PrvsInstgAgt1Acct
}

func (c *CreditTransferTransaction36) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt2
}

func (c *CreditTransferTransaction36) AddPrvsInstgAgt2Acct() *CashAccount38 {
	c.PrvsInstgAgt2Acct = new(CashAccount38)
	return c.PrvsInstgAgt2Acct
}

func (c *CreditTransferTransaction36) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt3
}

func (c *CreditTransferTransaction36) AddPrvsInstgAgt3Acct() *CashAccount38 {
	c.PrvsInstgAgt3Acct = new(CashAccount38)
	return c.PrvsInstgAgt3Acct
}

func (c *CreditTransferTransaction36) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.InstgAgt
}

func (c *CreditTransferTransaction36) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.InstdAgt
}

func (c *CreditTransferTransaction36) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt1
}

func (c *CreditTransferTransaction36) AddIntrmyAgt1Acct() *CashAccount38 {
	c.IntrmyAgt1Acct = new(CashAccount38)
	return c.IntrmyAgt1Acct
}

func (c *CreditTransferTransaction36) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt2
}

func (c *CreditTransferTransaction36) AddIntrmyAgt2Acct() *CashAccount38 {
	c.IntrmyAgt2Acct = new(CashAccount38)
	return c.IntrmyAgt2Acct
}

func (c *CreditTransferTransaction36) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt3
}

func (c *CreditTransferTransaction36) AddIntrmyAgt3Acct() *CashAccount38 {
	c.IntrmyAgt3Acct = new(CashAccount38)
	return c.IntrmyAgt3Acct
}

func (c *CreditTransferTransaction36) AddUltmtDbtr() *BranchAndFinancialInstitutionIdentification6 {
	c.UltmtDbtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.UltmtDbtr
}

func (c *CreditTransferTransaction36) AddDbtr() *BranchAndFinancialInstitutionIdentification6 {
	c.Dbtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.Dbtr
}

func (c *CreditTransferTransaction36) AddDbtrAcct() *CashAccount38 {
	c.DbtrAcct = new(CashAccount38)
	return c.DbtrAcct
}

func (c *CreditTransferTransaction36) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.DbtrAgt
}

func (c *CreditTransferTransaction36) AddDbtrAgtAcct() *CashAccount38 {
	c.DbtrAgtAcct = new(CashAccount38)
	return c.DbtrAgtAcct
}

func (c *CreditTransferTransaction36) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.CdtrAgt
}

func (c *CreditTransferTransaction36) AddCdtrAgtAcct() *CashAccount38 {
	c.CdtrAgtAcct = new(CashAccount38)
	return c.CdtrAgtAcct
}

func (c *CreditTransferTransaction36) AddCdtr() *BranchAndFinancialInstitutionIdentification6 {
	c.Cdtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.Cdtr
}

func (c *CreditTransferTransaction36) AddCdtrAcct() *CashAccount38 {
	c.CdtrAcct = new(CashAccount38)
	return c.CdtrAcct
}

func (c *CreditTransferTransaction36) AddUltmtCdtr() *BranchAndFinancialInstitutionIdentification6 {
	c.UltmtCdtr = new(BranchAndFinancialInstitutionIdentification6)
	return c.UltmtCdtr
}

func (c *CreditTransferTransaction36) AddInstrForCdtrAgt() *InstructionForCreditorAgent2 {
	newValue := new(InstructionForCreditorAgent2)
	c.InstrForCdtrAgt = append(c.InstrForCdtrAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction36) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstrForNxtAgt = append(c.InstrForNxtAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction36) AddPurp() *Purpose2Choice {
	c.Purp = new(Purpose2Choice)
	return c.Purp
}

func (c *CreditTransferTransaction36) AddRmtInf() *RemittanceInformation2 {
	c.RmtInf = new(RemittanceInformation2)
	return c.RmtInf
}

func (c *CreditTransferTransaction36) AddUndrlygCstmrCdtTrf() *CreditTransferTransaction37 {
	c.UndrlygCstmrCdtTrf = new(CreditTransferTransaction37)
	return c.UndrlygCstmrCdtTrf
}

func (c *CreditTransferTransaction36) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	c.SplmtryData = append(c.SplmtryData, newValue)
	return newValue
}

// Provides further details specific to the individual transaction(s) included in the message.
type CreditTransferTransaction37 struct {
	XMLName xml.Name

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltmtDbtr *PartyIdentification135 `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Party that initiates the payment.
	// Usage: This can be either the debtor or a party that initiates the credit transfer on behalf of the debtor.
	InitgPty *PartyIdentification135 `xml:"InitgPty,omitempty" json:"initg_pty" `

	// Party that owes an amount of money to the (ultimate) creditor.
	Dbtr *PartyIdentification135 `xml:"Dbtr" json:"dbtr" `

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DbtrAcct *CashAccount38 `xml:"DbtrAcct,omitempty" json:"dbtr_acct" `

	// Financial institution servicing an account for the debtor.
	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt" json:"dbtr_agt" `

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DbtrAgtAcct *CashAccount38 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"prvs_instg_agt_1" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt1Acct *CashAccount38 `xml:"PrvsInstgAgt1Acct,omitempty" json:"prvs_instg_agt_1_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"prvs_instg_agt_2" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt2Acct *CashAccount38 `xml:"PrvsInstgAgt2Acct,omitempty" json:"prvs_instg_agt_2_acct" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"prvs_instg_agt_3" `

	// Unambiguous identification of the account of the previous instructing agent at its servicing agent in the payment chain.
	PrvsInstgAgt3Acct *CashAccount38 `xml:"PrvsInstgAgt3Acct,omitempty" json:"prvs_instg_agt_3_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"intrmy_agt_1" `

	// Unambiguous identification of the account of the intermediary agent 1 at its servicing agent in the payment chain.
	IntrmyAgt1Acct *CashAccount38 `xml:"IntrmyAgt1Acct,omitempty" json:"intrmy_agt_1_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"intrmy_agt_2" `

	// Unambiguous identification of the account of the intermediary agent 2 at its servicing agent in the payment chain.
	IntrmyAgt2Acct *CashAccount38 `xml:"IntrmyAgt2Acct,omitempty" json:"intrmy_agt_2_acct" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"intrmy_agt_3" `

	// Unambiguous identification of the account of the intermediary agent 3 at its servicing agent in the payment chain.
	IntrmyAgt3Acct *CashAccount38 `xml:"IntrmyAgt3Acct,omitempty" json:"intrmy_agt_3_acct" `

	// Financial institution servicing an account for the creditor.
	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt" json:"cdtr_agt" `

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CdtrAgtAcct *CashAccount38 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct" `

	// Party to which an amount of money is due.
	Cdtr *PartyIdentification135 `xml:"Cdtr" json:"cdtr" `

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CdtrAcct *CashAccount38 `xml:"CdtrAcct,omitempty" json:"cdtr_acct" `

	// Ultimate party to which an amount of money is due.
	UltmtCdtr *PartyIdentification135 `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstrForCdtrAgt []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty" json:"instr_for_cdtr_agt" `

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	InstrForNxtAgt []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty" json:"instr_for_nxt_agt" `

	// Provides details on the tax.
	Tax *TaxInformation8 `xml:"Tax,omitempty" json:"tax" `

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RmtInf *RemittanceInformation16 `xml:"RmtInf,omitempty" json:"rmt_inf" `

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	InstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt,omitempty" json:"instd_amt" `
}

func (c *CreditTransferTransaction37) AddUltmtDbtr() *PartyIdentification135 {
	c.UltmtDbtr = new(PartyIdentification135)
	return c.UltmtDbtr
}

func (c *CreditTransferTransaction37) AddInitgPty() *PartyIdentification135 {
	c.InitgPty = new(PartyIdentification135)
	return c.InitgPty
}

func (c *CreditTransferTransaction37) AddDbtr() *PartyIdentification135 {
	c.Dbtr = new(PartyIdentification135)
	return c.Dbtr
}

func (c *CreditTransferTransaction37) AddDbtrAcct() *CashAccount38 {
	c.DbtrAcct = new(CashAccount38)
	return c.DbtrAcct
}

func (c *CreditTransferTransaction37) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.DbtrAgt
}

func (c *CreditTransferTransaction37) AddDbtrAgtAcct() *CashAccount38 {
	c.DbtrAgtAcct = new(CashAccount38)
	return c.DbtrAgtAcct
}

func (c *CreditTransferTransaction37) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt1
}

func (c *CreditTransferTransaction37) AddPrvsInstgAgt1Acct() *CashAccount38 {
	c.PrvsInstgAgt1Acct = new(CashAccount38)
	return c.PrvsInstgAgt1Acct
}

func (c *CreditTransferTransaction37) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt2
}

func (c *CreditTransferTransaction37) AddPrvsInstgAgt2Acct() *CashAccount38 {
	c.PrvsInstgAgt2Acct = new(CashAccount38)
	return c.PrvsInstgAgt2Acct
}

func (c *CreditTransferTransaction37) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.PrvsInstgAgt3
}

func (c *CreditTransferTransaction37) AddPrvsInstgAgt3Acct() *CashAccount38 {
	c.PrvsInstgAgt3Acct = new(CashAccount38)
	return c.PrvsInstgAgt3Acct
}

func (c *CreditTransferTransaction37) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt1
}

func (c *CreditTransferTransaction37) AddIntrmyAgt1Acct() *CashAccount38 {
	c.IntrmyAgt1Acct = new(CashAccount38)
	return c.IntrmyAgt1Acct
}

func (c *CreditTransferTransaction37) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt2
}

func (c *CreditTransferTransaction37) AddIntrmyAgt2Acct() *CashAccount38 {
	c.IntrmyAgt2Acct = new(CashAccount38)
	return c.IntrmyAgt2Acct
}

func (c *CreditTransferTransaction37) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification6 {
	c.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return c.IntrmyAgt3
}

func (c *CreditTransferTransaction37) AddIntrmyAgt3Acct() *CashAccount38 {
	c.IntrmyAgt3Acct = new(CashAccount38)
	return c.IntrmyAgt3Acct
}

func (c *CreditTransferTransaction37) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	c.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return c.CdtrAgt
}

func (c *CreditTransferTransaction37) AddCdtrAgtAcct() *CashAccount38 {
	c.CdtrAgtAcct = new(CashAccount38)
	return c.CdtrAgtAcct
}

func (c *CreditTransferTransaction37) AddCdtr() *PartyIdentification135 {
	c.Cdtr = new(PartyIdentification135)
	return c.Cdtr
}

func (c *CreditTransferTransaction37) AddCdtrAcct() *CashAccount38 {
	c.CdtrAcct = new(CashAccount38)
	return c.CdtrAcct
}

func (c *CreditTransferTransaction37) AddUltmtCdtr() *PartyIdentification135 {
	c.UltmtCdtr = new(PartyIdentification135)
	return c.UltmtCdtr
}

func (c *CreditTransferTransaction37) AddInstrForCdtrAgt() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	c.InstrForCdtrAgt = append(c.InstrForCdtrAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction37) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	c.InstrForNxtAgt = append(c.InstrForNxtAgt, newValue)
	return newValue
}

func (c *CreditTransferTransaction37) AddTax() *TaxInformation8 {
	c.Tax = new(TaxInformation8)
	return c.Tax
}

func (c *CreditTransferTransaction37) AddRmtInf() *RemittanceInformation16 {
	c.RmtInf = new(RemittanceInformation16)
	return c.RmtInf
}

func (c *CreditTransferTransaction37) SetInstdAmt(value, currency string) {
	c.InstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Reference information provided by the creditor to allow the identification of the underlying documents.
type CreditorReferenceInformation2 struct {
	XMLName xml.Name

	// Specifies the type of creditor reference.
	Tp *CreditorReferenceType2 `xml:"Tp,omitempty" json:"tp" `

	// Unique reference, as assigned by the creditor, to unambiguously refer to the payment transaction.
	//
	// Usage: If available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the amount of money.
	//
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	Ref *Max35Text `xml:"Ref,omitempty" json:"ref" `
}

func (c *CreditorReferenceInformation2) AddTp() *CreditorReferenceType2 {
	c.Tp = new(CreditorReferenceType2)
	return c.Tp
}

func (c *CreditorReferenceInformation2) SetRef(value string) {
	c.Ref = (*Max35Text)(&value)
}

// Specifies the type of document referred by the creditor.
type CreditorReferenceType1Choice struct {
	XMLName xml.Name

	// Type of creditor reference, in a coded form.
	Cd *DocumentType3Code `xml:"Cd" json:"cd" `

	// Creditor reference type, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (c *CreditorReferenceType1Choice) SetCd(value string) {
	c.Cd = (*DocumentType3Code)(&value)
}

func (c *CreditorReferenceType1Choice) SetPrtry(value string) {
	c.Prtry = (*Max35Text)(&value)
}

// Specifies the type of creditor reference.
type CreditorReferenceType2 struct {
	XMLName xml.Name

	// Coded or proprietary format creditor reference type.
	CdOrPrtry *CreditorReferenceType1Choice `xml:"CdOrPrtry" json:"cd_or_prtry" `

	// Entity that assigns the credit reference type.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (c *CreditorReferenceType2) AddCdOrPrtry() *CreditorReferenceType1Choice {
	c.CdOrPrtry = new(CreditorReferenceType1Choice)
	return c.CdOrPrtry
}

func (c *CreditorReferenceType2) SetIssr(value string) {
	c.Issr = (*Max35Text)(&value)
}

// Choice between a date or a date and time format.
type DateAndDateTime2Choice struct {
	XMLName xml.Name

	// Specified date.
	Dt *ISODate `xml:"Dt" json:"dt" `

	// Specified date and time.
	DtTm *ISODateTime `xml:"DtTm" json:"dt_tm" `
}

func (d *DateAndDateTime2Choice) SetDt(value string) {
	d.Dt = (*ISODate)(&value)
}

func (d *DateAndDateTime2Choice) SetDtTm(value string) {
	d.DtTm = (*ISODateTime)(&value)
}

// Date and place of birth of a person.
type DateAndPlaceOfBirth struct {
	XMLName xml.Name

	// Date on which a person is born.
	BirthDt *ISODate `xml:"BirthDt" json:"birth_dt" `

	// Province where a person was born.
	PrvcOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty" json:"prvc_of_birth" `

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth" json:"city_of_birth" `

	// Country where a person was born.
	CtryOfBirth *CountryCode `xml:"CtryOfBirth" json:"ctry_of_birth" `
}

func (d *DateAndPlaceOfBirth) SetBirthDt(value string) {
	d.BirthDt = (*ISODate)(&value)
}

func (d *DateAndPlaceOfBirth) SetPrvcOfBirth(value string) {
	d.PrvcOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth) SetCityOfBirth(value string) {
	d.CityOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth) SetCtryOfBirth(value string) {
	d.CtryOfBirth = (*CountryCode)(&value)
}

// Date and place of birth of a person.
type DateAndPlaceOfBirth1 struct {
	XMLName xml.Name

	// Date on which a person is born.
	BirthDt *ISODate `xml:"BirthDt" json:"birth_dt" `

	// Province where a person was born.
	PrvcOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty" json:"prvc_of_birth" `

	// City where a person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth" json:"city_of_birth" `

	// Country where a person was born.
	CtryOfBirth *CountryCode `xml:"CtryOfBirth" json:"ctry_of_birth" `
}

func (d *DateAndPlaceOfBirth1) SetBirthDt(value string) {
	d.BirthDt = (*ISODate)(&value)
}

func (d *DateAndPlaceOfBirth1) SetPrvcOfBirth(value string) {
	d.PrvcOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth1) SetCityOfBirth(value string) {
	d.CityOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth1) SetCtryOfBirth(value string) {
	d.CtryOfBirth = (*CountryCode)(&value)
}

// Range of time defined by a start date and an end date.
type DatePeriod2 struct {
	XMLName xml.Name

	// Start date of the range.
	FrDt *ISODate `xml:"FrDt" json:"fr_dt" `

	// End date of the range.
	ToDt *ISODate `xml:"ToDt" json:"to_dt" `
}

func (d *DatePeriod2) SetFrDt(value string) {
	d.FrDt = (*ISODate)(&value)
}

func (d *DatePeriod2) SetToDt(value string) {
	d.ToDt = (*ISODate)(&value)
}

type DecimalNumber string

// Specifies the amount with a specific type.
type DiscountAmountAndType1 struct {
	XMLName xml.Name

	// Specifies the type of the amount.
	Tp *DiscountAmountType1Choice `xml:"Tp,omitempty" json:"tp" `

	// Amount of money, which has been typed.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `
}

func (d *DiscountAmountAndType1) AddTp() *DiscountAmountType1Choice {
	d.Tp = new(DiscountAmountType1Choice)
	return d.Tp
}

func (d *DiscountAmountAndType1) SetAmt(value, currency string) {
	d.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Specifies the amount type.
type DiscountAmountType1Choice struct {
	XMLName xml.Name

	// Specifies the amount type, in a coded form.
	Cd *ExternalDiscountAmountType1Code `xml:"Cd" json:"cd" `

	// Specifies the amount type, in a free-text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (d *DiscountAmountType1Choice) SetCd(value string) {
	d.Cd = (*ExternalDiscountAmountType1Code)(&value)
}

func (d *DiscountAmountType1Choice) SetPrtry(value string) {
	d.Prtry = (*Max35Text)(&value)
}

// Set of elements used to provide information on the amount and reason of the document adjustment.
type DocumentAdjustment1 struct {
	XMLName xml.Name

	// Amount of money of the document adjustment.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `

	// Specifies whether the adjustment must be subtracted or added to the total amount.
	CdtDbtInd *CreditDebitCode `xml:"CdtDbtInd,omitempty" json:"cdt_dbt_ind" `

	// Specifies the reason for the adjustment.
	Rsn *Max4Text `xml:"Rsn,omitempty" json:"rsn" `

	// Provides further details on the document adjustment.
	AddtlInf *Max140Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (d *DocumentAdjustment1) SetAmt(value, currency string) {
	d.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DocumentAdjustment1) SetCdtDbtInd(value string) {
	d.CdtDbtInd = (*CreditDebitCode)(&value)
}

func (d *DocumentAdjustment1) SetRsn(value string) {
	d.Rsn = (*Max4Text)(&value)
}

func (d *DocumentAdjustment1) SetAddtlInf(value string) {
	d.AddtlInf = (*Max140Text)(&value)
}

// Identifies the documents referred to in the remittance information.
type DocumentLineIdentification1 struct {
	XMLName xml.Name

	// Specifies the type of referred document line identification.
	Tp *DocumentLineType1 `xml:"Tp,omitempty" json:"tp" `

	// Identification of the type specified for the referred document line.
	Nb *Max35Text `xml:"Nb,omitempty" json:"nb" `

	// Date associated with the referred document line.
	RltdDt *ISODate `xml:"RltdDt,omitempty" json:"rltd_dt" `
}

func (d *DocumentLineIdentification1) AddTp() *DocumentLineType1 {
	d.Tp = new(DocumentLineType1)
	return d.Tp
}

func (d *DocumentLineIdentification1) SetNb(value string) {
	d.Nb = (*Max35Text)(&value)
}

func (d *DocumentLineIdentification1) SetRltdDt(value string) {
	d.RltdDt = (*ISODate)(&value)
}

// Provides document line information.
//
type DocumentLineInformation1 struct {
	XMLName xml.Name

	// Provides identification of the document line.
	Id []*DocumentLineIdentification1 `xml:"Id" json:"id" `

	// Description associated with the document line.
	Desc *Max2048Text `xml:"Desc,omitempty" json:"desc" `

	// Provides details on the amounts of the document line.
	Amt *RemittanceAmount3 `xml:"Amt,omitempty" json:"amt" `
}

func (d *DocumentLineInformation1) AddId() *DocumentLineIdentification1 {
	newValue := new(DocumentLineIdentification1)
	d.Id = append(d.Id, newValue)
	return newValue
}

func (d *DocumentLineInformation1) SetDesc(value string) {
	d.Desc = (*Max2048Text)(&value)
}

func (d *DocumentLineInformation1) AddAmt() *RemittanceAmount3 {
	d.Amt = new(RemittanceAmount3)
	return d.Amt
}

// Specifies the type of the document line identification.
type DocumentLineType1 struct {
	XMLName xml.Name

	// Provides the type details of the referred document line identification.
	CdOrPrtry *DocumentLineType1Choice `xml:"CdOrPrtry" json:"cd_or_prtry" `

	// Identification of the issuer of the reference document line identificationtype.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (d *DocumentLineType1) AddCdOrPrtry() *DocumentLineType1Choice {
	d.CdOrPrtry = new(DocumentLineType1Choice)
	return d.CdOrPrtry
}

func (d *DocumentLineType1) SetIssr(value string) {
	d.Issr = (*Max35Text)(&value)
}

// Specifies the type of the document line identification.
type DocumentLineType1Choice struct {
	XMLName xml.Name

	// Line identification type in a coded form.
	Cd *ExternalDocumentLineType1Code `xml:"Cd" json:"cd" `

	// Proprietary identification of the type of the remittance document.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (d *DocumentLineType1Choice) SetCd(value string) {
	d.Cd = (*ExternalDocumentLineType1Code)(&value)
}

func (d *DocumentLineType1Choice) SetPrtry(value string) {
	d.Prtry = (*Max35Text)(&value)
}

type DocumentType3Code string

type DocumentType6Code string

// Amount of money to be moved between the debtor and creditor, expressed in the currency of the debtor's account, and the currency in which the amount is to be moved.
type EquivalentAmount2 struct {
	XMLName xml.Name

	// Amount of money to be moved between debtor and creditor, before deduction of charges, expressed in the currency of the debtor's account, and to be moved in a different currency.
	// Usage: The first agent will convert the equivalent amount into the amount to be moved.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `

	// Specifies the currency of the to be transferred amount, which is different from the currency of the debtor's account.
	CcyOfTrf *ActiveOrHistoricCurrencyCode `xml:"CcyOfTrf" json:"ccy_of_trf" `
}

func (e *EquivalentAmount2) SetAmt(value, currency string) {
	e.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *EquivalentAmount2) SetCcyOfTrf(value string) {
	e.CcyOfTrf = (*ActiveOrHistoricCurrencyCode)(&value)
}

type Exact2NumericText string

type Exact4AlphaNumericText string

type ExternalAccountIdentification1Code string

type ExternalAgentInstruction1Code string

type ExternalCancellationReason1Code string

type ExternalCashAccountType1Code string

type ExternalCashClearingSystem1Code string

type ExternalCategoryPurpose1Code string

type ExternalChargeType1Code string

type ExternalClaimNonReceiptRejection1Code string

type ExternalClearingSystemIdentification1Code string

type ExternalDiscountAmountType1Code string

type ExternalDocumentLineType1Code string

type ExternalFinancialInstitutionIdentification1Code string

type ExternalGarnishmentType1Code string

type ExternalInvestigationExecutionConfirmation1Code string

type ExternalLocalInstrument1Code string

type ExternalMandateSetupReason1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPaymentCancellationRejection1Code string

type ExternalPaymentCompensationReason1Code string

type ExternalPaymentGroupStatus1Code string

type ExternalPaymentModificationRejection1Code string

type ExternalPaymentTransactionStatus1Code string

type ExternalPersonIdentification1Code string

type ExternalProxyAccountType1Code string

type ExternalPurpose1Code string

type ExternalReturnReason1Code string

type ExternalServiceLevel1Code string

type ExternalStatusReason1Code string

type ExternalTaxAmountType1Code string

// Sets of elements to identify a name of the organisation identification scheme.
type FinancialIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	// Name of the identification scheme, in a coded form as published in an external list.
	Cd *ExternalFinancialInstitutionIdentification1Code `xml:"Cd" json:"cd" `

	// Name of the identification scheme, in a free text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (f *FinancialIdentificationSchemeName1Choice) SetCd(value string) {
	f.Cd = (*ExternalFinancialInstitutionIdentification1Code)(&value)
}

func (f *FinancialIdentificationSchemeName1Choice) SetPrtry(value string) {
	f.Prtry = (*Max35Text)(&value)
}

// Specifies the details to identify a financial institution.
type FinancialInstitutionIdentification18 struct {
	XMLName xml.Name

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIDec2014Identifier `xml:"BICFI,omitempty" json:"bicfi" `

	// Information used to identify a member within a clearing system.
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"clr_sys_mmb_id" `

	// Legal entity identifier of the financial institution.
	LEI *LEIIdentifier `xml:"LEI,omitempty" json:"lei" `

	// Name by which an agent is known and which is usually used to identify that agent.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress24 `xml:"PstlAdr,omitempty" json:"pstl_adr" `

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Othr *GenericFinancialIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (f *FinancialInstitutionIdentification18) SetBICFI(value string) {
	f.BICFI = (*BICFIDec2014Identifier)(&value)
}

func (f *FinancialInstitutionIdentification18) AddClrSysMmbId() *ClearingSystemMemberIdentification2 {
	f.ClrSysMmbId = new(ClearingSystemMemberIdentification2)
	return f.ClrSysMmbId
}

func (f *FinancialInstitutionIdentification18) SetLEI(value string) {
	f.LEI = (*LEIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification18) SetNm(value string) {
	f.Nm = (*Max140Text)(&value)
}

func (f *FinancialInstitutionIdentification18) AddPstlAdr() *PostalAddress24 {
	f.PstlAdr = new(PostalAddress24)
	return f.PstlAdr
}

func (f *FinancialInstitutionIdentification18) AddOthr() *GenericFinancialIdentification1 {
	f.Othr = new(GenericFinancialIdentification1)
	return f.Othr
}

// Set of elements used to identify a financial institution.
type FinancialInstitutionIdentification8 struct {
	XMLName xml.Name

	// Code allocated to a financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICFI *BICFIIdentifier `xml:"BICFI,omitempty" json:"bicfi" `

	// Information used to identify a member within a clearing system.
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"clr_sys_mmb_id" `

	// Name by which an agent is known and which is usually used to identify that agent.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress6 `xml:"PstlAdr,omitempty" json:"pstl_adr" `

	// Unique identification of an agent, as assigned by an institution, using an identification scheme.
	Othr *GenericFinancialIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (f *FinancialInstitutionIdentification8) SetBICFI(value string) {
	f.BICFI = (*BICFIIdentifier)(&value)
}

func (f *FinancialInstitutionIdentification8) AddClrSysMmbId() *ClearingSystemMemberIdentification2 {
	f.ClrSysMmbId = new(ClearingSystemMemberIdentification2)
	return f.ClrSysMmbId
}

func (f *FinancialInstitutionIdentification8) SetNm(value string) {
	f.Nm = (*Max140Text)(&value)
}

func (f *FinancialInstitutionIdentification8) AddPstlAdr() *PostalAddress6 {
	f.PstlAdr = new(PostalAddress6)
	return f.PstlAdr
}

func (f *FinancialInstitutionIdentification8) AddOthr() *GenericFinancialIdentification1 {
	f.Othr = new(GenericFinancialIdentification1)
	return f.Othr
}

// Choice of format for a frequency, for example, the frequency of payment.
type Frequency36Choice struct {
	XMLName xml.Name

	// Specifies a frequency in terms of a specified period type.
	Tp *Frequency6Code `xml:"Tp" json:"tp" `

	// Specifies a frequency in terms of a count per period within a specified period type.
	Prd *FrequencyPeriod1 `xml:"Prd" json:"prd" `

	// Specifies a frequency in terms of an exact point in time or moment within a specified period type.
	PtInTm *FrequencyAndMoment1 `xml:"PtInTm" json:"pt_in_tm" `
}

func (f *Frequency36Choice) SetTp(value string) {
	f.Tp = (*Frequency6Code)(&value)
}

func (f *Frequency36Choice) AddPrd() *FrequencyPeriod1 {
	f.Prd = new(FrequencyPeriod1)
	return f.Prd
}

func (f *Frequency36Choice) AddPtInTm() *FrequencyAndMoment1 {
	f.PtInTm = new(FrequencyAndMoment1)
	return f.PtInTm
}

type Frequency6Code string

// Defines a frequency in terms a specific moment within a specified period type.
type FrequencyAndMoment1 struct {
	XMLName xml.Name

	// Period for which the number of instructions are to be created and processed.
	Tp *Frequency6Code `xml:"Tp" json:"tp" `

	// Further information on the exact point in time the event should take place.
	PtInTm *Exact2NumericText `xml:"PtInTm" json:"pt_in_tm" `
}

func (f *FrequencyAndMoment1) SetTp(value string) {
	f.Tp = (*Frequency6Code)(&value)
}

func (f *FrequencyAndMoment1) SetPtInTm(value string) {
	f.PtInTm = (*Exact2NumericText)(&value)
}

// Defines a frequency in terms on counts per period for a specific period type.
type FrequencyPeriod1 struct {
	XMLName xml.Name

	// Period for which the number of instructions are to be created and processed.
	Tp *Frequency6Code `xml:"Tp" json:"tp" `

	// Number of instructions to be created and processed during the specified period.
	CntPerPrd *DecimalNumber `xml:"CntPerPrd" json:"cnt_per_prd" `
}

func (f *FrequencyPeriod1) SetTp(value string) {
	f.Tp = (*Frequency6Code)(&value)
}

func (f *FrequencyPeriod1) SetCntPerPrd(value string) {
	f.CntPerPrd = (*DecimalNumber)(&value)
}

// Provides remittance information about a payment for garnishment-related purposes.
type Garnishment2 struct {
	XMLName xml.Name

	// Specifies the type of garnishment.
	Tp *GarnishmentType1 `xml:"Tp" json:"tp" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor, in this case, to the garnisher.
	Grnshee *PartyIdentification125 `xml:"Grnshee,omitempty" json:"grnshee" `

	// Party on the credit side of the transaction who administers the garnishment on behalf of the ultimate beneficiary.
	GrnshmtAdmstr *PartyIdentification125 `xml:"GrnshmtAdmstr,omitempty" json:"grnshmt_admstr" `

	// Reference information that is specific to the agency receiving the garnishment.
	RefNb *Max140Text `xml:"RefNb,omitempty" json:"ref_nb" `

	// Date of payment which garnishment was taken from.
	Dt *ISODate `xml:"Dt,omitempty" json:"dt" `

	// Amount of money remitted for the referred document.
	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"rmtd_amt" `

	// Indicates if the person to whom the garnishment applies (that is, the ultimate debtor) has family medical insurance coverage available.
	FmlyMdclInsrncInd *TrueFalseIndicator `xml:"FmlyMdclInsrncInd,omitempty" json:"fmly_mdcl_insrnc_ind" `

	// Indicates if the employment of the person to whom the garnishment applies (that is, the ultimate debtor) has been terminated.
	MplyeeTermntnInd *TrueFalseIndicator `xml:"MplyeeTermntnInd,omitempty" json:"mplyee_termntn_ind" `
}

func (g *Garnishment2) AddTp() *GarnishmentType1 {
	g.Tp = new(GarnishmentType1)
	return g.Tp
}

func (g *Garnishment2) AddGrnshee() *PartyIdentification125 {
	g.Grnshee = new(PartyIdentification125)
	return g.Grnshee
}

func (g *Garnishment2) AddGrnshmtAdmstr() *PartyIdentification125 {
	g.GrnshmtAdmstr = new(PartyIdentification125)
	return g.GrnshmtAdmstr
}

func (g *Garnishment2) SetRefNb(value string) {
	g.RefNb = (*Max140Text)(&value)
}

func (g *Garnishment2) SetDt(value string) {
	g.Dt = (*ISODate)(&value)
}

func (g *Garnishment2) SetRmtdAmt(value, currency string) {
	g.RmtdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (g *Garnishment2) SetFmlyMdclInsrncInd(value string) {
	g.FmlyMdclInsrncInd = (*TrueFalseIndicator)(&value)
}

func (g *Garnishment2) SetMplyeeTermntnInd(value string) {
	g.MplyeeTermntnInd = (*TrueFalseIndicator)(&value)
}

// Provides remittance information about a payment for garnishment-related purposes.
type Garnishment3 struct {
	XMLName xml.Name

	// Specifies the type of garnishment.
	Tp *GarnishmentType1 `xml:"Tp" json:"tp" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor, in this case, to the garnisher.
	Grnshee *PartyIdentification135 `xml:"Grnshee,omitempty" json:"grnshee" `

	// Party on the credit side of the transaction who administers the garnishment on behalf of the ultimate beneficiary.
	GrnshmtAdmstr *PartyIdentification135 `xml:"GrnshmtAdmstr,omitempty" json:"grnshmt_admstr" `

	// Reference information that is specific to the agency receiving the garnishment.
	RefNb *Max140Text `xml:"RefNb,omitempty" json:"ref_nb" `

	// Date of payment which garnishment was taken from.
	Dt *ISODate `xml:"Dt,omitempty" json:"dt" `

	// Amount of money remitted for the referred document.
	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"rmtd_amt" `

	// Indicates if the person to whom the garnishment applies (that is, the ultimate debtor) has family medical insurance coverage available.
	FmlyMdclInsrncInd *TrueFalseIndicator `xml:"FmlyMdclInsrncInd,omitempty" json:"fmly_mdcl_insrnc_ind" `

	// Indicates if the employment of the person to whom the garnishment applies (that is, the ultimate debtor) has been terminated.
	MplyeeTermntnInd *TrueFalseIndicator `xml:"MplyeeTermntnInd,omitempty" json:"mplyee_termntn_ind" `
}

func (g *Garnishment3) AddTp() *GarnishmentType1 {
	g.Tp = new(GarnishmentType1)
	return g.Tp
}

func (g *Garnishment3) AddGrnshee() *PartyIdentification135 {
	g.Grnshee = new(PartyIdentification135)
	return g.Grnshee
}

func (g *Garnishment3) AddGrnshmtAdmstr() *PartyIdentification135 {
	g.GrnshmtAdmstr = new(PartyIdentification135)
	return g.GrnshmtAdmstr
}

func (g *Garnishment3) SetRefNb(value string) {
	g.RefNb = (*Max140Text)(&value)
}

func (g *Garnishment3) SetDt(value string) {
	g.Dt = (*ISODate)(&value)
}

func (g *Garnishment3) SetRmtdAmt(value, currency string) {
	g.RmtdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (g *Garnishment3) SetFmlyMdclInsrncInd(value string) {
	g.FmlyMdclInsrncInd = (*TrueFalseIndicator)(&value)
}

func (g *Garnishment3) SetMplyeeTermntnInd(value string) {
	g.MplyeeTermntnInd = (*TrueFalseIndicator)(&value)
}

// Specifies the type of garnishment.
type GarnishmentType1 struct {
	XMLName xml.Name

	// Provides the type details of the garnishment.
	CdOrPrtry *GarnishmentType1Choice `xml:"CdOrPrtry" json:"cd_or_prtry" `

	// Identification of the issuer of the garnishment type.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (g *GarnishmentType1) AddCdOrPrtry() *GarnishmentType1Choice {
	g.CdOrPrtry = new(GarnishmentType1Choice)
	return g.CdOrPrtry
}

func (g *GarnishmentType1) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

// Specifies the type of garnishment.
type GarnishmentType1Choice struct {
	XMLName xml.Name

	// Garnishment type in a coded form.
	// Would suggest this to be an External Code List to contain:
	// GNCS    Garnishment from a third party payer for Child Support
	// GNDP    Garnishment from a Direct Payer for Child Support
	// GTPP     Garnishment from a third party payer to taxing agency.
	Cd *ExternalGarnishmentType1Code `xml:"Cd" json:"cd" `

	// Proprietary identification of the type of garnishment.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (g *GarnishmentType1Choice) SetCd(value string) {
	g.Cd = (*ExternalGarnishmentType1Code)(&value)
}

func (g *GarnishmentType1Choice) SetPrtry(value string) {
	g.Prtry = (*Max35Text)(&value)
}

// Information related to a generic account identification.
type GenericAccountIdentification1 struct {
	XMLName xml.Name

	// Identification assigned by an institution.
	Id *Max34Text `xml:"Id" json:"id" `

	// Name of the identification scheme.
	SchmeNm *AccountSchemeName1Choice `xml:"SchmeNm,omitempty" json:"schme_nm" `

	// Entity that assigns the identification.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (g *GenericAccountIdentification1) SetId(value string) {
	g.Id = (*Max34Text)(&value)
}

func (g *GenericAccountIdentification1) AddSchmeNm() *AccountSchemeName1Choice {
	g.SchmeNm = new(AccountSchemeName1Choice)
	return g.SchmeNm
}

func (g *GenericAccountIdentification1) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

// Information related to an identification of a financial institution.
type GenericFinancialIdentification1 struct {
	XMLName xml.Name

	// Unique and unambiguous identification of a person.
	Id *Max35Text `xml:"Id" json:"id" `

	// Name of the identification scheme.
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"schme_nm" `

	// Entity that assigns the identification.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (g *GenericFinancialIdentification1) SetId(value string) {
	g.Id = (*Max35Text)(&value)
}

func (g *GenericFinancialIdentification1) AddSchmeNm() *FinancialIdentificationSchemeName1Choice {
	g.SchmeNm = new(FinancialIdentificationSchemeName1Choice)
	return g.SchmeNm
}

func (g *GenericFinancialIdentification1) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification3 struct {
	XMLName xml.Name

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Id *Max35Text `xml:"Id" json:"id" `

	// Entity that assigns the identification.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (g *GenericIdentification3) SetId(value string) {
	g.Id = (*Max35Text)(&value)
}

func (g *GenericIdentification3) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification30 struct {
	XMLName xml.Name

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Id *Exact4AlphaNumericText `xml:"Id" json:"id" `

	// Entity that assigns the identification.
	Issr *Max35Text `xml:"Issr" json:"issr" `

	// Short textual description of the scheme.
	SchmeNm *Max35Text `xml:"SchmeNm,omitempty" json:"schme_nm" `
}

func (g *GenericIdentification30) SetId(value string) {
	g.Id = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification30) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

func (g *GenericIdentification30) SetSchmeNm(value string) {
	g.SchmeNm = (*Max35Text)(&value)
}

// Information related to an identification of an organisation.
type GenericOrganisationIdentification1 struct {
	XMLName xml.Name

	// Identification assigned by an institution.
	Id *Max35Text `xml:"Id" json:"id" `

	// Name of the identification scheme.
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"schme_nm" `

	// Entity that assigns the identification.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (g *GenericOrganisationIdentification1) SetId(value string) {
	g.Id = (*Max35Text)(&value)
}

func (g *GenericOrganisationIdentification1) AddSchmeNm() *OrganisationIdentificationSchemeName1Choice {
	g.SchmeNm = new(OrganisationIdentificationSchemeName1Choice)
	return g.SchmeNm
}

func (g *GenericOrganisationIdentification1) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

// Information related to an identification of a person.
type GenericPersonIdentification1 struct {
	XMLName xml.Name

	// Unique and unambiguous identification of a person.
	Id *Max35Text `xml:"Id" json:"id" `

	// Name of the identification scheme.
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"schme_nm" `

	// Entity that assigns the identification.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (g *GenericPersonIdentification1) SetId(value string) {
	g.Id = (*Max35Text)(&value)
}

func (g *GenericPersonIdentification1) AddSchmeNm() *PersonIdentificationSchemeName1Choice {
	g.SchmeNm = new(PersonIdentificationSchemeName1Choice)
	return g.SchmeNm
}

func (g *GenericPersonIdentification1) SetIssr(value string) {
	g.Issr = (*Max35Text)(&value)
}

type GroupCancellationIndicator string

type GroupCancellationStatus1Code string

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader53 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MsgId *Max35Text `xml:"MsgId" json:"msg_id" `

	// Date and time at which the message was created.
	CreDtTm *ISODateTime `xml:"CreDtTm" json:"cre_dt_tm" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstgAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstdAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty" json:"instd_agt" `
}

func (g *GroupHeader53) SetMsgId(value string) {
	g.MsgId = (*Max35Text)(&value)
}

func (g *GroupHeader53) SetCreDtTm(value string) {
	g.CreDtTm = (*ISODateTime)(&value)
}

func (g *GroupHeader53) AddInstgAgt() *BranchAndFinancialInstitutionIdentification5 {
	g.InstgAgt = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstgAgt
}

func (g *GroupHeader53) AddInstdAgt() *BranchAndFinancialInstitutionIdentification5 {
	g.InstdAgt = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstdAgt
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader70 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MsgId *Max35Text `xml:"MsgId" json:"msg_id" `

	// Date and time at which the message was created.
	CreDtTm *ISODateTime `xml:"CreDtTm" json:"cre_dt_tm" `

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BtchBookg *BatchBookingIndicator `xml:"BtchBookg,omitempty" json:"btch_bookg" `

	// Number of individual transactions contained in the message.
	NbOfTxs *Max15NumericText `xml:"NbOfTxs" json:"nb_of_txs" `

	// Total of all individual amounts included in the message, irrespective of currencies.
	CtrlSum *DecimalNumber `xml:"CtrlSum,omitempty" json:"ctrl_sum" `

	// Total amount of money moved between the instructing agent and the instructed agent.
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty" json:"ttl_intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SttlmInf *SettlementInstruction4 `xml:"SttlmInf" json:"sttlm_inf" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation21 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstgAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstdAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty" json:"instd_agt" `
}

func (g *GroupHeader70) SetMsgId(value string) {
	g.MsgId = (*Max35Text)(&value)
}

func (g *GroupHeader70) SetCreDtTm(value string) {
	g.CreDtTm = (*ISODateTime)(&value)
}

func (g *GroupHeader70) SetBtchBookg(value string) {
	g.BtchBookg = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader70) SetNbOfTxs(value string) {
	g.NbOfTxs = (*Max15NumericText)(&value)
}

func (g *GroupHeader70) SetCtrlSum(value string) {
	g.CtrlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader70) SetTtlIntrBkSttlmAmt(value, currency string) {
	g.TtlIntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader70) SetIntrBkSttlmDt(value string) {
	g.IntrBkSttlmDt = (*ISODate)(&value)
}

func (g *GroupHeader70) AddSttlmInf() *SettlementInstruction4 {
	g.SttlmInf = new(SettlementInstruction4)
	return g.SttlmInf
}

func (g *GroupHeader70) AddPmtTpInf() *PaymentTypeInformation21 {
	g.PmtTpInf = new(PaymentTypeInformation21)
	return g.PmtTpInf
}

func (g *GroupHeader70) AddInstgAgt() *BranchAndFinancialInstitutionIdentification5 {
	g.InstgAgt = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstgAgt
}

func (g *GroupHeader70) AddInstdAgt() *BranchAndFinancialInstitutionIdentification5 {
	g.InstdAgt = new(BranchAndFinancialInstitutionIdentification5)
	return g.InstdAgt
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader90 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the instructing party and sent to the next party in the chain, to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MsgId *Max35Text `xml:"MsgId" json:"msg_id" `

	// Date and time at which the message was created.
	CreDtTm *ISODateTime `xml:"CreDtTm" json:"cre_dt_tm" `

	// User identification or any user key to be used to check whether the initiating party is allowed to initiate transactions from the account specified in the message.
	//
	// Usage: The content is not of a technical nature, but reflects the organisational structure at the initiating side.
	// The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a party different from the initiating party.
	Authstn []*Authorisation1Choice `xml:"Authstn,omitempty" json:"authstn" `

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BtchBookg *BatchBookingIndicator `xml:"BtchBookg,omitempty" json:"btch_bookg" `

	// Number of individual transactions contained in the message.
	NbOfTxs *Max15NumericText `xml:"NbOfTxs" json:"nb_of_txs" `

	// Total of all individual amounts included in the message, irrespective of currencies.
	CtrlSum *DecimalNumber `xml:"CtrlSum,omitempty" json:"ctrl_sum" `

	// Indicates whether the return applies to the whole group of transactions or to individual transactions within the original group(s).
	GrpRtr *TrueFalseIndicator `xml:"GrpRtr,omitempty" json:"grp_rtr" `

	// Total amount of money moved between the instructing agent and the instructed agent in the return message.
	TtlRtrdIntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"TtlRtrdIntrBkSttlmAmt,omitempty" json:"ttl_rtrd_intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SttlmInf *SettlementInstruction7 `xml:"SttlmInf" json:"sttlm_inf" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"instd_agt" `
}

func (g *GroupHeader90) SetMsgId(value string) {
	g.MsgId = (*Max35Text)(&value)
}

func (g *GroupHeader90) SetCreDtTm(value string) {
	g.CreDtTm = (*ISODateTime)(&value)
}

func (g *GroupHeader90) AddAuthstn() *Authorisation1Choice {
	newValue := new(Authorisation1Choice)
	g.Authstn = append(g.Authstn, newValue)
	return newValue
}

func (g *GroupHeader90) SetBtchBookg(value string) {
	g.BtchBookg = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader90) SetNbOfTxs(value string) {
	g.NbOfTxs = (*Max15NumericText)(&value)
}

func (g *GroupHeader90) SetCtrlSum(value string) {
	g.CtrlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader90) SetGrpRtr(value string) {
	g.GrpRtr = (*TrueFalseIndicator)(&value)
}

func (g *GroupHeader90) SetTtlRtrdIntrBkSttlmAmt(value, currency string) {
	g.TtlRtrdIntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader90) SetIntrBkSttlmDt(value string) {
	g.IntrBkSttlmDt = (*ISODate)(&value)
}

func (g *GroupHeader90) AddSttlmInf() *SettlementInstruction7 {
	g.SttlmInf = new(SettlementInstruction7)
	return g.SttlmInf
}

func (g *GroupHeader90) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstgAgt
}

func (g *GroupHeader90) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstdAgt
}

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader93 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the instructing party, and sent to the next party in the chain to unambiguously identify the message.
	// Usage: The instructing party has to make sure that MessageIdentification is unique per instructed party for a pre-agreed period.
	MsgId *Max35Text `xml:"MsgId" json:"msg_id" `

	// Date and time at which the message was created.
	CreDtTm *ISODateTime `xml:"CreDtTm" json:"cre_dt_tm" `

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions within the group of a message is requested.
	// Usage: Batch booking is used to request and not order a possible batch booking.
	BtchBookg *BatchBookingIndicator `xml:"BtchBookg,omitempty" json:"btch_bookg" `

	// Number of individual transactions contained in the message.
	NbOfTxs *Max15NumericText `xml:"NbOfTxs" json:"nb_of_txs" `

	// Total of all individual amounts included in the message, irrespective of currencies.
	CtrlSum *DecimalNumber `xml:"CtrlSum,omitempty" json:"ctrl_sum" `

	// Total amount of money moved between the instructing agent and the instructed agent.
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty" json:"ttl_intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SttlmInf *SettlementInstruction7 `xml:"SttlmInf" json:"sttlm_inf" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation28 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"instd_agt" `
}

func (g *GroupHeader93) SetMsgId(value string) {
	g.MsgId = (*Max35Text)(&value)
}

func (g *GroupHeader93) SetCreDtTm(value string) {
	g.CreDtTm = (*ISODateTime)(&value)
}

func (g *GroupHeader93) SetBtchBookg(value string) {
	g.BtchBookg = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader93) SetNbOfTxs(value string) {
	g.NbOfTxs = (*Max15NumericText)(&value)
}

func (g *GroupHeader93) SetCtrlSum(value string) {
	g.CtrlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader93) SetTtlIntrBkSttlmAmt(value, currency string) {
	g.TtlIntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (g *GroupHeader93) SetIntrBkSttlmDt(value string) {
	g.IntrBkSttlmDt = (*ISODate)(&value)
}

func (g *GroupHeader93) AddSttlmInf() *SettlementInstruction7 {
	g.SttlmInf = new(SettlementInstruction7)
	return g.SttlmInf
}

func (g *GroupHeader93) AddPmtTpInf() *PaymentTypeInformation28 {
	g.PmtTpInf = new(PaymentTypeInformation28)
	return g.PmtTpInf
}

func (g *GroupHeader93) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstgAgt
}

func (g *GroupHeader93) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	g.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return g.InstdAgt
}

type IBAN2007Identifier string

type ISODate string

type ISODateTime string

type ISONormalisedDateTime string

type ISOTime string

type Instruction3Code string

type Instruction4Code string

type Instruction5Code string

// Further information related to the processing of the investigation that may need to be acted upon by the assignee.
type InstructionForAssignee1 struct {
	XMLName xml.Name

	// Coded information related to the processing of the investigation instruction, provided by the assigner, and intended for the assignee.
	Cd *ExternalAgentInstruction1Code `xml:"Cd,omitempty" json:"cd" `

	// Further information complementing the coded instruction or instruction to the assignee.
	InstrInf *Max140Text `xml:"InstrInf,omitempty" json:"instr_inf" `
}

func (i *InstructionForAssignee1) SetCd(value string) {
	i.Cd = (*ExternalAgentInstruction1Code)(&value)
}

func (i *InstructionForAssignee1) SetInstrInf(value string) {
	i.InstrInf = (*Max140Text)(&value)
}

// Further information related to the processing of the payment instruction that may need to be acted upon by the creditor's agent. The instruction may relate to a level of service, or may be an instruction that has to be executed by the creditor's agent, or may be information required by the creditor's agent.
type InstructionForCreditorAgent1 struct {
	XMLName xml.Name

	// Coded information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor's agent.
	Cd *Instruction3Code `xml:"Cd,omitempty" json:"cd" `

	// Further information complementing the coded instruction or instruction to the creditor's agent that is bilaterally agreed or specific to a user community.
	InstrInf *Max140Text `xml:"InstrInf,omitempty" json:"instr_inf" `
}

func (i *InstructionForCreditorAgent1) SetCd(value string) {
	i.Cd = (*Instruction3Code)(&value)
}

func (i *InstructionForCreditorAgent1) SetInstrInf(value string) {
	i.InstrInf = (*Max140Text)(&value)
}

// Further information related to the processing of the payment instruction that may need to be acted upon by the creditor's agent. The instruction may relate to a level of service, or may be an instruction that has to be executed by the creditor's agent, or may be information required by the creditor's agent.
type InstructionForCreditorAgent2 struct {
	XMLName xml.Name

	// Coded information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor's agent.
	Cd *Instruction5Code `xml:"Cd,omitempty" json:"cd" `

	// Further information complementing the coded instruction or instruction to the creditor's agent that is bilaterally agreed or specific to a user community.
	InstrInf *Max140Text `xml:"InstrInf,omitempty" json:"instr_inf" `
}

func (i *InstructionForCreditorAgent2) SetCd(value string) {
	i.Cd = (*Instruction5Code)(&value)
}

func (i *InstructionForCreditorAgent2) SetInstrInf(value string) {
	i.InstrInf = (*Max140Text)(&value)
}

// Further information related to the processing of the payment instruction that may need to be acted upon by an other agent. The instruction may relate to a level of service, or may be an instruction that has to be executed by the creditor's agent, or may be information required by the other agent.
type InstructionForNextAgent1 struct {
	XMLName xml.Name

	// Coded information related to the processing of the payment instruction, provided by the initiating party, and intended for the next agent in the payment chain.
	Cd *Instruction4Code `xml:"Cd,omitempty" json:"cd" `

	// Further information complementing the coded instruction or instruction to the next agent that is bilaterally agreed or specific to a user community.
	InstrInf *Max140Text `xml:"InstrInf,omitempty" json:"instr_inf" `
}

func (i *InstructionForNextAgent1) SetCd(value string) {
	i.Cd = (*Instruction4Code)(&value)
}

func (i *InstructionForNextAgent1) SetInstrInf(value string) {
	i.InstrInf = (*Max140Text)(&value)
}

// Specifies the status of an investigation case.
type InvestigationStatus5Choice struct {
	XMLName xml.Name

	// Specifies the status of the investigation, in a coded form.
	Conf *ExternalInvestigationExecutionConfirmation1Code `xml:"Conf" json:"conf" `

	// Reason for the rejection of a modification request, in a coded form.
	RjctdMod []*ModificationStatusReason1Choice `xml:"RjctdMod" json:"rjctd_mod" `

	// Indicates a duplicated case.
	// Usage: When present, the case identified in the message must be closed. The case identified as duplicated (in this component) will be pursued.
	DplctOf *Case5 `xml:"DplctOf" json:"dplct_of" `

	// Indicates whether the cancellation of the assignment is confirmed or rejected.
	// Usage: If yes, the cancellation of the assignment is confirmed.
	// If no, the cancellation of the assignment is rejected and the investigation process will continue.
	AssgnmtCxlConf *YesNoIndicator `xml:"AssgnmtCxlConf" json:"assgnmt_cxl_conf" `
}

func (i *InvestigationStatus5Choice) SetConf(value string) {
	i.Conf = (*ExternalInvestigationExecutionConfirmation1Code)(&value)
}

func (i *InvestigationStatus5Choice) AddRjctdMod() *ModificationStatusReason1Choice {
	newValue := new(ModificationStatusReason1Choice)
	i.RjctdMod = append(i.RjctdMod, newValue)
	return newValue
}

func (i *InvestigationStatus5Choice) AddDplctOf() *Case5 {
	i.DplctOf = new(Case5)
	return i.DplctOf
}

func (i *InvestigationStatus5Choice) SetAssgnmtCxlConf(value string) {
	i.AssgnmtCxlConf = (*YesNoIndicator)(&value)
}

type LEIIdentifier string

// Set of elements that further identifies the type of local instruments being requested by the initiating party.
type LocalInstrument2Choice struct {
	XMLName xml.Name

	// Specifies the local instrument, as published in an external local instrument code list.
	Cd *ExternalLocalInstrument1Code `xml:"Cd" json:"cd" `

	// Specifies the local instrument, as a proprietary code.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (l *LocalInstrument2Choice) SetCd(value string) {
	l.Cd = (*ExternalLocalInstrument1Code)(&value)
}

func (l *LocalInstrument2Choice) SetPrtry(value string) {
	l.Prtry = (*Max35Text)(&value)
}

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation12 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MndtId *Max35Text `xml:"MndtId,omitempty" json:"mndt_id" `

	// Date on which the direct debit mandate has been signed by the debtor.
	DtOfSgntr *ISODate `xml:"DtOfSgntr,omitempty" json:"dt_of_sgntr" `

	// Indicator notifying whether the underlying mandate is amended or not.
	AmdmntInd *TrueFalseIndicator `xml:"AmdmntInd,omitempty" json:"amdmnt_ind" `

	// List of mandate elements that have been modified.
	AmdmntInfDtls *AmendmentInformationDetails12 `xml:"AmdmntInfDtls,omitempty" json:"amdmnt_inf_dtls" `

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElctrncSgntr *Max1025Text `xml:"ElctrncSgntr,omitempty" json:"elctrnc_sgntr" `

	// Date of the first collection of a direct debit as per the mandate.
	FrstColltnDt *ISODate `xml:"FrstColltnDt,omitempty" json:"frst_colltn_dt" `

	// Date of the final collection of a direct debit as per the mandate.
	FnlColltnDt *ISODate `xml:"FnlColltnDt,omitempty" json:"fnl_colltn_dt" `

	// Regularity with which direct debit instructions are to be created and processed.
	Frqcy *Frequency36Choice `xml:"Frqcy,omitempty" json:"frqcy" `

	// Reason for the direct debit mandate to allow the user to distinguish between different mandates for the same creditor.
	Rsn *MandateSetupReason1Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Specifies the number of days the direct debit instruction must be tracked.
	TrckgDays *Exact2NumericText `xml:"TrckgDays,omitempty" json:"trckg_days" `
}

func (m *MandateRelatedInformation12) SetMndtId(value string) {
	m.MndtId = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation12) SetDtOfSgntr(value string) {
	m.DtOfSgntr = (*ISODate)(&value)
}

func (m *MandateRelatedInformation12) SetAmdmntInd(value string) {
	m.AmdmntInd = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation12) AddAmdmntInfDtls() *AmendmentInformationDetails12 {
	m.AmdmntInfDtls = new(AmendmentInformationDetails12)
	return m.AmdmntInfDtls
}

func (m *MandateRelatedInformation12) SetElctrncSgntr(value string) {
	m.ElctrncSgntr = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation12) SetFrstColltnDt(value string) {
	m.FrstColltnDt = (*ISODate)(&value)
}

func (m *MandateRelatedInformation12) SetFnlColltnDt(value string) {
	m.FnlColltnDt = (*ISODate)(&value)
}

func (m *MandateRelatedInformation12) AddFrqcy() *Frequency36Choice {
	m.Frqcy = new(Frequency36Choice)
	return m.Frqcy
}

func (m *MandateRelatedInformation12) AddRsn() *MandateSetupReason1Choice {
	m.Rsn = new(MandateSetupReason1Choice)
	return m.Rsn
}

func (m *MandateRelatedInformation12) SetTrckgDays(value string) {
	m.TrckgDays = (*Exact2NumericText)(&value)
}

// Provides further details related to a direct debit mandate signed between the creditor and the debtor.
type MandateRelatedInformation14 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MndtId *Max35Text `xml:"MndtId,omitempty" json:"mndt_id" `

	// Date on which the direct debit mandate has been signed by the debtor.
	DtOfSgntr *ISODate `xml:"DtOfSgntr,omitempty" json:"dt_of_sgntr" `

	// Indicator notifying whether the underlying mandate is amended or not.
	AmdmntInd *TrueFalseIndicator `xml:"AmdmntInd,omitempty" json:"amdmnt_ind" `

	// List of mandate elements that have been modified.
	AmdmntInfDtls *AmendmentInformationDetails13 `xml:"AmdmntInfDtls,omitempty" json:"amdmnt_inf_dtls" `

	// Additional security provisions, such as a digital signature, as provided by the debtor.
	ElctrncSgntr *Max1025Text `xml:"ElctrncSgntr,omitempty" json:"elctrnc_sgntr" `

	// Date of the first collection of a direct debit as per the mandate.
	FrstColltnDt *ISODate `xml:"FrstColltnDt,omitempty" json:"frst_colltn_dt" `

	// Date of the final collection of a direct debit as per the mandate.
	FnlColltnDt *ISODate `xml:"FnlColltnDt,omitempty" json:"fnl_colltn_dt" `

	// Regularity with which direct debit instructions are to be created and processed.
	Frqcy *Frequency36Choice `xml:"Frqcy,omitempty" json:"frqcy" `

	// Reason for the direct debit mandate to allow the user to distinguish between different mandates for the same creditor.
	Rsn *MandateSetupReason1Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Specifies the number of days the direct debit instruction must be tracked.
	TrckgDays *Exact2NumericText `xml:"TrckgDays,omitempty" json:"trckg_days" `
}

func (m *MandateRelatedInformation14) SetMndtId(value string) {
	m.MndtId = (*Max35Text)(&value)
}

func (m *MandateRelatedInformation14) SetDtOfSgntr(value string) {
	m.DtOfSgntr = (*ISODate)(&value)
}

func (m *MandateRelatedInformation14) SetAmdmntInd(value string) {
	m.AmdmntInd = (*TrueFalseIndicator)(&value)
}

func (m *MandateRelatedInformation14) AddAmdmntInfDtls() *AmendmentInformationDetails13 {
	m.AmdmntInfDtls = new(AmendmentInformationDetails13)
	return m.AmdmntInfDtls
}

func (m *MandateRelatedInformation14) SetElctrncSgntr(value string) {
	m.ElctrncSgntr = (*Max1025Text)(&value)
}

func (m *MandateRelatedInformation14) SetFrstColltnDt(value string) {
	m.FrstColltnDt = (*ISODate)(&value)
}

func (m *MandateRelatedInformation14) SetFnlColltnDt(value string) {
	m.FnlColltnDt = (*ISODate)(&value)
}

func (m *MandateRelatedInformation14) AddFrqcy() *Frequency36Choice {
	m.Frqcy = new(Frequency36Choice)
	return m.Frqcy
}

func (m *MandateRelatedInformation14) AddRsn() *MandateSetupReason1Choice {
	m.Rsn = new(MandateSetupReason1Choice)
	return m.Rsn
}

func (m *MandateRelatedInformation14) SetTrckgDays(value string) {
	m.TrckgDays = (*Exact2NumericText)(&value)
}

// Specifies the reason for the setup of the mandate.
type MandateSetupReason1Choice struct {
	XMLName xml.Name

	// Reason for the return, as published in an external reason code list.
	Cd *ExternalMandateSetupReason1Code `xml:"Cd" json:"cd" `

	// Reason for the return, in a proprietary form.
	Prtry *Max70Text `xml:"Prtry" json:"prtry" `
}

func (m *MandateSetupReason1Choice) SetCd(value string) {
	m.Cd = (*ExternalMandateSetupReason1Code)(&value)
}

func (m *MandateSetupReason1Choice) SetPrtry(value string) {
	m.Prtry = (*Max70Text)(&value)
}

type Max1025Text string

type Max105Text string

type Max10Text string

type Max128Text string

type Max140Text string

type Max15NumericText string

type Max16Text string

type Max2048Text string

type Max34Text string

type Max350Text string

type Max35Text string

type Max4Text string

type Max70Text string

// Set of elements used to provide further information on the reason for the unable to apply investigation.
type MissingOrIncorrectInformation3 struct {
	XMLName xml.Name

	// Indicates whether the request is related to an AML (Anti Money Laundering) investigation or not.
	AMLReq *AMLIndicator `xml:"AMLReq,omitempty" json:"a_m_l_req" `

	// Indicates the missing information.
	MssngInf []*UnableToApplyMissing1 `xml:"MssngInf,omitempty" json:"mssng_inf" `

	// Indicates, in a coded form, the incorrect information.
	IncrrctInf []*UnableToApplyIncorrect1 `xml:"IncrrctInf,omitempty" json:"incrrct_inf" `
}

func (m *MissingOrIncorrectInformation3) SetAMLReq(value string) {
	m.AMLReq = (*AMLIndicator)(&value)
}

func (m *MissingOrIncorrectInformation3) AddMssngInf() *UnableToApplyMissing1 {
	newValue := new(UnableToApplyMissing1)
	m.MssngInf = append(m.MssngInf, newValue)
	return newValue
}

func (m *MissingOrIncorrectInformation3) AddIncrrctInf() *UnableToApplyIncorrect1 {
	newValue := new(UnableToApplyIncorrect1)
	m.IncrrctInf = append(m.IncrrctInf, newValue)
	return newValue
}

// Specifies the reason for the transaction modification status.
type ModificationStatusReason1Choice struct {
	XMLName xml.Name

	// Reason for the modification status, in a coded form.
	Cd *ExternalPaymentModificationRejection1Code `xml:"Cd" json:"cd" `

	// Reason for the status, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (m *ModificationStatusReason1Choice) SetCd(value string) {
	m.Cd = (*ExternalPaymentModificationRejection1Code)(&value)
}

func (m *ModificationStatusReason1Choice) SetPrtry(value string) {
	m.Prtry = (*Max35Text)(&value)
}

// Provides further details on the status of the cancellation request.
type ModificationStatusReason2 struct {
	XMLName xml.Name

	// Party that issues the cancellation status.
	Orgtr *PartyIdentification135 `xml:"Orgtr,omitempty" json:"orgtr" `

	// Specifies the reason for the status report.
	Rsn *ModificationStatusReason1Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Further details on the cancellation status reason.
	AddtlInf []*Max105Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (m *ModificationStatusReason2) AddOrgtr() *PartyIdentification135 {
	m.Orgtr = new(PartyIdentification135)
	return m.Orgtr
}

func (m *ModificationStatusReason2) AddRsn() *ModificationStatusReason1Choice {
	m.Rsn = new(ModificationStatusReason1Choice)
	return m.Rsn
}

func (m *ModificationStatusReason2) AddAddtlInf(value string) {
	m.AddtlInf = append(m.AddtlInf, (*Max105Text)(&value))
}

// Information that locates and identifies a party.
type NameAndAddress10 struct {
	XMLName xml.Name

	// Name by which a party is known and is usually used to identify that party.
	Nm *Max140Text `xml:"Nm" json:"nm" `

	// Postal address of a party.
	Adr *PostalAddress6 `xml:"Adr" json:"adr" `
}

func (n *NameAndAddress10) SetNm(value string) {
	n.Nm = (*Max140Text)(&value)
}

func (n *NameAndAddress10) AddAdr() *PostalAddress6 {
	n.Adr = new(PostalAddress6)
	return n.Adr
}

type NamePrefix1Code string

type NamePrefix2Code string

type Number string

// Set of elements used to provide detailed information on the number of transactions that are reported with a specific cancellation status.
type NumberOfCancellationsPerStatus1 struct {
	XMLName xml.Name

	// Number of individual cancellation requests contained in the message, detailed per status.
	DtldNbOfTxs *Max15NumericText `xml:"DtldNbOfTxs" json:"dtld_nb_of_txs" `

	// Common cancellation request status for all individual cancellation requests reported.
	DtldSts *CancellationIndividualStatus1Code `xml:"DtldSts" json:"dtld_sts" `

	// Total of all individual amounts included in the message, irrespective of currencies, detailed per status.
	DtldCtrlSum *DecimalNumber `xml:"DtldCtrlSum,omitempty" json:"dtld_ctrl_sum" `
}

func (n *NumberOfCancellationsPerStatus1) SetDtldNbOfTxs(value string) {
	n.DtldNbOfTxs = (*Max15NumericText)(&value)
}

func (n *NumberOfCancellationsPerStatus1) SetDtldSts(value string) {
	n.DtldSts = (*CancellationIndividualStatus1Code)(&value)
}

func (n *NumberOfCancellationsPerStatus1) SetDtldCtrlSum(value string) {
	n.DtldCtrlSum = (*DecimalNumber)(&value)
}

// Set of additional elements to provide detailed information on the number of transactions that are reported with a specific transaction status.
type NumberOfTransactionsPerStatus1 struct {
	XMLName xml.Name

	// Number of individual transactions contained in the message, detailed per status.
	DtldNbOfTxs *Max15NumericText `xml:"DtldNbOfTxs" json:"dtld_nb_of_txs" `

	// Common transaction status for all individual transactions reported with the same status.
	DtldSts *TransactionIndividualStatus1Code `xml:"DtldSts" json:"dtld_sts" `

	// Total of all individual amounts included in the message, irrespective of currencies, detailed per status.
	DtldCtrlSum *DecimalNumber `xml:"DtldCtrlSum,omitempty" json:"dtld_ctrl_sum" `
}

func (n *NumberOfTransactionsPerStatus1) SetDtldNbOfTxs(value string) {
	n.DtldNbOfTxs = (*Max15NumericText)(&value)
}

func (n *NumberOfTransactionsPerStatus1) SetDtldSts(value string) {
	n.DtldSts = (*TransactionIndividualStatus1Code)(&value)
}

func (n *NumberOfTransactionsPerStatus1) SetDtldCtrlSum(value string) {
	n.DtldCtrlSum = (*DecimalNumber)(&value)
}

// Set of elements used to provide detailed information on the number of transactions that are reported with a specific transaction status.
type NumberOfTransactionsPerStatus5 struct {
	XMLName xml.Name

	// Number of individual transactions contained in the message, detailed per status.
	DtldNbOfTxs *Max15NumericText `xml:"DtldNbOfTxs" json:"dtld_nb_of_txs" `

	// Common transaction status for all individual transactions reported.
	DtldSts *ExternalPaymentTransactionStatus1Code `xml:"DtldSts" json:"dtld_sts" `

	// Total of all individual amounts included in the message, irrespective of currencies, detailed per status.
	DtldCtrlSum *DecimalNumber `xml:"DtldCtrlSum,omitempty" json:"dtld_ctrl_sum" `
}

func (n *NumberOfTransactionsPerStatus5) SetDtldNbOfTxs(value string) {
	n.DtldNbOfTxs = (*Max15NumericText)(&value)
}

func (n *NumberOfTransactionsPerStatus5) SetDtldSts(value string) {
	n.DtldSts = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (n *NumberOfTransactionsPerStatus5) SetDtldCtrlSum(value string) {
	n.DtldCtrlSum = (*DecimalNumber)(&value)
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification29 struct {
	XMLName xml.Name

	// Business identification code of the organisation.
	AnyBIC *AnyBICDec2014Identifier `xml:"AnyBIC,omitempty" json:"any_b_i_c" `

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty" json:"lei" `

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Othr []*GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (o *OrganisationIdentification29) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICDec2014Identifier)(&value)
}

func (o *OrganisationIdentification29) SetLEI(value string) {
	o.LEI = (*LEIIdentifier)(&value)
}

func (o *OrganisationIdentification29) AddOthr() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Othr = append(o.Othr, newValue)
	return newValue
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification7 struct {
	XMLName xml.Name

	// Code allocated to an institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty" json:"any_b_i_c" `

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Othr []*GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (o *OrganisationIdentification7) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification7) AddOthr() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Othr = append(o.Othr, newValue)
	return newValue
}

// Unique and unambiguous way to identify an organisation.
type OrganisationIdentification8 struct {
	XMLName xml.Name

	// Code allocated to a financial institution or non financial institution by the ISO 9362 Registration Authority as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	AnyBIC *AnyBICIdentifier `xml:"AnyBIC,omitempty" json:"any_b_i_c" `

	// Unique identification of an organisation, as assigned by an institution, using an identification scheme.
	Othr []*GenericOrganisationIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (o *OrganisationIdentification8) SetAnyBIC(value string) {
	o.AnyBIC = (*AnyBICIdentifier)(&value)
}

func (o *OrganisationIdentification8) AddOthr() *GenericOrganisationIdentification1 {
	newValue := new(GenericOrganisationIdentification1)
	o.Othr = append(o.Othr, newValue)
	return newValue
}

// Sets of elements to identify a name of the organisation identification scheme.
type OrganisationIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	// Name of the identification scheme, in a coded form as published in an external list.
	Cd *ExternalOrganisationIdentification1Code `xml:"Cd" json:"cd" `

	// Name of the identification scheme, in a free text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (o *OrganisationIdentificationSchemeName1Choice) SetCd(value string) {
	o.Cd = (*ExternalOrganisationIdentification1Code)(&value)
}

func (o *OrganisationIdentificationSchemeName1Choice) SetPrtry(value string) {
	o.Prtry = (*Max35Text)(&value)
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader13 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OrgnlMsgId *Max35Text `xml:"OrgnlMsgId" json:"orgnl_msg_id" `

	// Specifies the original message name identifier to which the message refers.
	OrgnlMsgNmId *Max35Text `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id" `

	// Date and time at which the original message was created.
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm" `

	// Number of individual transactions contained in the original message.
	OrgnlNbOfTxs *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty" json:"orgnl_nb_of_txs" `

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OrgnlCtrlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty" json:"orgnl_ctrl_sum" `

	// Specifies the status of a group of transactions.
	GrpSts *ExternalPaymentGroupStatus1Code `xml:"GrpSts,omitempty" json:"grp_sts" `

	// Provides detailed information on the status reason.
	StsRsnInf []*StatusReasonInformation11 `xml:"StsRsnInf,omitempty" json:"sts_rsn_inf" `

	// Detailed information on the number of transactions for each identical transaction status.
	NbOfTxsPerSts []*NumberOfTransactionsPerStatus5 `xml:"NbOfTxsPerSts,omitempty" json:"nb_of_txs_per_sts" `
}

func (o *OriginalGroupHeader13) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader13) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader13) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader13) SetOrgnlNbOfTxs(value string) {
	o.OrgnlNbOfTxs = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader13) SetOrgnlCtrlSum(value string) {
	o.OrgnlCtrlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader13) SetGrpSts(value string) {
	o.GrpSts = (*ExternalPaymentGroupStatus1Code)(&value)
}

func (o *OriginalGroupHeader13) AddStsRsnInf() *StatusReasonInformation11 {
	newValue := new(StatusReasonInformation11)
	o.StsRsnInf = append(o.StsRsnInf, newValue)
	return newValue
}

func (o *OriginalGroupHeader13) AddNbOfTxsPerSts() *NumberOfTransactionsPerStatus5 {
	newValue := new(NumberOfTransactionsPerStatus5)
	o.NbOfTxsPerSts = append(o.NbOfTxsPerSts, newValue)
	return newValue
}

// Provides the details on the original group, to which the message refers.
type OriginalGroupHeader14 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original group cancellation request.
	OrgnlGrpCxlId *Max35Text `xml:"OrgnlGrpCxlId,omitempty" json:"orgnl_grp_cxl_id" `

	// Identifies the case.
	RslvdCase *Case5 `xml:"RslvdCase,omitempty" json:"rslvd_case" `

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OrgnlMsgId *Max35Text `xml:"OrgnlMsgId" json:"orgnl_msg_id" `

	// Specifies the original message name identifier to which the message refers.
	OrgnlMsgNmId *Max35Text `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id" `

	// Date and time at which the original message was created.
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm" `

	// Number of individual transactions contained in the original message.
	OrgnlNbOfTxs *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty" json:"orgnl_nb_of_txs" `

	// Total of all individual amounts included in the original message, irrespective of currencies.
	OrgnlCtrlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty" json:"orgnl_ctrl_sum" `

	// Specifies the status of a group cancellation request.
	GrpCxlSts *GroupCancellationStatus1Code `xml:"GrpCxlSts,omitempty" json:"grp_cxl_sts" `

	// Provides detailed information on the cancellation status reason.
	CxlStsRsnInf []*CancellationStatusReason4 `xml:"CxlStsRsnInf,omitempty" json:"cxl_sts_rsn_inf" `

	// Detailed information on the number of transactions for each identical cancellation status.
	NbOfTxsPerCxlSts []*NumberOfTransactionsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty" json:"nb_of_txs_per_cxl_sts" `
}

func (o *OriginalGroupHeader14) SetOrgnlGrpCxlId(value string) {
	o.OrgnlGrpCxlId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader14) AddRslvdCase() *Case5 {
	o.RslvdCase = new(Case5)
	return o.RslvdCase
}

func (o *OriginalGroupHeader14) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader14) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader14) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader14) SetOrgnlNbOfTxs(value string) {
	o.OrgnlNbOfTxs = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader14) SetOrgnlCtrlSum(value string) {
	o.OrgnlCtrlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader14) SetGrpCxlSts(value string) {
	o.GrpCxlSts = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalGroupHeader14) AddCxlStsRsnInf() *CancellationStatusReason4 {
	newValue := new(CancellationStatusReason4)
	o.CxlStsRsnInf = append(o.CxlStsRsnInf, newValue)
	return newValue
}

func (o *OriginalGroupHeader14) AddNbOfTxsPerCxlSts() *NumberOfTransactionsPerStatus1 {
	newValue := new(NumberOfTransactionsPerStatus1)
	o.NbOfTxsPerCxlSts = append(o.NbOfTxsPerCxlSts, newValue)
	return newValue
}

// Provides details on the original group, to which the message refers.
type OriginalGroupHeader15 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the assigner, to unambiguously identify the group cancellation request.
	//
	// Usage: The group cancellation request identification can be used for reconciliation or to link tasks related to the cancellation request.
	GrpCxlId *Max35Text `xml:"GrpCxlId,omitempty" json:"grp_cxl_id" `

	// Uniquely and unambiguously identifies an exception or an investigation workflow.
	Case *Case5 `xml:"Case,omitempty" json:"case" `

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OrgnlMsgId *Max35Text `xml:"OrgnlMsgId" json:"orgnl_msg_id" `

	// Specifies the original message name identifier to which the message refers.
	OrgnlMsgNmId *Max35Text `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id" `

	// Date and time at which the original message was created.
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm" `

	// Number of individual transactions contained in the original message.
	NbOfTxs *Max15NumericText `xml:"NbOfTxs,omitempty" json:"nb_of_txs" `

	// Total of all individual amounts included in the message, irrespective of currencies.
	CtrlSum *DecimalNumber `xml:"CtrlSum,omitempty" json:"ctrl_sum" `

	// Indicates whether the cancellation request applies to a whole group of transactions or to individual transactions within an original group.
	GrpCxl *GroupCancellationIndicator `xml:"GrpCxl,omitempty" json:"grp_cxl" `

	// Provides detailed information on the cancellation reason.
	CxlRsnInf []*PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty" json:"cxl_rsn_inf" `
}

func (o *OriginalGroupHeader15) SetGrpCxlId(value string) {
	o.GrpCxlId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader15) AddCase() *Case5 {
	o.Case = new(Case5)
	return o.Case
}

func (o *OriginalGroupHeader15) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader15) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader15) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader15) SetNbOfTxs(value string) {
	o.NbOfTxs = (*Max15NumericText)(&value)
}

func (o *OriginalGroupHeader15) SetCtrlSum(value string) {
	o.CtrlSum = (*DecimalNumber)(&value)
}

func (o *OriginalGroupHeader15) SetGrpCxl(value string) {
	o.GrpCxl = (*GroupCancellationIndicator)(&value)
}

func (o *OriginalGroupHeader15) AddCxlRsnInf() *PaymentCancellationReason5 {
	newValue := new(PaymentCancellationReason5)
	o.CxlRsnInf = append(o.CxlRsnInf, newValue)
	return newValue
}

// Provides information on the original group, to which the message refers.
type OriginalGroupHeader18 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OrgnlMsgId *Max35Text `xml:"OrgnlMsgId" json:"orgnl_msg_id" `

	// Specifies the original message name identifier to which the message refers.
	OrgnlMsgNmId *Max35Text `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id" `

	// Date and time at which the original message was created.
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm" `

	// Provides detailed information on the return reason.
	RtrRsnInf []*PaymentReturnReason6 `xml:"RtrRsnInf,omitempty" json:"rtr_rsn_inf" `
}

func (o *OriginalGroupHeader18) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader18) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupHeader18) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

func (o *OriginalGroupHeader18) AddRtrRsnInf() *PaymentReturnReason6 {
	newValue := new(PaymentReturnReason6)
	o.RtrRsnInf = append(o.RtrRsnInf, newValue)
	return newValue
}

// Unique and unambiguous identifier of the group of transactions as assigned by the original instructing party.
type OriginalGroupInformation29 struct {
	XMLName xml.Name

	// Point to point reference assigned by the original instructing party to unambiguously identify the original message.
	OrgnlMsgId *Max35Text `xml:"OrgnlMsgId" json:"orgnl_msg_id" `

	// Specifies the original message name identifier to which the message refers, for example, pacs.003.001.01 or MT103.
	OrgnlMsgNmId *Max35Text `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id" `

	// Original date and time at which the message was created.
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm" `
}

func (o *OriginalGroupInformation29) SetOrgnlMsgId(value string) {
	o.OrgnlMsgId = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation29) SetOrgnlMsgNmId(value string) {
	o.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (o *OriginalGroupInformation29) SetOrgnlCreDtTm(value string) {
	o.OrgnlCreDtTm = (*ISODateTime)(&value)
}

// Provides the details on the reference and status of the original transactions, included in the original instruction, to which the cancellation request message applies.
type OriginalPaymentInstruction30 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the original assigner, to unambiguously identify the original payment information cancellation request.
	OrgnlPmtInfCxlId *Max35Text `xml:"OrgnlPmtInfCxlId,omitempty" json:"orgnl_pmt_inf_cxl_id" `

	// Identifies the resolved case.
	RslvdCase *Case5 `xml:"RslvdCase,omitempty" json:"rslvd_case" `

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OrgnlPmtInfId *Max35Text `xml:"OrgnlPmtInfId" json:"orgnl_pmt_inf_id" `

	// Provides information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Number of individual transactions contained in the original payment information group.
	OrgnlNbOfTxs *Max15NumericText `xml:"OrgnlNbOfTxs,omitempty" json:"orgnl_nb_of_txs" `

	// Total of all individual amounts included in the original payment information group, irrespective of currencies.
	OrgnlCtrlSum *DecimalNumber `xml:"OrgnlCtrlSum,omitempty" json:"orgnl_ctrl_sum" `

	// Specifies the status of a cancellation request, related to a payment information group.
	PmtInfCxlSts *GroupCancellationStatus1Code `xml:"PmtInfCxlSts,omitempty" json:"pmt_inf_cxl_sts" `

	// Provides detailed information on the cancellation status reason.
	CxlStsRsnInf []*CancellationStatusReason4 `xml:"CxlStsRsnInf,omitempty" json:"cxl_sts_rsn_inf" `

	// Detailed information on the number of transactions for each identical cancellation status.
	NbOfTxsPerCxlSts []*NumberOfCancellationsPerStatus1 `xml:"NbOfTxsPerCxlSts,omitempty" json:"nb_of_txs_per_cxl_sts" `

	// Provides information on the original transactions to which the cancellation request message refers.
	TxInfAndSts []*PaymentTransaction103 `xml:"TxInfAndSts,omitempty" json:"tx_inf_and_sts" `
}

func (o *OriginalPaymentInstruction30) SetOrgnlPmtInfCxlId(value string) {
	o.OrgnlPmtInfCxlId = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction30) AddRslvdCase() *Case5 {
	o.RslvdCase = new(Case5)
	return o.RslvdCase
}

func (o *OriginalPaymentInstruction30) SetOrgnlPmtInfId(value string) {
	o.OrgnlPmtInfId = (*Max35Text)(&value)
}

func (o *OriginalPaymentInstruction30) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	o.OrgnlGrpInf = new(OriginalGroupInformation29)
	return o.OrgnlGrpInf
}

func (o *OriginalPaymentInstruction30) SetOrgnlNbOfTxs(value string) {
	o.OrgnlNbOfTxs = (*Max15NumericText)(&value)
}

func (o *OriginalPaymentInstruction30) SetOrgnlCtrlSum(value string) {
	o.OrgnlCtrlSum = (*DecimalNumber)(&value)
}

func (o *OriginalPaymentInstruction30) SetPmtInfCxlSts(value string) {
	o.PmtInfCxlSts = (*GroupCancellationStatus1Code)(&value)
}

func (o *OriginalPaymentInstruction30) AddCxlStsRsnInf() *CancellationStatusReason4 {
	newValue := new(CancellationStatusReason4)
	o.CxlStsRsnInf = append(o.CxlStsRsnInf, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction30) AddNbOfTxsPerCxlSts() *NumberOfCancellationsPerStatus1 {
	newValue := new(NumberOfCancellationsPerStatus1)
	o.NbOfTxsPerCxlSts = append(o.NbOfTxsPerCxlSts, newValue)
	return newValue
}

func (o *OriginalPaymentInstruction30) AddTxInfAndSts() *PaymentTransaction103 {
	newValue := new(PaymentTransaction103)
	o.TxInfAndSts = append(o.TxInfAndSts, newValue)
	return newValue
}

// Key elements used to refer the original transaction.
type OriginalTransactionReference27 struct {
	XMLName xml.Name

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:"intr_bk_sttlm_amt" `

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amt *AmountType4Choice `xml:"Amt,omitempty" json:"amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	ReqdColltnDt *ISODate `xml:"ReqdColltnDt,omitempty" json:"reqd_colltn_dt" `

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty" json:"reqd_exctn_dt" `

	// Credit party that signs the mandate.
	CdtrSchmeId *PartyIdentification125 `xml:"CdtrSchmeId,omitempty" json:"cdtr_schme_id" `

	// Specifies the details on how the settlement of the original transaction(s) between the instructing agent and the instructed agent was completed.
	SttlmInf *SettlementInstruction4 `xml:"SttlmInf,omitempty" json:"sttlm_inf" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation25 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Specifies the means of payment that will be used to move the amount of money.
	PmtMtd *PaymentMethod4Code `xml:"PmtMtd,omitempty" json:"pmt_mtd" `

	// Provides further details of the mandate signed between the creditor and the debtor.
	MndtRltdInf *MandateRelatedInformation12 `xml:"MndtRltdInf,omitempty" json:"mndt_rltd_inf" `

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RmtInf *RemittanceInformation15 `xml:"RmtInf,omitempty" json:"rmt_inf" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltmtDbtr *Party35Choice `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Party that owes an amount of money to the (ultimate) creditor.
	Dbtr *Party35Choice `xml:"Dbtr,omitempty" json:"dbtr" `

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DbtrAcct *CashAccount24 `xml:"DbtrAcct,omitempty" json:"dbtr_acct" `

	// Financial institution servicing an account for the debtor.
	DbtrAgt *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty" json:"dbtr_agt" `

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DbtrAgtAcct *CashAccount24 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct" `

	// Financial institution servicing an account for the creditor.
	CdtrAgt *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty" json:"cdtr_agt" `

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CdtrAgtAcct *CashAccount24 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct" `

	// Party to which an amount of money is due.
	Cdtr *Party35Choice `xml:"Cdtr,omitempty" json:"cdtr" `

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CdtrAcct *CashAccount24 `xml:"CdtrAcct,omitempty" json:"cdtr_acct" `

	// Ultimate party to which an amount of money is due.
	UltmtCdtr *Party35Choice `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `

	// Underlying reason for the payment transaction.
	//
	// Usage:
	// Purpose is used by the end customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purp *Purpose2Choice `xml:"Purp,omitempty" json:"purp" `
}

func (o *OriginalTransactionReference27) SetIntrBkSttlmAmt(value, currency string) {
	o.IntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalTransactionReference27) AddAmt() *AmountType4Choice {
	o.Amt = new(AmountType4Choice)
	return o.Amt
}

func (o *OriginalTransactionReference27) SetIntrBkSttlmDt(value string) {
	o.IntrBkSttlmDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference27) SetReqdColltnDt(value string) {
	o.ReqdColltnDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference27) AddReqdExctnDt() *DateAndDateTime2Choice {
	o.ReqdExctnDt = new(DateAndDateTime2Choice)
	return o.ReqdExctnDt
}

func (o *OriginalTransactionReference27) AddCdtrSchmeId() *PartyIdentification125 {
	o.CdtrSchmeId = new(PartyIdentification125)
	return o.CdtrSchmeId
}

func (o *OriginalTransactionReference27) AddSttlmInf() *SettlementInstruction4 {
	o.SttlmInf = new(SettlementInstruction4)
	return o.SttlmInf
}

func (o *OriginalTransactionReference27) AddPmtTpInf() *PaymentTypeInformation25 {
	o.PmtTpInf = new(PaymentTypeInformation25)
	return o.PmtTpInf
}

func (o *OriginalTransactionReference27) SetPmtMtd(value string) {
	o.PmtMtd = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference27) AddMndtRltdInf() *MandateRelatedInformation12 {
	o.MndtRltdInf = new(MandateRelatedInformation12)
	return o.MndtRltdInf
}

func (o *OriginalTransactionReference27) AddRmtInf() *RemittanceInformation15 {
	o.RmtInf = new(RemittanceInformation15)
	return o.RmtInf
}

func (o *OriginalTransactionReference27) AddUltmtDbtr() *Party35Choice {
	o.UltmtDbtr = new(Party35Choice)
	return o.UltmtDbtr
}

func (o *OriginalTransactionReference27) AddDbtr() *Party35Choice {
	o.Dbtr = new(Party35Choice)
	return o.Dbtr
}

func (o *OriginalTransactionReference27) AddDbtrAcct() *CashAccount24 {
	o.DbtrAcct = new(CashAccount24)
	return o.DbtrAcct
}

func (o *OriginalTransactionReference27) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification5 {
	o.DbtrAgt = new(BranchAndFinancialInstitutionIdentification5)
	return o.DbtrAgt
}

func (o *OriginalTransactionReference27) AddDbtrAgtAcct() *CashAccount24 {
	o.DbtrAgtAcct = new(CashAccount24)
	return o.DbtrAgtAcct
}

func (o *OriginalTransactionReference27) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification5 {
	o.CdtrAgt = new(BranchAndFinancialInstitutionIdentification5)
	return o.CdtrAgt
}

func (o *OriginalTransactionReference27) AddCdtrAgtAcct() *CashAccount24 {
	o.CdtrAgtAcct = new(CashAccount24)
	return o.CdtrAgtAcct
}

func (o *OriginalTransactionReference27) AddCdtr() *Party35Choice {
	o.Cdtr = new(Party35Choice)
	return o.Cdtr
}

func (o *OriginalTransactionReference27) AddCdtrAcct() *CashAccount24 {
	o.CdtrAcct = new(CashAccount24)
	return o.CdtrAcct
}

func (o *OriginalTransactionReference27) AddUltmtCdtr() *Party35Choice {
	o.UltmtCdtr = new(Party35Choice)
	return o.UltmtCdtr
}

func (o *OriginalTransactionReference27) AddPurp() *Purpose2Choice {
	o.Purp = new(Purpose2Choice)
	return o.Purp
}

// Key elements used to refer the original transaction.
type OriginalTransactionReference28 struct {
	XMLName xml.Name

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:"intr_bk_sttlm_amt" `

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amt *AmountType4Choice `xml:"Amt,omitempty" json:"amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	ReqdColltnDt *ISODate `xml:"ReqdColltnDt,omitempty" json:"reqd_colltn_dt" `

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty" json:"reqd_exctn_dt" `

	// Credit party that signs the mandate.
	CdtrSchmeId *PartyIdentification135 `xml:"CdtrSchmeId,omitempty" json:"cdtr_schme_id" `

	// Specifies the details on how the settlement of the original transaction(s) between the instructing agent and the instructed agent was completed.
	SttlmInf *SettlementInstruction7 `xml:"SttlmInf,omitempty" json:"sttlm_inf" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation27 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Specifies the means of payment that will be used to move the amount of money.
	PmtMtd *PaymentMethod4Code `xml:"PmtMtd,omitempty" json:"pmt_mtd" `

	// Provides further details of the mandate signed between the creditor and the debtor.
	MndtRltdInf *MandateRelatedInformation14 `xml:"MndtRltdInf,omitempty" json:"mndt_rltd_inf" `

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RmtInf *RemittanceInformation16 `xml:"RmtInf,omitempty" json:"rmt_inf" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltmtDbtr *Party40Choice `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Party that owes an amount of money to the (ultimate) creditor.
	Dbtr *Party40Choice `xml:"Dbtr,omitempty" json:"dbtr" `

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DbtrAcct *CashAccount38 `xml:"DbtrAcct,omitempty" json:"dbtr_acct" `

	// Financial institution servicing an account for the debtor.
	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"dbtr_agt" `

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DbtrAgtAcct *CashAccount38 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct" `

	// Financial institution servicing an account for the creditor.
	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"cdtr_agt" `

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CdtrAgtAcct *CashAccount38 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct" `

	// Party to which an amount of money is due.
	Cdtr *Party40Choice `xml:"Cdtr,omitempty" json:"cdtr" `

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CdtrAcct *CashAccount38 `xml:"CdtrAcct,omitempty" json:"cdtr_acct" `

	// Ultimate party to which an amount of money is due.
	UltmtCdtr *Party40Choice `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `

	// Underlying reason for the payment transaction.
	//
	// Usage:
	// Purpose is used by the end customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purp *Purpose2Choice `xml:"Purp,omitempty" json:"purp" `
}

func (o *OriginalTransactionReference28) SetIntrBkSttlmAmt(value, currency string) {
	o.IntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalTransactionReference28) AddAmt() *AmountType4Choice {
	o.Amt = new(AmountType4Choice)
	return o.Amt
}

func (o *OriginalTransactionReference28) SetIntrBkSttlmDt(value string) {
	o.IntrBkSttlmDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) SetReqdColltnDt(value string) {
	o.ReqdColltnDt = (*ISODate)(&value)
}

func (o *OriginalTransactionReference28) AddReqdExctnDt() *DateAndDateTime2Choice {
	o.ReqdExctnDt = new(DateAndDateTime2Choice)
	return o.ReqdExctnDt
}

func (o *OriginalTransactionReference28) AddCdtrSchmeId() *PartyIdentification135 {
	o.CdtrSchmeId = new(PartyIdentification135)
	return o.CdtrSchmeId
}

func (o *OriginalTransactionReference28) AddSttlmInf() *SettlementInstruction7 {
	o.SttlmInf = new(SettlementInstruction7)
	return o.SttlmInf
}

func (o *OriginalTransactionReference28) AddPmtTpInf() *PaymentTypeInformation27 {
	o.PmtTpInf = new(PaymentTypeInformation27)
	return o.PmtTpInf
}

func (o *OriginalTransactionReference28) SetPmtMtd(value string) {
	o.PmtMtd = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference28) AddMndtRltdInf() *MandateRelatedInformation14 {
	o.MndtRltdInf = new(MandateRelatedInformation14)
	return o.MndtRltdInf
}

func (o *OriginalTransactionReference28) AddRmtInf() *RemittanceInformation16 {
	o.RmtInf = new(RemittanceInformation16)
	return o.RmtInf
}

func (o *OriginalTransactionReference28) AddUltmtDbtr() *Party40Choice {
	o.UltmtDbtr = new(Party40Choice)
	return o.UltmtDbtr
}

func (o *OriginalTransactionReference28) AddDbtr() *Party40Choice {
	o.Dbtr = new(Party40Choice)
	return o.Dbtr
}

func (o *OriginalTransactionReference28) AddDbtrAcct() *CashAccount38 {
	o.DbtrAcct = new(CashAccount38)
	return o.DbtrAcct
}

func (o *OriginalTransactionReference28) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	o.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return o.DbtrAgt
}

func (o *OriginalTransactionReference28) AddDbtrAgtAcct() *CashAccount38 {
	o.DbtrAgtAcct = new(CashAccount38)
	return o.DbtrAgtAcct
}

func (o *OriginalTransactionReference28) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	o.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return o.CdtrAgt
}

func (o *OriginalTransactionReference28) AddCdtrAgtAcct() *CashAccount38 {
	o.CdtrAgtAcct = new(CashAccount38)
	return o.CdtrAgtAcct
}

func (o *OriginalTransactionReference28) AddCdtr() *Party40Choice {
	o.Cdtr = new(Party40Choice)
	return o.Cdtr
}

func (o *OriginalTransactionReference28) AddCdtrAcct() *CashAccount38 {
	o.CdtrAcct = new(CashAccount38)
	return o.CdtrAcct
}

func (o *OriginalTransactionReference28) AddUltmtCdtr() *Party40Choice {
	o.UltmtCdtr = new(Party40Choice)
	return o.UltmtCdtr
}

func (o *OriginalTransactionReference28) AddPurp() *Purpose2Choice {
	o.Purp = new(Purpose2Choice)
	return o.Purp
}

// Communication device number or electronic address used for communication.
type OtherContact1 struct {
	XMLName xml.Name

	// Method used to contact the financial institution’s contact for the specific tax region.
	ChanlTp *Max4Text `xml:"ChanlTp" json:"chanl_tp" `

	// Communication value such as phone number or email address.
	Id *Max128Text `xml:"Id,omitempty" json:"id" `
}

func (o *OtherContact1) SetChanlTp(value string) {
	o.ChanlTp = (*Max4Text)(&value)
}

func (o *OtherContact1) SetId(value string) {
	o.Id = (*Max128Text)(&value)
}

// Nature or use of the account.
type Party10Choice struct {
	XMLName xml.Name

	// Unique and unambiguous way to identify an organisation.
	OrgId *OrganisationIdentification7 `xml:"OrgId" json:"org_id" `

	// Unique and unambiguous identification of a person, eg, passport.
	PrvtId *PersonIdentification5 `xml:"PrvtId" json:"prvt_id" `
}

func (p *Party10Choice) AddOrgId() *OrganisationIdentification7 {
	p.OrgId = new(OrganisationIdentification7)
	return p.OrgId
}

func (p *Party10Choice) AddPrvtId() *PersonIdentification5 {
	p.PrvtId = new(PersonIdentification5)
	return p.PrvtId
}

// Nature or use of the account.
type Party34Choice struct {
	XMLName xml.Name

	// Unique and unambiguous way to identify an organisation.
	OrgId *OrganisationIdentification8 `xml:"OrgId" json:"org_id" `

	// Unique and unambiguous identification of a person, for example a passport.
	PrvtId *PersonIdentification13 `xml:"PrvtId" json:"prvt_id" `
}

func (p *Party34Choice) AddOrgId() *OrganisationIdentification8 {
	p.OrgId = new(OrganisationIdentification8)
	return p.OrgId
}

func (p *Party34Choice) AddPrvtId() *PersonIdentification13 {
	p.PrvtId = new(PersonIdentification13)
	return p.PrvtId
}

// Identification of a person, an organisation or a financial institution.
type Party35Choice struct {
	XMLName xml.Name

	// Identification of a person or an organisation.
	Pty *PartyIdentification125 `xml:"Pty" json:"pty" `

	// Identification of a financial institution.
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"Agt" json:"agt" `
}

func (p *Party35Choice) AddPty() *PartyIdentification125 {
	p.Pty = new(PartyIdentification125)
	return p.Pty
}

func (p *Party35Choice) AddAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agt
}

// Nature or use of the account.
type Party38Choice struct {
	XMLName xml.Name

	// Unique and unambiguous way to identify an organisation.
	OrgId *OrganisationIdentification29 `xml:"OrgId" json:"org_id" `

	// Unique and unambiguous identification of a person, for example a passport.
	PrvtId *PersonIdentification13 `xml:"PrvtId" json:"prvt_id" `
}

func (p *Party38Choice) AddOrgId() *OrganisationIdentification29 {
	p.OrgId = new(OrganisationIdentification29)
	return p.OrgId
}

func (p *Party38Choice) AddPrvtId() *PersonIdentification13 {
	p.PrvtId = new(PersonIdentification13)
	return p.PrvtId
}

// Identification of a person, an organisation or a financial institution.
type Party40Choice struct {
	XMLName xml.Name

	// Identification of a person or an organisation.
	Pty *PartyIdentification135 `xml:"Pty" json:"pty" `

	// Identification of a financial institution.
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"Agt" json:"agt" `
}

func (p *Party40Choice) AddPty() *PartyIdentification135 {
	p.Pty = new(PartyIdentification135)
	return p.Pty
}

func (p *Party40Choice) AddAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.Agt = new(BranchAndFinancialInstitutionIdentification6)
	return p.Agt
}

// Identification of a person, an organisation or a financial institution.
type Party9Choice struct {
	XMLName xml.Name

	// Identification of a person or an organisation.
	OrgId *PartyIdentification42 `xml:"OrgId" json:"org_id" `

	// Identification of a financial institution.
	FIId *BranchAndFinancialInstitutionIdentification5 `xml:"FIId" json:"f_i_id" `
}

func (p *Party9Choice) AddOrgId() *PartyIdentification42 {
	p.OrgId = new(PartyIdentification42)
	return p.OrgId
}

func (p *Party9Choice) AddFIId() *BranchAndFinancialInstitutionIdentification5 {
	p.FIId = new(BranchAndFinancialInstitutionIdentification5)
	return p.FIId
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification125 struct {
	XMLName xml.Name

	// Name by which a party is known and which is usually used to identify that party.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress6 `xml:"PstlAdr,omitempty" json:"pstl_adr" `

	// Unique and unambiguous identification of a party.
	Id *Party34Choice `xml:"Id,omitempty" json:"id" `

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CtryOfRes *CountryCode `xml:"CtryOfRes,omitempty" json:"ctry_of_res" `

	// Set of elements used to indicate how to contact the party.
	CtctDtls *ContactDetails2 `xml:"CtctDtls,omitempty" json:"ctct_dtls" `
}

func (p *PartyIdentification125) SetNm(value string) {
	p.Nm = (*Max140Text)(&value)
}

func (p *PartyIdentification125) AddPstlAdr() *PostalAddress6 {
	p.PstlAdr = new(PostalAddress6)
	return p.PstlAdr
}

func (p *PartyIdentification125) AddId() *Party34Choice {
	p.Id = new(Party34Choice)
	return p.Id
}

func (p *PartyIdentification125) SetCtryOfRes(value string) {
	p.CtryOfRes = (*CountryCode)(&value)
}

func (p *PartyIdentification125) AddCtctDtls() *ContactDetails2 {
	p.CtctDtls = new(ContactDetails2)
	return p.CtctDtls
}

// Specifies the identification of a person or an organisation.
type PartyIdentification135 struct {
	XMLName xml.Name

	// Name by which a party is known and which is usually used to identify that party.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress24 `xml:"PstlAdr,omitempty" json:"pstl_adr" `

	// Unique and unambiguous identification of a party.
	Id *Party38Choice `xml:"Id,omitempty" json:"id" `

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CtryOfRes *CountryCode `xml:"CtryOfRes,omitempty" json:"ctry_of_res" `

	// Set of elements used to indicate how to contact the party.
	CtctDtls *Contact4 `xml:"CtctDtls,omitempty" json:"ctct_dtls" `
}

func (p *PartyIdentification135) SetNm(value string) {
	p.Nm = (*Max140Text)(&value)
}

func (p *PartyIdentification135) AddPstlAdr() *PostalAddress24 {
	p.PstlAdr = new(PostalAddress24)
	return p.PstlAdr
}

func (p *PartyIdentification135) AddId() *Party38Choice {
	p.Id = new(Party38Choice)
	return p.Id
}

func (p *PartyIdentification135) SetCtryOfRes(value string) {
	p.CtryOfRes = (*CountryCode)(&value)
}

func (p *PartyIdentification135) AddCtctDtls() *Contact4 {
	p.CtctDtls = new(Contact4)
	return p.CtctDtls
}

// Set of elements used to identify a person or an organisation.
type PartyIdentification42 struct {
	XMLName xml.Name

	// Name by which a party is known and which is usually used to identify that party.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Information that locates and identifies a specific address, as defined by postal services.
	PstlAdr *PostalAddress6 `xml:"PstlAdr,omitempty" json:"pstl_adr" `

	// Unique and unambiguous identification of a party.
	Id *Party10Choice `xml:"Id,omitempty" json:"id" `

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CtryOfRes *CountryCode `xml:"CtryOfRes,omitempty" json:"ctry_of_res" `

	// Set of elements used to indicate how to contact the party.
	CtctDtls *ContactDetails2 `xml:"CtctDtls,omitempty" json:"ctct_dtls" `
}

func (p *PartyIdentification42) SetNm(value string) {
	p.Nm = (*Max140Text)(&value)
}

func (p *PartyIdentification42) AddPstlAdr() *PostalAddress6 {
	p.PstlAdr = new(PostalAddress6)
	return p.PstlAdr
}

func (p *PartyIdentification42) AddId() *Party10Choice {
	p.Id = new(Party10Choice)
	return p.Id
}

func (p *PartyIdentification42) SetCtryOfRes(value string) {
	p.CtryOfRes = (*CountryCode)(&value)
}

func (p *PartyIdentification42) AddCtctDtls() *ContactDetails2 {
	p.CtctDtls = new(ContactDetails2)
	return p.CtctDtls
}

// Provides further details on the reason of the cancellation request.
type PaymentCancellationReason5 struct {
	XMLName xml.Name

	// Party that issues the cancellation request.
	Orgtr *PartyIdentification135 `xml:"Orgtr,omitempty" json:"orgtr" `

	// Specifies the reason for the cancellation.
	Rsn *CancellationReason33Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Further details on the cancellation request reason.
	AddtlInf []*Max105Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (p *PaymentCancellationReason5) AddOrgtr() *PartyIdentification135 {
	p.Orgtr = new(PartyIdentification135)
	return p.Orgtr
}

func (p *PaymentCancellationReason5) AddRsn() *CancellationReason33Choice {
	p.Rsn = new(CancellationReason33Choice)
	return p.Rsn
}

func (p *PaymentCancellationReason5) AddAddtlInf(value string) {
	p.AddtlInf = append(p.AddtlInf, (*Max105Text)(&value))
}

// Set of elements used to provide further means of referencing a payment transaction.
type PaymentIdentification3 struct {
	XMLName xml.Name

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstrId *Max35Text `xml:"InstrId,omitempty" json:"instr_id" `

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndId *Max35Text `xml:"EndToEndId" json:"end_to_end_id" `

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TxId *Max35Text `xml:"TxId" json:"tx_id" `

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClrSysRef *Max35Text `xml:"ClrSysRef,omitempty" json:"clr_sys_ref" `
}

func (p *PaymentIdentification3) SetInstrId(value string) {
	p.InstrId = (*Max35Text)(&value)
}

func (p *PaymentIdentification3) SetEndToEndId(value string) {
	p.EndToEndId = (*Max35Text)(&value)
}

func (p *PaymentIdentification3) SetTxId(value string) {
	p.TxId = (*Max35Text)(&value)
}

func (p *PaymentIdentification3) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

// Provides further means of referencing a payment transaction.
type PaymentIdentification7 struct {
	XMLName xml.Name

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstrId *Max35Text `xml:"InstrId,omitempty" json:"instr_id" `

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	//
	// Usage: In case there are technical limitations to pass on multiple references, the end-to-end identification must be passed on throughout the entire end-to-end chain.
	EndToEndId *Max35Text `xml:"EndToEndId" json:"end_to_end_id" `

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level.
	// Usage: The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TxId *Max35Text `xml:"TxId,omitempty" json:"tx_id" `

	// Universally unique identifier to provide an end-to-end reference of a payment transaction.
	UETR *UUIDv4Identifier `xml:"UETR,omitempty" json:"uetr" `

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClrSysRef *Max35Text `xml:"ClrSysRef,omitempty" json:"clr_sys_ref" `
}

func (p *PaymentIdentification7) SetInstrId(value string) {
	p.InstrId = (*Max35Text)(&value)
}

func (p *PaymentIdentification7) SetEndToEndId(value string) {
	p.EndToEndId = (*Max35Text)(&value)
}

func (p *PaymentIdentification7) SetTxId(value string) {
	p.TxId = (*Max35Text)(&value)
}

func (p *PaymentIdentification7) SetUETR(value string) {
	p.UETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentIdentification7) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

type PaymentMethod4Code string

// Provides further details on the reason of the return of the transaction.
type PaymentReturnReason6 struct {
	XMLName xml.Name

	// Party that issues the return.
	Orgtr *PartyIdentification135 `xml:"Orgtr,omitempty" json:"orgtr" `

	// Specifies the reason for the return.
	Rsn *ReturnReason5Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Further details on the return reason.
	AddtlInf []*Max105Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (p *PaymentReturnReason6) AddOrgtr() *PartyIdentification135 {
	p.Orgtr = new(PartyIdentification135)
	return p.Orgtr
}

func (p *PaymentReturnReason6) AddRsn() *ReturnReason5Choice {
	p.Rsn = new(ReturnReason5Choice)
	return p.Rsn
}

func (p *PaymentReturnReason6) AddAddtlInf(value string) {
	p.AddtlInf = append(p.AddtlInf, (*Max105Text)(&value))
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction102 struct {
	XMLName xml.Name

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CxlStsId *Max35Text `xml:"CxlStsId,omitempty" json:"cxl_sts_id" `

	// Identifies the resolved case.
	RslvdCase *Case5 `xml:"RslvdCase,omitempty" json:"rslvd_case" `

	// Provides information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OrgnlTxId *Max35Text `xml:"OrgnlTxId,omitempty" json:"orgnl_tx_id" `

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OrgnlClrSysRef *Max35Text `xml:"OrgnlClrSysRef,omitempty" json:"orgnl_clr_sys_ref" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `

	// Specifies the status of the transaction cancellation request.
	TxCxlSts *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty" json:"tx_cxl_sts" `

	// Provides detailed information on the cancellation status reason.
	CxlStsRsnInf []*CancellationStatusReason4 `xml:"CxlStsRsnInf,omitempty" json:"cxl_sts_rsn_inf" `

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	RsltnRltdInf *ResolutionData1 `xml:"RsltnRltdInf,omitempty" json:"rsltn_rltd_inf" `

	// Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:"orgnl_intr_bk_sttlm_amt" `

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OrgnlIntrBkSttlmDt *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty" json:"orgnl_intr_bk_sttlm_dt" `

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assgnr *Party40Choice `xml:"Assgnr,omitempty" json:"assgnr" `

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assgne *Party40Choice `xml:"Assgne,omitempty" json:"assgne" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `
}

func (p *PaymentTransaction102) SetCxlStsId(value string) {
	p.CxlStsId = (*Max35Text)(&value)
}

func (p *PaymentTransaction102) AddRslvdCase() *Case5 {
	p.RslvdCase = new(Case5)
	return p.RslvdCase
}

func (p *PaymentTransaction102) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

func (p *PaymentTransaction102) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction102) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction102) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

func (p *PaymentTransaction102) SetOrgnlClrSysRef(value string) {
	p.OrgnlClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction102) SetOrgnlUETR(value string) {
	p.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentTransaction102) SetTxCxlSts(value string) {
	p.TxCxlSts = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction102) AddCxlStsRsnInf() *CancellationStatusReason4 {
	newValue := new(CancellationStatusReason4)
	p.CxlStsRsnInf = append(p.CxlStsRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction102) AddRsltnRltdInf() *ResolutionData1 {
	p.RsltnRltdInf = new(ResolutionData1)
	return p.RsltnRltdInf
}

func (p *PaymentTransaction102) SetOrgnlIntrBkSttlmAmt(value, currency string) {
	p.OrgnlIntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction102) SetOrgnlIntrBkSttlmDt(value string) {
	p.OrgnlIntrBkSttlmDt = (*ISODate)(&value)
}

func (p *PaymentTransaction102) AddAssgnr() *Party40Choice {
	p.Assgnr = new(Party40Choice)
	return p.Assgnr
}

func (p *PaymentTransaction102) AddAssgne() *Party40Choice {
	p.Assgne = new(Party40Choice)
	return p.Assgne
}

func (p *PaymentTransaction102) AddOrgnlTxRef() *OriginalTransactionReference28 {
	p.OrgnlTxRef = new(OriginalTransactionReference28)
	return p.OrgnlTxRef
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction103 struct {
	XMLName xml.Name

	// Unique and unambiguous identifier of a cancellation request status, as assigned by the assigner.
	//
	// Usage: The cancellation status identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CxlStsId *Max35Text `xml:"CxlStsId,omitempty" json:"cxl_sts_id" `

	// Identifies the resolved case.
	RslvdCase *Case5 `xml:"RslvdCase,omitempty" json:"rslvd_case" `

	// Unique identification, as assigned by the original instructing party for the original instructed party to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Universally unique identifier to provide an end-to-end reference of a payment transaction.
	UETR *UUIDv4Identifier `xml:"UETR,omitempty" json:"uetr" `

	// Specifies the status of the transaction cancellation request.
	TxCxlSts *CancellationIndividualStatus1Code `xml:"TxCxlSts,omitempty" json:"tx_cxl_sts" `

	// Provides detailed information on the cancellation status reason.
	CxlStsRsnInf []*CancellationStatusReason4 `xml:"CxlStsRsnInf,omitempty" json:"cxl_sts_rsn_inf" `

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	OrgnlInstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt,omitempty" json:"orgnl_instd_amt" `

	// Date at which the initiating party originally requested the clearing agent to process the payment.
	OrgnlReqdExctnDt *DateAndDateTime2Choice `xml:"OrgnlReqdExctnDt,omitempty" json:"orgnl_reqd_exctn_dt" `

	// Date at which the creditor originally requested the collection of the amount of money from the debtor.
	OrgnlReqdColltnDt *ISODate `xml:"OrgnlReqdColltnDt,omitempty" json:"orgnl_reqd_colltn_dt" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `
}

func (p *PaymentTransaction103) SetCxlStsId(value string) {
	p.CxlStsId = (*Max35Text)(&value)
}

func (p *PaymentTransaction103) AddRslvdCase() *Case5 {
	p.RslvdCase = new(Case5)
	return p.RslvdCase
}

func (p *PaymentTransaction103) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction103) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction103) SetUETR(value string) {
	p.UETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentTransaction103) SetTxCxlSts(value string) {
	p.TxCxlSts = (*CancellationIndividualStatus1Code)(&value)
}

func (p *PaymentTransaction103) AddCxlStsRsnInf() *CancellationStatusReason4 {
	newValue := new(CancellationStatusReason4)
	p.CxlStsRsnInf = append(p.CxlStsRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction103) SetOrgnlInstdAmt(value, currency string) {
	p.OrgnlInstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction103) AddOrgnlReqdExctnDt() *DateAndDateTime2Choice {
	p.OrgnlReqdExctnDt = new(DateAndDateTime2Choice)
	return p.OrgnlReqdExctnDt
}

func (p *PaymentTransaction103) SetOrgnlReqdColltnDt(value string) {
	p.OrgnlReqdColltnDt = (*ISODate)(&value)
}

func (p *PaymentTransaction103) AddOrgnlTxRef() *OriginalTransactionReference28 {
	p.OrgnlTxRef = new(OriginalTransactionReference28)
	return p.OrgnlTxRef
}

// Provides reference and status information on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction106 struct {
	XMLName xml.Name

	// Unique and unambiguous identifier of a cancellation request, as assigned by the assigner.
	//
	// Usage: The cancellation request identification can be used for reconciliation or to link tasks relating to the cancellation request.
	CxlId *Max35Text `xml:"CxlId,omitempty" json:"cxl_id" `

	// Set of elements to uniquely and unambiguously identify an exception or an investigation workflow.
	Case *Case5 `xml:"Case,omitempty" json:"case" `

	// Provides information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OrgnlTxId *Max35Text `xml:"OrgnlTxId,omitempty" json:"orgnl_tx_id" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OrgnlClrSysRef *Max35Text `xml:"OrgnlClrSysRef,omitempty" json:"orgnl_clr_sys_ref" `

	// Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:"orgnl_intr_bk_sttlm_amt" `

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OrgnlIntrBkSttlmDt *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty" json:"orgnl_intr_bk_sttlm_dt" `

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assgnr *BranchAndFinancialInstitutionIdentification6 `xml:"Assgnr,omitempty" json:"assgnr" `

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assgne *BranchAndFinancialInstitutionIdentification6 `xml:"Assgne,omitempty" json:"assgne" `

	// Provides detailed information on the cancellation reason.
	CxlRsnInf []*PaymentCancellationReason5 `xml:"CxlRsnInf,omitempty" json:"cxl_rsn_inf" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"SplmtryData,omitempty" json:"splmtry_data" `
}

func (p *PaymentTransaction106) SetCxlId(value string) {
	p.CxlId = (*Max35Text)(&value)
}

func (p *PaymentTransaction106) AddCase() *Case5 {
	p.Case = new(Case5)
	return p.Case
}

func (p *PaymentTransaction106) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

func (p *PaymentTransaction106) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction106) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction106) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

func (p *PaymentTransaction106) SetOrgnlUETR(value string) {
	p.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentTransaction106) SetOrgnlClrSysRef(value string) {
	p.OrgnlClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction106) SetOrgnlIntrBkSttlmAmt(value, currency string) {
	p.OrgnlIntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction106) SetOrgnlIntrBkSttlmDt(value string) {
	p.OrgnlIntrBkSttlmDt = (*ISODate)(&value)
}

func (p *PaymentTransaction106) AddAssgnr() *BranchAndFinancialInstitutionIdentification6 {
	p.Assgnr = new(BranchAndFinancialInstitutionIdentification6)
	return p.Assgnr
}

func (p *PaymentTransaction106) AddAssgne() *BranchAndFinancialInstitutionIdentification6 {
	p.Assgne = new(BranchAndFinancialInstitutionIdentification6)
	return p.Assgne
}

func (p *PaymentTransaction106) AddCxlRsnInf() *PaymentCancellationReason5 {
	newValue := new(PaymentCancellationReason5)
	p.CxlRsnInf = append(p.CxlRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction106) AddOrgnlTxRef() *OriginalTransactionReference28 {
	p.OrgnlTxRef = new(OriginalTransactionReference28)
	return p.OrgnlTxRef
}

func (p *PaymentTransaction106) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the cancellation request message applies.
type PaymentTransaction107 struct {
	XMLName xml.Name

	// Unique and unambiguous identifier of a modification request status, as assigned by the assigner.
	//
	// Usage: The modification status identification can be used for reconciliation or to link tasks relating to the modification request.
	ModStsId *Max35Text `xml:"ModStsId,omitempty" json:"mod_sts_id" `

	// Identifies the resolved case.
	RslvdCase *Case5 `xml:"RslvdCase,omitempty" json:"rslvd_case" `

	// Provides information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OrgnlPmtInfId *Max35Text `xml:"OrgnlPmtInfId,omitempty" json:"orgnl_pmt_inf_id" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OrgnlTxId *Max35Text `xml:"OrgnlTxId,omitempty" json:"orgnl_tx_id" `

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OrgnlClrSysRef *Max35Text `xml:"OrgnlClrSysRef,omitempty" json:"orgnl_clr_sys_ref" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `

	// Provides detailed information on the modification status reason.
	ModStsRsnInf []*ModificationStatusReason2 `xml:"ModStsRsnInf,omitempty" json:"mod_sts_rsn_inf" `

	// Reference of a return or a reversal transaction that is initiated to fix the case under investigation as part of the resolution.
	RsltnRltdInf *ResolutionData1 `xml:"RsltnRltdInf,omitempty" json:"rsltn_rltd_inf" `

	// Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:"orgnl_intr_bk_sttlm_amt" `

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OrgnlIntrBkSttlmDt *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty" json:"orgnl_intr_bk_sttlm_dt" `

	// Party who assigns the case.
	// Usage: This is also the agent that instructs the next party in the chain to carry out the (set of) cancellation request(s).
	Assgnr *Party40Choice `xml:"Assgnr,omitempty" json:"assgnr" `

	// Party to which the case is assigned.
	// Usage: This is also the agent that is instructed by the previous party in the chain to carry out the (set of) cancellation request(s).
	Assgne *Party40Choice `xml:"Assgne,omitempty" json:"assgne" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `
}

func (p *PaymentTransaction107) SetModStsId(value string) {
	p.ModStsId = (*Max35Text)(&value)
}

func (p *PaymentTransaction107) AddRslvdCase() *Case5 {
	p.RslvdCase = new(Case5)
	return p.RslvdCase
}

func (p *PaymentTransaction107) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

func (p *PaymentTransaction107) SetOrgnlPmtInfId(value string) {
	p.OrgnlPmtInfId = (*Max35Text)(&value)
}

func (p *PaymentTransaction107) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction107) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction107) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

func (p *PaymentTransaction107) SetOrgnlClrSysRef(value string) {
	p.OrgnlClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction107) SetOrgnlUETR(value string) {
	p.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentTransaction107) AddModStsRsnInf() *ModificationStatusReason2 {
	newValue := new(ModificationStatusReason2)
	p.ModStsRsnInf = append(p.ModStsRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction107) AddRsltnRltdInf() *ResolutionData1 {
	p.RsltnRltdInf = new(ResolutionData1)
	return p.RsltnRltdInf
}

func (p *PaymentTransaction107) SetOrgnlIntrBkSttlmAmt(value, currency string) {
	p.OrgnlIntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction107) SetOrgnlIntrBkSttlmDt(value string) {
	p.OrgnlIntrBkSttlmDt = (*ISODate)(&value)
}

func (p *PaymentTransaction107) AddAssgnr() *Party40Choice {
	p.Assgnr = new(Party40Choice)
	return p.Assgnr
}

func (p *PaymentTransaction107) AddAssgne() *Party40Choice {
	p.Assgne = new(Party40Choice)
	return p.Assgne
}

func (p *PaymentTransaction107) AddOrgnlTxRef() *OriginalTransactionReference28 {
	p.OrgnlTxRef = new(OriginalTransactionReference28)
	return p.OrgnlTxRef
}

// Provides further details on the reference and status on the original transactions, included in the original instruction, to which the return message applies.
type PaymentTransaction112 struct {
	XMLName xml.Name

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the returned transaction.
	// Usage: The instructing party is the party sending the return message and not the party that sent the original instruction that is being returned.
	RtrId *Max35Text `xml:"RtrId,omitempty" json:"rtr_id" `

	// Provides information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OrgnlTxId *Max35Text `xml:"OrgnlTxId,omitempty" json:"orgnl_tx_id" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `

	// Unique reference, as assigned by the original clearing system, to unambiguously identify the original instruction.
	OrgnlClrSysRef *Max35Text `xml:"OrgnlClrSysRef,omitempty" json:"orgnl_clr_sys_ref" `

	// Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt,omitempty" json:"orgnl_intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the OriginalInterbankSettlementDate is the interbank settlement date of the original instruction return message, and not of the return message.
	OrgnlIntrBkSttlmDt *ISODate `xml:"OrgnlIntrBkSttlmDt,omitempty" json:"orgnl_intr_bk_sttlm_dt" `

	// Amount of money to be moved between the instructing agent and the instructed agent in the returned instruction.
	RtrdIntrBkSttlmAmt *ActiveCurrencyAndAmount `xml:"RtrdIntrBkSttlmAmt" json:"rtrd_intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	//
	// Usage: the InterbankSettlementDate is the interbank settlement date of the return message, and not of the original instruction.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the settlement instruction.
	//
	//
	// Usage: the SettlementPriority is the settlement priority of the return message, and not of the original instruction.
	SttlmPrty *Priority3Code `xml:"SttlmPrty,omitempty" json:"sttlm_prty" `

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SttlmTmIndctn *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty" json:"sttlm_tm_indctn" `

	// Amount of money to be moved between the debtor and the creditor, before deduction of charges, in the returned transaction.
	// Usage: This amount has to be transported unchanged through the transaction chain.
	RtrdInstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RtrdInstdAmt,omitempty" json:"rtrd_instd_amt" `

	// Factor used to convert an amount from one currency into another. This reflects the price at which one currency was bought with another currency.
	XchgRate *BaseOneRate `xml:"XchgRate,omitempty" json:"xchg_rate" `

	// Amount of money asked or paid as compensation for the processing of the instruction.
	CompstnAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CompstnAmt,omitempty" json:"compstn_amt" `

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	//
	// Usage: The ChargeBearer applies to the return message, not to the original instruction.
	ChrgBr *ChargeBearerType1Code `xml:"ChrgBr,omitempty" json:"chrg_br" `

	// Provides information on the charges to be paid by the charge bearer(s) related to the processing of the return transaction.
	ChrgsInf []*Charges7 `xml:"ChrgsInf,omitempty" json:"chrgs_inf" `

	// Unique reference, as assigned by the clearing system, to unambiguously identify the return instruction.
	ClrSysRef *Max35Text `xml:"ClrSysRef,omitempty" json:"clr_sys_ref" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the return message and not the party that sent the original instruction that is being returned.
	InstgAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the return message and not the party that received the original instruction that is being returned.
	InstdAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdAgt,omitempty" json:"instd_agt" `

	// Provides all parties (agents and non-agents) involved in a return transaction.
	RtrChain *TransactionParties7 `xml:"RtrChain,omitempty" json:"rtr_chain" `

	// Provides detailed information on the return reason.
	RtrRsnInf []*PaymentReturnReason6 `xml:"RtrRsnInf,omitempty" json:"rtr_rsn_inf" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"SplmtryData,omitempty" json:"splmtry_data" `
}

func (p *PaymentTransaction112) SetRtrId(value string) {
	p.RtrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction112) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

func (p *PaymentTransaction112) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction112) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction112) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

func (p *PaymentTransaction112) SetOrgnlUETR(value string) {
	p.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (p *PaymentTransaction112) SetOrgnlClrSysRef(value string) {
	p.OrgnlClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction112) SetOrgnlIntrBkSttlmAmt(value, currency string) {
	p.OrgnlIntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction112) SetOrgnlIntrBkSttlmDt(value string) {
	p.OrgnlIntrBkSttlmDt = (*ISODate)(&value)
}

func (p *PaymentTransaction112) SetRtrdIntrBkSttlmAmt(value, currency string) {
	p.RtrdIntrBkSttlmAmt = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction112) SetIntrBkSttlmDt(value string) {
	p.IntrBkSttlmDt = (*ISODate)(&value)
}

func (p *PaymentTransaction112) SetSttlmPrty(value string) {
	p.SttlmPrty = (*Priority3Code)(&value)
}

func (p *PaymentTransaction112) AddSttlmTmIndctn() *SettlementDateTimeIndication1 {
	p.SttlmTmIndctn = new(SettlementDateTimeIndication1)
	return p.SttlmTmIndctn
}

func (p *PaymentTransaction112) SetRtrdInstdAmt(value, currency string) {
	p.RtrdInstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction112) SetXchgRate(value string) {
	p.XchgRate = (*BaseOneRate)(&value)
}

func (p *PaymentTransaction112) SetCompstnAmt(value, currency string) {
	p.CompstnAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction112) SetChrgBr(value string) {
	p.ChrgBr = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentTransaction112) AddChrgsInf() *Charges7 {
	newValue := new(Charges7)
	p.ChrgsInf = append(p.ChrgsInf, newValue)
	return newValue
}

func (p *PaymentTransaction112) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction112) AddInstgAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.InstgAgt = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstgAgt
}

func (p *PaymentTransaction112) AddInstdAgt() *BranchAndFinancialInstitutionIdentification6 {
	p.InstdAgt = new(BranchAndFinancialInstitutionIdentification6)
	return p.InstdAgt
}

func (p *PaymentTransaction112) AddRtrChain() *TransactionParties7 {
	p.RtrChain = new(TransactionParties7)
	return p.RtrChain
}

func (p *PaymentTransaction112) AddRtrRsnInf() *PaymentReturnReason6 {
	newValue := new(PaymentReturnReason6)
	p.RtrRsnInf = append(p.RtrRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction112) AddOrgnlTxRef() *OriginalTransactionReference28 {
	p.OrgnlTxRef = new(OriginalTransactionReference28)
	return p.OrgnlTxRef
}

func (p *PaymentTransaction112) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

// Provides further details on the original transactions, to which the status report message refers.
type PaymentTransaction91 struct {
	XMLName xml.Name

	// Unique identification, as assigned by an instructing party for an instructed party, to unambiguously identify the reported status.
	// Usage: The instructing party is the party sending the status message and not the party that sent the original instruction that is being reported on.
	StsId *Max35Text `xml:"StsId,omitempty" json:"sts_id" `

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OrgnlTxId *Max35Text `xml:"OrgnlTxId,omitempty" json:"orgnl_tx_id" `

	// Specifies the status of a transaction, in a coded form.
	TxSts *ExternalPaymentTransactionStatus1Code `xml:"TxSts,omitempty" json:"tx_sts" `

	// Provides detailed information on the status reason.
	StsRsnInf []*StatusReasonInformation11 `xml:"StsRsnInf,omitempty" json:"sts_rsn_inf" `

	// Provides information on the charges related to the processing of the rejection of the instruction.
	//
	// Usage: This is passed on for information purposes only. Settlement of the charges will be done separately.
	ChrgsInf []*Charges2 `xml:"ChrgsInf,omitempty" json:"chrgs_inf" `

	// Point in time when the payment order from the initiating party meets the processing conditions of the account servicing agent. This means that the account servicing agent has received the payment order and has applied checks such as authorisation, availability of funds.
	AccptncDtTm *ISODateTime `xml:"AccptncDtTm,omitempty" json:"accptnc_dt_tm" `

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the instruction.
	AcctSvcrRef *Max35Text `xml:"AcctSvcrRef,omitempty" json:"acct_svcr_ref" `

	// Unique reference, as assigned by a clearing system, to unambiguously identify the instruction.
	ClrSysRef *Max35Text `xml:"ClrSysRef,omitempty" json:"clr_sys_ref" `

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructing agent is the party sending the status message and not the party that sent the original instruction that is being reported on.
	InstgAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgAgt,omitempty" json:"instg_agt" `

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	//
	// Usage: The instructed agent is the party receiving the status message and not the party that received the original instruction that is being reported on.
	InstdAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstdAgt,omitempty" json:"instd_agt" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference27 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"SplmtryData,omitempty" json:"splmtry_data" `
}

func (p *PaymentTransaction91) SetStsId(value string) {
	p.StsId = (*Max35Text)(&value)
}

func (p *PaymentTransaction91) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	p.OrgnlGrpInf = new(OriginalGroupInformation29)
	return p.OrgnlGrpInf
}

func (p *PaymentTransaction91) SetOrgnlInstrId(value string) {
	p.OrgnlInstrId = (*Max35Text)(&value)
}

func (p *PaymentTransaction91) SetOrgnlEndToEndId(value string) {
	p.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (p *PaymentTransaction91) SetOrgnlTxId(value string) {
	p.OrgnlTxId = (*Max35Text)(&value)
}

func (p *PaymentTransaction91) SetTxSts(value string) {
	p.TxSts = (*ExternalPaymentTransactionStatus1Code)(&value)
}

func (p *PaymentTransaction91) AddStsRsnInf() *StatusReasonInformation11 {
	newValue := new(StatusReasonInformation11)
	p.StsRsnInf = append(p.StsRsnInf, newValue)
	return newValue
}

func (p *PaymentTransaction91) AddChrgsInf() *Charges2 {
	newValue := new(Charges2)
	p.ChrgsInf = append(p.ChrgsInf, newValue)
	return newValue
}

func (p *PaymentTransaction91) SetAccptncDtTm(value string) {
	p.AccptncDtTm = (*ISODateTime)(&value)
}

func (p *PaymentTransaction91) SetAcctSvcrRef(value string) {
	p.AcctSvcrRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction91) SetClrSysRef(value string) {
	p.ClrSysRef = (*Max35Text)(&value)
}

func (p *PaymentTransaction91) AddInstgAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.InstgAgt = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstgAgt
}

func (p *PaymentTransaction91) AddInstdAgt() *BranchAndFinancialInstitutionIdentification5 {
	p.InstdAgt = new(BranchAndFinancialInstitutionIdentification5)
	return p.InstdAgt
}

func (p *PaymentTransaction91) AddOrgnlTxRef() *OriginalTransactionReference27 {
	p.OrgnlTxRef = new(OriginalTransactionReference27)
	return p.OrgnlTxRef
}

func (p *PaymentTransaction91) AddSplmtryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation21 struct {
	XMLName xml.Name

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstrPrty *Priority2Code `xml:"InstrPrty,omitempty" json:"instr_prty" `

	// Specifies the clearing channel to be used to process the payment instruction.
	ClrChanl *ClearingChannel2Code `xml:"ClrChanl,omitempty" json:"clr_chanl" `

	// Agreement under which or rules under which the transaction should be processed.
	SvcLvl *ServiceLevel8Choice `xml:"SvcLvl,omitempty" json:"svc_lvl" `

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"lcl_instrm" `

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"ctgy_purp" `
}

func (p *PaymentTypeInformation21) SetInstrPrty(value string) {
	p.InstrPrty = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation21) SetClrChanl(value string) {
	p.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation21) AddSvcLvl() *ServiceLevel8Choice {
	p.SvcLvl = new(ServiceLevel8Choice)
	return p.SvcLvl
}

func (p *PaymentTypeInformation21) AddLclInstrm() *LocalInstrument2Choice {
	p.LclInstrm = new(LocalInstrument2Choice)
	return p.LclInstrm
}

func (p *PaymentTypeInformation21) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

// Set of elements used to provide further details of the type of payment.
type PaymentTypeInformation25 struct {
	XMLName xml.Name

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstrPrty *Priority2Code `xml:"InstrPrty,omitempty" json:"instr_prty" `

	// Specifies the clearing channel to be used to process the payment instruction.
	ClrChanl *ClearingChannel2Code `xml:"ClrChanl,omitempty" json:"clr_chanl" `

	// Agreement under which or rules under which the transaction should be processed.
	SvcLvl *ServiceLevel8Choice `xml:"SvcLvl,omitempty" json:"svc_lvl" `

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"lcl_instrm" `

	// Identifies the direct debit sequence, such as first, recurrent, final or one-off.
	SeqTp *SequenceType3Code `xml:"SeqTp,omitempty" json:"seq_tp" `

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"ctgy_purp" `
}

func (p *PaymentTypeInformation25) SetInstrPrty(value string) {
	p.InstrPrty = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation25) SetClrChanl(value string) {
	p.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation25) AddSvcLvl() *ServiceLevel8Choice {
	p.SvcLvl = new(ServiceLevel8Choice)
	return p.SvcLvl
}

func (p *PaymentTypeInformation25) AddLclInstrm() *LocalInstrument2Choice {
	p.LclInstrm = new(LocalInstrument2Choice)
	return p.LclInstrm
}

func (p *PaymentTypeInformation25) SetSeqTp(value string) {
	p.SeqTp = (*SequenceType3Code)(&value)
}

func (p *PaymentTypeInformation25) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

// Provides further details of the type of payment.
type PaymentTypeInformation27 struct {
	XMLName xml.Name

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstrPrty *Priority2Code `xml:"InstrPrty,omitempty" json:"instr_prty" `

	// Specifies the clearing channel to be used to process the payment instruction.
	ClrChanl *ClearingChannel2Code `xml:"ClrChanl,omitempty" json:"clr_chanl" `

	// Agreement under which or rules under which the transaction should be processed.
	SvcLvl []*ServiceLevel8Choice `xml:"SvcLvl,omitempty" json:"svc_lvl" `

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"lcl_instrm" `

	// Identifies the direct debit sequence, such as first, recurrent, final or one-off.
	SeqTp *SequenceType3Code `xml:"SeqTp,omitempty" json:"seq_tp" `

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"ctgy_purp" `
}

func (p *PaymentTypeInformation27) SetInstrPrty(value string) {
	p.InstrPrty = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation27) SetClrChanl(value string) {
	p.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation27) AddSvcLvl() *ServiceLevel8Choice {
	newValue := new(ServiceLevel8Choice)
	p.SvcLvl = append(p.SvcLvl, newValue)
	return newValue
}

func (p *PaymentTypeInformation27) AddLclInstrm() *LocalInstrument2Choice {
	p.LclInstrm = new(LocalInstrument2Choice)
	return p.LclInstrm
}

func (p *PaymentTypeInformation27) SetSeqTp(value string) {
	p.SeqTp = (*SequenceType3Code)(&value)
}

func (p *PaymentTypeInformation27) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

// Provides further details of the type of payment.
type PaymentTypeInformation28 struct {
	XMLName xml.Name

	// Indicator of the urgency or order of importance that the instructing party would like the instructed party to apply to the processing of the instruction.
	InstrPrty *Priority2Code `xml:"InstrPrty,omitempty" json:"instr_prty" `

	// Specifies the clearing channel to be used to process the payment instruction.
	ClrChanl *ClearingChannel2Code `xml:"ClrChanl,omitempty" json:"clr_chanl" `

	// Agreement under which or rules under which the transaction should be processed.
	SvcLvl []*ServiceLevel8Choice `xml:"SvcLvl,omitempty" json:"svc_lvl" `

	// User community specific instrument.
	//
	// Usage: This element is used to specify a local instrument, local clearing option and/or further qualify the service or service level.
	LclInstrm *LocalInstrument2Choice `xml:"LclInstrm,omitempty" json:"lcl_instrm" `

	// Specifies the high level purpose of the instruction based on a set of pre-defined categories.
	// Usage: This is used by the initiating party to provide information concerning the processing of the payment. It is likely to trigger special processing by any of the agents involved in the payment chain.
	CtgyPurp *CategoryPurpose1Choice `xml:"CtgyPurp,omitempty" json:"ctgy_purp" `
}

func (p *PaymentTypeInformation28) SetInstrPrty(value string) {
	p.InstrPrty = (*Priority2Code)(&value)
}

func (p *PaymentTypeInformation28) SetClrChanl(value string) {
	p.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (p *PaymentTypeInformation28) AddSvcLvl() *ServiceLevel8Choice {
	newValue := new(ServiceLevel8Choice)
	p.SvcLvl = append(p.SvcLvl, newValue)
	return newValue
}

func (p *PaymentTypeInformation28) AddLclInstrm() *LocalInstrument2Choice {
	p.LclInstrm = new(LocalInstrument2Choice)
	return p.LclInstrm
}

func (p *PaymentTypeInformation28) AddCtgyPurp() *CategoryPurpose1Choice {
	p.CtgyPurp = new(CategoryPurpose1Choice)
	return p.CtgyPurp
}

type PercentageRate string

// Unique and unambiguous way to identify a person.
type PersonIdentification13 struct {
	XMLName xml.Name

	// Date and place of birth of a person.
	DtAndPlcOfBirth *DateAndPlaceOfBirth1 `xml:"DtAndPlcOfBirth,omitempty" json:"dt_and_plc_of_birth" `

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Othr []*GenericPersonIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (p *PersonIdentification13) AddDtAndPlcOfBirth() *DateAndPlaceOfBirth1 {
	p.DtAndPlcOfBirth = new(DateAndPlaceOfBirth1)
	return p.DtAndPlcOfBirth
}

func (p *PersonIdentification13) AddOthr() *GenericPersonIdentification1 {
	newValue := new(GenericPersonIdentification1)
	p.Othr = append(p.Othr, newValue)
	return newValue
}

// Unique and unambiguous way to identify a person.
type PersonIdentification5 struct {
	XMLName xml.Name

	// Date and place of birth of a person.
	DtAndPlcOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty" json:"dt_and_plc_of_birth" `

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Othr []*GenericPersonIdentification1 `xml:"Othr,omitempty" json:"othr" `
}

func (p *PersonIdentification5) AddDtAndPlcOfBirth() *DateAndPlaceOfBirth {
	p.DtAndPlcOfBirth = new(DateAndPlaceOfBirth)
	return p.DtAndPlcOfBirth
}

func (p *PersonIdentification5) AddOthr() *GenericPersonIdentification1 {
	newValue := new(GenericPersonIdentification1)
	p.Othr = append(p.Othr, newValue)
	return newValue
}

// Sets of elements to identify a name of the identification scheme.
type PersonIdentificationSchemeName1Choice struct {
	XMLName xml.Name

	// Name of the identification scheme, in a coded form as published in an external list.
	Cd *ExternalPersonIdentification1Code `xml:"Cd" json:"cd" `

	// Name of the identification scheme, in a free text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (p *PersonIdentificationSchemeName1Choice) SetCd(value string) {
	p.Cd = (*ExternalPersonIdentification1Code)(&value)
}

func (p *PersonIdentificationSchemeName1Choice) SetPrtry(value string) {
	p.Prtry = (*Max35Text)(&value)
}

type PhoneNumber string

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress24 struct {
	XMLName xml.Name

	// Identifies the nature of the postal address.
	AdrTp *AddressType3Choice `xml:"AdrTp,omitempty" json:"adr_tp" `

	// Identification of a division of a large organisation or building.
	Dept *Max70Text `xml:"Dept,omitempty" json:"dept" `

	// Identification of a sub-division of a large organisation or building.
	SubDept *Max70Text `xml:"SubDept,omitempty" json:"sub_dept" `

	// Name of a street or thoroughfare.
	StrtNm *Max70Text `xml:"StrtNm,omitempty" json:"strt_nm" `

	// Number that identifies the position of a building on a street.
	BldgNb *Max16Text `xml:"BldgNb,omitempty" json:"bldg_nb" `

	// Name of the building or house.
	BldgNm *Max35Text `xml:"BldgNm,omitempty" json:"bldg_nm" `

	// Floor or storey within a building.
	Flr *Max70Text `xml:"Flr,omitempty" json:"flr" `

	// Numbered box in a post office, assigned to a person or organisation, where letters are kept until called for.
	PstBx *Max16Text `xml:"PstBx,omitempty" json:"pst_bx" `

	// Building room number.
	Room *Max70Text `xml:"Room,omitempty" json:"room" `

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PstCd *Max16Text `xml:"PstCd,omitempty" json:"pst_cd" `

	// Name of a built-up area, with defined boundaries, and a local government.
	TwnNm *Max35Text `xml:"TwnNm,omitempty" json:"twn_nm" `

	// Specific location name within the town.
	TwnLctnNm *Max35Text `xml:"TwnLctnNm,omitempty" json:"twn_lctn_nm" `

	// Identifies a subdivision within a country sub-division.
	DstrctNm *Max35Text `xml:"DstrctNm,omitempty" json:"dstrct_nm" `

	// Identifies a subdivision of a country such as state, region, county.
	CtrySubDvsn *Max35Text `xml:"CtrySubDvsn,omitempty" json:"ctry_sub_dvsn" `

	// Nation with its own government.
	Ctry *CountryCode `xml:"Ctry,omitempty" json:"ctry" `

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AdrLine []*Max70Text `xml:"AdrLine,omitempty" json:"adr_line" `
}

func (p *PostalAddress24) AddAdrTp() *AddressType3Choice {
	p.AdrTp = new(AddressType3Choice)
	return p.AdrTp
}

func (p *PostalAddress24) SetDept(value string) {
	p.Dept = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetSubDept(value string) {
	p.SubDept = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetStrtNm(value string) {
	p.StrtNm = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetBldgNb(value string) {
	p.BldgNb = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetBldgNm(value string) {
	p.BldgNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetFlr(value string) {
	p.Flr = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetPstBx(value string) {
	p.PstBx = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetRoom(value string) {
	p.Room = (*Max70Text)(&value)
}

func (p *PostalAddress24) SetPstCd(value string) {
	p.PstCd = (*Max16Text)(&value)
}

func (p *PostalAddress24) SetTwnNm(value string) {
	p.TwnNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetTwnLctnNm(value string) {
	p.TwnLctnNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetDstrctNm(value string) {
	p.DstrctNm = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetCtrySubDvsn(value string) {
	p.CtrySubDvsn = (*Max35Text)(&value)
}

func (p *PostalAddress24) SetCtry(value string) {
	p.Ctry = (*CountryCode)(&value)
}

func (p *PostalAddress24) AddAdrLine(value string) {
	p.AdrLine = append(p.AdrLine, (*Max70Text)(&value))
}

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress6 struct {
	XMLName xml.Name

	// Identifies the nature of the postal address.
	AdrTp *AddressType2Code `xml:"AdrTp,omitempty" json:"adr_tp" `

	// Identification of a division of a large organisation or building.
	Dept *Max70Text `xml:"Dept,omitempty" json:"dept" `

	// Identification of a sub-division of a large organisation or building.
	SubDept *Max70Text `xml:"SubDept,omitempty" json:"sub_dept" `

	// Name of a street or thoroughfare.
	StrtNm *Max70Text `xml:"StrtNm,omitempty" json:"strt_nm" `

	// Number that identifies the position of a building on a street.
	BldgNb *Max16Text `xml:"BldgNb,omitempty" json:"bldg_nb" `

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PstCd *Max16Text `xml:"PstCd,omitempty" json:"pst_cd" `

	// Name of a built-up area, with defined boundaries, and a local government.
	TwnNm *Max35Text `xml:"TwnNm,omitempty" json:"twn_nm" `

	// Identifies a subdivision of a country such as state, region, county.
	CtrySubDvsn *Max35Text `xml:"CtrySubDvsn,omitempty" json:"ctry_sub_dvsn" `

	// Nation with its own government.
	Ctry *CountryCode `xml:"Ctry,omitempty" json:"ctry" `

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AdrLine []*Max70Text `xml:"AdrLine,omitempty" json:"adr_line" `
}

func (p *PostalAddress6) SetAdrTp(value string) {
	p.AdrTp = (*AddressType2Code)(&value)
}

func (p *PostalAddress6) SetDept(value string) {
	p.Dept = (*Max70Text)(&value)
}

func (p *PostalAddress6) SetSubDept(value string) {
	p.SubDept = (*Max70Text)(&value)
}

func (p *PostalAddress6) SetStrtNm(value string) {
	p.StrtNm = (*Max70Text)(&value)
}

func (p *PostalAddress6) SetBldgNb(value string) {
	p.BldgNb = (*Max16Text)(&value)
}

func (p *PostalAddress6) SetPstCd(value string) {
	p.PstCd = (*Max16Text)(&value)
}

func (p *PostalAddress6) SetTwnNm(value string) {
	p.TwnNm = (*Max35Text)(&value)
}

func (p *PostalAddress6) SetCtrySubDvsn(value string) {
	p.CtrySubDvsn = (*Max35Text)(&value)
}

func (p *PostalAddress6) SetCtry(value string) {
	p.Ctry = (*CountryCode)(&value)
}

func (p *PostalAddress6) AddAdrLine(value string) {
	p.AdrLine = append(p.AdrLine, (*Max70Text)(&value))
}

type PreferredContactMethod1Code string

type Priority2Code string

type Priority3Code string

// Information related to a proxy  identification of the account.
type ProxyAccountIdentification1 struct {
	XMLName xml.Name

	// Type of the proxy identification.
	Tp *ProxyAccountType1Choice `xml:"Tp,omitempty" json:"tp" `

	// Identification used to indicate the account identification under another specified name.
	Id *Max2048Text `xml:"Id" json:"id" `
}

func (p *ProxyAccountIdentification1) AddTp() *ProxyAccountType1Choice {
	p.Tp = new(ProxyAccountType1Choice)
	return p.Tp
}

func (p *ProxyAccountIdentification1) SetId(value string) {
	p.Id = (*Max2048Text)(&value)
}

// Specifies the scheme used for the identification of an account alias.
type ProxyAccountType1Choice struct {
	XMLName xml.Name

	// Name of the identification scheme, in a coded form as published in an external list.
	Cd *ExternalProxyAccountType1Code `xml:"Cd" json:"cd" `

	// Name of the identification scheme, in a free text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (p *ProxyAccountType1Choice) SetCd(value string) {
	p.Cd = (*ExternalProxyAccountType1Code)(&value)
}

func (p *ProxyAccountType1Choice) SetPrtry(value string) {
	p.Prtry = (*Max35Text)(&value)
}

// Specifies the underlying reason for the payment transaction.
// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
type Purpose2Choice struct {
	XMLName xml.Name

	// Underlying reason for the payment transaction, as published in an external purpose code list.
	Cd *ExternalPurpose1Code `xml:"Cd" json:"cd" `

	// Purpose, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (p *Purpose2Choice) SetCd(value string) {
	p.Cd = (*ExternalPurpose1Code)(&value)
}

func (p *Purpose2Choice) SetPrtry(value string) {
	p.Prtry = (*Max35Text)(&value)
}

// Set of elements used to identify the documents referred to in the remittance information.
type ReferredDocumentInformation7 struct {
	XMLName xml.Name

	// Specifies the type of referred document.
	Tp *ReferredDocumentType4 `xml:"Tp,omitempty" json:"tp" `

	// Unique and unambiguous identification of the referred document.
	Nb *Max35Text `xml:"Nb,omitempty" json:"nb" `

	// Date associated with the referred document.
	RltdDt *ISODate `xml:"RltdDt,omitempty" json:"rltd_dt" `

	// Set of elements used to provide the content of the referred document line.
	LineDtls []*DocumentLineInformation1 `xml:"LineDtls,omitempty" json:"line_dtls" `
}

func (r *ReferredDocumentInformation7) AddTp() *ReferredDocumentType4 {
	r.Tp = new(ReferredDocumentType4)
	return r.Tp
}

func (r *ReferredDocumentInformation7) SetNb(value string) {
	r.Nb = (*Max35Text)(&value)
}

func (r *ReferredDocumentInformation7) SetRltdDt(value string) {
	r.RltdDt = (*ISODate)(&value)
}

func (r *ReferredDocumentInformation7) AddLineDtls() *DocumentLineInformation1 {
	newValue := new(DocumentLineInformation1)
	r.LineDtls = append(r.LineDtls, newValue)
	return newValue
}

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType3Choice struct {
	XMLName xml.Name

	// Document type in a coded form.
	Cd *DocumentType6Code `xml:"Cd" json:"cd" `

	// Proprietary identification of the type of the remittance document.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (r *ReferredDocumentType3Choice) SetCd(value string) {
	r.Cd = (*DocumentType6Code)(&value)
}

func (r *ReferredDocumentType3Choice) SetPrtry(value string) {
	r.Prtry = (*Max35Text)(&value)
}

// Specifies the type of the document referred in the remittance information.
type ReferredDocumentType4 struct {
	XMLName xml.Name

	// Provides the type details of the referred document.
	CdOrPrtry *ReferredDocumentType3Choice `xml:"CdOrPrtry" json:"cd_or_prtry" `

	// Identification of the issuer of the reference document type.
	Issr *Max35Text `xml:"Issr,omitempty" json:"issr" `
}

func (r *ReferredDocumentType4) AddCdOrPrtry() *ReferredDocumentType3Choice {
	r.CdOrPrtry = new(ReferredDocumentType3Choice)
	return r.CdOrPrtry
}

func (r *ReferredDocumentType4) SetIssr(value string) {
	r.Issr = (*Max35Text)(&value)
}

// Entity requiring the regulatory reporting information.
type RegulatoryAuthority2 struct {
	XMLName xml.Name

	// Name of the entity requiring the regulatory reporting information.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `

	// Country of the entity that requires the regulatory reporting information.
	Ctry *CountryCode `xml:"Ctry,omitempty" json:"ctry" `
}

func (r *RegulatoryAuthority2) SetNm(value string) {
	r.Nm = (*Max140Text)(&value)
}

func (r *RegulatoryAuthority2) SetCtry(value string) {
	r.Ctry = (*CountryCode)(&value)
}

// Information needed due to regulatory and/or statutory requirements.
type RegulatoryReporting3 struct {
	XMLName xml.Name

	// Identifies whether the regulatory reporting information applies to the debit side, to the credit side or to both debit and credit sides of the transaction.
	DbtCdtRptgInd *RegulatoryReportingType1Code `xml:"DbtCdtRptgInd,omitempty" json:"dbt_cdt_rptg_ind" `

	// Entity requiring the regulatory reporting information.
	Authrty *RegulatoryAuthority2 `xml:"Authrty,omitempty" json:"authrty" `

	// Set of elements used to provide details on the regulatory reporting information.
	Dtls []*StructuredRegulatoryReporting3 `xml:"Dtls,omitempty" json:"dtls" `
}

func (r *RegulatoryReporting3) SetDbtCdtRptgInd(value string) {
	r.DbtCdtRptgInd = (*RegulatoryReportingType1Code)(&value)
}

func (r *RegulatoryReporting3) AddAuthrty() *RegulatoryAuthority2 {
	r.Authrty = new(RegulatoryAuthority2)
	return r.Authrty
}

func (r *RegulatoryReporting3) AddDtls() *StructuredRegulatoryReporting3 {
	newValue := new(StructuredRegulatoryReporting3)
	r.Dtls = append(r.Dtls, newValue)
	return newValue
}

type RegulatoryReportingType1Code string

// Nature of the amount and currency on a document referred to in the remittance section, typically either the original amount due/payable or the amount actually remitted for the referenced document.
type RemittanceAmount2 struct {
	XMLName xml.Name

	// Amount specified is the exact amount due and payable to the creditor.
	DuePyblAmt *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:"due_pybl_amt" `

	// Amount specified for the referred document is the amount of discount to be applied to the amount due and payable to the creditor.
	DscntApldAmt []*DiscountAmountAndType1 `xml:"DscntApldAmt,omitempty" json:"dscnt_apld_amt" `

	// Amount specified for the referred document is the amount of a credit note.
	CdtNoteAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:"cdt_note_amt" `

	// Quantity of cash resulting from the calculation of the tax.
	TaxAmt []*TaxAmountAndType1 `xml:"TaxAmt,omitempty" json:"tax_amt" `

	// Specifies detailed information on the amount and reason of the document adjustment.
	AdjstmntAmtAndRsn []*DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn,omitempty" json:"adjstmnt_amt_and_rsn" `

	// Amount of money remitted for the referred document.
	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"rmtd_amt" `
}

func (r *RemittanceAmount2) SetDuePyblAmt(value, currency string) {
	r.DuePyblAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount2) AddDscntApldAmt() *DiscountAmountAndType1 {
	newValue := new(DiscountAmountAndType1)
	r.DscntApldAmt = append(r.DscntApldAmt, newValue)
	return newValue
}

func (r *RemittanceAmount2) SetCdtNoteAmt(value, currency string) {
	r.CdtNoteAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount2) AddTaxAmt() *TaxAmountAndType1 {
	newValue := new(TaxAmountAndType1)
	r.TaxAmt = append(r.TaxAmt, newValue)
	return newValue
}

func (r *RemittanceAmount2) AddAdjstmntAmtAndRsn() *DocumentAdjustment1 {
	newValue := new(DocumentAdjustment1)
	r.AdjstmntAmtAndRsn = append(r.AdjstmntAmtAndRsn, newValue)
	return newValue
}

func (r *RemittanceAmount2) SetRmtdAmt(value, currency string) {
	r.RmtdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Nature of the amount and currency on a document referred to in the remittance section, typically either the original amount due/payable or the amount actually remitted for the referenced document.
type RemittanceAmount3 struct {
	XMLName xml.Name

	// Amount specified is the exact amount due and payable to the creditor.
	DuePyblAmt *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty" json:"due_pybl_amt" `

	// Amount of discount to be applied to the amount due and payable to the creditor.
	DscntApldAmt []*DiscountAmountAndType1 `xml:"DscntApldAmt,omitempty" json:"dscnt_apld_amt" `

	// Amount of a credit note.
	CdtNoteAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty" json:"cdt_note_amt" `

	// Amount of the tax.
	TaxAmt []*TaxAmountAndType1 `xml:"TaxAmt,omitempty" json:"tax_amt" `

	// Specifies detailed information on the amount and reason of the adjustment.
	AdjstmntAmtAndRsn []*DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn,omitempty" json:"adjstmnt_amt_and_rsn" `

	// Amount of money remitted.
	RmtdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty" json:"rmtd_amt" `
}

func (r *RemittanceAmount3) SetDuePyblAmt(value, currency string) {
	r.DuePyblAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount3) AddDscntApldAmt() *DiscountAmountAndType1 {
	newValue := new(DiscountAmountAndType1)
	r.DscntApldAmt = append(r.DscntApldAmt, newValue)
	return newValue
}

func (r *RemittanceAmount3) SetCdtNoteAmt(value, currency string) {
	r.CdtNoteAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount3) AddTaxAmt() *TaxAmountAndType1 {
	newValue := new(TaxAmountAndType1)
	r.TaxAmt = append(r.TaxAmt, newValue)
	return newValue
}

func (r *RemittanceAmount3) AddAdjstmntAmtAndRsn() *DocumentAdjustment1 {
	newValue := new(DocumentAdjustment1)
	r.AdjstmntAmtAndRsn = append(r.AdjstmntAmtAndRsn, newValue)
	return newValue
}

func (r *RemittanceAmount3) SetRmtdAmt(value, currency string) {
	r.RmtdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation15 struct {
	XMLName xml.Name

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Ustrd []*Max140Text `xml:"Ustrd,omitempty" json:"ustrd" `

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Strd []*StructuredRemittanceInformation15 `xml:"Strd,omitempty" json:"strd" `
}

func (r *RemittanceInformation15) AddUstrd(value string) {
	r.Ustrd = append(r.Ustrd, (*Max140Text)(&value))
}

func (r *RemittanceInformation15) AddStrd() *StructuredRemittanceInformation15 {
	newValue := new(StructuredRemittanceInformation15)
	r.Strd = append(r.Strd, newValue)
	return newValue
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation16 struct {
	XMLName xml.Name

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Ustrd []*Max140Text `xml:"Ustrd,omitempty" json:"ustrd" `

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Strd []*StructuredRemittanceInformation16 `xml:"Strd,omitempty" json:"strd" `
}

func (r *RemittanceInformation16) AddUstrd(value string) {
	r.Ustrd = append(r.Ustrd, (*Max140Text)(&value))
}

func (r *RemittanceInformation16) AddStrd() *StructuredRemittanceInformation16 {
	newValue := new(StructuredRemittanceInformation16)
	r.Strd = append(r.Strd, newValue)
	return newValue
}

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle.
type RemittanceInformation2 struct {
	XMLName xml.Name

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, for example, commercial invoices in an accounts' receivable system in an unstructured form.
	Ustrd []*Max140Text `xml:"Ustrd,omitempty" json:"ustrd" `
}

func (r *RemittanceInformation2) AddUstrd(value string) {
	r.Ustrd = append(r.Ustrd, (*Max140Text)(&value))
}

// Set of elements used to provide information on the remittance advice.
type RemittanceLocation4 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the initiating party, to unambiguously identify the remittance information sent separately from the payment instruction, such as a remittance advice.
	RmtId *Max35Text `xml:"RmtId,omitempty" json:"rmt_id" `

	// Set of elements used to provide information on the location and/or delivery of the remittance information.
	RmtLctnDtls []*RemittanceLocationDetails1 `xml:"RmtLctnDtls,omitempty" json:"rmt_lctn_dtls" `
}

func (r *RemittanceLocation4) SetRmtId(value string) {
	r.RmtId = (*Max35Text)(&value)
}

func (r *RemittanceLocation4) AddRmtLctnDtls() *RemittanceLocationDetails1 {
	newValue := new(RemittanceLocationDetails1)
	r.RmtLctnDtls = append(r.RmtLctnDtls, newValue)
	return newValue
}

// Provides information on the remittance advice.
type RemittanceLocationDetails1 struct {
	XMLName xml.Name

	// Method used to deliver the remittance advice information.
	Mtd *RemittanceLocationMethod2Code `xml:"Mtd" json:"mtd" `

	// Electronic address to which an agent is to send the remittance information.
	ElctrncAdr *Max2048Text `xml:"ElctrncAdr,omitempty" json:"elctrnc_adr" `

	// Postal address to which an agent is to send the remittance information.
	PstlAdr *NameAndAddress10 `xml:"PstlAdr,omitempty" json:"pstl_adr" `
}

func (r *RemittanceLocationDetails1) SetMtd(value string) {
	r.Mtd = (*RemittanceLocationMethod2Code)(&value)
}

func (r *RemittanceLocationDetails1) SetElctrncAdr(value string) {
	r.ElctrncAdr = (*Max2048Text)(&value)
}

func (r *RemittanceLocationDetails1) AddPstlAdr() *NameAndAddress10 {
	r.PstlAdr = new(NameAndAddress10)
	return r.PstlAdr
}

type RemittanceLocationMethod2Code string

// Specifies generic information about an investigation report.
type ReportHeader5 struct {
	XMLName xml.Name

	// Point to point reference as assigned by the case assigner to unambiguously identify the case status report.
	Id *Max35Text `xml:"Id" json:"id" `

	// Party reporting the status of the investigation case.
	Fr *Party40Choice `xml:"Fr" json:"fr" `

	// Party to which the status of the case is reported.
	To *Party40Choice `xml:"To" json:"to" `

	// Date and time at which the message was created.
	CreDtTm *ISODateTime `xml:"CreDtTm" json:"cre_dt_tm" `
}

func (r *ReportHeader5) SetId(value string) {
	r.Id = (*Max35Text)(&value)
}

func (r *ReportHeader5) AddFr() *Party40Choice {
	r.Fr = new(Party40Choice)
	return r.Fr
}

func (r *ReportHeader5) AddTo() *Party40Choice {
	r.To = new(Party40Choice)
	return r.To
}

func (r *ReportHeader5) SetCreDtTm(value string) {
	r.CreDtTm = (*ISODateTime)(&value)
}

// Provide further details on the requested modifications of the underlying payment instruction.
type RequestedModification8 struct {
	XMLName xml.Name

	// Unique identification as assigned by an instructing party for an instructed party to unambiguously identify the instruction.
	//
	// Usage: The instruction identification is a point to point reference that can be used between the instructing party and the instructed party to refer to the individual instruction. It can be included in several messages related to the instruction.
	InstrId *Max35Text `xml:"InstrId,omitempty" json:"instr_id" `

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	//
	// Usage: The end-to-end identification can be used for reconciliation or to link tasks relating to the transaction. It can be included in several messages related to the transaction.
	EndToEndId *Max35Text `xml:"EndToEndId,omitempty" json:"end_to_end_id" `

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	// Usage: The transaction identification can be used for reconciliation, tracking or to link tasks relating to the transaction on the interbank level. The instructing agent has to make sure that the transaction identification is unique for a pre-agreed period.
	TxId *Max35Text `xml:"TxId,omitempty" json:"tx_id" `

	// Set of elements used to further specify the type of transaction.
	PmtTpInf *PaymentTypeInformation27 `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf" `

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty" json:"reqd_exctn_dt" `

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	ReqdColltnDt *ISODate `xml:"ReqdColltnDt,omitempty" json:"reqd_colltn_dt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Provides information on the occurred settlement time(s) of the payment transaction.
	SttlmTmIndctn *SettlementDateTimeIndication1 `xml:"SttlmTmIndctn,omitempty" json:"sttlm_tm_indctn" `

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amt *AmountType4Choice `xml:"Amt,omitempty" json:"amt" `

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:"intr_bk_sttlm_amt" `

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChrgBr *ChargeBearerType1Code `xml:"ChrgBr,omitempty" json:"chrg_br" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltmtDbtr *PartyIdentification135 `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Party that owes an amount of money to the (ultimate) creditor. In the context of the payment model, the debtor is also the debit account owner.
	Dbtr *PartyIdentification135 `xml:"Dbtr,omitempty" json:"dbtr" `

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DbtrAcct *CashAccount38 `xml:"DbtrAcct,omitempty" json:"dbtr_acct" `

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DbtrAgtAcct *CashAccount38 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct" `

	// Instruction between two clearing agents stipulating the cash transfer characteristics between the two parties.
	SttlmInf *SettlementInstruction6 `xml:"SttlmInf,omitempty" json:"sttlm_inf" `

	// Unambiguous identification of the account of the creditor agent at its servicing agent to which a credit entry will be made as a result of the payment transaction.
	CdtrAgtAcct *CashAccount38 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct" `

	// Party to which an amount of money is due.
	Cdtr *PartyIdentification135 `xml:"Cdtr,omitempty" json:"cdtr" `

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CdtrAcct *CashAccount38 `xml:"CdtrAcct,omitempty" json:"cdtr_acct" `

	// Ultimate party to which an amount of money is due.
	UltmtCdtr *PartyIdentification135 `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purp *Purpose2Choice `xml:"Purp,omitempty" json:"purp" `

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the debtor agent.
	InstrForDbtrAgt *Max140Text `xml:"InstrForDbtrAgt,omitempty" json:"instr_for_dbtr_agt" `

	// Further information related to the processing of the payment instruction that may need to be acted upon by the next agent.
	//
	// Usage: The next agent may not be the creditor agent.
	// The instruction can relate to a level of service, can be an instruction that has to be executed by the agent, or can be information required by the next agent.
	InstrForNxtAgt []*InstructionForNextAgent1 `xml:"InstrForNxtAgt,omitempty" json:"instr_for_nxt_agt" `

	// Further information related to the processing of the payment instruction, provided by the initiating party, and intended for the creditor agent.
	InstrForCdtrAgt []*InstructionForCreditorAgent1 `xml:"InstrForCdtrAgt,omitempty" json:"instr_for_cdtr_agt" `

	// Structured information that enables the matching, that is reconciliation, of a payment with the items that the payment is intended to settle, such as commercial invoices in an account receivable system.
	RmtInf *RemittanceInformation16 `xml:"RmtInf,omitempty" json:"rmt_inf" `
}

func (r *RequestedModification8) SetInstrId(value string) {
	r.InstrId = (*Max35Text)(&value)
}

func (r *RequestedModification8) SetEndToEndId(value string) {
	r.EndToEndId = (*Max35Text)(&value)
}

func (r *RequestedModification8) SetTxId(value string) {
	r.TxId = (*Max35Text)(&value)
}

func (r *RequestedModification8) AddPmtTpInf() *PaymentTypeInformation27 {
	r.PmtTpInf = new(PaymentTypeInformation27)
	return r.PmtTpInf
}

func (r *RequestedModification8) AddReqdExctnDt() *DateAndDateTime2Choice {
	r.ReqdExctnDt = new(DateAndDateTime2Choice)
	return r.ReqdExctnDt
}

func (r *RequestedModification8) SetReqdColltnDt(value string) {
	r.ReqdColltnDt = (*ISODate)(&value)
}

func (r *RequestedModification8) SetIntrBkSttlmDt(value string) {
	r.IntrBkSttlmDt = (*ISODate)(&value)
}

func (r *RequestedModification8) AddSttlmTmIndctn() *SettlementDateTimeIndication1 {
	r.SttlmTmIndctn = new(SettlementDateTimeIndication1)
	return r.SttlmTmIndctn
}

func (r *RequestedModification8) AddAmt() *AmountType4Choice {
	r.Amt = new(AmountType4Choice)
	return r.Amt
}

func (r *RequestedModification8) SetIntrBkSttlmAmt(value, currency string) {
	r.IntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RequestedModification8) SetChrgBr(value string) {
	r.ChrgBr = (*ChargeBearerType1Code)(&value)
}

func (r *RequestedModification8) AddUltmtDbtr() *PartyIdentification135 {
	r.UltmtDbtr = new(PartyIdentification135)
	return r.UltmtDbtr
}

func (r *RequestedModification8) AddDbtr() *PartyIdentification135 {
	r.Dbtr = new(PartyIdentification135)
	return r.Dbtr
}

func (r *RequestedModification8) AddDbtrAcct() *CashAccount38 {
	r.DbtrAcct = new(CashAccount38)
	return r.DbtrAcct
}

func (r *RequestedModification8) AddDbtrAgtAcct() *CashAccount38 {
	r.DbtrAgtAcct = new(CashAccount38)
	return r.DbtrAgtAcct
}

func (r *RequestedModification8) AddSttlmInf() *SettlementInstruction6 {
	r.SttlmInf = new(SettlementInstruction6)
	return r.SttlmInf
}

func (r *RequestedModification8) AddCdtrAgtAcct() *CashAccount38 {
	r.CdtrAgtAcct = new(CashAccount38)
	return r.CdtrAgtAcct
}

func (r *RequestedModification8) AddCdtr() *PartyIdentification135 {
	r.Cdtr = new(PartyIdentification135)
	return r.Cdtr
}

func (r *RequestedModification8) AddCdtrAcct() *CashAccount38 {
	r.CdtrAcct = new(CashAccount38)
	return r.CdtrAcct
}

func (r *RequestedModification8) AddUltmtCdtr() *PartyIdentification135 {
	r.UltmtCdtr = new(PartyIdentification135)
	return r.UltmtCdtr
}

func (r *RequestedModification8) AddPurp() *Purpose2Choice {
	r.Purp = new(Purpose2Choice)
	return r.Purp
}

func (r *RequestedModification8) SetInstrForDbtrAgt(value string) {
	r.InstrForDbtrAgt = (*Max140Text)(&value)
}

func (r *RequestedModification8) AddInstrForNxtAgt() *InstructionForNextAgent1 {
	newValue := new(InstructionForNextAgent1)
	r.InstrForNxtAgt = append(r.InstrForNxtAgt, newValue)
	return newValue
}

func (r *RequestedModification8) AddInstrForCdtrAgt() *InstructionForCreditorAgent1 {
	newValue := new(InstructionForCreditorAgent1)
	r.InstrForCdtrAgt = append(r.InstrForCdtrAgt, newValue)
	return newValue
}

func (r *RequestedModification8) AddRmtInf() *RemittanceInformation16 {
	r.RmtInf = new(RemittanceInformation16)
	return r.RmtInf
}

// Specifies additional information as expected by the party that the investigation performs the expected actions for its resolution.
type ResolutionData1 struct {
	XMLName xml.Name

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	EndToEndId *Max35Text `xml:"EndToEndId,omitempty" json:"end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	TxId *Max35Text `xml:"TxId,omitempty" json:"tx_id" `

	// Universally unique identifier to provide an end-to-end reference of a payment transaction.
	UETR *UUIDv4Identifier `xml:"UETR,omitempty" json:"uetr" `

	// Amount of money moved between the instructing agent and the instructed agent.
	IntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty" json:"intr_bk_sttlm_amt" `

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	IntrBkSttlmDt *ISODate `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt" `

	// Specifies the clearing channel to be used to process the payment instruction.
	ClrChanl *ClearingChannel2Code `xml:"ClrChanl,omitempty" json:"clr_chanl" `

	// Provides the details of the compensation made due to the modification or cancellation of a previous payment.
	Compstn *Compensation2 `xml:"Compstn,omitempty" json:"compstn" `

	// Provides information on the charges to be paid by the charge bearer(s) related to the payment transaction.
	Chrgs []*Charges7 `xml:"Chrgs,omitempty" json:"chrgs" `
}

func (r *ResolutionData1) SetEndToEndId(value string) {
	r.EndToEndId = (*Max35Text)(&value)
}

func (r *ResolutionData1) SetTxId(value string) {
	r.TxId = (*Max35Text)(&value)
}

func (r *ResolutionData1) SetUETR(value string) {
	r.UETR = (*UUIDv4Identifier)(&value)
}

func (r *ResolutionData1) SetIntrBkSttlmAmt(value, currency string) {
	r.IntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *ResolutionData1) SetIntrBkSttlmDt(value string) {
	r.IntrBkSttlmDt = (*ISODate)(&value)
}

func (r *ResolutionData1) SetClrChanl(value string) {
	r.ClrChanl = (*ClearingChannel2Code)(&value)
}

func (r *ResolutionData1) AddCompstn() *Compensation2 {
	r.Compstn = new(Compensation2)
	return r.Compstn
}

func (r *ResolutionData1) AddChrgs() *Charges7 {
	newValue := new(Charges7)
	r.Chrgs = append(r.Chrgs, newValue)
	return newValue
}

// Specifies the reason for the return of the transaction.
type ReturnReason5Choice struct {
	XMLName xml.Name

	// Reason for the return, as published in an external reason code list.
	Cd *ExternalReturnReason1Code `xml:"Cd" json:"cd" `

	// Reason for the return, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (r *ReturnReason5Choice) SetCd(value string) {
	r.Cd = (*ExternalReturnReason1Code)(&value)
}

func (r *ReturnReason5Choice) SetPrtry(value string) {
	r.Prtry = (*Max35Text)(&value)
}

type SequenceType3Code string

// Specifies the service level of the transaction.
type ServiceLevel8Choice struct {
	XMLName xml.Name

	// Specifies a pre-agreed service or level of service between the parties, as published in an external service level code list.
	Cd *ExternalServiceLevel1Code `xml:"Cd" json:"cd" `

	// Specifies a pre-agreed service or level of service between the parties, as a proprietary code.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (s *ServiceLevel8Choice) SetCd(value string) {
	s.Cd = (*ExternalServiceLevel1Code)(&value)
}

func (s *ServiceLevel8Choice) SetPrtry(value string) {
	s.Prtry = (*Max35Text)(&value)
}

// Information on the occurred settlement time(s) of the payment transaction.
type SettlementDateTimeIndication1 struct {
	XMLName xml.Name

	// Date and time at which a payment has been debited at the transaction administrator. In the case of TARGET, the date and time at which the payment has been debited at the central bank, expressed in Central European Time (CET).
	DbtDtTm *ISODateTime `xml:"DbtDtTm,omitempty" json:"dbt_dt_tm" `

	// Date and time at which a payment has been credited at the transaction administrator. In the case of TARGET, the date and time at which the payment has been credited at the receiving central bank, expressed in Central European Time (CET).
	CdtDtTm *ISODateTime `xml:"CdtDtTm,omitempty" json:"cdt_dt_tm" `
}

func (s *SettlementDateTimeIndication1) SetDbtDtTm(value string) {
	s.DbtDtTm = (*ISODateTime)(&value)
}

func (s *SettlementDateTimeIndication1) SetCdtDtTm(value string) {
	s.CdtDtTm = (*ISODateTime)(&value)
}

// Provides further details on the settlement of the instruction.
type SettlementInstruction4 struct {
	XMLName xml.Name

	// Method used to settle the (batch of) payment instructions.
	SttlmMtd *SettlementMethod1Code `xml:"SttlmMtd" json:"sttlm_mtd" `

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SttlmAcct *CashAccount24 `xml:"SttlmAcct,omitempty" json:"sttlm_acct" `

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClrSys *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty" json:"clr_sys" `

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstgRmbrsmntAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstgRmbrsmntAgt,omitempty" json:"instg_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstgRmbrsmntAgtAcct *CashAccount24 `xml:"InstgRmbrsmntAgtAcct,omitempty" json:"instg_rmbrsmnt_agt_acct" `

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification5 `xml:"InstdRmbrsmntAgt,omitempty" json:"instd_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstdRmbrsmntAgtAcct *CashAccount24 `xml:"InstdRmbrsmntAgtAcct,omitempty" json:"instd_rmbrsmnt_agt_acct" `

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThrdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification5 `xml:"ThrdRmbrsmntAgt,omitempty" json:"thrd_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThrdRmbrsmntAgtAcct *CashAccount24 `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:"thrd_rmbrsmnt_agt_acct" `
}

func (s *SettlementInstruction4) SetSttlmMtd(value string) {
	s.SttlmMtd = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInstruction4) AddSttlmAcct() *CashAccount24 {
	s.SttlmAcct = new(CashAccount24)
	return s.SttlmAcct
}

func (s *SettlementInstruction4) AddClrSys() *ClearingSystemIdentification3Choice {
	s.ClrSys = new(ClearingSystemIdentification3Choice)
	return s.ClrSys
}

func (s *SettlementInstruction4) AddInstgRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification5 {
	s.InstgRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstgRmbrsmntAgt
}

func (s *SettlementInstruction4) AddInstgRmbrsmntAgtAcct() *CashAccount24 {
	s.InstgRmbrsmntAgtAcct = new(CashAccount24)
	return s.InstgRmbrsmntAgtAcct
}

func (s *SettlementInstruction4) AddInstdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification5 {
	s.InstdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification5)
	return s.InstdRmbrsmntAgt
}

func (s *SettlementInstruction4) AddInstdRmbrsmntAgtAcct() *CashAccount24 {
	s.InstdRmbrsmntAgtAcct = new(CashAccount24)
	return s.InstdRmbrsmntAgtAcct
}

func (s *SettlementInstruction4) AddThrdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification5 {
	s.ThrdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification5)
	return s.ThrdRmbrsmntAgt
}

func (s *SettlementInstruction4) AddThrdRmbrsmntAgtAcct() *CashAccount24 {
	s.ThrdRmbrsmntAgtAcct = new(CashAccount24)
	return s.ThrdRmbrsmntAgtAcct
}

// Provides further details on the settlement of the instruction.
type SettlementInstruction6 struct {
	XMLName xml.Name

	// Agent through which the instructing agent will reimburse the instructed agent.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstgRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:"instg_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstgRmbrsmntAgtAcct *CashAccount38 `xml:"InstgRmbrsmntAgtAcct,omitempty" json:"instg_rmbrsmnt_agt_acct" `

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:"instd_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstdRmbrsmntAgtAcct *CashAccount38 `xml:"InstdRmbrsmntAgtAcct,omitempty" json:"instd_rmbrsmnt_agt_acct" `
}

func (s *SettlementInstruction6) AddInstgRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.InstgRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.InstgRmbrsmntAgt
}

func (s *SettlementInstruction6) AddInstgRmbrsmntAgtAcct() *CashAccount38 {
	s.InstgRmbrsmntAgtAcct = new(CashAccount38)
	return s.InstgRmbrsmntAgtAcct
}

func (s *SettlementInstruction6) AddInstdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.InstdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.InstdRmbrsmntAgt
}

func (s *SettlementInstruction6) AddInstdRmbrsmntAgtAcct() *CashAccount38 {
	s.InstdRmbrsmntAgtAcct = new(CashAccount38)
	return s.InstdRmbrsmntAgtAcct
}

// Provides further details on the settlement of the instruction.
type SettlementInstruction7 struct {
	XMLName xml.Name

	// Method used to settle the (batch of) payment instructions.
	SttlmMtd *SettlementMethod1Code `xml:"SttlmMtd" json:"sttlm_mtd" `

	// A specific purpose account used to post debit and credit entries as a result of the transaction.
	SttlmAcct *CashAccount38 `xml:"SttlmAcct,omitempty" json:"sttlm_acct" `

	// Specification of a pre-agreed offering between clearing agents or the channel through which the payment instruction is processed.
	ClrSys *ClearingSystemIdentification3Choice `xml:"ClrSys,omitempty" json:"clr_sys" `

	// Agent through which the instructing agent will reimburse the instructed agent.
	//
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstgRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstgRmbrsmntAgt,omitempty" json:"instg_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the instructing reimbursement agent account at its servicing agent in the payment chain.
	InstgRmbrsmntAgtAcct *CashAccount38 `xml:"InstgRmbrsmntAgtAcct,omitempty" json:"instg_rmbrsmnt_agt_acct" `

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If InstructedReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	// Usage: If InstructingAgent and InstructedAgent have the same reimbursement agent, then only InstructingReimbursementAgent must be used.
	InstdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"InstdRmbrsmntAgt,omitempty" json:"instd_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the instructed reimbursement agent account at its servicing agent in the payment chain.
	InstdRmbrsmntAgtAcct *CashAccount38 `xml:"InstdRmbrsmntAgtAcct,omitempty" json:"instd_rmbrsmnt_agt_acct" `

	// Agent at which the instructed agent will be reimbursed.
	// Usage: If ThirdReimbursementAgent contains a branch of the InstructedAgent, then the party in InstructedAgent will claim reimbursement from that branch/will be paid by that branch.
	ThrdRmbrsmntAgt *BranchAndFinancialInstitutionIdentification6 `xml:"ThrdRmbrsmntAgt,omitempty" json:"thrd_rmbrsmnt_agt" `

	// Unambiguous identification of the account of the third reimbursement agent account at its servicing agent in the payment chain.
	ThrdRmbrsmntAgtAcct *CashAccount38 `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:"thrd_rmbrsmnt_agt_acct" `
}

func (s *SettlementInstruction7) SetSttlmMtd(value string) {
	s.SttlmMtd = (*SettlementMethod1Code)(&value)
}

func (s *SettlementInstruction7) AddSttlmAcct() *CashAccount38 {
	s.SttlmAcct = new(CashAccount38)
	return s.SttlmAcct
}

func (s *SettlementInstruction7) AddClrSys() *ClearingSystemIdentification3Choice {
	s.ClrSys = new(ClearingSystemIdentification3Choice)
	return s.ClrSys
}

func (s *SettlementInstruction7) AddInstgRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.InstgRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.InstgRmbrsmntAgt
}

func (s *SettlementInstruction7) AddInstgRmbrsmntAgtAcct() *CashAccount38 {
	s.InstgRmbrsmntAgtAcct = new(CashAccount38)
	return s.InstgRmbrsmntAgtAcct
}

func (s *SettlementInstruction7) AddInstdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.InstdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.InstdRmbrsmntAgt
}

func (s *SettlementInstruction7) AddInstdRmbrsmntAgtAcct() *CashAccount38 {
	s.InstdRmbrsmntAgtAcct = new(CashAccount38)
	return s.InstdRmbrsmntAgtAcct
}

func (s *SettlementInstruction7) AddThrdRmbrsmntAgt() *BranchAndFinancialInstitutionIdentification6 {
	s.ThrdRmbrsmntAgt = new(BranchAndFinancialInstitutionIdentification6)
	return s.ThrdRmbrsmntAgt
}

func (s *SettlementInstruction7) AddThrdRmbrsmntAgtAcct() *CashAccount38 {
	s.ThrdRmbrsmntAgtAcct = new(CashAccount38)
	return s.ThrdRmbrsmntAgtAcct
}

type SettlementMethod1Code string

// Provides information on the requested settlement time(s) of the payment instruction.
type SettlementTimeRequest2 struct {
	XMLName xml.Name

	// Time by which the amount of money must be credited, with confirmation, to the CLS Bank's account at the central bank.
	// Usage: Time must be expressed in Central European Time (CET).
	CLSTm *ISOTime `xml:"CLSTm,omitempty" json:"c_l_s_tm" `

	// Time until when the payment may be settled.
	TillTm *ISOTime `xml:"TillTm,omitempty" json:"till_tm" `

	// Time as from when the payment may be settled.
	FrTm *ISOTime `xml:"FrTm,omitempty" json:"fr_tm" `

	// Time by when the payment must be settled to avoid rejection.
	RjctTm *ISOTime `xml:"RjctTm,omitempty" json:"rjct_tm" `
}

func (s *SettlementTimeRequest2) SetCLSTm(value string) {
	s.CLSTm = (*ISOTime)(&value)
}

func (s *SettlementTimeRequest2) SetTillTm(value string) {
	s.TillTm = (*ISOTime)(&value)
}

func (s *SettlementTimeRequest2) SetFrTm(value string) {
	s.FrTm = (*ISOTime)(&value)
}

func (s *SettlementTimeRequest2) SetRjctTm(value string) {
	s.RjctTm = (*ISOTime)(&value)
}

type SignatureEnvelope string

// Provides further details on the statement entry for the resolution of the investigation.
type StatementResolutionEntry4 struct {
	XMLName xml.Name

	// Provides information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the account servicer, to unambiguously identify the original statement.
	OrgnlStmtId *Max35Text `xml:"OrgnlStmtId,omitempty" json:"orgnl_stmt_id" `

	// Universally unique identifier to provide an end-to-end reference of a payment transaction.
	UETR *UUIDv4Identifier `xml:"UETR,omitempty" json:"uetr" `

	// Unique reference, as assigned by the account servicing institution, to unambiguously identify the entry.
	AcctSvcrRef *Max35Text `xml:"AcctSvcrRef,omitempty" json:"acct_svcr_ref" `

	// Corrected debit or credit amount, compared to the original entry where the amount is incorrect.
	//
	// Usage: This amount may only be present if an incorrect statement entry has been reported.
	CrrctdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"CrrctdAmt,omitempty" json:"crrctd_amt" `

	// Provides information on the charges included in the original entry amount.
	Chrgs []*Charges6 `xml:"Chrgs,omitempty" json:"chrgs" `

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purp *Purpose2Choice `xml:"Purp,omitempty" json:"purp" `
}

func (s *StatementResolutionEntry4) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	s.OrgnlGrpInf = new(OriginalGroupInformation29)
	return s.OrgnlGrpInf
}

func (s *StatementResolutionEntry4) SetOrgnlStmtId(value string) {
	s.OrgnlStmtId = (*Max35Text)(&value)
}

func (s *StatementResolutionEntry4) SetUETR(value string) {
	s.UETR = (*UUIDv4Identifier)(&value)
}

func (s *StatementResolutionEntry4) SetAcctSvcrRef(value string) {
	s.AcctSvcrRef = (*Max35Text)(&value)
}

func (s *StatementResolutionEntry4) SetCrrctdAmt(value, currency string) {
	s.CrrctdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *StatementResolutionEntry4) AddChrgs() *Charges6 {
	newValue := new(Charges6)
	s.Chrgs = append(s.Chrgs, newValue)
	return newValue
}

func (s *StatementResolutionEntry4) AddPurp() *Purpose2Choice {
	s.Purp = new(Purpose2Choice)
	return s.Purp
}

// Specifies the reason for the status of the transaction.
type StatusReason6Choice struct {
	XMLName xml.Name

	// Reason for the status, as published in an external reason code list.
	Cd *ExternalStatusReason1Code `xml:"Cd" json:"cd" `

	// Reason for the status, in a proprietary form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (s *StatusReason6Choice) SetCd(value string) {
	s.Cd = (*ExternalStatusReason1Code)(&value)
}

func (s *StatusReason6Choice) SetPrtry(value string) {
	s.Prtry = (*Max35Text)(&value)
}

// Set of elements used to provide information on the status reason of the transaction.
type StatusReasonInformation11 struct {
	XMLName xml.Name

	// Party that issues the status.
	Orgtr *PartyIdentification125 `xml:"Orgtr,omitempty" json:"orgtr" `

	// Specifies the reason for the status report.
	Rsn *StatusReason6Choice `xml:"Rsn,omitempty" json:"rsn" `

	// Further details on the status reason.
	//
	// Usage: Additional information can be used for several purposes such as the reporting of repaired information.
	AddtlInf []*Max105Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (s *StatusReasonInformation11) AddOrgtr() *PartyIdentification125 {
	s.Orgtr = new(PartyIdentification125)
	return s.Orgtr
}

func (s *StatusReasonInformation11) AddRsn() *StatusReason6Choice {
	s.Rsn = new(StatusReason6Choice)
	return s.Rsn
}

func (s *StatusReasonInformation11) AddAddtlInf(value string) {
	s.AddtlInf = append(s.AddtlInf, (*Max105Text)(&value))
}

// Information needed due to regulatory and statutory requirements.
type StructuredRegulatoryReporting3 struct {
	XMLName xml.Name

	// Specifies the type of the information supplied in the regulatory reporting details.
	Tp *Max35Text `xml:"Tp,omitempty" json:"tp" `

	// Date related to the specified type of regulatory reporting details.
	Dt *ISODate `xml:"Dt,omitempty" json:"dt" `

	// Country related to the specified type of regulatory reporting details.
	Ctry *CountryCode `xml:"Ctry,omitempty" json:"ctry" `

	// Specifies the nature, purpose, and reason for the transaction to be reported for regulatory and statutory requirements in a coded form.
	Cd *Max10Text `xml:"Cd,omitempty" json:"cd" `

	// Amount of money to be reported for regulatory and statutory requirements.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:"amt" `

	// Additional details that cater for specific domestic regulatory requirements.
	Inf []*Max35Text `xml:"Inf,omitempty" json:"inf" `
}

func (s *StructuredRegulatoryReporting3) SetTp(value string) {
	s.Tp = (*Max35Text)(&value)
}

func (s *StructuredRegulatoryReporting3) SetDt(value string) {
	s.Dt = (*ISODate)(&value)
}

func (s *StructuredRegulatoryReporting3) SetCtry(value string) {
	s.Ctry = (*CountryCode)(&value)
}

func (s *StructuredRegulatoryReporting3) SetCd(value string) {
	s.Cd = (*Max10Text)(&value)
}

func (s *StructuredRegulatoryReporting3) SetAmt(value, currency string) {
	s.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *StructuredRegulatoryReporting3) AddInf(value string) {
	s.Inf = append(s.Inf, (*Max35Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation15 struct {
	XMLName xml.Name

	// Provides the identification and the content of the referred document.
	RfrdDocInf []*ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:"rfrd_doc_inf" `

	// Provides details on the amounts of the referred document.
	RfrdDocAmt *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty" json:"rfrd_doc_amt" `

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CdtrRefInf *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:"cdtr_ref_inf" `

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invcr *PartyIdentification125 `xml:"Invcr,omitempty" json:"invcr" `

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invcee *PartyIdentification125 `xml:"Invcee,omitempty" json:"invcee" `

	// Provides remittance information about a payment made for tax-related purposes.
	TaxRmt *TaxInformation7 `xml:"TaxRmt,omitempty" json:"tax_rmt" `

	// Provides remittance information about a payment for garnishment-related purposes.
	GrnshmtRmt *Garnishment2 `xml:"GrnshmtRmt,omitempty" json:"grnshmt_rmt" `

	// Additional information, in free text form, to complement the structured remittance information.
	AddtlRmtInf []*Max140Text `xml:"AddtlRmtInf,omitempty" json:"addtl_rmt_inf" `
}

func (s *StructuredRemittanceInformation15) AddRfrdDocInf() *ReferredDocumentInformation7 {
	newValue := new(ReferredDocumentInformation7)
	s.RfrdDocInf = append(s.RfrdDocInf, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation15) AddRfrdDocAmt() *RemittanceAmount2 {
	s.RfrdDocAmt = new(RemittanceAmount2)
	return s.RfrdDocAmt
}

func (s *StructuredRemittanceInformation15) AddCdtrRefInf() *CreditorReferenceInformation2 {
	s.CdtrRefInf = new(CreditorReferenceInformation2)
	return s.CdtrRefInf
}

func (s *StructuredRemittanceInformation15) AddInvcr() *PartyIdentification125 {
	s.Invcr = new(PartyIdentification125)
	return s.Invcr
}

func (s *StructuredRemittanceInformation15) AddInvcee() *PartyIdentification125 {
	s.Invcee = new(PartyIdentification125)
	return s.Invcee
}

func (s *StructuredRemittanceInformation15) AddTaxRmt() *TaxInformation7 {
	s.TaxRmt = new(TaxInformation7)
	return s.TaxRmt
}

func (s *StructuredRemittanceInformation15) AddGrnshmtRmt() *Garnishment2 {
	s.GrnshmtRmt = new(Garnishment2)
	return s.GrnshmtRmt
}

func (s *StructuredRemittanceInformation15) AddAddtlRmtInf(value string) {
	s.AddtlRmtInf = append(s.AddtlRmtInf, (*Max140Text)(&value))
}

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation16 struct {
	XMLName xml.Name

	// Provides the identification and the content of the referred document.
	RfrdDocInf []*ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty" json:"rfrd_doc_inf" `

	// Provides details on the amounts of the referred document.
	RfrdDocAmt *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty" json:"rfrd_doc_amt" `

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CdtrRefInf *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty" json:"cdtr_ref_inf" `

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invcr *PartyIdentification135 `xml:"Invcr,omitempty" json:"invcr" `

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invcee *PartyIdentification135 `xml:"Invcee,omitempty" json:"invcee" `

	// Provides remittance information about a payment made for tax-related purposes.
	TaxRmt *TaxInformation7 `xml:"TaxRmt,omitempty" json:"tax_rmt" `

	// Provides remittance information about a payment for garnishment-related purposes.
	GrnshmtRmt *Garnishment3 `xml:"GrnshmtRmt,omitempty" json:"grnshmt_rmt" `

	// Additional information, in free text form, to complement the structured remittance information.
	AddtlRmtInf []*Max140Text `xml:"AddtlRmtInf,omitempty" json:"addtl_rmt_inf" `
}

func (s *StructuredRemittanceInformation16) AddRfrdDocInf() *ReferredDocumentInformation7 {
	newValue := new(ReferredDocumentInformation7)
	s.RfrdDocInf = append(s.RfrdDocInf, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation16) AddRfrdDocAmt() *RemittanceAmount2 {
	s.RfrdDocAmt = new(RemittanceAmount2)
	return s.RfrdDocAmt
}

func (s *StructuredRemittanceInformation16) AddCdtrRefInf() *CreditorReferenceInformation2 {
	s.CdtrRefInf = new(CreditorReferenceInformation2)
	return s.CdtrRefInf
}

func (s *StructuredRemittanceInformation16) AddInvcr() *PartyIdentification135 {
	s.Invcr = new(PartyIdentification135)
	return s.Invcr
}

func (s *StructuredRemittanceInformation16) AddInvcee() *PartyIdentification135 {
	s.Invcee = new(PartyIdentification135)
	return s.Invcee
}

func (s *StructuredRemittanceInformation16) AddTaxRmt() *TaxInformation7 {
	s.TaxRmt = new(TaxInformation7)
	return s.TaxRmt
}

func (s *StructuredRemittanceInformation16) AddGrnshmtRmt() *Garnishment3 {
	s.GrnshmtRmt = new(Garnishment3)
	return s.GrnshmtRmt
}

func (s *StructuredRemittanceInformation16) AddAddtlRmtInf(value string) {
	s.AddtlRmtInf = append(s.AddtlRmtInf, (*Max140Text)(&value))
}

// Additional information that can not be captured in the structured fields and/or any other specific block.
type SupplementaryData1 struct {
	XMLName xml.Name

	// Unambiguous reference to the location where the supplementary data must be inserted in the message instance.
	// In the case of XML, this is expressed by a valid XPath.
	PlcAndNm *Max350Text `xml:"PlcAndNm,omitempty" json:"plc_and_nm" `

	// Technical element wrapping the supplementary data.
	Envlp *SupplementaryDataEnvelope1 `xml:"Envlp" json:"envlp" `
}

func (s *SupplementaryData1) SetPlcAndNm(value string) {
	s.PlcAndNm = (*Max350Text)(&value)
}

func (s *SupplementaryData1) AddEnvlp() *SupplementaryDataEnvelope1 {
	s.Envlp = new(SupplementaryDataEnvelope1)
	return s.Envlp
}

// Technical component that contains the validated supplementary data information. This technical envelope allows to segregate the supplementary data information from any other information.
type SupplementaryDataEnvelope1 struct {
	XMLName xml.Name

	// Identification.
	Id *Max34Text `xml:"Id" json:"id" `
}

func (s *SupplementaryDataEnvelope1) SetId(value string) {
	s.Id = (*Max34Text)(&value)
}

// Set of elements used to provide information on the tax amount(s) of tax record.
type TaxAmount2 struct {
	XMLName xml.Name

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty" json:"rate" `

	// Amount of money on which the tax is based.
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TaxblBaseAmt,omitempty" json:"taxbl_base_amt" `

	// Total amount that is the result of the calculation of the tax for the record.
	TtlAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlAmt,omitempty" json:"ttl_amt" `

	// Set of elements used to provide details on the tax period and amount.
	Dtls []*TaxRecordDetails2 `xml:"Dtls,omitempty" json:"dtls" `
}

func (t *TaxAmount2) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxAmount2) SetTaxblBaseAmt(value, currency string) {
	t.TaxblBaseAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxAmount2) SetTtlAmt(value, currency string) {
	t.TtlAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxAmount2) AddDtls() *TaxRecordDetails2 {
	newValue := new(TaxRecordDetails2)
	t.Dtls = append(t.Dtls, newValue)
	return newValue
}

// Specifies the amount with a specific type.
type TaxAmountAndType1 struct {
	XMLName xml.Name

	// Specifies the type of the amount.
	Tp *TaxAmountType1Choice `xml:"Tp,omitempty" json:"tp" `

	// Amount of money, which has been typed.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `
}

func (t *TaxAmountAndType1) AddTp() *TaxAmountType1Choice {
	t.Tp = new(TaxAmountType1Choice)
	return t.Tp
}

func (t *TaxAmountAndType1) SetAmt(value, currency string) {
	t.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Specifies the amount type.
type TaxAmountType1Choice struct {
	XMLName xml.Name

	// Specifies the amount type, in a coded form.
	Cd *ExternalTaxAmountType1Code `xml:"Cd" json:"cd" `

	// Specifies the amount type, in a free-text form.
	Prtry *Max35Text `xml:"Prtry" json:"prtry" `
}

func (t *TaxAmountType1Choice) SetCd(value string) {
	t.Cd = (*ExternalTaxAmountType1Code)(&value)
}

func (t *TaxAmountType1Choice) SetPrtry(value string) {
	t.Prtry = (*Max35Text)(&value)
}

// Details of the authorised tax paying party.
type TaxAuthorisation1 struct {
	XMLName xml.Name

	// Title or position of debtor or the debtor's authorised representative.
	Titl *Max35Text `xml:"Titl,omitempty" json:"titl" `

	// Name of the debtor or the debtor's authorised representative.
	Nm *Max140Text `xml:"Nm,omitempty" json:"nm" `
}

func (t *TaxAuthorisation1) SetTitl(value string) {
	t.Titl = (*Max35Text)(&value)
}

func (t *TaxAuthorisation1) SetNm(value string) {
	t.Nm = (*Max140Text)(&value)
}

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type TaxCharges2 struct {
	XMLName xml.Name

	// Unique reference to unambiguously identify the nature of the tax levied, such as Value Added Tax (VAT).
	Id *Max35Text `xml:"Id,omitempty" json:"id" `

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty" json:"rate" `

	// Amount of money resulting from the calculation of the tax.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty" json:"amt" `
}

func (t *TaxCharges2) SetId(value string) {
	t.Id = (*Max35Text)(&value)
}

func (t *TaxCharges2) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCharges2) SetAmt(value, currency string) {
	t.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation6 struct {
	XMLName xml.Name

	// Party on the credit side of the transaction to which the tax applies.
	Cdtr *TaxParty1 `xml:"Cdtr,omitempty" json:"cdtr" `

	// Set of elements used to identify the party on the debit side of the transaction to which the tax applies.
	Dbtr *TaxParty2 `xml:"Dbtr,omitempty" json:"dbtr" `

	// Territorial part of a country to which the tax payment is related.
	AdmstnZn *Max35Text `xml:"AdmstnZn,omitempty" json:"admstn_zn" `

	// Tax reference information that is specific to a taxing agency.
	RefNb *Max140Text `xml:"RefNb,omitempty" json:"ref_nb" `

	// Method used to indicate the underlying business or how the tax is paid.
	Mtd *Max35Text `xml:"Mtd,omitempty" json:"mtd" `

	// Total amount of money on which the tax is based.
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"ttl_taxbl_base_amt" `

	// Total amount of money as result of the calculation of the tax.
	TtlTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"ttl_tax_amt" `

	// Date by which tax is due.
	Dt *ISODate `xml:"Dt,omitempty" json:"dt" `

	// Sequential number of the tax report.
	SeqNb *Number `xml:"SeqNb,omitempty" json:"seq_nb" `

	// Record of tax details.
	Rcrd []*TaxRecord2 `xml:"Rcrd,omitempty" json:"rcrd" `
}

func (t *TaxInformation6) AddCdtr() *TaxParty1 {
	t.Cdtr = new(TaxParty1)
	return t.Cdtr
}

func (t *TaxInformation6) AddDbtr() *TaxParty2 {
	t.Dbtr = new(TaxParty2)
	return t.Dbtr
}

func (t *TaxInformation6) SetAdmstnZn(value string) {
	t.AdmstnZn = (*Max35Text)(&value)
}

func (t *TaxInformation6) SetRefNb(value string) {
	t.RefNb = (*Max140Text)(&value)
}

func (t *TaxInformation6) SetMtd(value string) {
	t.Mtd = (*Max35Text)(&value)
}

func (t *TaxInformation6) SetTtlTaxblBaseAmt(value, currency string) {
	t.TtlTaxblBaseAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation6) SetTtlTaxAmt(value, currency string) {
	t.TtlTaxAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation6) SetDt(value string) {
	t.Dt = (*ISODate)(&value)
}

func (t *TaxInformation6) SetSeqNb(value string) {
	t.SeqNb = (*Number)(&value)
}

func (t *TaxInformation6) AddRcrd() *TaxRecord2 {
	newValue := new(TaxRecord2)
	t.Rcrd = append(t.Rcrd, newValue)
	return newValue
}

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation7 struct {
	XMLName xml.Name

	// Party on the credit side of the transaction to which the tax applies.
	Cdtr *TaxParty1 `xml:"Cdtr,omitempty" json:"cdtr" `

	// Identifies the party on the debit side of the transaction to which the tax applies.
	Dbtr *TaxParty2 `xml:"Dbtr,omitempty" json:"dbtr" `

	// Ultimate party that owes an amount of money to the (ultimate) creditor, in this case, to the taxing authority.
	UltmtDbtr *TaxParty2 `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Territorial part of a country to which the tax payment is related.
	AdmstnZone *Max35Text `xml:"AdmstnZone,omitempty" json:"admstn_zone" `

	// Tax reference information that is specific to a taxing agency.
	RefNb *Max140Text `xml:"RefNb,omitempty" json:"ref_nb" `

	// Method used to indicate the underlying business or how the tax is paid.
	Mtd *Max35Text `xml:"Mtd,omitempty" json:"mtd" `

	// Total amount of money on which the tax is based.
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"ttl_taxbl_base_amt" `

	// Total amount of money as result of the calculation of the tax.
	TtlTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"ttl_tax_amt" `

	// Date by which tax is due.
	Dt *ISODate `xml:"Dt,omitempty" json:"dt" `

	// Sequential number of the tax report.
	SeqNb *Number `xml:"SeqNb,omitempty" json:"seq_nb" `

	// Record of tax details.
	Rcrd []*TaxRecord2 `xml:"Rcrd,omitempty" json:"rcrd" `
}

func (t *TaxInformation7) AddCdtr() *TaxParty1 {
	t.Cdtr = new(TaxParty1)
	return t.Cdtr
}

func (t *TaxInformation7) AddDbtr() *TaxParty2 {
	t.Dbtr = new(TaxParty2)
	return t.Dbtr
}

func (t *TaxInformation7) AddUltmtDbtr() *TaxParty2 {
	t.UltmtDbtr = new(TaxParty2)
	return t.UltmtDbtr
}

func (t *TaxInformation7) SetAdmstnZone(value string) {
	t.AdmstnZone = (*Max35Text)(&value)
}

func (t *TaxInformation7) SetRefNb(value string) {
	t.RefNb = (*Max140Text)(&value)
}

func (t *TaxInformation7) SetMtd(value string) {
	t.Mtd = (*Max35Text)(&value)
}

func (t *TaxInformation7) SetTtlTaxblBaseAmt(value, currency string) {
	t.TtlTaxblBaseAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation7) SetTtlTaxAmt(value, currency string) {
	t.TtlTaxAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation7) SetDt(value string) {
	t.Dt = (*ISODate)(&value)
}

func (t *TaxInformation7) SetSeqNb(value string) {
	t.SeqNb = (*Number)(&value)
}

func (t *TaxInformation7) AddRcrd() *TaxRecord2 {
	newValue := new(TaxRecord2)
	t.Rcrd = append(t.Rcrd, newValue)
	return newValue
}

// Details about tax paid, or to be paid, to the government in accordance with the law, including pre-defined parameters such as thresholds and type of account.
type TaxInformation8 struct {
	XMLName xml.Name

	// Party on the credit side of the transaction to which the tax applies.
	Cdtr *TaxParty1 `xml:"Cdtr,omitempty" json:"cdtr" `

	// Party on the debit side of the transaction to which the tax applies.
	Dbtr *TaxParty2 `xml:"Dbtr,omitempty" json:"dbtr" `

	// Territorial part of a country to which the tax payment is related.
	AdmstnZone *Max35Text `xml:"AdmstnZone,omitempty" json:"admstn_zone" `

	// Tax reference information that is specific to a taxing agency.
	RefNb *Max140Text `xml:"RefNb,omitempty" json:"ref_nb" `

	// Method used to indicate the underlying business or how the tax is paid.
	Mtd *Max35Text `xml:"Mtd,omitempty" json:"mtd" `

	// Total amount of money on which the tax is based.
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxblBaseAmt,omitempty" json:"ttl_taxbl_base_amt" `

	// Total amount of money as result of the calculation of the tax.
	TtlTaxAmt *ActiveOrHistoricCurrencyAndAmount `xml:"TtlTaxAmt,omitempty" json:"ttl_tax_amt" `

	// Date by which tax is due.
	Dt *ISODate `xml:"Dt,omitempty" json:"dt" `

	// Sequential number of the tax report.
	SeqNb *Number `xml:"SeqNb,omitempty" json:"seq_nb" `

	// Record of tax details.
	Rcrd []*TaxRecord2 `xml:"Rcrd,omitempty" json:"rcrd" `
}

func (t *TaxInformation8) AddCdtr() *TaxParty1 {
	t.Cdtr = new(TaxParty1)
	return t.Cdtr
}

func (t *TaxInformation8) AddDbtr() *TaxParty2 {
	t.Dbtr = new(TaxParty2)
	return t.Dbtr
}

func (t *TaxInformation8) SetAdmstnZone(value string) {
	t.AdmstnZone = (*Max35Text)(&value)
}

func (t *TaxInformation8) SetRefNb(value string) {
	t.RefNb = (*Max140Text)(&value)
}

func (t *TaxInformation8) SetMtd(value string) {
	t.Mtd = (*Max35Text)(&value)
}

func (t *TaxInformation8) SetTtlTaxblBaseAmt(value, currency string) {
	t.TtlTaxblBaseAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation8) SetTtlTaxAmt(value, currency string) {
	t.TtlTaxAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (t *TaxInformation8) SetDt(value string) {
	t.Dt = (*ISODate)(&value)
}

func (t *TaxInformation8) SetSeqNb(value string) {
	t.SeqNb = (*Number)(&value)
}

func (t *TaxInformation8) AddRcrd() *TaxRecord2 {
	newValue := new(TaxRecord2)
	t.Rcrd = append(t.Rcrd, newValue)
	return newValue
}

// Details about the entity involved in the tax paid or to be paid.
type TaxParty1 struct {
	XMLName xml.Name

	// Tax identification number of the creditor.
	TaxId *Max35Text `xml:"TaxId,omitempty" json:"tax_id" `

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegnId *Max35Text `xml:"RegnId,omitempty" json:"regn_id" `

	// Type of tax payer.
	TaxTp *Max35Text `xml:"TaxTp,omitempty" json:"tax_tp" `
}

func (t *TaxParty1) SetTaxId(value string) {
	t.TaxId = (*Max35Text)(&value)
}

func (t *TaxParty1) SetRegnId(value string) {
	t.RegnId = (*Max35Text)(&value)
}

func (t *TaxParty1) SetTaxTp(value string) {
	t.TaxTp = (*Max35Text)(&value)
}

// Details about the entity involved in the tax paid or to be paid.
type TaxParty2 struct {
	XMLName xml.Name

	// Tax identification number of the debtor.
	TaxId *Max35Text `xml:"TaxId,omitempty" json:"tax_id" `

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	RegnId *Max35Text `xml:"RegnId,omitempty" json:"regn_id" `

	// Type of tax payer.
	TaxTp *Max35Text `xml:"TaxTp,omitempty" json:"tax_tp" `

	// Details of the authorised tax paying party.
	Authstn *TaxAuthorisation1 `xml:"Authstn,omitempty" json:"authstn" `
}

func (t *TaxParty2) SetTaxId(value string) {
	t.TaxId = (*Max35Text)(&value)
}

func (t *TaxParty2) SetRegnId(value string) {
	t.RegnId = (*Max35Text)(&value)
}

func (t *TaxParty2) SetTaxTp(value string) {
	t.TaxTp = (*Max35Text)(&value)
}

func (t *TaxParty2) AddAuthstn() *TaxAuthorisation1 {
	t.Authstn = new(TaxAuthorisation1)
	return t.Authstn
}

// Period of time details related to the tax payment.
type TaxPeriod2 struct {
	XMLName xml.Name

	// Year related to the tax payment.
	Yr *ISODate `xml:"Yr,omitempty" json:"yr" `

	// Identification of the period related to the tax payment.
	Tp *TaxRecordPeriod1Code `xml:"Tp,omitempty" json:"tp" `

	// Range of time between a start date and an end date for which the tax report is provided.
	FrToDt *DatePeriod2 `xml:"FrToDt,omitempty" json:"fr_to_dt" `
}

func (t *TaxPeriod2) SetYr(value string) {
	t.Yr = (*ISODate)(&value)
}

func (t *TaxPeriod2) SetTp(value string) {
	t.Tp = (*TaxRecordPeriod1Code)(&value)
}

func (t *TaxPeriod2) AddFrToDt() *DatePeriod2 {
	t.FrToDt = new(DatePeriod2)
	return t.FrToDt
}

// Set of elements used to define the tax record.
type TaxRecord2 struct {
	XMLName xml.Name

	// High level code to identify the type of tax details.
	Tp *Max35Text `xml:"Tp,omitempty" json:"tp" `

	// Specifies the tax code as published by the tax authority.
	Ctgy *Max35Text `xml:"Ctgy,omitempty" json:"ctgy" `

	// Provides further details of the category tax code.
	CtgyDtls *Max35Text `xml:"CtgyDtls,omitempty" json:"ctgy_dtls" `

	// Code provided by local authority to identify the status of the party that has drawn up the settlement document.
	DbtrSts *Max35Text `xml:"DbtrSts,omitempty" json:"dbtr_sts" `

	// Identification number of the tax report as assigned by the taxing authority.
	CertId *Max35Text `xml:"CertId,omitempty" json:"cert_id" `

	// Identifies, in a coded form, on which template the tax report is to be provided.
	FrmsCd *Max35Text `xml:"FrmsCd,omitempty" json:"frms_cd" `

	// Set of elements used to provide details on the period of time related to the tax payment.
	Prd *TaxPeriod2 `xml:"Prd,omitempty" json:"prd" `

	// Set of elements used to provide information on the amount of the tax record.
	TaxAmt *TaxAmount2 `xml:"TaxAmt,omitempty" json:"tax_amt" `

	// Further details of the tax record.
	AddtlInf *Max140Text `xml:"AddtlInf,omitempty" json:"addtl_inf" `
}

func (t *TaxRecord2) SetTp(value string) {
	t.Tp = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetCtgy(value string) {
	t.Ctgy = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetCtgyDtls(value string) {
	t.CtgyDtls = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetDbtrSts(value string) {
	t.DbtrSts = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetCertId(value string) {
	t.CertId = (*Max35Text)(&value)
}

func (t *TaxRecord2) SetFrmsCd(value string) {
	t.FrmsCd = (*Max35Text)(&value)
}

func (t *TaxRecord2) AddPrd() *TaxPeriod2 {
	t.Prd = new(TaxPeriod2)
	return t.Prd
}

func (t *TaxRecord2) AddTaxAmt() *TaxAmount2 {
	t.TaxAmt = new(TaxAmount2)
	return t.TaxAmt
}

func (t *TaxRecord2) SetAddtlInf(value string) {
	t.AddtlInf = (*Max140Text)(&value)
}

// Provides information on the individual tax amount(s) per period of the tax record.
type TaxRecordDetails2 struct {
	XMLName xml.Name

	// Set of elements used to provide details on the period of time related to the tax payment.
	Prd *TaxPeriod2 `xml:"Prd,omitempty" json:"prd" `

	// Underlying tax amount related to the specified period.
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"Amt" json:"amt" `
}

func (t *TaxRecordDetails2) AddPrd() *TaxPeriod2 {
	t.Prd = new(TaxPeriod2)
	return t.Prd
}

func (t *TaxRecordDetails2) SetAmt(value, currency string) {
	t.Amt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

type TaxRecordPeriod1Code string

type TransactionIndividualStatus1Code string

// Provides further details on the parties specific to the individual transaction.
type TransactionParties7 struct {
	XMLName xml.Name

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltmtDbtr *Party40Choice `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr" `

	// Party that owes an amount of money to the (ultimate) creditor.
	Dbtr *Party40Choice `xml:"Dbtr" json:"dbtr" `

	// Party that initiates the payment.
	// Usage: This can be either the debtor or a party that initiates the credit transfer on behalf of the debtor.
	InitgPty *Party40Choice `xml:"InitgPty,omitempty" json:"initg_pty" `

	// Financial institution servicing an account for the debtor.
	DbtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"DbtrAgt,omitempty" json:"dbtr_agt" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt1,omitempty" json:"prvs_instg_agt_1" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt2,omitempty" json:"prvs_instg_agt_2" `

	// Agent immediately prior to the instructing agent.
	PrvsInstgAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"PrvsInstgAgt3,omitempty" json:"prvs_instg_agt_3" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than one intermediary agent is present, then IntermediaryAgent1 identifies the agent between the DebtorAgent and the IntermediaryAgent2.
	IntrmyAgt1 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt1,omitempty" json:"intrmy_agt_1" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If more than two intermediary agents are present, then IntermediaryAgent2 identifies the agent between the IntermediaryAgent1 and the IntermediaryAgent3.
	IntrmyAgt2 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt2,omitempty" json:"intrmy_agt_2" `

	// Agent between the debtor's agent and the creditor's agent.
	//
	// Usage: If IntermediaryAgent3 is present, then it identifies the agent between the IntermediaryAgent 2 and the CreditorAgent.
	IntrmyAgt3 *BranchAndFinancialInstitutionIdentification6 `xml:"IntrmyAgt3,omitempty" json:"intrmy_agt_3" `

	// Financial institution servicing an account for the creditor.
	CdtrAgt *BranchAndFinancialInstitutionIdentification6 `xml:"CdtrAgt,omitempty" json:"cdtr_agt" `

	// Party to which an amount of money is due.
	Cdtr *Party40Choice `xml:"Cdtr" json:"cdtr" `

	// Ultimate party to which an amount of money is due.
	UltmtCdtr *Party40Choice `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr" `
}

func (t *TransactionParties7) AddUltmtDbtr() *Party40Choice {
	t.UltmtDbtr = new(Party40Choice)
	return t.UltmtDbtr
}

func (t *TransactionParties7) AddDbtr() *Party40Choice {
	t.Dbtr = new(Party40Choice)
	return t.Dbtr
}

func (t *TransactionParties7) AddInitgPty() *Party40Choice {
	t.InitgPty = new(Party40Choice)
	return t.InitgPty
}

func (t *TransactionParties7) AddDbtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	t.DbtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return t.DbtrAgt
}

func (t *TransactionParties7) AddPrvsInstgAgt1() *BranchAndFinancialInstitutionIdentification6 {
	t.PrvsInstgAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return t.PrvsInstgAgt1
}

func (t *TransactionParties7) AddPrvsInstgAgt2() *BranchAndFinancialInstitutionIdentification6 {
	t.PrvsInstgAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return t.PrvsInstgAgt2
}

func (t *TransactionParties7) AddPrvsInstgAgt3() *BranchAndFinancialInstitutionIdentification6 {
	t.PrvsInstgAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return t.PrvsInstgAgt3
}

func (t *TransactionParties7) AddIntrmyAgt1() *BranchAndFinancialInstitutionIdentification6 {
	t.IntrmyAgt1 = new(BranchAndFinancialInstitutionIdentification6)
	return t.IntrmyAgt1
}

func (t *TransactionParties7) AddIntrmyAgt2() *BranchAndFinancialInstitutionIdentification6 {
	t.IntrmyAgt2 = new(BranchAndFinancialInstitutionIdentification6)
	return t.IntrmyAgt2
}

func (t *TransactionParties7) AddIntrmyAgt3() *BranchAndFinancialInstitutionIdentification6 {
	t.IntrmyAgt3 = new(BranchAndFinancialInstitutionIdentification6)
	return t.IntrmyAgt3
}

func (t *TransactionParties7) AddCdtrAgt() *BranchAndFinancialInstitutionIdentification6 {
	t.CdtrAgt = new(BranchAndFinancialInstitutionIdentification6)
	return t.CdtrAgt
}

func (t *TransactionParties7) AddCdtr() *Party40Choice {
	t.Cdtr = new(Party40Choice)
	return t.Cdtr
}

func (t *TransactionParties7) AddUltmtCdtr() *Party40Choice {
	t.UltmtCdtr = new(Party40Choice)
	return t.UltmtCdtr
}

type TrueFalseIndicator string

type UUIDv4Identifier string

// Specifies the details of incorrect information.
type UnableToApplyIncorrect1 struct {
	XMLName xml.Name

	// Specifies the missing information in a coded form.
	Cd *UnableToApplyIncorrectInformation4Code `xml:"Cd" json:"cd" `

	// Further details about the incorrect information.
	AddtlIncrrctInf *Max140Text `xml:"AddtlIncrrctInf,omitempty" json:"addtl_incrrct_inf" `
}

func (u *UnableToApplyIncorrect1) SetCd(value string) {
	u.Cd = (*UnableToApplyIncorrectInformation4Code)(&value)
}

func (u *UnableToApplyIncorrect1) SetAddtlIncrrctInf(value string) {
	u.AddtlIncrrctInf = (*Max140Text)(&value)
}

type UnableToApplyIncorrectInformation4Code string

// Specifies the details of missing or incorrect information or the complete set of available information.
type UnableToApplyJustification3Choice struct {
	XMLName xml.Name

	// Indicates whether or not all available information on the underlying payment instruction is requested.
	AnyInf *YesNoIndicator `xml:"AnyInf" json:"any_inf" `

	// Set of elements used to indicate which information is missing or incorrect.
	MssngOrIncrrctInf *MissingOrIncorrectInformation3 `xml:"MssngOrIncrrctInf" json:"mssng_or_incrrct_inf" `

	// Indicates whether or not the referred item is a possible duplicate of a previous instruction or entry.
	PssblDplctInstr *TrueFalseIndicator `xml:"PssblDplctInstr" json:"pssbl_dplct_instr" `
}

func (u *UnableToApplyJustification3Choice) SetAnyInf(value string) {
	u.AnyInf = (*YesNoIndicator)(&value)
}

func (u *UnableToApplyJustification3Choice) AddMssngOrIncrrctInf() *MissingOrIncorrectInformation3 {
	u.MssngOrIncrrctInf = new(MissingOrIncorrectInformation3)
	return u.MssngOrIncrrctInf
}

func (u *UnableToApplyJustification3Choice) SetPssblDplctInstr(value string) {
	u.PssblDplctInstr = (*TrueFalseIndicator)(&value)
}

// Specifies the details of missing information.
type UnableToApplyMissing1 struct {
	XMLName xml.Name

	// Specifies the missing information in a coded form.
	Cd *UnableToApplyMissingInformation3Code `xml:"Cd" json:"cd" `

	// Further details about the missing information.
	AddtlMssngInf *Max140Text `xml:"AddtlMssngInf,omitempty" json:"addtl_mssng_inf" `
}

func (u *UnableToApplyMissing1) SetCd(value string) {
	u.Cd = (*UnableToApplyMissingInformation3Code)(&value)
}

func (u *UnableToApplyMissing1) SetAddtlMssngInf(value string) {
	u.AddtlMssngInf = (*Max140Text)(&value)
}

type UnableToApplyMissingInformation3Code string

// Unique identification, as assigned by the first instructing agent, to unambiguously identify the group of transactions.
type UnderlyingGroupInformation1 struct {
	XMLName xml.Name

	// Point to point reference, as assigned by the original instructing party, to unambiguously identify the original message.
	OrgnlMsgId *Max35Text `xml:"OrgnlMsgId" json:"orgnl_msg_id" `

	// Specifies the original message name identifier to which the message refers.
	OrgnlMsgNmId *Max35Text `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id" `

	// Date and time at which the original message was created.
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm" `

	// Original channel used for the delivery of the message, to allow the receiver of the request to locate the payment with greater ease.
	OrgnlMsgDlvryChanl *Max35Text `xml:"OrgnlMsgDlvryChanl,omitempty" json:"orgnl_msg_dlvry_chanl" `
}

func (u *UnderlyingGroupInformation1) SetOrgnlMsgId(value string) {
	u.OrgnlMsgId = (*Max35Text)(&value)
}

func (u *UnderlyingGroupInformation1) SetOrgnlMsgNmId(value string) {
	u.OrgnlMsgNmId = (*Max35Text)(&value)
}

func (u *UnderlyingGroupInformation1) SetOrgnlCreDtTm(value string) {
	u.OrgnlCreDtTm = (*ISODateTime)(&value)
}

func (u *UnderlyingGroupInformation1) SetOrgnlMsgDlvryChanl(value string) {
	u.OrgnlMsgDlvryChanl = (*Max35Text)(&value)
}

// Provides further details on the original payment initiation, to which the investigation message refers.
type UnderlyingPaymentInstruction5 struct {
	XMLName xml.Name

	// Set of elements used to provide information on the original message.
	OrgnlGrpInf *UnderlyingGroupInformation1 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original sending party, to unambiguously identify the original payment information group.
	OrgnlPmtInfId *Max35Text `xml:"OrgnlPmtInfId,omitempty" json:"orgnl_pmt_inf_id" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `

	// Amount of money, as provided in the original transaction, to be moved between the debtor and the creditor, before deduction of charges, expressed in the currency, as ordered by the original initiating party.
	OrgnlInstdAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlInstdAmt" json:"orgnl_instd_amt" `

	// Date at which the initiating party requests the clearing agent to process the payment.
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	ReqdExctnDt *DateAndDateTime2Choice `xml:"ReqdExctnDt,omitempty" json:"reqd_exctn_dt" `

	// Date at which the creditor requests the amount of money to be collected from the debtor.
	ReqdColltnDt *ISODate `xml:"ReqdColltnDt,omitempty" json:"reqd_colltn_dt" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `
}

func (u *UnderlyingPaymentInstruction5) AddOrgnlGrpInf() *UnderlyingGroupInformation1 {
	u.OrgnlGrpInf = new(UnderlyingGroupInformation1)
	return u.OrgnlGrpInf
}

func (u *UnderlyingPaymentInstruction5) SetOrgnlPmtInfId(value string) {
	u.OrgnlPmtInfId = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction5) SetOrgnlInstrId(value string) {
	u.OrgnlInstrId = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction5) SetOrgnlEndToEndId(value string) {
	u.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentInstruction5) SetOrgnlUETR(value string) {
	u.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (u *UnderlyingPaymentInstruction5) SetOrgnlInstdAmt(value, currency string) {
	u.OrgnlInstdAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnderlyingPaymentInstruction5) AddReqdExctnDt() *DateAndDateTime2Choice {
	u.ReqdExctnDt = new(DateAndDateTime2Choice)
	return u.ReqdExctnDt
}

func (u *UnderlyingPaymentInstruction5) SetReqdColltnDt(value string) {
	u.ReqdColltnDt = (*ISODate)(&value)
}

func (u *UnderlyingPaymentInstruction5) AddOrgnlTxRef() *OriginalTransactionReference28 {
	u.OrgnlTxRef = new(OriginalTransactionReference28)
	return u.OrgnlTxRef
}

// Provides further details on the original payment transaction, to which the investigation message refers.
type UnderlyingPaymentTransaction4 struct {
	XMLName xml.Name

	// Set of elements used to provide information on the original message.
	OrgnlGrpInf *UnderlyingGroupInformation1 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the original instructing party for the original instructed party, to unambiguously identify the original instruction.
	OrgnlInstrId *Max35Text `xml:"OrgnlInstrId,omitempty" json:"orgnl_instr_id" `

	// Unique identification, as assigned by the original initiating party, to unambiguously identify the original transaction.
	OrgnlEndToEndId *Max35Text `xml:"OrgnlEndToEndId,omitempty" json:"orgnl_end_to_end_id" `

	// Unique identification, as assigned by the original first instructing agent, to unambiguously identify the transaction.
	OrgnlTxId *Max35Text `xml:"OrgnlTxId,omitempty" json:"orgnl_tx_id" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `

	// Amount of money moved between the instructing agent and the instructed agent, as provided in the original instruction.
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount `xml:"OrgnlIntrBkSttlmAmt" json:"orgnl_intr_bk_sttlm_amt" `

	// Date, as provided in the original transaction, on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	OrgnlIntrBkSttlmDt *ISODate `xml:"OrgnlIntrBkSttlmDt" json:"orgnl_intr_bk_sttlm_dt" `

	// Key elements used to identify the original transaction that is being referred to.
	OrgnlTxRef *OriginalTransactionReference28 `xml:"OrgnlTxRef,omitempty" json:"orgnl_tx_ref" `
}

func (u *UnderlyingPaymentTransaction4) AddOrgnlGrpInf() *UnderlyingGroupInformation1 {
	u.OrgnlGrpInf = new(UnderlyingGroupInformation1)
	return u.OrgnlGrpInf
}

func (u *UnderlyingPaymentTransaction4) SetOrgnlInstrId(value string) {
	u.OrgnlInstrId = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentTransaction4) SetOrgnlEndToEndId(value string) {
	u.OrgnlEndToEndId = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentTransaction4) SetOrgnlTxId(value string) {
	u.OrgnlTxId = (*Max35Text)(&value)
}

func (u *UnderlyingPaymentTransaction4) SetOrgnlUETR(value string) {
	u.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

func (u *UnderlyingPaymentTransaction4) SetOrgnlIntrBkSttlmAmt(value, currency string) {
	u.OrgnlIntrBkSttlmAmt = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (u *UnderlyingPaymentTransaction4) SetOrgnlIntrBkSttlmDt(value string) {
	u.OrgnlIntrBkSttlmDt = (*ISODate)(&value)
}

func (u *UnderlyingPaymentTransaction4) AddOrgnlTxRef() *OriginalTransactionReference28 {
	u.OrgnlTxRef = new(OriginalTransactionReference28)
	return u.OrgnlTxRef
}

// Provides further details on the underlying statement entry, to which the investigation message refers.
type UnderlyingStatementEntry3 struct {
	XMLName xml.Name

	// Set of elements used to provide information on the original message.
	OrgnlGrpInf *OriginalGroupInformation29 `xml:"OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `

	// Unique identification, as assigned by the account servicer, to unambiguously identify the original statement.
	OrgnlStmtId *Max35Text `xml:"OrgnlStmtId,omitempty" json:"orgnl_stmt_id" `

	// Original unique identification, as assigned by the account servicer, to unambiguously identify the original entry.
	OrgnlNtryId *Max35Text `xml:"OrgnlNtryId,omitempty" json:"orgnl_ntry_id" `

	// Universally unique identifier to provide the original end-to-end reference of a payment transaction.
	OrgnlUETR *UUIDv4Identifier `xml:"OrgnlUETR,omitempty" json:"orgnl_u_e_t_r" `
}

func (u *UnderlyingStatementEntry3) AddOrgnlGrpInf() *OriginalGroupInformation29 {
	u.OrgnlGrpInf = new(OriginalGroupInformation29)
	return u.OrgnlGrpInf
}

func (u *UnderlyingStatementEntry3) SetOrgnlStmtId(value string) {
	u.OrgnlStmtId = (*Max35Text)(&value)
}

func (u *UnderlyingStatementEntry3) SetOrgnlNtryId(value string) {
	u.OrgnlNtryId = (*Max35Text)(&value)
}

func (u *UnderlyingStatementEntry3) SetOrgnlUETR(value string) {
	u.OrgnlUETR = (*UUIDv4Identifier)(&value)
}

// Identifies the underlying (group of) transaction(s) to which the resolution of investigation applies.
type UnderlyingTransaction22 struct {
	XMLName xml.Name

	// Provides information on the original cancellation message, to which the resolution refers.
	OrgnlGrpInfAndSts *OriginalGroupHeader14 `xml:"OrgnlGrpInfAndSts,omitempty" json:"orgnl_grp_inf_and_sts" `

	// Provides information on the original (group of) transactions, to which the cancellation status refers.
	OrgnlPmtInfAndSts []*OriginalPaymentInstruction30 `xml:"OrgnlPmtInfAndSts,omitempty" json:"orgnl_pmt_inf_and_sts" `

	// Provides details on the original transactions to which the cancellation request message refers.
	TxInfAndSts []*PaymentTransaction102 `xml:"TxInfAndSts,omitempty" json:"tx_inf_and_sts" `
}

func (u *UnderlyingTransaction22) AddOrgnlGrpInfAndSts() *OriginalGroupHeader14 {
	u.OrgnlGrpInfAndSts = new(OriginalGroupHeader14)
	return u.OrgnlGrpInfAndSts
}

func (u *UnderlyingTransaction22) AddOrgnlPmtInfAndSts() *OriginalPaymentInstruction30 {
	newValue := new(OriginalPaymentInstruction30)
	u.OrgnlPmtInfAndSts = append(u.OrgnlPmtInfAndSts, newValue)
	return newValue
}

func (u *UnderlyingTransaction22) AddTxInfAndSts() *PaymentTransaction102 {
	newValue := new(PaymentTransaction102)
	u.TxInfAndSts = append(u.TxInfAndSts, newValue)
	return newValue
}

// Identifies the underlying (group of) transaction(s) to which the investigation applies.
type UnderlyingTransaction23 struct {
	XMLName xml.Name

	// Provides information on the original message, to which the cancellation refers.
	OrgnlGrpInfAndCxl *OriginalGroupHeader15 `xml:"OrgnlGrpInfAndCxl,omitempty" json:"orgnl_grp_inf_and_cxl" `

	// Provides information on the original transactions to which the cancellation request message refers.
	TxInf []*PaymentTransaction106 `xml:"TxInf,omitempty" json:"tx_inf" `
}

func (u *UnderlyingTransaction23) AddOrgnlGrpInfAndCxl() *OriginalGroupHeader15 {
	u.OrgnlGrpInfAndCxl = new(OriginalGroupHeader15)
	return u.OrgnlGrpInfAndCxl
}

func (u *UnderlyingTransaction23) AddTxInf() *PaymentTransaction106 {
	newValue := new(PaymentTransaction106)
	u.TxInf = append(u.TxInf, newValue)
	return newValue
}

// Provides details of the underlying transaction, on which the investigation is processed.
type UnderlyingTransaction5Choice struct {
	XMLName xml.Name

	// Set of elements used to reference the details of the original payment initiation.
	Initn *UnderlyingPaymentInstruction5 `xml:"Initn" json:"initn" `

	// Set of elements used to reference the details of the original interbank payment transaction.
	IntrBk *UnderlyingPaymentTransaction4 `xml:"IntrBk" json:"intr_bk" `

	// Reference details on the underlying statement cash entry.
	StmtNtry *UnderlyingStatementEntry3 `xml:"StmtNtry" json:"stmt_ntry" `
}

func (u *UnderlyingTransaction5Choice) AddInitn() *UnderlyingPaymentInstruction5 {
	u.Initn = new(UnderlyingPaymentInstruction5)
	return u.Initn
}

func (u *UnderlyingTransaction5Choice) AddIntrBk() *UnderlyingPaymentTransaction4 {
	u.IntrBk = new(UnderlyingPaymentTransaction4)
	return u.IntrBk
}

func (u *UnderlyingTransaction5Choice) AddStmtNtry() *UnderlyingStatementEntry3 {
	u.StmtNtry = new(UnderlyingStatementEntry3)
	return u.StmtNtry
}

type UnicodeChartsCode string

type YesNoIndicator string
