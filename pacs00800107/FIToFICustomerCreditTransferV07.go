package pacs00800107

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *FIToFICustomerCreditTransferV07 `xml:"FIToFICstmrCdtTrf" json:"body"`
}

func (d *Message) AddMessage() *FIToFICustomerCreditTransferV07 {
	d.Body = new(FIToFICustomerCreditTransferV07)
	return d.Body
}

// Scope
// The FinancialInstitutionToFinancialInstitutionCustomerCreditTransfer message is sent by the debtor agent to the creditor agent, directly or through other agents and/or a payment clearing and settlement system. It is used to move funds from a debtor account to a creditor.
// Usage
// The FIToFICustomerCreditTransfer message is exchanged between agents and can contain one or more customer credit transfer instructions.
// The FIToFICustomerCreditTransfer message does not allow for grouping: a CreditTransferTransactionInformation block must be present for each credit transfer transaction.
// The FIToFICustomerCreditTransfer message can be used in different ways:
// - If the instructing agent and the instructed agent wish to use their direct account relationship in the currency of the transfer then the message contains both the funds for the customer transfer(s) as well as the payment details;
// - If the instructing agent and the instructed agent have no direct account relationship in the currency of the transfer, or do not wish to use their account relationship, then other (reimbursement) agents will be involved to cover for the customer transfer(s). The FIToFICustomerCreditTransfer contains only the payment details and the instructing agent must cover the customer transfer by sending a FinancialInstitutionCreditTransfer to a reimbursement agent. This payment method is called the Cover method;
// - If more than two financial institutions are involved in the payment chain and if the FIToFICustomerCreditTransfer is sent from one financial institution to the next financial institution in the payment chain, then the payment method is called the Serial method.
// The FIToFICustomerCreditTransfer message can be used in domestic and cross-border scenarios.
type FIToFICustomerCreditTransferV07 struct {
	XMLName xml.Name
	
		// Set of characteristics shared by all individual transactions included in the message.
	GrpHdr *GroupHeader70 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07 GrpHdr" json:"grp_hdr" `
	
		// Set of elements providing information specific to the individual credit transfer(s).
	CdtTrfTxInf []*CreditTransferTransaction30 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07 CdtTrfTxInf" json:"cdt_trf_tx_inf" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.07 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (f *FIToFICustomerCreditTransferV07) AddGrpHdr() *GroupHeader70 {
	f.GrpHdr = new(GroupHeader70)
	return f.GrpHdr
}

func (f *FIToFICustomerCreditTransferV07) AddCdtTrfTxInf() *CreditTransferTransaction30 {
	newValue := new (CreditTransferTransaction30)
	f.CdtTrfTxInf = append(f.CdtTrfTxInf, newValue)
	return newValue
}

func (f *FIToFICustomerCreditTransferV07) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	f.SplmtryData = append(f.SplmtryData, newValue)
	return newValue
}

