/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-COMMON-IEs"
 * 	found in "../v1/ORAN-WG3.E2SM-v02.00.02.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "GUAMI.h"

asn_TYPE_member_t asn_MBR_GUAMI_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct GUAMI, pLMNIdentity),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PLMNIdentity,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"pLMNIdentity"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct GUAMI, aMFRegionID),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_AMFRegionID,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"aMFRegionID"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct GUAMI, aMFSetID),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_AMFSetID,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"aMFSetID"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct GUAMI, aMFPointer),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_AMFPointer,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"aMFPointer"
		},
};
static const ber_tlv_tag_t asn_DEF_GUAMI_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_GUAMI_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* pLMNIdentity */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* aMFRegionID */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* aMFSetID */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 } /* aMFPointer */
};
asn_SEQUENCE_specifics_t asn_SPC_GUAMI_specs_1 = {
	sizeof(struct GUAMI),
	offsetof(struct GUAMI, _asn_ctx),
	asn_MAP_GUAMI_tag2el_1,
	4,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	4,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_GUAMI = {
	"GUAMI",
	"GUAMI",
	&asn_OP_SEQUENCE,
	asn_DEF_GUAMI_tags_1,
	sizeof(asn_DEF_GUAMI_tags_1)
		/sizeof(asn_DEF_GUAMI_tags_1[0]), /* 1 */
	asn_DEF_GUAMI_tags_1,	/* Same as above */
	sizeof(asn_DEF_GUAMI_tags_1)
		/sizeof(asn_DEF_GUAMI_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_GUAMI_1,
	4,	/* Elements count */
	&asn_SPC_GUAMI_specs_1	/* Additional specs */
};

