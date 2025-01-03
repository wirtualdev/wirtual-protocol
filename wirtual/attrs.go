package wirtual

// Names of participant attributes for SIP.
const (
	// AttrSIPPrefix is shared for all SIP attributes.
	AttrSIPPrefix = "sip."
	// AttrSIPCallID attribute contains Wirtual SIP call ID.
	AttrSIPCallID = AttrSIPPrefix + "callID"
	// AttrSIPTrunkID attribute contains Wirtual SIP Trunk ID used for the call.
	AttrSIPTrunkID = AttrSIPPrefix + "trunkID"
	// AttrSIPDispatchRuleID attribute contains Wirtual SIP DispatchRule ID used for the inbound call.
	AttrSIPDispatchRuleID = AttrSIPPrefix + "ruleID"
	// AttrSIPTrunkNumber attribute contains number associate with Wirtual SIP Trunk.
	// This attribute will be omitted if HidePhoneNumber is set.
	AttrSIPTrunkNumber = AttrSIPPrefix + "trunkPhoneNumber"
	// AttrSIPPhoneNumber attribute contains number external to Wirtual SIP (caller for inbound and called number for outbound).
	// This attribute will be omitted if HidePhoneNumber is set.
	AttrSIPPhoneNumber = AttrSIPPrefix + "phoneNumber"
	// AttrSIPCallStatus attribute contains current call status for a SIP call associated with the participant.
	//
	// SIP participant is ready when it reaches "active" status.
	AttrSIPCallStatus = AttrSIPPrefix + "callStatus"
	// AttrSIPHeaderPrefix is a prefix for automatically mapped SIP header attributes.
	AttrSIPHeaderPrefix = AttrSIPPrefix + "h."
)
