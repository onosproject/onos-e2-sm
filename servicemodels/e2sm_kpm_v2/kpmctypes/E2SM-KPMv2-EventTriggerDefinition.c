/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-changed.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "E2SM-KPMv2-EventTriggerDefinition.h"

#include "E2SM-KPMv2-EventTriggerDefinition-Format1.h"
static asn_per_constraints_t asn_PER_type_eventDefinition_formats_constr_2 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  0,  0,  0,  0 }	/* (0..0,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_TYPE_member_t asn_MBR_eventDefinition_formats_2[] = {
	{ ATF_POINTER, 0, offsetof(struct E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats, choice.eventDefinition_Format1),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_E2SM_KPMv2_EventTriggerDefinition_Format1,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"eventDefinition-Format1"
		},
};
static const asn_TYPE_tag2member_t asn_MAP_eventDefinition_formats_tag2el_2[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 } /* eventDefinition-Format1 */
};
static asn_CHOICE_specifics_t asn_SPC_eventDefinition_formats_specs_2 = {
	sizeof(struct E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats),
	offsetof(struct E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats, _asn_ctx),
	offsetof(struct E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats, present),
	sizeof(((struct E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats *)0)->present),
	asn_MAP_eventDefinition_formats_tag2el_2,
	1,	/* Count of tags in the map */
	0, 0,
	1	/* Extensions start */
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_eventDefinition_formats_2 = {
	"eventDefinition-formats",
	"eventDefinition-formats",
	&asn_OP_CHOICE,
	0,	/* No effective tags (pointer) */
	0,	/* No effective tags (count) */
	0,	/* No tags (pointer) */
	0,	/* No tags (count) */
	{ 0, &asn_PER_type_eventDefinition_formats_constr_2, CHOICE_constraint },
	asn_MBR_eventDefinition_formats_2,
	1,	/* Elements count */
	&asn_SPC_eventDefinition_formats_specs_2	/* Additional specs */
};

static asn_TYPE_member_t asn_MBR_E2SM_KPMv2_EventTriggerDefinition_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct E2SM_KPMv2_EventTriggerDefinition, eventDefinition_formats),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_eventDefinition_formats_2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"eventDefinition-formats"
		},
};
static const ber_tlv_tag_t asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_E2SM_KPMv2_EventTriggerDefinition_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 } /* eventDefinition-formats */
};
static asn_SEQUENCE_specifics_t asn_SPC_E2SM_KPMv2_EventTriggerDefinition_specs_1 = {
	sizeof(struct E2SM_KPMv2_EventTriggerDefinition),
	offsetof(struct E2SM_KPMv2_EventTriggerDefinition, _asn_ctx),
	asn_MAP_E2SM_KPMv2_EventTriggerDefinition_tag2el_1,
	1,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_E2SM_KPMv2_EventTriggerDefinition = {
	"E2SM-KPMv2-EventTriggerDefinition",
	"E2SM-KPMv2-EventTriggerDefinition",
	&asn_OP_SEQUENCE,
	asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1,
	sizeof(asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1)
		/sizeof(asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1[0]), /* 1 */
	asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1,	/* Same as above */
	sizeof(asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1)
		/sizeof(asn_DEF_E2SM_KPMv2_EventTriggerDefinition_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_E2SM_KPMv2_EventTriggerDefinition_1,
	1,	/* Elements count */
	&asn_SPC_E2SM_KPMv2_EventTriggerDefinition_specs_1	/* Additional specs */
};

