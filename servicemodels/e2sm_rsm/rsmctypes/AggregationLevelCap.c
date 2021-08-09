/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RSM-IEs"
 * 	found in "../v1/e2sm-rsm.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "AggregationLevelCap.h"

/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
asn_per_constraints_t asn_PER_type_AggregationLevelCap_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 3,  3,  0,  4 }	/* (0..4) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static const asn_INTEGER_enum_map_t asn_MAP_AggregationLevelCap_value2enum_1[] = {
	{ 0,	3,	"one" },
	{ 1,	3,	"two" },
	{ 2,	4,	"four" },
	{ 3,	5,	"eight" },
	{ 4,	7,	"sixteen" }
};
static const unsigned int asn_MAP_AggregationLevelCap_enum2value_1[] = {
	3,	/* eight(3) */
	2,	/* four(2) */
	0,	/* one(0) */
	4,	/* sixteen(4) */
	1	/* two(1) */
};
const asn_INTEGER_specifics_t asn_SPC_AggregationLevelCap_specs_1 = {
	asn_MAP_AggregationLevelCap_value2enum_1,	/* "tag" => N; sorted by tag */
	asn_MAP_AggregationLevelCap_enum2value_1,	/* N => "tag"; sorted by N */
	5,	/* Number of elements in the maps */
	0,	/* Enumeration is not extensible */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_AggregationLevelCap_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
asn_TYPE_descriptor_t asn_DEF_AggregationLevelCap = {
	"AggregationLevelCap",
	"AggregationLevelCap",
	&asn_OP_NativeEnumerated,
	asn_DEF_AggregationLevelCap_tags_1,
	sizeof(asn_DEF_AggregationLevelCap_tags_1)
		/sizeof(asn_DEF_AggregationLevelCap_tags_1[0]), /* 1 */
	asn_DEF_AggregationLevelCap_tags_1,	/* Same as above */
	sizeof(asn_DEF_AggregationLevelCap_tags_1)
		/sizeof(asn_DEF_AggregationLevelCap_tags_1[0]), /* 1 */
	{ 0, &asn_PER_type_AggregationLevelCap_constr_1, NativeEnumerated_constraint },
	0, 0,	/* Defined elsewhere */
	&asn_SPC_AggregationLevelCap_specs_1	/* Additional specs */
};

