package camt05600108

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *FIToFIPaymentCancellationRequestV08 `xml:"FIToFIPmtCxlReq" json:"body"`
}

func (d *Message) AddMessage() *FIToFIPaymentCancellationRequestV08 {
	d.Body = new(FIToFIPaymentCancellationRequestV08)
	return d.Body
}

// Scope
// The FIToFIPaymentCancellationRequest message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the cancellation of an original payment instruction. The FIToFIPaymentCancellationRequest message is exchanged between the instructing agent and the instructed agent to request the cancellation of a interbank payment message previously sent (such as FIToFICustomerCreditTransfer, FIToFICustomerDirectDebit or FinancialInstitutionCreditTransfer).
// 
// The FIToFIPaymentCancellationRequest message supports both the request for cancellation (the instructed agent - or assignee - has not yet processed and forwarded the payment instruction) as well as the request for refund (payment has been fully processed already by the instructed agent - or assignee).
// 
// Usage
// The FIToFIPaymentCancellationRequest message must be answered with a:
// - ResolutionOfInvestigation message with a positive final outcome when the case assignee can perform the requested cancellation;
// - ResolutionOfInvestigation message with a negative final outcome when the case assignee may perform the requested cancellation but fails to do so (too late, irrevocable instruction);
// - RejectInvestigation message when the case assignee is unable or not authorised to perform the requested cancellation;
// - NotificationOfCaseAssignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// A FIToFIPaymentCancellationRequest message concerns one and only one original payment instruction at a time.
// When a case assignee successfully performs a cancellation, it must return the corresponding funds to the case assigner. It may provide some details about the return in the ResolutionOfInvestigation message.
// The processing of a FIToFIPaymentCancellationRequest message case may lead to a DebitAuthorisationRequest message sent to the creditor by its account servicing institution.
// The FIToFIPaymentCancellationRequest message may be used to escalate a case after an unsuccessful request to modify the payment. In this scenario, the case identification remains the same as in the original FIToFIPaymentCancellationRequest message and the element ReopenCaseIndication is set to 'Yes' or 'true'.
// The FIToFIPaymentCancellationRequest message has the following main characteristics: the case creator assigns a unique case identification and the reason code for the cancellation request. This information will be passed unchanged to all subsequent case assignee(s). 
// For the FIToFIPaymentCancellationRequest message the case has been made optional, as the message might be used outside of a case management environment where the case identification is not relevant.
// Moreover, the case identification may be present at different levels:
// - One unique case is defined per cancellation request message: If multiple underlying groups or transactions are present in the message and the case assignee has already forwarded the transaction for which the cancellation is requested, the case cannot be forwarded to the next party in the chain (see rule on uniqueness of the case) and the case creator will have to issue individual cancellation requests for each underlying individual transaction. In response to this cancellation request, the case must also be present at the message level in the Resolution of Investigation message;
// - One case per original group or transaction present in the cancellation request: For each group or transaction, a unique case has been assigned. This means, when a payment instruction has already been forwarded by the case assignee, the cancellation request may be forwarded to next party in the payment chain, with the unique case assigned to the transaction. When the group can only be cancelled partially, new cancellation requests need however to be issued for the individual transactions within the group for which the cancellation request has not been successful. In response to this cancellation request, the case must be present in the cancellation details identifying the original group or transaction in the Resolution of Investigation message;
// - No case used in cancellation request message.
// Cancellation of a cover payment:
// The cancellation of a payment instruction for which cover is provided by a separate instruction always results in the cancellation of the whole transaction, including the cover. The case assignee performing the cancellation must initiate the return of funds to the case creator. The case assigner must not request the cancellation of the cover separately.
// Cancellation request initiators:
// The cancellation of a payment instruction can be initiated by either the debtor/creditor or any subsequent agent in the payment instruction processing chain.
type FIToFIPaymentCancellationRequestV08 struct {
	XMLName xml.Name
	
		// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The assigner must be the sender of this confirmation and the assignee must be the receiver.
	Assgnmt *CaseAssignment5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Assgnmt" json:"assgnmt" `
	
		// Identifies the investigation case.
	Case *Case5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Case,omitempty" json:"case" `
	
		// Provides details on the number of transactions and the control sum of the message.
	CtrlData *ControlData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 CtrlData,omitempty" json:"ctrl_data" `
	
		// Identifies the payment instruction to be cancelled.
	Undrlyg []*UnderlyingTransaction23 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 Undrlyg" json:"undrlyg" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.08 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (f *FIToFIPaymentCancellationRequestV08) AddAssgnmt() *CaseAssignment5 {
	f.Assgnmt = new(CaseAssignment5)
	return f.Assgnmt
}

func (f *FIToFIPaymentCancellationRequestV08) AddCase() *Case5 {
	f.Case = new(Case5)
	return f.Case
}

func (f *FIToFIPaymentCancellationRequestV08) AddCtrlData() *ControlData1 {
	f.CtrlData = new(ControlData1)
	return f.CtrlData
}

func (f *FIToFIPaymentCancellationRequestV08) AddUndrlyg() *UnderlyingTransaction23 {
	newValue := new (UnderlyingTransaction23)
	f.Undrlyg = append(f.Undrlyg, newValue)
	return newValue
}

func (f *FIToFIPaymentCancellationRequestV08) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	f.SplmtryData = append(f.SplmtryData, newValue)
	return newValue
}

