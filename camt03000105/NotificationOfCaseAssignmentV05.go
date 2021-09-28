package camt03000105

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *NotificationOfCaseAssignmentV05 `xml:"NtfctnOfCaseAssgnmt" json:"body"`
}

func (d *Message) AddMessage() *NotificationOfCaseAssignmentV05 {
	d.Body = new(NotificationOfCaseAssignmentV05)
	return d.Body
}

// Scope
// The NotificationOfCaseAssignment message is sent by a case assignee to a case creator/case assigner.
// This message is used to inform the case assigner that:
// - the assignee is reassigning the case to the next agent in the transaction processing chain for further action
// - the assignee will work on the case himself, without re-assigning it to another party, and therefore indicating that the re-assignment has reached its end-point
// Usage
// The NotificationOfCaseAssignment message is used to notify the case creator or case assigner of further action undertaken by the case assignee in a:
// - request to cancel payment case;
// - request to modify payment case;
// - unable to apply case;
// - claim non receipt case.
// The NotificationOfCaseAssignment message:
// - covers one and only one case at a time. If the case assignee needs to inform a case creator or case assigner about several cases, then multiple Notification Of Case Assignment messages must be sent;
// - except in the case where it is used to indicate that an agent is doing the correction himself, this message must be forwarded by all subsequent case assigner(s) until it reaches the case creator;
// - must not be used in place of a Resolution Of Investigation or a Case Status Report message.
// When the assignee does not reassign the case to another party (that is responding with a NotificationOfCaseAssignment message with notification MINE - The case is processed by the assignee), the case assignment should contain the case assignment elements as received in the original query.
type NotificationOfCaseAssignmentV05 struct {
	XMLName xml.Name
	
		// Specifies generic information about the notification.
	// The receiver of a notification must be the party which assigned the case to the sender.
	Hdr *ReportHeader5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 Hdr" json:"hdr" `
	
		// Identifies the investigation case.
	Case *Case5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 Case" json:"case" `
	
		// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The assigner must be the sender of this confirmation and the assignee must be the receiver.
	Assgnmt *CaseAssignment5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 Assgnmt" json:"assgnmt" `
	
		// Information about the type of action taken.
	Ntfctn *CaseForwardingNotification3 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 Ntfctn" json:"ntfctn" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.030.001.05 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (n *NotificationOfCaseAssignmentV05) AddHdr() *ReportHeader5 {
	n.Hdr = new(ReportHeader5)
	return n.Hdr
}

func (n *NotificationOfCaseAssignmentV05) AddCase() *Case5 {
	n.Case = new(Case5)
	return n.Case
}

func (n *NotificationOfCaseAssignmentV05) AddAssgnmt() *CaseAssignment5 {
	n.Assgnmt = new(CaseAssignment5)
	return n.Assgnmt
}

func (n *NotificationOfCaseAssignmentV05) AddNtfctn() *CaseForwardingNotification3 {
	n.Ntfctn = new(CaseForwardingNotification3)
	return n.Ntfctn
}

func (n *NotificationOfCaseAssignmentV05) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	n.SplmtryData = append(n.SplmtryData, newValue)
	return newValue
}

