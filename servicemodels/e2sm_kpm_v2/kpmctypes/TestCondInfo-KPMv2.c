/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-changed.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TestCondInfo-KPMv2.h"

asn_TYPE_member_t asn_MBR_TestCondInfo_KPMv2_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TestCondInfo_KPMv2, testType),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_TestCond_Type_KPMv2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"testType"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCondInfo_KPMv2, testExpr),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_TestCond_Expression_KPMv2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"testExpr"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TestCondInfo_KPMv2, testValue),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		+1,	/* EXPLICIT tag at current level */
		&asn_DEF_TestCond_Value_KPMv2,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"testValue"
		},
};
static const ber_tlv_tag_t asn_DEF_TestCondInfo_KPMv2_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_TestCondInfo_KPMv2_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* testType */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* testExpr */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 } /* testValue */
};
asn_SEQUENCE_specifics_t asn_SPC_TestCondInfo_KPMv2_specs_1 = {
	sizeof(struct TestCondInfo_KPMv2),
	offsetof(struct TestCondInfo_KPMv2, _asn_ctx),
	asn_MAP_TestCondInfo_KPMv2_tag2el_1,
	3,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	3,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_TestCondInfo_KPMv2 = {
	"TestCondInfo-KPMv2",
	"TestCondInfo-KPMv2",
	&asn_OP_SEQUENCE,
	asn_DEF_TestCondInfo_KPMv2_tags_1,
	sizeof(asn_DEF_TestCondInfo_KPMv2_tags_1)
		/sizeof(asn_DEF_TestCondInfo_KPMv2_tags_1[0]), /* 1 */
	asn_DEF_TestCondInfo_KPMv2_tags_1,	/* Same as above */
	sizeof(asn_DEF_TestCondInfo_KPMv2_tags_1)
		/sizeof(asn_DEF_TestCondInfo_KPMv2_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_TestCondInfo_KPMv2_1,
	3,	/* Elements count */
	&asn_SPC_TestCondInfo_KPMv2_specs_1	/* Additional specs */
};

