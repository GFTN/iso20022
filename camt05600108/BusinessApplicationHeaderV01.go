package camt05600108

import (
	"encoding/xml"
)

type Message1 struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"body"`
}

func (d *Message1) AddMessage() *BusinessApplicationHeaderV01 {
	d.Body = new(BusinessApplicationHeaderV01)
	return d.Body
}

// The Business Layer deals with Business Messages. The behaviour of the Business Messages is fully described by the Business Transaction and the structure of the Business Messages is fully described by the Message Definitions and related Message Rules, Rules and Market Practices. All of which are registered in the ISO 20022 Repository.
// A single new Business Message (with its accompagnying business application header) is created - by the sending MessagingEndpoint - for each business event; that is each interaction in a Business Transaction. A Business Message adheres to the following principles:
// " A Business Message (and its business application header) must not contain information about the Message Transport System or the mechanics or mechanism of message sending, transportation, or receipt. 
// " A Business Message must be comprehensible outside of the context of the Transport Message. That is the Business Message must not require knowledge of the Transport Message to be understood.
// " A Business Message may contain headers, footers, and envelopes that are meaningful for the business. When present, they are treated as any other message content, which means that they are considered part of the Message Definition of the Business Message and as such will be part of the ISO 20022 Repository.
// " A Business Message refers to Business Actors by their Name. Each instance of a Business Actor has one Name. The Business Actor must not be referred to in the Transport Layer.
// Specific usage of this BusinessMessageHeader may be defined by the relevant SEG.
type BusinessApplicationHeaderV01 struct {
	XMLName xml.Name
	
		// Contains the character set of the text-based elements used in the Business Message.
	CharSet *UnicodeChartsCode `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CharSet,omitempty" json:"char_set" `
	
		// The sending MessagingEndpoint that has created this Business Message for the receiving MessagingEndpoint that will process this Business Message.
	// 
	// Note	the sending MessagingEndpoint might be different from the sending address potentially contained in the transport header (as defined in the transport layer).
	Fr *Party9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Fr" json:"fr" `
	
		// The MessagingEndpoint designated by the sending MessagingEndpoint to be the recipient who will ultimately process this Business Message.
	// 
	// Note the receiving MessagingEndpoint might be different from the receiving address potentially contained in the transport header (as defined in the transport layer).
	To *Party9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 To" json:"to" `
	
		// Unambiguously identifies the Business Message to the MessagingEndpoint that has created the Business Message.
	BizMsgIdr *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 BizMsgIdr" json:"biz_msg_idr" `
	
		// Contains the MessageIdentifier that defines the BusinessMessage.
	// It must contain a MessageIdentifier published on the ISO 20022 website.
	// 
	// example	camt.001.001.03.
	MsgDefIdr *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 MsgDefIdr" json:"msg_def_idr" `
	
		// Specifies the business service agreed between the two MessagingEndpoints under which rules this Business Message is exchanged.
	//  To be used when there is a choice of processing services or processing service levels.
	// Example: E&I.
	BizSvc *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 BizSvc,omitempty" json:"biz_svc" `
	
		// Date and time when this Business Message (header) was created.
	// Note Times must be normalized, using the "Z" annotation.
	CreDt *ISONormalisedDateTime `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CreDt" json:"cre_dt" `
	
		// Indicates whether the message is a Copy, a Duplicate or a copy of a duplicate of a previously sent ISO 20022 Message.
	CpyDplct *CopyDuplicate1Code `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 CpyDplct,omitempty" json:"cpy_dplct" `
	
		// Flag indicating if the Business Message exchanged between the MessagingEndpoints is possibly a duplicate. 
	// If the receiving MessagingEndpoint did not receive the original, then this Business Message should be processed as if it were the original. 
	// 
	// If the receiving MessagingEndpoint did receive the original, then it should perform necessary actions to avoid processing this Business Message again.
	// 
	// This will guarantee business idempotent behaviour.
	// 
	// NOTE: this is named "PossResend" in FIX - this is an application level resend not a network level retransmission.
	PssblDplct *YesNoIndicator `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 PssblDplct,omitempty" json:"pssbl_dplct" `
	
		// Relative indication of the processing precedence of the message over a (set of) Business Messages with assigned priorities.
	Prty *BusinessMessagePriorityCode `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Prty,omitempty" json:"prty" `
	
		// Contains the digital signature of the Business Entity authorised to sign this Business Message.
	Sgntr *SignatureEnvelope `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Sgntr,omitempty" json:"sgntr" `
	
		// Specifies the Business Application Header of the Business Message to which this Business Message relates.
	// Can be used when replying to a query; can also be used when canceling or amending.
	Rltd *BusinessApplicationHeader1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.01 Rltd,omitempty" json:"rltd" `
	
}


func (b *BusinessApplicationHeaderV01) SetCharSet(value string) {
	b.CharSet = (*UnicodeChartsCode)(&value)
}

func (b *BusinessApplicationHeaderV01) AddFr() *Party9Choice {
	b.Fr = new(Party9Choice)
	return b.Fr
}

func (b *BusinessApplicationHeaderV01) AddTo() *Party9Choice {
	b.To = new(Party9Choice)
	return b.To
}

func (b *BusinessApplicationHeaderV01) SetBizMsgIdr(value string) {
	b.BizMsgIdr = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeaderV01) SetMsgDefIdr(value string) {
	b.MsgDefIdr = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeaderV01) SetBizSvc(value string) {
	b.BizSvc = (*Max35Text)(&value)
}

func (b *BusinessApplicationHeaderV01) SetCreDt(value string) {
	b.CreDt = (*ISONormalisedDateTime)(&value)
}

func (b *BusinessApplicationHeaderV01) SetCpyDplct(value string) {
	b.CpyDplct = (*CopyDuplicate1Code)(&value)
}

func (b *BusinessApplicationHeaderV01) SetPssblDplct(value string) {
	b.PssblDplct = (*YesNoIndicator)(&value)
}

func (b *BusinessApplicationHeaderV01) SetPrty(value string) {
	b.Prty = (*BusinessMessagePriorityCode)(&value)
}

func (b *BusinessApplicationHeaderV01) AddSgntr() *SignatureEnvelope {
	b.Sgntr = new(SignatureEnvelope)
	return b.Sgntr
}

func (b *BusinessApplicationHeaderV01) AddRltd() *BusinessApplicationHeader1 {
	b.Rltd = new(BusinessApplicationHeader1)
	return b.Rltd
}

