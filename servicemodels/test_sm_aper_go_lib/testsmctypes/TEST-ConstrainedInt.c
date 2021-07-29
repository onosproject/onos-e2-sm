/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TEST-ConstrainedInt.h"

static int
attrCiB_3_constraint(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	unsigned long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const unsigned long *)sptr;
	
	if((value >= 10 && value <= 4294967295)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

/*
 * This type is implemented using NativeInteger,
 * so here we adjust the DEF accordingly.
 */
static int
memb_attrCiA_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value >= 10 && value <= 100)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_attrCiB_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	unsigned long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const unsigned long *)sptr;
	
	if((value >= 10 && value <= 4294967295)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_attrCiC_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value <= 100)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_attrCiD_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value >= 10 && value <= 20)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_attrCiE_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value == 10)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int
memb_attrCiF_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	long value;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	value = *(const long *)sptr;
	
	if((value == 10)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static asn_per_constraints_t asn_PER_type_attrCiB_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 32, -1,  10,  4294967295 }	/* (10..4294967295) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_attrCiA_constr_2 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  10,  100 }	/* (10..100) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_attrCiB_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 32, -1,  10,  4294967295 }	/* (10..4294967295) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_attrCiC_constr_4 CC_NOTUSED = {
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 }	/* (MIN..100) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_attrCiD_constr_5 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 4,  4,  10,  20 }	/* (10..20) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_attrCiE_constr_6 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 0,  0,  10,  10 }	/* (10..10) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static asn_per_constraints_t asn_PER_memb_attrCiF_constr_7 CC_NOTUSED = {
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  0,  0,  10,  10 }	/* (10..10,...) */,
	{ APC_UNCONSTRAINED,	-1, -1,  0,  0 },
	0, 0	/* No PER value map */
};
static const asn_INTEGER_specifics_t asn_SPC_attrCiB_specs_3 = {
	0,	0,	0,	0,	0,
	0,	/* Native long size */
	1	/* Unsigned representation */
};
static const ber_tlv_tag_t asn_DEF_attrCiB_tags_3[] = {
	(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
	(ASN_TAG_CLASS_UNIVERSAL | (2 << 2))
};
static /* Use -fall-defs-global to expose */
asn_TYPE_descriptor_t asn_DEF_attrCiB_3 = {
	"attrCiB",
	"attrCiB",
	&asn_OP_NativeInteger,
	asn_DEF_attrCiB_tags_3,
	sizeof(asn_DEF_attrCiB_tags_3)
		/sizeof(asn_DEF_attrCiB_tags_3[0]) - 1, /* 1 */
	asn_DEF_attrCiB_tags_3,	/* Same as above */
	sizeof(asn_DEF_attrCiB_tags_3)
		/sizeof(asn_DEF_attrCiB_tags_3[0]), /* 2 */
	{ 0, &asn_PER_type_attrCiB_constr_3, attrCiB_3_constraint },
	0, 0,	/* No members */
	&asn_SPC_attrCiB_specs_3	/* Additional specs */
};

static asn_TYPE_member_t asn_MBR_TEST_ConstrainedInt_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_ConstrainedInt, attrCiA),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, &asn_PER_memb_attrCiA_constr_2,  memb_attrCiA_constraint_1 },
		0, 0, /* No default value */
		"attrCiA"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_ConstrainedInt, attrCiB),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_attrCiB_3,
		0,
		{ 0, &asn_PER_memb_attrCiB_constr_3,  memb_attrCiB_constraint_1 },
		0, 0, /* No default value */
		"attrCiB"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_ConstrainedInt, attrCiC),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, &asn_PER_memb_attrCiC_constr_4,  memb_attrCiC_constraint_1 },
		0, 0, /* No default value */
		"attrCiC"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_ConstrainedInt, attrCiD),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, &asn_PER_memb_attrCiD_constr_5,  memb_attrCiD_constraint_1 },
		0, 0, /* No default value */
		"attrCiD"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_ConstrainedInt, attrCiE),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, &asn_PER_memb_attrCiE_constr_6,  memb_attrCiE_constraint_1 },
		0, 0, /* No default value */
		"attrCiE"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_ConstrainedInt, attrCiF),
		(ASN_TAG_CLASS_CONTEXT | (5 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_NativeInteger,
		0,
		{ 0, &asn_PER_memb_attrCiF_constr_7,  memb_attrCiF_constraint_1 },
		0, 0, /* No default value */
		"attrCiF"
		},
};
static const ber_tlv_tag_t asn_DEF_TEST_ConstrainedInt_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_TEST_ConstrainedInt_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* attrCiA */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* attrCiB */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* attrCiC */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* attrCiD */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 }, /* attrCiE */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 5, 0, 0 } /* attrCiF */
};
static asn_SEQUENCE_specifics_t asn_SPC_TEST_ConstrainedInt_specs_1 = {
	sizeof(struct TEST_ConstrainedInt),
	offsetof(struct TEST_ConstrainedInt, _asn_ctx),
	asn_MAP_TEST_ConstrainedInt_tag2el_1,
	6,	/* Count of tags in the map */
	0, 0, 0,	/* Optional elements (not needed) */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_TEST_ConstrainedInt = {
	"TEST-ConstrainedInt",
	"TEST-ConstrainedInt",
	&asn_OP_SEQUENCE,
	asn_DEF_TEST_ConstrainedInt_tags_1,
	sizeof(asn_DEF_TEST_ConstrainedInt_tags_1)
		/sizeof(asn_DEF_TEST_ConstrainedInt_tags_1[0]), /* 1 */
	asn_DEF_TEST_ConstrainedInt_tags_1,	/* Same as above */
	sizeof(asn_DEF_TEST_ConstrainedInt_tags_1)
		/sizeof(asn_DEF_TEST_ConstrainedInt_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_TEST_ConstrainedInt_1,
	6,	/* Elements count */
	&asn_SPC_TEST_ConstrainedInt_specs_1	/* Additional specs */
};

