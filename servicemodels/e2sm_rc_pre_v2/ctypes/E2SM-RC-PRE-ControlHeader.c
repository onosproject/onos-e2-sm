/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v2.asn1"
 */

#include "E2SM-RC-PRE-ControlHeader.h"

static asn_oer_constraints_t asn_OER_type_E2SM_RC_PRE_ControlHeader_constr_1 CC_NOTUSED = {
	{ 0, 0 },
	-1};
static asn_per_constraints_t asn_PER_type_E2SM_RC_PRE_ControlHeader_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  0,  0,  0,  0 }	/* (0..0,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_TYPE_member_t asn_MBR_E2SM_RC_PRE_ControlHeader_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct E2SM_RC_PRE_ControlHeader, choice.controlHeader_Format1),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_E2SM_RC_PRE_ControlHeader_Format1,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"controlHeader-Format1"
		},
};
static const asn_TYPE_tag2member_t asn_MAP_E2SM_RC_PRE_ControlHeader_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 } /* controlHeader-Format1 */
};
static asn_CHOICE_specifics_t asn_SPC_E2SM_RC_PRE_ControlHeader_specs_1 = {
	sizeof(struct E2SM_RC_PRE_ControlHeader),
	offsetof(struct E2SM_RC_PRE_ControlHeader, _asn_ctx),
	offsetof(struct E2SM_RC_PRE_ControlHeader, present),
	sizeof(((struct E2SM_RC_PRE_ControlHeader *)0)->present),
	asn_MAP_E2SM_RC_PRE_ControlHeader_tag2el_1,
	1,	/* Count of tags in the map */
	0, 0,
	1	/* Extensions start */
};
asn_TYPE_descriptor_t asn_DEF_E2SM_RC_PRE_ControlHeader = {
	"E2SM-RC-PRE-ControlHeader",
	"E2SM-RC-PRE-ControlHeader",
	&asn_OP_CHOICE,
	0,	/* No effective tags (pointer) */
	0,	/* No effective tags (count) */
	0,	/* No tags (pointer) */
	0,	/* No tags (count) */
	{ &asn_OER_type_E2SM_RC_PRE_ControlHeader_constr_1, &asn_PER_type_E2SM_RC_PRE_ControlHeader_constr_1, CHOICE_constraint },
	asn_MBR_E2SM_RC_PRE_ControlHeader_1,
	1,	/* Elements count */
	&asn_SPC_E2SM_RC_PRE_ControlHeader_specs_1	/* Additional specs */
};

