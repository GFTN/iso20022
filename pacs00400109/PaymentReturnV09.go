package pacs00400109

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *PaymentReturnV09 `xml:"PmtRtr" json:"body"`
}

func (d *Message) AddMessage() *PaymentReturnV09 {
	d.Body = new(PaymentReturnV09)
	return d.Body
}

// Scope
// The PaymentReturn message is sent by an agent to the previous agent in the payment chain to undo a payment previously settled.
// Usage
// The PaymentReturn message is exchanged between agents to return funds after settlement of credit transfer instructions (that is FIToFICustomerCreditTransfer message and FinancialInstitutionCreditTransfer message) or direct debit instructions (FIToFICustomerDirectDebit message).
// The PaymentReturn message should not be used between agents and non-financial institution customers. Non-financial institution customers will be informed about a debit or a credit on their account(s) through a BankToCustomerDebitCreditNotification message ('notification') and/or BankToCustomerAccountReport/BankToCustomerStatement message ('statement').
// The PaymentReturn message can be used to return single instructions or multiple instructions from one or different files.
// The PaymentReturn message can be used in domestic and cross-border scenarios.
// The PaymentReturn message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
type PaymentReturnV09 struct {
	XMLName xml.Name
	
		// Set of characteristics shared by all individual transactions included in the message.
	GrpHdr *GroupHeader90 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09 GrpHdr" json:"grp_hdr" `
	
		// Information concerning the original group of transactions, to which the message refers.
	OrgnlGrpInf *OriginalGroupHeader18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09 OrgnlGrpInf,omitempty" json:"orgnl_grp_inf" `
	
		// Information concerning the original transactions, to which the return message refers.
	TxInf []*PaymentTransaction112 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09 TxInf,omitempty" json:"tx_inf" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.09 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (p *PaymentReturnV09) AddGrpHdr() *GroupHeader90 {
	p.GrpHdr = new(GroupHeader90)
	return p.GrpHdr
}

func (p *PaymentReturnV09) AddOrgnlGrpInf() *OriginalGroupHeader18 {
	p.OrgnlGrpInf = new(OriginalGroupHeader18)
	return p.OrgnlGrpInf
}

func (p *PaymentReturnV09) AddTxInf() *PaymentTransaction112 {
	newValue := new (PaymentTransaction112)
	p.TxInf = append(p.TxInf, newValue)
	return newValue
}

func (p *PaymentReturnV09) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	p.SplmtryData = append(p.SplmtryData, newValue)
	return newValue
}

