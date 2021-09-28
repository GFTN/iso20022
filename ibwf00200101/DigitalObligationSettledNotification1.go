package ibwf00200101

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification3Choice struct {
	IBAN      *IBANIdentifier                   `xml:"IBAN" json:"iban"`
	BBAN      *BBANIdentifier                   `xml:"BBAN" json:"bban"`
	UPIC      *UPICIdentifier                   `xml:"UPIC" json:"upic"`
	PrtryAcct *SimpleIdentificationInformation2 `xml:"PrtryAcct" json:"prtry_acct"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AmendmentInformationDetails1 struct {
	OrgnlMndtId      *Max35Text                                    `xml:"OrgnlMndtId,omitempty" json:"orgnl_mndt_id"`
	OrgnlCdtrSchmeId *PartyIdentification8                         `xml:"OrgnlCdtrSchmeId,omitempty" json:"orgnl_cdtr_schme_id"`
	OrgnlCdtrAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlCdtrAgt,omitempty" json:"orgnl_cdtr_agt"`
	OrgnlCdtrAgtAcct *CashAccount7                                 `xml:"OrgnlCdtrAgtAcct,omitempty" json:"orgnl_cdtr_agt_acct"`
	OrgnlDbtr        *PartyIdentification8                         `xml:"OrgnlDbtr,omitempty" json:"orgnl_dbtr"`
	OrgnlDbtrAcct    *CashAccount7                                 `xml:"OrgnlDbtrAcct,omitempty" json:"orgnl_dbtr_acct"`
	OrgnlDbtrAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"OrgnlDbtrAgt,omitempty" json:"orgnl_dbtr_agt"`
	OrgnlDbtrAgtAcct *CashAccount7                                 `xml:"OrgnlDbtrAgtAcct,omitempty" json:"orgnl_dbtr_agt_acct"`
	OrgnlFnlColltnDt *ISODate                                      `xml:"OrgnlFnlColltnDt,omitempty" json:"orgnl_fnl_colltn_dt"`
	OrgnlFrqcy       *Frequency1Code                               `xml:"OrgnlFrqcy,omitempty" json:"orgnl_frqcy"`
}

type AmountType2Choice struct {
	InstdAmt *CurrencyAndAmount `xml:"InstdAmt" json:"instd_amt"`
	EqvtAmt  *EquivalentAmount  `xml:"EqvtAmt" json:"eqvt_amt"`
}

// Must match the pattern [a-zA-Z0-9]{1,30}
type BBANIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BEIIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification3 struct {
	FinInstnId *FinancialInstitutionIdentification5Choice `xml:"FinInstnId" json:"fin_instn_id"`
	BrnchId    *BranchData                                `xml:"BrnchId,omitempty" json:"brnch_id"`
}

type BranchData struct {
	Id      *Max35Text      `xml:"Id,omitempty" json:"id"`
	Nm      *Max35Text      `xml:"Nm,omitempty" json:"nm"`
	PstlAdr *PostalAddress1 `xml:"PstlAdr,omitempty" json:"pstl_adr"`
}

// Must match the pattern CH[0-9]{6,6}
type CHIPSUniversalIdentifier string

type CashAccount7 struct {
	Id  *AccountIdentification3Choice `xml:"Id" json:"id"`
	Tp  *CashAccountType2             `xml:"Tp,omitempty" json:"tp"`
	Ccy *CurrencyCode                 `xml:"Ccy,omitempty" json:"ccy"`
	Nm  *Max70Text                    `xml:"Nm,omitempty" json:"nm"`
}

type CashAccountType2 struct {
	Cd    *CashAccountType4Code `xml:"Cd,omitempty" json:"cd"`
	Prtry *Max35Text            `xml:"Prtry,omitempty" json:"prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

// May be one of ABE, ART, AVP, AZM, BAP, BEL, BOF, BRL, CAD, CAM, CBJ, CHP, DKC, RTP, EBA, ELS, ERP, XCT, HRK, HRM, HUF, LGS, LVL, MUP, NOC, PCH, PDS, PEG, PNS, PVE, SEC, SIT, SLB, SPG, SSK, TBF, TGT, TOP, FDW, BOJ, FEY, ZEN, DDK, AIP, BCC, BDS, BGN, BHS, BIS, BSP, EPM, EPN, FDA, GIS, INC, JOD, KPS, LKB, MEP, MRS, NAM, PTR, ROL, ROS, SCP, STG, THB, TIS, TTD, UIS, MOS, ZET, ZIS, CHI, COP
type CashClearingSystem3Code string

// May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification1Choice struct {
	ClrSysId *CashClearingSystem3Code `xml:"ClrSysId,omitempty" json:"clr_sys_id"`
	Prtry    *Max35Text               `xml:"Prtry,omitempty" json:"prtry"`
}

type ClearingSystemMemberIdentification3Choice struct {
	Id    *ExternalClearingSystemMemberCode `xml:"Id" json:"id"`
	Prtry *Max35Text                        `xml:"Prtry,omitempty" json:"prtry"`
}

type CreditorReferenceInformation1 struct {
	CdtrRefTp *CreditorReferenceType1 `xml:"CdtrRefTp,omitempty" json:"cdtr_ref_tp"`
	CdtrRef   *Max35Text              `xml:"CdtrRef,omitempty" json:"cdtr_ref"`
}

type CreditorReferenceType1 struct {
	Cd    *DocumentType3Code `xml:"Cd,omitempty" json:"cd"`
	Prtry *Max35Text         `xml:"Prtry,omitempty" json:"prtry"`
	Issr  *Max35Text         `xml:"Issr,omitempty" json:"issr"`
}

type CurrencyAndAmount struct {
	Value float64       `xml:",chardata"`
	Ccy   *CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type DigitalObligationSettledNotification1 struct {
	GrpHdr     *GroupHeader70            `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.002.001.01 GrpHdr" json:"grp_hdr"`
	SttlOblInf []*SettledObligationInfo1 `xml:"urn:iso:std:iso:20022:tech:xsd:ibwf.002.001.01 SttlOblInf" json:"sttl_obl_inf"`
}

type Document struct {
	DigOblSetNotif *DigitalObligationSettledNotification1 `xml:"DigOblSetNotif" json:"dig_obl_set_notif"`
}

type Message struct {
	XMLName        xml.Name                               `xml:"Message"`    /////
	XMLns          string                                 `xml:"xmlns,attr"` /////
	Head           *BusinessApplicationHeaderV01          `xml:"AppHdr" json:"head"`
	DigOblSetNotif *DigitalObligationSettledNotification1 `xml:"DigOblSetNotif,omitempty" json:"body"`
}

// May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP
type DocumentType2Code string

// Must match the pattern [0-9]{9,9}
type DunsIdentifier string

// Must match the pattern [0-9]{13,13}
type EANGLNIdentifier string

type EquivalentAmount struct {
	Amt      *CurrencyAndAmount `xml:"Amt" json:"amt"`
	CcyOfTrf *CurrencyCode      `xml:"CcyOfTrf" json:"ccy_of_trf"`
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
type ExternalClearingSystemMemberCode string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalLocalInstrument1Code string

// Must be at least 1 items long
type ExternalLocalInstrumentCode string

// Must be at least 1 items long
type ExternalServiceLevel1Code string
type FinancialInstitutionIdentification3 struct {
	BIC         *BICIdentifier                             `xml:"BIC,omitempty" json:"bic"`
	ClrSysMmbId *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId,omitempty" json:"clr_sys_mmb_id"`
	Nm          *Max70Text                                 `xml:"Nm,omitempty" json:"nm"`
	PstlAdr     *PostalAddress1                            `xml:"PstlAdr,omitempty" json:"pstl_adr"`
	PrtryId     *GenericIdentification3                    `xml:"PrtryId,omitempty" json:"prtry_id"`
}

type FinancialInstitutionIdentification5Choice struct {
	BIC         *BICIdentifier                             `xml:"BIC" json:"bic"`
	ClrSysMmbId *ClearingSystemMemberIdentification3Choice `xml:"ClrSysMmbId,omitempty" json:"clr_sys_mmb_id"`
	NmAndAdr    *NameAndAddress7                           `xml:"NmAndAdr" json:"nm_and_adr"`
	PrtryId     *GenericIdentification3                    `xml:"PrtryId" json:"prtry_id"`
	CmbndId     *FinancialInstitutionIdentification3       `xml:"CmbndId" json:"cmbnd_id"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA
type Frequency1Code string

type GenericIdentification4 struct {
	Id   Max35Text `xml:"Id" json:"id"`
	IdTp Max35Text `xml:"IdTp" json:"id_tp"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

// Must match the pattern [a-zA-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBANIdentifier string

// Must match the pattern [A-Z]{2,2}[B-DF-HJ-NP-TV-XZ0-9]{7,7}[0-9]{1,1}
type IBEIIdentifier string

type ISODate string

type ISODateTime string

type LocalInstrument1Choice struct {
	Cd    ExternalLocalInstrumentCode `xml:"Cd,omitempty" json:"cd"`
	Prtry Max35Text                   `xml:"Prtry,omitempty" json:"prtry"`
}

type MandateRelatedInformation1 struct {
	MndtId        *Max35Text                    `xml:"MndtId,omitempty" json:"mndt_id"`
	DtOfSgntr     *ISODate                      `xml:"DtOfSgntr,omitempty" json:"dt_of_sgntr"`
	AmdmntInd     bool                          `xml:"AmdmntInd,omitempty" json:"amdmnt_ind"`
	AmdmntInfDtls *AmendmentInformationDetails1 `xml:"AmdmntInfDtls,omitempty" json:"amdmnt_inf_dtls"`
	ElctrncSgntr  *Max1025Text                  `xml:"ElctrncSgntr,omitempty" json:"elctrnc_sgntr"`
	FrstColltnDt  *ISODate                      `xml:"FrstColltnDt,omitempty" json:"frst_colltn_dt"`
	FnlColltnDt   *ISODate                      `xml:"FnlColltnDt,omitempty" json:"fnl_colltn_dt"`
	Frqcy         *Frequency1Code               `xml:"Frqcy,omitempty" json:"frqcy"`
}

// Must be at least 1 items long
type Max1025Text string

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

type NameAndAddress7 struct {
	Nm      *Max70Text      `xml:"Nm"`
	PstlAdr *PostalAddress1 `xml:"PstlAdr"`
}

type OrganisationIdentification2 struct {
	BIC     *BICIdentifier            `xml:"BIC,omitempty" json:"bic"`
	IBEI    *IBEIIdentifier           `xml:"IBEI,omitempty" json:"ibei"`
	BEI     *BEIIdentifier            `xml:"BEI,omitempty" json:"bei"`
	EANGLN  *EANGLNIdentifier         `xml:"EANGLN,omitempty" json:"eangln"`
	USCHU   *CHIPSUniversalIdentifier `xml:"USCHU,omitempty" json:"uschu"`
	DUNS    *DunsIdentifier           `xml:"DUNS,omitempty" json:"duns"`
	BkPtyId *Max35Text                `xml:"BkPtyId,omitempty" json:"bk_pty_id"`
	TaxIdNb *Max35Text                `xml:"TaxIdNb,omitempty" json:"tax_id_nb"`
	PrtryId *GenericIdentification3   `xml:"PrtryId,omitempty" json:"prtry_id"`
}

type OriginalGroupInformation3 struct {
	OrgnlMsgId   *Max35Text   `xml:"OrgnlMsgId" json:"orgnl_msg_id"`
	OrgnlMsgNmId *Max35Text   `xml:"OrgnlMsgNmId" json:"orgnl_msg_nm_id"`
	OrgnlCreDtTm *ISODateTime `xml:"OrgnlCreDtTm,omitempty" json:"orgnl_cre_dt_tm"`
}

type OriginalTransactionReference1 struct {
	IntrBkSttlmAmt *CurrencyAndAmount                            `xml:"IntrBkSttlmAmt" json:"intr_bk_sttlm_amt"`
	Amt            *AmountType2Choice                            `xml:"Amt,omitempty" json:"amt"`
	IntrBkSttlmDt  *ISODate                                      `xml:"IntrBkSttlmDt,omitempty" json:"intr_bk_sttlm_dt"`
	ReqdExctnDt    *ISODate                                      `xml:"ReqdExctnDt,omitempty" json:"reqd_exctn_dt"`
	ReqdColltnDt   *ISODate                                      `xml:"ReqdColltnDt,omitempty" json:"reqd_colltn_dt"`
	CdtrSchmeId    *PartyIdentification8                         `xml:"CdtrSchmeId,omitempty" json:"cdtr_schme_id"`
	SttlmInf       *SettlementInformation3                       `xml:"SttlmInf,omitempty" json:"sttlm_inf"`
	PmtTpInf       *PaymentTypeInformation6                      `xml:"PmtTpInf,omitempty" json:"pmt_tp_inf"`
	PmtMtd         *PaymentMethod4Code                           `xml:"PmtMtd,omitempty" json:"pmt_mtd"`
	MndtRltdInf    *MandateRelatedInformation1                   `xml:"MndtRltdInf,omitempty" json:"mndt_rltd_inf"`
	RmtInf         *RemittanceInformation1                       `xml:"RmtInf,omitempty" json:"rmt_inf"`
	UltmtDbtr      *PartyIdentification8                         `xml:"UltmtDbtr,omitempty" json:"ultmt_dbtr"`
	Dbtr           *PartyIdentification8                         `xml:"Dbtr,omitempty" json:"dbtr"`
	DbtrAcct       *CashAccount7                                 `xml:"DbtrAcct,omitempty" json:"dbtr_acct"`
	DbtrAgt        *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt,omitempty" json:"dbtr_agt"`
	DbtrAgtAcct    *CashAccount7                                 `xml:"DbtrAgtAcct,omitempty" json:"dbtr_agt_acct"`
	CdtrAgt        *BranchAndFinancialInstitutionIdentification3 `xml:"CdtrAgt,omitempty" json:"cdtr_agt"`
	CdtrAgtAcct    *CashAccount7                                 `xml:"CdtrAgtAcct,omitempty" json:"cdtr_agt_acct"`
	Cdtr           *PartyIdentification8                         `xml:"Cdtr,omitempty" json:"cdtr"`
	CdtrAcct       *CashAccount7                                 `xml:"CdtrAcct,omitempty" json:"cdtr_acct"`
	UltmtCdtr      *PartyIdentification8                         `xml:"UltmtCdtr,omitempty" json:"ultmt_cdtr"`
}

type Party2Choice struct {
	OrgId  *OrganisationIdentification2 `xml:"OrgId" json:"org_id"`
	PrvtId []*PersonIdentification3     `xml:"PrvtId" json:"prvt_id"`
}

type PartyIdentification8 struct {
	Nm        *Max70Text      `xml:"Nm,omitempty" json:"nm"`
	PstlAdr   *PostalAddress1 `xml:"PstlAdr,omitempty" json:"pstl_adr"`
	Id        *Party2Choice   `xml:"Id,omitempty" json:"id"`
	CtryOfRes *CountryCode    `xml:"CtryOfRes,omitempty" json:"ctry_of_res"`
}

// May be one of CORT, SALA, TREA, CASH, DIVI, GOVT, INTE, LOAN, PENS, SECU, SSBE, SUPP, TAXS, TRAD, VATX, HEDG, INTC, WHLD
type PaymentCategoryPurpose1Code string

// May be one of CHK, TRF, DD, TRA
type PaymentMethod4Code string

type PaymentTransactionInformation2 struct {
	NtfId           *Max35Text                     `xml:"NtfId" json:"ntf_id"`
	OrgnlInstrId    *Max35Text                     `xml:"OrgnlInstrId" json:"orgnl_instr_id"`
	OrgnlEndToEndId *Max35Text                     `xml:"OrgnlEndToEndId" json:"orgnl_end_to_end_id"`
	OrgnlTxId       *Max35Text                     `xml:"OrgnlTxId" json:"orgnl_tx_id"`
	ActSttldAmt     *CurrencyAndAmount             `xml:"ActSttldAmt" json:"act_sttld_amt"`
	OrgnlTxRef      *OriginalTransactionReference1 `xml:"OrgnlTxRef" json:"orgnl_tx_ref"`
}

type PaymentTypeInformation6 struct {
	InstrPrty *Priority2Code               `xml:"InstrPrty,omitempty" json:"instr_prty"`
	SvcLvl    *ServiceLevel2Choice         `xml:"SvcLvl,omitempty" json:"svc_lvl"`
	ClrChanl  *ClearingChannel2Code        `xml:"ClrChanl,omitempty" json:"clr_chanl"`
	LclInstrm *LocalInstrument1Choice      `xml:"LclInstrm,omitempty" json:"lcl_instrm"`
	SeqTp     *SequenceType1Code           `xml:"SeqTp,omitempty" json:"seq_tp"`
	CtgyPurp  *PaymentCategoryPurpose1Code `xml:"CtgyPurp,omitempty" json:"ctgy_purp"`
}

type PersonIdentification3 struct {
	DrvrsLicNb      *Max35Text              `xml:"DrvrsLicNb" json:"drvrs_lic_nb"`
	CstmrNb         *Max35Text              `xml:"CstmrNb" json:"cstmr_nb"`
	SclSctyNb       *Max35Text              `xml:"SclSctyNb" json:"scl_scty_nb"`
	AlnRegnNb       *Max35Text              `xml:"AlnRegnNb" json:"aln_regn_nb"`
	PsptNb          *Max35Text              `xml:"PsptNb" json:"pspt_nb"`
	TaxIdNb         *Max35Text              `xml:"TaxIdNb" json:"tax_id_nb"`
	IdntyCardNb     *Max35Text              `xml:"IdntyCardNb" json:"idnty_card_nb"`
	MplyrIdNb       *Max35Text              `xml:"MplyrIdNb" json:"mplyr_id_nb"`
	DtAndPlcOfBirth *DateAndPlaceOfBirth    `xml:"DtAndPlcOfBirth" json:"dt_and_plc_of_birth"`
	OthrId          *GenericIdentification4 `xml:"OthrId" json:"othr_id"`
	Issr            *Max35Text              `xml:"Issr,omitempty" json:"issr"`
}

type PostalAddress1 struct {
	AdrTp       *AddressType2Code `xml:"AdrTp,omitempty" json:"adr_tp"`
	AdrLine     []*Max70Text      `xml:"AdrLine,omitempty" json:"adr_line"`
	StrtNm      *Max70Text        `xml:"StrtNm,omitempty" json:"strt_nm"`
	BldgNb      *Max16Text        `xml:"BldgNb,omitempty" json:"bldg_nb"`
	PstCd       *Max16Text        `xml:"PstCd,omitempty" json:"pst_cd"`
	TwnNm       *Max35Text        `xml:"TwnNm,omitempty" json:"twn_nm"`
	CtrySubDvsn *Max35Text        `xml:"CtrySubDvsn,omitempty" json:"ctry_sub_dvsn"`
	Ctry        *CountryCode      `xml:"Ctry" json:"ctry"`
}

// May be one of HIGH, NORM
type Priority2Code string

type ReferredDocumentAmount1Choice struct {
	DuePyblAmt   *CurrencyAndAmount `xml:"DuePyblAmt" json:"due_pybl_amt"`
	DscntApldAmt *CurrencyAndAmount `xml:"DscntApldAmt" json:"dscnt_apld_amt"`
	RmtdAmt      *CurrencyAndAmount `xml:"RmtdAmt" json:"rmtd_amt"`
	CdtNoteAmt   *CurrencyAndAmount `xml:"CdtNoteAmt" json:"cdt_note_amt"`
	TaxAmt       *CurrencyAndAmount `xml:"TaxAmt" json:"tax_amt"`
}

type ReferredDocumentInformation1 struct {
	RfrdDocTp *ReferredDocumentType1 `xml:"RfrdDocTp,omitempty" json:"rfrd_doc_tp"`
	RfrdDocNb *Max35Text             `xml:"RfrdDocNb,omitempty" json:"rfrd_doc_nb"`
}

type ReferredDocumentType1 struct {
	Cd    *DocumentType2Code `xml:"Cd,omitempty" json:"cd"`
	Prtry *Max35Text         `xml:"Prtry,omitempty" json:"prtry"`
	Issr  *Max35Text         `xml:"Issr,omitempty" json:"issr"`
}

type RemittanceInformation1 struct {
	Ustrd []*Max140Text                       `xml:"Ustrd,omitempty" json:"ustrd"`
	Strd  []*StructuredRemittanceInformation6 `xml:"Strd,omitempty" json:"strd"`
}

// May be one of FRST, RCUR, FNAL, OOFF
type SequenceType1Code string

// May be one of SEPA, SDVA, PRPT
type ServiceLevel1Code string

type ServiceLevel2Choice struct {
	Cd    *ServiceLevel1Code `xml:"Cd,omitempty" json:"cd"`
	Prtry *Max35Text         `xml:"Prtry,omitempty" json:"prtry"`
}

type SettledObligationInfo1 struct {
	OrgnlGrpInf *OriginalGroupInformation3        `xml:"OrgnlGrpInf" json:"orgnl_grp_inf"`
	TxInf       []*PaymentTransactionInformation2 `xml:"TxInf" json:"tx_inf"`
}

type SettlementInformation3 struct {
	SttlmMtd             *SettlementMethod1Code                        `xml:"SttlmMtd" json:"sttlm_mtd"`
	SttlmAcct            *CashAccount7                                 `xml:"SttlmAcct,omitempty" json:"sttlm_acct"`
	ClrSys               *ClearingSystemIdentification1Choice          `xml:"ClrSys,omitempty" json:"clr_sys"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"InstgRmbrsmntAgt,omitempty" json:"instg_rmbrsmnt_agt"`
	InstgRmbrsmntAgtAcct *CashAccount7                                 `xml:"InstgRmbrsmntAgtAcct,omitempty" json:"instg_rmbrsmnt_agt_acct"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification3 `xml:"InstdRmbrsmntAgt,omitempty" json:"instd_rmbrsmnt_agt"`
	InstdRmbrsmntAgtAcct *CashAccount7                                 `xml:"InstdRmbrsmntAgtAcct,omitempty" json:"instd_rmbrsmnt_agt_acct"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification3 `xml:"ThrdRmbrsmntAgt,omitempty" json:"thrd_rmbrsmnt_agt"`
	ThrdRmbrsmntAgtAcct  *CashAccount7                                 `xml:"ThrdRmbrsmntAgtAcct,omitempty" json:"thrd_rmbrsmnt_agt_acct"`
}

// May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type SimpleIdentificationInformation2 struct {
	Id *Max34Text `xml:"Id" json:"id"`
}

type StructuredRemittanceInformation6 struct {
	RfrdDocInf    *ReferredDocumentInformation1    `xml:"RfrdDocInf,omitempty" json:"rfrd_doc_inf"`
	RfrdDocRltdDt *ISODate                         `xml:"RfrdDocRltdDt,omitempty" json:"rfrd_doc_rltd_dt"`
	RfrdDocAmt    []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt,omitempty" json:"rfrd_doc_amt"`
	CdtrRefInf    *CreditorReferenceInformation1   `xml:"CdtrRefInf,omitempty" json:"cdtr_ref_inf"`
	Invcr         *PartyIdentification8            `xml:"Invcr,omitempty" json:"invcr"`
	Invcee        *PartyIdentification8            `xml:"Invcee,omitempty" json:"invcee"`
	AddtlRmtInf   *Max140Text                      `xml:"AddtlRmtInf,omitempty" json:"addtl_rmt_inf"`
}

// Must match the pattern [0-9]{8,17}
type UPICIdentifier string

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
