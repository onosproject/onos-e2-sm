/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TEST-EnumeratedExtensible.h"

/*
 * This type is implemented using NativeEnumerated,
 * so here we adjust the DEF accordingly.
 */
asn_per_constraints_t asn_PER_type_TEST_EnumeratedExtensible_constr_1 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  3,  3,  0,  5 }	/* (0..5,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static const asn_INTEGER_enum_map_t asn_MAP_TEST_EnumeratedExtensible_value2enum_1[] = {
	{ 0,	5,	"enum1" },
	{ 1,	5,	"enum2" },
	{ 2,	5,	"enum3" },
	{ 3,	5,	"enum4" },
	{ 4,	5,	"enum5" },
	{ 5,	5,	"enum6" }
	/* This list is extensible */
};
static const unsigned int asn_MAP_TEST_EnumeratedExtensible_enum2value_1[] = {
	0,	/* enum1(0) */
	1,	/* enum2(1) */
	2,	/* enum3(2) */
	3,	/* enum4(3) */
	4,	/* enum5(4) */
	5	/* enum6(5) */
	/* This list is extensible */
};
const asn_INTEGER_specifics_t asn_SPC_TEST_EnumeratedExtensible_specs_1 = {
	asn_MAP_TEST_EnumeratedExtensible_value2enum_1,	/* "tag" => N; sorted by tag */
	asn_MAP_TEST_EnumeratedExtensible_enum2value_1,	/* N => "tag"; sorted by N */
	6,	/* Number of elements in the maps */
	7,	/* Extensions before this member */
	1,	/* Strict enumeration */
	0,	/* Native long size */
	0
};
static const ber_tlv_tag_t asn_DEF_TEST_EnumeratedExtensible_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (10 << 2))
};
asn_TYPE_descriptor_t asn_DEF_TEST_EnumeratedExtensible = {
	"TEST-EnumeratedExtensible",
	"TEST-EnumeratedExtensible",
	&asn_OP_NativeEnumerated,
	asn_DEF_TEST_EnumeratedExtensible_tags_1,
	sizeof(asn_DEF_TEST_EnumeratedExtensible_tags_1)
		/sizeof(asn_DEF_TEST_EnumeratedExtensible_tags_1[0]), /* 1 */
	asn_DEF_TEST_EnumeratedExtensible_tags_1,	/* Same as above */
	sizeof(asn_DEF_TEST_EnumeratedExtensible_tags_1)
		/sizeof(asn_DEF_TEST_EnumeratedExtensible_tags_1[0]), /* 1 */
	{ 0, &asn_PER_type_TEST_EnumeratedExtensible_constr_1, NativeEnumerated_constraint },
	0, 0,	/* Defined elsewhere */
	&asn_SPC_TEST_EnumeratedExtensible_specs_1	/* Additional specs */
};

