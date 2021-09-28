package camt02900109

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *ResolutionOfInvestigationV09 `xml:"RsltnOfInvstgtn" json:"body"`
}

func (d *Message) AddMessage() *ResolutionOfInvestigationV09 {
	d.Body = new(ResolutionOfInvestigationV09)
	return d.Body
}

// Scope
// The ResolutionOfInvestigation message is sent by a case assignee to a case creator/case assigner.
// This message is used to inform of the resolution of a case, and optionally provides details about.
// - the corrective action undertaken by the case assignee;
// - information on the return where applicable.
// Usage
// The ResolutionOfInvestigation message is used by the case assignee to inform a case creator or case assigner about the resolution of a:
// - request to cancel payment case;
// - request to modify payment case;
// - unable to apply case;
// - claim non receipt case.
// The ResolutionOfInvestigation message covers one and only one case at a time. If the case assignee needs to communicate about several cases, then several Resolution Of Investigation messages must be sent.
// The ResolutionOfInvestigation message provides:
// - the final outcome of the case, whether positive or negative;
// - optionally, the details of the corrective action undertaken by the case assignee and the information of the return.
// Whenever a payment instruction has been generated to solve the case under investigation following a claim non receipt or an unable to apply, the optional CorrectionTransaction component present in the message must be completed.
// Whenever the action of modifying or cancelling a payment results in funds being returned or reversed, an investigating agent may provide the details in the resolution related investigation component, to identify the return or reversal transaction. These details will facilitate the account reconciliations at the initiating bank and the intermediaries. It must be stressed that the return or reversal of funds is outside the scope of this Exceptions and Investigation service. The features given here is only meant to transmit the information of return or reversal when it is available through the resolution of the case.
// The ResolutionOfInvestigation message must:
// - be forwarded by all subsequent case assignee(s) until it reaches the case creator;
// - not be used in place of a RejectCaseAssignment or CaseStatusReport or NotificationOfCaseAssignment message.
// Take note of an exceptional rule that allows the use of ResolutionOfInvestigation in lieu of a CaseStatusReport. CaseStatusReport is a response-message to a CaseStatusReportRequest. The latter which is sent when the assigner has reached its own time-out threshold to receive a response. However it may happen that when the request arrives, the investigating agent has just obtained a resolution. In such a situation, it would be redundant to send a CaseStatusReport when then followed immediately by a ResolutionOfInvestigation. It is therefore quite acceptable for the investigating agent, the assignee, to skip the Case Status Report and send the ResolutionOfInvestigation message directly.
// The ResolutionOfInvestigation message should be the sole message to respond to a cancellation request. Details of the underlying transactions and the related statuses for which the cancellation request has been issued may be provided in the CancellationDetails component.
type ResolutionOfInvestigationV09 struct {
	XMLName xml.Name
	
		// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The assigner must be the sender of this confirmation and the assignee must be the receiver.
	Assgnmt *CaseAssignment5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Assgnmt" json:"assgnmt" `
	
		// Identifies a resolved case.
	RslvdCase *Case5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RslvdCase,omitempty" json:"rslvd_case" `
	
		// Indicates the status of the investigation.
	Sts *InvestigationStatus5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 Sts" json:"sts" `
	
		// Specifies the details of the underlying transactions being cancelled.
	CxlDtls []*UnderlyingTransaction22 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CxlDtls,omitempty" json:"cxl_dtls" `
	
		// Specifies the details of the underlying transaction being modified.
	ModDtls *PaymentTransaction107 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ModDtls,omitempty" json:"mod_dtls" `
	
		// Specifies the details of the underlying transaction for which a claim non receipt has been initiated.
	ClmNonRctDtls *ClaimNonReceipt2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 ClmNonRctDtls,omitempty" json:"clm_non_rct_dtls" `
	
		// Details on the underlying statement entry.
	StmtDtls *StatementResolutionEntry4 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 StmtDtls,omitempty" json:"stmt_dtls" `
	
		// References a transaction initiated to fix the case under investigation.
	CrrctnTx *CorrectiveTransaction4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 CrrctnTx,omitempty" json:"crrctn_tx" `
	
		// Reference to fix the case under investigation as part of the resolution.
	RsltnRltdInf *ResolutionData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 RsltnRltdInf,omitempty" json:"rsltn_rltd_inf" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.029.001.09 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (r *ResolutionOfInvestigationV09) AddAssgnmt() *CaseAssignment5 {
	r.Assgnmt = new(CaseAssignment5)
	return r.Assgnmt
}

func (r *ResolutionOfInvestigationV09) AddRslvdCase() *Case5 {
	r.RslvdCase = new(Case5)
	return r.RslvdCase
}

func (r *ResolutionOfInvestigationV09) AddSts() *InvestigationStatus5Choice {
	r.Sts = new(InvestigationStatus5Choice)
	return r.Sts
}

func (r *ResolutionOfInvestigationV09) AddCxlDtls() *UnderlyingTransaction22 {
	newValue := new (UnderlyingTransaction22)
	r.CxlDtls = append(r.CxlDtls, newValue)
	return newValue
}

func (r *ResolutionOfInvestigationV09) AddModDtls() *PaymentTransaction107 {
	r.ModDtls = new(PaymentTransaction107)
	return r.ModDtls
}

func (r *ResolutionOfInvestigationV09) AddClmNonRctDtls() *ClaimNonReceipt2Choice {
	r.ClmNonRctDtls = new(ClaimNonReceipt2Choice)
	return r.ClmNonRctDtls
}

func (r *ResolutionOfInvestigationV09) AddStmtDtls() *StatementResolutionEntry4 {
	r.StmtDtls = new(StatementResolutionEntry4)
	return r.StmtDtls
}

func (r *ResolutionOfInvestigationV09) AddCrrctnTx() *CorrectiveTransaction4Choice {
	r.CrrctnTx = new(CorrectiveTransaction4Choice)
	return r.CrrctnTx
}

func (r *ResolutionOfInvestigationV09) AddRsltnRltdInf() *ResolutionData1 {
	r.RsltnRltdInf = new(ResolutionData1)
	return r.RsltnRltdInf
}

func (r *ResolutionOfInvestigationV09) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	r.SplmtryData = append(r.SplmtryData, newValue)
	return newValue
}

