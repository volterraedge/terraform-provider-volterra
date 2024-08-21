// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package explain_log_record

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.ai_assistant.explain_log_record.Bot"] = BotValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.BotDefenseEventDetails"] = BotDefenseEventDetailsValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.ExplainLogRecordResponse"] = ExplainLogRecordResponseValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.RequestDetails"] = RequestDetailsValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.Signature"] = SignatureValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.SvcPolicyEventDetails"] = SvcPolicyEventDetailsValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.ThreatCampaign"] = ThreatCampaignValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.Violation"] = ViolationValidator()
	vr["ves.io.schema.ai_assistant.explain_log_record.WAFEventDetails"] = WAFEventDetailsValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

}

func InitializeMDRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	initializeEntryRegistry(mdr)
	initializeValidatorRegistry(mdr.ValidatorRegistry)

	initializeCRUDServiceRegistry(mdr, isExternal)
	if isExternal {
		return
	}

	initializeRPCRegistry(mdr)
	initializeAPIGwServiceSlugsRegistry(mdr.APIGwServiceSlugs)
	initializeP0PolicyRegistry(mdr.P0PolicyRegistry)

}