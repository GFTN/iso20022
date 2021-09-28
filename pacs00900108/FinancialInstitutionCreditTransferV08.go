package pacs00900108

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *FinancialInstitutionCreditTransferV08 `xml:"FICdtTrf" json:"body"`
}

func (d *Message) AddMessage() *FinancialInstitutionCreditTransferV08 {
	d.Body = new(FinancialInstitutionCreditTransferV08)
	return d.Body
}

// Scope
// The FinancialInstitutionCreditTransfer message is sent by a debtor financial institution to a creditor financial institution, directly or through other agents and/or a payment clearing and settlement system.
// It is used to move funds from a debtor account to a creditor, where both debtor and creditor are financial institutions.
// Usage
// The FinancialInstitutionCreditTransfer message is exchanged between agents and can contain one or more credit transfer instructions where debtor and creditor are both financial institutions.
// The FinancialInstitutionCreditTransfer message does not allow for grouping: a CreditTransferTransactionInformation block must be present for each credit transfer transaction.
// The FinancialInstitutionCreditTransfer message can be used in domestic and cross-border scenarios.
type FinancialInstitutionCreditTransferV08 struct {
	XMLName xml.Name
	
		// Set of characteristics shared by all individual transactions included in the message.
	GrpHdr *GroupHeader93 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 GrpHdr" json:"grp_hdr" `
	
		// Set of elements providing information specific to the individual credit transfer(s).
	CdtTrfTxInf []*CreditTransferTransaction36 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtTrfTxInf" json:"cdt_trf_tx_inf" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (f *FinancialInstitutionCreditTransferV08) AddGrpHdr() *GroupHeader93 {
	f.GrpHdr = new(GroupHeader93)
	return f.GrpHdr
}

func (f *FinancialInstitutionCreditTransferV08) AddCdtTrfTxInf() *CreditTransferTransaction36 {
	newValue := new (CreditTransferTransaction36)
	f.CdtTrfTxInf = append(f.CdtTrfTxInf, newValue)
	return newValue
}

func (f *FinancialInstitutionCreditTransferV08) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	f.SplmtryData = append(f.SplmtryData, newValue)
	return newValue
}

