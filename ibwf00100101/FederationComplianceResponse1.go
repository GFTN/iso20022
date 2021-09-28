package ibwf00100101

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier             `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 IBAN,omitempty" json:"iban,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Othr,omitempty" json:"othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ,chardata" json:"value,omitempty"`
	Ccy   ActiveCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Ccy,attr" json:"ccy,omitempty"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId *FinancialInstitutionIdentification8 `xml:"FinInstnId,omitempty" json:"fin_instn_id,omitempty"`
	BrnchId    *BranchData2                         `xml:"BrnchId,omitempty" json:"brnch_id,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Id,omitempty" json:"id,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Nm,omitempty" json:"nm,omitempty"`
	PstlAdr *PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 PstlAdr,omitempty" json:"pstl_adr,omitempty"`
}

type CashAccount24 struct {
	Id  *AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Id,omitempty" json:"id,omitempty"`
	Tp  *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Tp,omitempty" json:"tp,omitempty"`
	Ccy ActiveOrHistoricCurrencyCode  `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Ccy,omitempty" json:"ccy,omitempty"`
	Nm  Max70Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Nm,omitempty" json:"nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

type CategoryPurpose1Choice struct {
	Cd    ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ClrSysId,omitempty" json:"clr_sys_id,omitempty"`
	MmbId    Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 MmbId,omitempty" json:"mmb_id,omitempty"`
}

type ComplianceResponse1 struct {
	InfSts InfoStatus1             `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InfSts,omitempty" json:"inf_sts,omitempty"`
	TxnSts TransactionStatus1      `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 TxnSts,omitempty" json:"txn_sts,omitempty"`
	DstInf DestinationInfo1        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 DstInf,omitempty" json:"dst_inf,omitempty"`
	PmtId  *PaymentIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 PmtId,omitempty" json:"pmt_id,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [0-9]{4}-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])
type DateString string

// Must match the pattern [0-9]{4}-(1[0-2]|0[1-9])-(3[01]|0[1-9]|[12][0-9])T(2[0-3]|[01][0-9]):([0-5][0-9]):([0-5][0-9])
type DateTimeString string

// Must be at least 1 items long
type DestinationInfo1 string

type Message struct {
	XMLName xml.Name                       `xml:"Message"`    /////
	XMLns   string                         `xml:"xmlns,attr"` /////
	Head    *BusinessApplicationHeaderV01  `xml:"AppHdr" json:"head"`
	Body    *FederationComplianceResponse1 `xml:"FedCompRes,omitempty" json:"body"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalCashClearingSystem1Code string

// Must be at least 1 items long
type ExternalCategoryPurpose1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalServiceLevel1Code string

type FederationComplianceResponse1 struct {
	GrpHdr  *GroupHeader70   `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 GrpHdr" json:"grp_hdr,omitempty"`
	ResBody []*ResponseBody1 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ResBody" json:"res_body,omitempty"`
}

type ResponseBody1 struct {
	Id     Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Id" json:"id,omitempty"`
	FedRes FederationResponse1 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 FedRes" json:"fed_res,omitempty"`
	CmpRes ComplianceResponse1 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 CmpRes" json:"cmp_res,omitempty"`
}

type FederationResponse1 struct {
	StrAdr Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 StrAdr,omitempty" json:"str_adr,omitempty"`
	AccId  Max140Text              `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 AccId,omitempty" json:"acc_id,omitempty"`
	MeTyp  MemoType1               `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 MeTyp,omitempty" json:"me_typ,omitempty"`
	Mem    Memo1                   `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Mem,omitempty" json:"mem,omitempty"`
	FedSts FederationStatus1       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 FedSts,omitempty" json:"fed_sts,omitempty"`
	PmtId  *PaymentIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 PmtId,omitempty" json:"pmt_id,omitempty"`
}

// May be one of ACTC, RJCT
type FederationStatus1 string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *BICFIIdentifier                     `xml:"BICFI,omitempty" json:"bicfi,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"ClrSysMmbId,omitempty" json:"clr_sys_mmb_id,omitempty"`
	Nm          Max140Text                           `xml:"Nm,omitempty" json:"nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"PstlAdr,omitempty" json:"pstl_adr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"Othr,omitempty" json:"othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Id,omitempty" json:"id,omitempty"`
	SchmeNm *AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 SchmeNm,omitempty" json:"schme_nm,omitempty"`
	Issr    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Issr,omitempty" json:"issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      *Max35Text                                `xml:"Id,omitempty" json:"id,omitempty"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty" json:"schme_nm,omitempty"`
	Issr    Max35Text                                 `xml:"Issr,omitempty" json:"issr,omitempty"`
}

type GroupHeader70 struct {
	MsgId             Max35Text                                     `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 MsgId,omitempty" json:"msg_id,omitempty"`
	CreDtTm           DateTimeString                                `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 CreDtTm,omitempty" json:"cre_dt_tm,omitempty"`
	BtchBookg         bool                                          `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 BtchBookg,omitempty" json:"btch_bookg,omitempty"`
	NbOfTxs           Max15NumericText                              `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 NbOfTxs,omitempty" json:"nb_of_txs,omitempty"`
	CtrlSum           float64                                       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 CtrlSum,omitempty" json:"ctrl_sum,omitempty"`
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 TtlIntrBkSttlmAmt,omitempty" json:"ttl_intr_bk_sttlm_amt,omitempty"`
	IntrBkSttlmDt     DateString                                    `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt,omitempty"`
	SttlmInf          *SettlementInstruction4                       `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 SttlmInf,omitempty" json:"sttlm_inf,omitempty"`
	PmtTpInf          *PaymentTypeInformation21                     `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 PmtTpInf,omitempty" json:"pmt_tp_inf,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstgAgt,omitempty" json:"instg_agt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstdAgt,omitempty" json:"instd_agt,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// May be one of ACTC, RJCT, PDNG
type InfoStatus1 string

type LocalInstrument2Choice struct {
	Cd    ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// Must be at least 1 items long
type Memo1 string

// May be one of TEST, ID, HASH
type MemoType1 string

type PaymentIdentification3 struct {
	InstrId    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstrId,omitempty" json:"instr_id,omitempty"`
	EndToEndId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 EndToEndId,omitempty" json:"end_to_end_id,omitempty"`
	TxId       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 TxId,omitempty" json:"tx_id,omitempty"`
	ClrSysRef  Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ClrSysRef,omitempty" json:"clr_sys_ref,omitempty"`
}

type PaymentTypeInformation21 struct {
	InstrPrty Priority2Code           `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstrPrty,omitempty" json:"instr_prty,omitempty"`
	ClrChanl  ClearingChannel2Code    `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ClrChanl,omitempty" json:"clr_chanl,omitempty"`
	SvcLvl    *ServiceLevel8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 SvcLvl,omitempty" json:"svc_lvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 LclInstrm,omitempty" json:"lcl_instrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 CtgyPurp,omitempty" json:"ctgy_purp,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 AdrTp,omitempty" json:"adr_tp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Dept,omitempty" json:"dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 SubDept,omitempty" json:"sub_dept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 StrtNm,omitempty" json:"strt_nm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 BldgNb,omitempty" json:"bldg_nb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 PstCd,omitempty" json:"pst_cd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 TwnNm,omitempty" json:"twn_nm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 CtrySubDvsn,omitempty" json:"ctry_sub_dvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Ctry,omitempty" json:"ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 AdrLine,omitempty" json:"adr_line,omitempty"`
}

// May be one of HIGH, NORM
type Priority2Code string

type ServiceLevel8Choice struct {
	Cd    ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Cd,omitempty" json:"cd,omitempty"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 Prtry,omitempty" json:"prtry,omitempty"`
}

type SettlementInstruction4 struct {
	SttlmMtd             SettlementMethod1Code                         `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 SttlmMtd,omitempty" json:"sttlm_mtd,omitempty"`
	SttlmAcct            *CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 SttlmAcct,omitempty" json:"sttlm_acct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ClrSys,omitempty" json:"clr_sys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstgRmbrsmntAgt,omitempty" json:"instg_rmbrsmnt_agt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstgRmbrsmntAgtAcct,omitempty" json:"instg_rmbrsmnt_agt_acct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstdRmbrsmntAgt,omitempty" json:"instd_rmbrsmnt_agt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 InstdRmbrsmntAgtAcct,omitempty" json:"instd_rmbrsmnt_agt_acct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ThrdRmbrsmntAgt,omitempty" json:"thrd_rmbrsmnt_agt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.001.001.01 ThrdRmbrsmntAgtAcct,omitempty" json:"thrd_rmbrsmnt_agt_acct,omitempty"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

// May be one of ACTC, RJCT, PDNG
type TransactionStatus1 string

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02")), nil
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
