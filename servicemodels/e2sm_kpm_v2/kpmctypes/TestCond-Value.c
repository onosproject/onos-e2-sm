/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-rm.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TestCond-Value.h"

asn_per_constraints_t asn_PER_type_TestCond_Value_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  3,  3,  0,  5 }	/* (0..5,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
asn_TYPE_member_t asn_MBR_TestCond_Value_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TestCond_Value, choice.valueInt),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"valueInt"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCond_Value, choice.valueEnum),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"valueEnum"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCond_Value, choice.valueBool),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BOOLEAN,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"valueBool"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCond_Value, choice.valueBitS),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_BIT_STRING,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"valueBitS"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCond_Value, choice.valueOctS),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_OCTET_STRING,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"valueOctS"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCond_Value, choice.valuePrtS),
		(ASN_TAG_CLASS_CONTEXT | (5 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"valuePrtS"
		},
};
static const asn_TYPE_tag2member_t asn_MAP_TestCond_Value_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* valueInt */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* valueEnum */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* valueBool */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* valueBitS */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 }, /* valueOctS */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 5, 0, 0 } /* valuePrtS */
};
asn_CHOICE_specifics_t asn_SPC_TestCond_Value_specs_1 = {
	sizeof(struct TestCond_Value),
	offsetof(struct TestCond_Value, _asn_ctx),
	offsetof(struct TestCond_Value, present),
	sizeof(((struct TestCond_Value *)0)->present),
	asn_MAP_TestCond_Value_tag2el_1,
	6,	/* Count of tags in the map */
	0, 0,
	6	/* Extensions start */
};
asn_TYPE_descriptor_t asn_DEF_TestCond_Value = {
	"TestCond-Value",
	"TestCond-Value",
	&asn_OP_CHOICE,
	0,	/* No effective tags (pointer) */
	0,	/* No effective tags (count) */
	0,	/* No tags (pointer) */
	0,	/* No tags (count) */
	{ 0, &asn_PER_type_TestCond_Value_constr_1, CHOICE_constraint },
	asn_MBR_TestCond_Value_1,
	6,	/* Elements count */
	&asn_SPC_TestCond_Value_specs_1	/* Additional specs */
};

