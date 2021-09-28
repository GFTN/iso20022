package pacs00200109

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *FIToFIPaymentStatusReportV09 `xml:"FIToFIPmtStsRpt" json:"body"`
}

func (d *Message) AddMessage() *FIToFIPaymentStatusReportV09 {
	d.Body = new(FIToFIPaymentStatusReportV09)
	return d.Body
}

// Scope
// The FIToFIPaymentStatusReport message is sent by an instructed agent to the previous party in the payment chain. It is used to inform this party about the positive or negative status of an instruction (either single or file). It is also used to report on a pending instruction.
// Usage
// The FIToFIPaymentStatusReport message is exchanged between agents to provide status information about instructions previously sent. Its usage will always be governed by a bilateral agreement between the agents.
// The FIToFIPaymentStatusReport message can be used to provide information about the status (e.g. rejection, acceptance) of a credit transfer instruction, a direct debit instruction, as well as other intra-agent instructions (for example FIToFIPaymentCancellationRequest).
// The FIToFIPaymentStatusReport message refers to the original instruction(s) by means of references only or by means of references and a set of elements from the original instruction.
// The FIToFIPaymentStatusReport message can be used in domestic and cross-border scenarios.
// The FIToFIPaymentStatusReport may also be sent to the receiver of the payment in a real time payment scenario, as both sides of the transactions must be informed of the status of the transaction (e.g. either the beneficiary is credited, or the transaction is rejected).
type FIToFIPaymentStatusReportV09 struct {
	XMLName xml.Name
	
		// Set of characteristics shared by all individual transactions included in the status report message.
	GrpHdr *GroupHeader53 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09 GrpHdr" json:"grp_hdr" `
	
		// Original group information concerning the group of transactions, to which the status report message refers to.
	OrgnlGrpInfAndSts []*OriginalGroupHeader13 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09 OrgnlGrpInfAndSts,omitempty" json:"orgnl_grp_inf_and_sts" `
	
		// Information concerning the original transactions, to which the status report message refers.
	TxInfAndSts []*PaymentTransaction91 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09 TxInfAndSts,omitempty" json:"tx_inf_and_sts" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (f *FIToFIPaymentStatusReportV09) AddGrpHdr() *GroupHeader53 {
	f.GrpHdr = new(GroupHeader53)
	return f.GrpHdr
}

func (f *FIToFIPaymentStatusReportV09) AddOrgnlGrpInfAndSts() *OriginalGroupHeader13 {
	newValue := new (OriginalGroupHeader13)
	f.OrgnlGrpInfAndSts = append(f.OrgnlGrpInfAndSts, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV09) AddTxInfAndSts() *PaymentTransaction91 {
	newValue := new (PaymentTransaction91)
	f.TxInfAndSts = append(f.TxInfAndSts, newValue)
	return newValue
}

func (f *FIToFIPaymentStatusReportV09) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	f.SplmtryData = append(f.SplmtryData, newValue)
	return newValue
}

