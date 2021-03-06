package camt02600107

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *UnableToApplyV07 `xml:"UblToApply" json:"body"`
}

func (d *Message) AddMessage() *UnableToApplyV07 {
	d.Body = new(UnableToApplyV07)
	return d.Body
}

// Scope
// The UnableToApply message is sent by a case creator or a case assigner to a case assignee. This message is used to initiate an investigation of a payment instruction that cannot be executed or reconciled.
// Usage
// The Unable To Apply case occurs in two situations:
// - an agent cannot execute the payment instruction due to missing or incorrect information;
// - a creditor cannot reconcile the payment entry as it is received unexpectedly, or missing or incorrect information prevents reconciliation.
// The UnableToApply message:
// - always follows the reverse route of the original payment instruction;
// - must be forwarded to the preceding agents in the payment processing chain, where appropriate;
// - covers one and only one payment instruction (or payment entry) at a time; if several payment instructions cannot be executed or several payment entries cannot be reconciled, then multiple UnableToApply messages must be sent.
// Depending on what stage the payment is and what has been done to it, for example incorrect routing, errors/omissions when processing the instruction or even the absence of any error, the unable to apply case may lead to a:
// - AdditionalPaymentInformation message, sent to the case creator/case assigner, if a truncation or omission in a payment instruction prevented reconciliation;
// - CustomerPaymentCancellationRequest or FIToFIPaymentCancellationRequest message, sent to the subsequent agent in the payment processing chain, if the original payment instruction has been incorrectly routed through the chain of agents (this also entails a new corrected payment instruction being issued). Prior to sending the payment cancellation request, the agent should first send a resolution indicating that a cancellation will follow (CWFW);
// - RequestToModifyPayment message, sent to the subsequent agent in the payment processing chain, if a truncation or omission has occurred during the processing of the original payment instruction. Prior to sending the modify payment request, the agent should first send a resolution indicating that a modification will follow (MWFW);
// - DebitAuthorisationRequest message, sent to the case creator/case assigner, if the payment instruction has reached an incorrect creditor.
// The UnableToApply message has the following main characteristics:
// The case creator (the instructed party/creditor of a payment instruction) assigns a unique case identification and optionally the reason code for the UnableToApply message. This information will be passed unchanged to all subsequent case assignees.
// The case creator specifies the identification of the underlying payment instruction. This identification needs to be updated by the subsequent case assigner(s) in order to match the one used with their case assignee(s).
// The UnableToApply Justification element allows the assigner to indicate whether a specific element causes the unable to apply situation (incorrect element and/or mismatched element can be listed) or whether any supplementary information needs to be forwarded.
type UnableToApplyV07 struct {
	XMLName xml.Name
	
		// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The assigner must be the sender of this confirmation and the assignee must be the receiver.
	Assgnmt *CaseAssignment5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Assgnmt" json:"assgnmt" `
	
		// Identifies the investigation case.
	Case *Case5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Case,omitempty" json:"case" `
	
		// References the payment instruction or statement entry that a party is unable to execute or unable to reconcile.
	Undrlyg *UnderlyingTransaction5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Undrlyg" json:"undrlyg" `
	
		// Explains the reason why the case creator is unable to apply the instruction.
	Justfn *UnableToApplyJustification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 Justfn" json:"justfn" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.026.001.07 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (u *UnableToApplyV07) AddAssgnmt() *CaseAssignment5 {
	u.Assgnmt = new(CaseAssignment5)
	return u.Assgnmt
}

func (u *UnableToApplyV07) AddCase() *Case5 {
	u.Case = new(Case5)
	return u.Case
}

func (u *UnableToApplyV07) AddUndrlyg() *UnderlyingTransaction5Choice {
	u.Undrlyg = new(UnderlyingTransaction5Choice)
	return u.Undrlyg
}

func (u *UnableToApplyV07) AddJustfn() *UnableToApplyJustification3Choice {
	u.Justfn = new(UnableToApplyJustification3Choice)
	return u.Justfn
}

func (u *UnableToApplyV07) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	u.SplmtryData = append(u.SplmtryData, newValue)
	return newValue
}

