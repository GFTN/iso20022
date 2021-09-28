package ibwf00100101

import "encoding/xml"

type UnicodeChartsCode string

// Identification of a person, an organisation or a financial institution.
type Party9Choice struct {
	XMLName xml.Name

	// Identification of a person or an organisation.
	OrgId *PartyIdentification42 `xml:"OrgId" json:"org_id" `

	// Identification of a financial institution.
	FIId *BranchAndFinancialInstitutionIdentification5 `xml:"FIId" json:"fiid" `
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

type PhoneNumber string
type Max2048Text string
type NamePrefix1Code string
type ISONormalisedDateTime string
type CopyDuplicate1Code string
type YesNoIndicator string
type BusinessMessagePriorityCode string
type SignatureEnvelope string

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

type AnyBICIdentifier string
type Authorisation1Code string
type Max128Text string

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

type ExternalOrganisationIdentification1Code string

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

func (d *DateAndPlaceOfBirth) SetPrvcOfBirth(value string) {
	d.PrvcOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth) SetCityOfBirth(value string) {
	d.CityOfBirth = (*Max35Text)(&value)
}

func (d *DateAndPlaceOfBirth) SetCtryOfBirth(value string) {
	d.CtryOfBirth = (*CountryCode)(&value)
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

type ExternalPersonIdentification1Code string

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
