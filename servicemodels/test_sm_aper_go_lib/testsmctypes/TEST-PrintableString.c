/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#include "TEST-PrintableString.h"

static const int permitted_alphabet_table_3[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_3[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_3(const void *sptr) {
	const int *table = permitted_alphabet_table_3;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static const int permitted_alphabet_table_4[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_4[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_4(const void *sptr) {
	const int *table = permitted_alphabet_table_4;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static const int permitted_alphabet_table_5[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_5[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_5(const void *sptr) {
	const int *table = permitted_alphabet_table_5;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static const int permitted_alphabet_table_6[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_6[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_6(const void *sptr) {
	const int *table = permitted_alphabet_table_6;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static const int permitted_alphabet_table_7[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_7[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_7(const void *sptr) {
	const int *table = permitted_alphabet_table_7;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static const int permitted_alphabet_table_8[256] = {
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,	/*                  */
 1, 0, 0, 0, 0, 0, 0, 2, 3, 4, 0, 5, 6, 7, 8, 9,	/* .      '() +,-./ */
10,11,12,13,14,15,16,17,18,19,20, 0, 0,21, 0,22,	/* 0123456789:  = ? */
 0,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,	/*  ABCDEFGHIJKLMNO */
38,39,40,41,42,43,44,45,46,47,48, 0, 0, 0, 0, 0,	/* PQRSTUVWXYZ      */
 0,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,	/*  abcdefghijklmno */
64,65,66,67,68,69,70,71,72,73,74, 0, 0, 0, 0, 0,	/* pqrstuvwxyz      */
};
static const int permitted_alphabet_code2value_8[74] = {
32,39,40,41,43,44,45,46,47,48,49,50,51,52,53,54,
55,56,57,58,61,63,65,66,67,68,69,70,71,72,73,74,
75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,
97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,
113,114,115,116,117,118,119,120,121,122,};


static int check_permitted_alphabet_8(const void *sptr) {
	const int *table = permitted_alphabet_table_8;
	/* The underlying type is PrintableString */
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	const uint8_t *ch = st->buf;
	const uint8_t *end = ch + st->size;
	
	for(; ch < end; ch++) {
		uint8_t cv = *ch;
		if(!table[cv]) return -1;
	}
	return 0;
}

static int
memb_attrPs2_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size == 2)
		 && !check_permitted_alphabet_3(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_attrPs2_3_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_3)/sizeof(permitted_alphabet_table_3[0]))
		return -1;
	return permitted_alphabet_table_3[value] - 1;
}
static int asn_PER_MAP_attrPs2_3_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_3)/sizeof(permitted_alphabet_code2value_3[0]))
		return -1;
	return permitted_alphabet_code2value_3[code];
}
static int
memb_attrPs3_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size == 2)
		 && !check_permitted_alphabet_4(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_attrPs3_4_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_4)/sizeof(permitted_alphabet_table_4[0]))
		return -1;
	return permitted_alphabet_table_4[value] - 1;
}
static int asn_PER_MAP_attrPs3_4_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_4)/sizeof(permitted_alphabet_code2value_4[0]))
		return -1;
	return permitted_alphabet_code2value_4[code];
}
static int
memb_attrPs4_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size <= 3)
		 && !check_permitted_alphabet_5(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_attrPs4_5_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_5)/sizeof(permitted_alphabet_table_5[0]))
		return -1;
	return permitted_alphabet_table_5[value] - 1;
}
static int asn_PER_MAP_attrPs4_5_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_5)/sizeof(permitted_alphabet_code2value_5[0]))
		return -1;
	return permitted_alphabet_code2value_5[code];
}
static int
memb_attrPs5_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size >= 2 && size <= 3)
		 && !check_permitted_alphabet_6(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_attrPs5_6_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_6)/sizeof(permitted_alphabet_table_6[0]))
		return -1;
	return permitted_alphabet_table_6[value] - 1;
}
static int asn_PER_MAP_attrPs5_6_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_6)/sizeof(permitted_alphabet_code2value_6[0]))
		return -1;
	return permitted_alphabet_code2value_6[code];
}
static int
memb_attrPs6_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size >= 1 && size <= 3)
		 && !check_permitted_alphabet_7(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_attrPs6_7_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_7)/sizeof(permitted_alphabet_table_7[0]))
		return -1;
	return permitted_alphabet_table_7[value] - 1;
}
static int asn_PER_MAP_attrPs6_7_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_7)/sizeof(permitted_alphabet_code2value_7[0]))
		return -1;
	return permitted_alphabet_code2value_7[code];
}
static int
memb_attrPs7_constraint_1(const asn_TYPE_descriptor_t *td, const void *sptr,
			asn_app_constraint_failed_f *ctfailcb, void *app_key) {
	const PrintableString_t *st = (const PrintableString_t *)sptr;
	size_t size;
	
	if(!sptr) {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: value not given (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
	
	size = st->size;
	
	if((size >= 2 && size <= 6)
		 && !check_permitted_alphabet_8(st)) {
		/* Constraint check succeeded */
		return 0;
	} else {
		ASN__CTFAIL(app_key, td, sptr,
			"%s: constraint failed (%s:%d)",
			td->name, __FILE__, __LINE__);
		return -1;
	}
}

static int asn_PER_MAP_attrPs7_8_v2c(unsigned int value) {
	if(value >= sizeof(permitted_alphabet_table_8)/sizeof(permitted_alphabet_table_8[0]))
		return -1;
	return permitted_alphabet_table_8[value] - 1;
}
static int asn_PER_MAP_attrPs7_8_c2v(unsigned int code) {
	if(code >= sizeof(permitted_alphabet_code2value_8)/sizeof(permitted_alphabet_code2value_8[0]))
		return -1;
	return permitted_alphabet_code2value_8[code];
}
static asn_per_constraints_t asn_PER_memb_attrPs2_constr_3 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED,	 0,  0,  2,  2 }	/* (SIZE(2..2)) */,
	asn_PER_MAP_attrPs2_3_v2c,	/* Value to PER code map */
	asn_PER_MAP_attrPs2_3_c2v	/* PER code to value map */
};
static asn_per_constraints_t asn_PER_memb_attrPs3_constr_4 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  0,  0,  2,  2 }	/* (SIZE(2..2,...)) */,
	asn_PER_MAP_attrPs3_4_v2c,	/* Value to PER code map */
	asn_PER_MAP_attrPs3_4_c2v	/* PER code to value map */
};
static asn_per_constraints_t asn_PER_memb_attrPs4_constr_5 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED,	 2,  2,  0,  3 }	/* (SIZE(0..3)) */,
	asn_PER_MAP_attrPs4_5_v2c,	/* Value to PER code map */
	asn_PER_MAP_attrPs4_5_c2v	/* PER code to value map */
};
static asn_per_constraints_t asn_PER_memb_attrPs5_constr_6 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED,	 1,  1,  2,  3 }	/* (SIZE(2..3)) */,
	asn_PER_MAP_attrPs5_6_v2c,	/* Value to PER code map */
	asn_PER_MAP_attrPs5_6_c2v	/* PER code to value map */
};
static asn_per_constraints_t asn_PER_memb_attrPs6_constr_7 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED | APC_EXTENSIBLE,  2,  2,  1,  3 }	/* (SIZE(1..3,...)) */,
	asn_PER_MAP_attrPs6_7_v2c,	/* Value to PER code map */
	asn_PER_MAP_attrPs6_7_c2v	/* PER code to value map */
};
static asn_per_constraints_t asn_PER_memb_attrPs7_constr_8 CC_NOTUSED = {
	{ APC_CONSTRAINED,	 7,  7,  32,  122 }	/* (32..122) */,
	{ APC_CONSTRAINED,	 3,  3,  2,  6 }	/* (SIZE(2..6)) */,
	asn_PER_MAP_attrPs7_8_v2c,	/* Value to PER code map */
	asn_PER_MAP_attrPs7_8_c2v	/* PER code to value map */
};
static asn_TYPE_member_t asn_MBR_TEST_PrintableString_1[] = {
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_PrintableString, attrPs1),
		(ASN_TAG_CLASS_CONTEXT | (0 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, 0, 0 },
		0, 0, /* No default value */
		"attrPs1"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_PrintableString, attrPs2),
		(ASN_TAG_CLASS_CONTEXT | (1 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, &asn_PER_memb_attrPs2_constr_3,  memb_attrPs2_constraint_1 },
		0, 0, /* No default value */
		"attrPs2"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_PrintableString, attrPs3),
		(ASN_TAG_CLASS_CONTEXT | (2 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, &asn_PER_memb_attrPs3_constr_4,  memb_attrPs3_constraint_1 },
		0, 0, /* No default value */
		"attrPs3"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_PrintableString, attrPs4),
		(ASN_TAG_CLASS_CONTEXT | (3 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, &asn_PER_memb_attrPs4_constr_5,  memb_attrPs4_constraint_1 },
		0, 0, /* No default value */
		"attrPs4"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_PrintableString, attrPs5),
		(ASN_TAG_CLASS_CONTEXT | (4 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, &asn_PER_memb_attrPs5_constr_6,  memb_attrPs5_constraint_1 },
		0, 0, /* No default value */
		"attrPs5"
		},
	{ ATF_NOFLAGS, 0, offsetof(struct TEST_PrintableString, attrPs6),
		(ASN_TAG_CLASS_CONTEXT | (5 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, &asn_PER_memb_attrPs6_constr_7,  memb_attrPs6_constraint_1 },
		0, 0, /* No default value */
		"attrPs6"
		},
	{ ATF_POINTER, 1, offsetof(struct TEST_PrintableString, attrPs7),
		(ASN_TAG_CLASS_CONTEXT | (6 << 2)),
		-1,	/* IMPLICIT tag at current level */
		&asn_DEF_PrintableString,
		0,
		{ 0, &asn_PER_memb_attrPs7_constr_8,  memb_attrPs7_constraint_1 },
		0, 0, /* No default value */
		"attrPs7"
		},
};
static const int asn_MAP_TEST_PrintableString_oms_1[] = { 6 };
static const ber_tlv_tag_t asn_DEF_TEST_PrintableString_tags_1[] = {
	(ASN_TAG_CLASS_UNIVERSAL | (16 << 2))
};
static const asn_TYPE_tag2member_t asn_MAP_TEST_PrintableString_tag2el_1[] = {
    { (ASN_TAG_CLASS_CONTEXT | (0 << 2)), 0, 0, 0 }, /* attrPs1 */
    { (ASN_TAG_CLASS_CONTEXT | (1 << 2)), 1, 0, 0 }, /* attrPs2 */
    { (ASN_TAG_CLASS_CONTEXT | (2 << 2)), 2, 0, 0 }, /* attrPs3 */
    { (ASN_TAG_CLASS_CONTEXT | (3 << 2)), 3, 0, 0 }, /* attrPs4 */
    { (ASN_TAG_CLASS_CONTEXT | (4 << 2)), 4, 0, 0 }, /* attrPs5 */
    { (ASN_TAG_CLASS_CONTEXT | (5 << 2)), 5, 0, 0 }, /* attrPs6 */
    { (ASN_TAG_CLASS_CONTEXT | (6 << 2)), 6, 0, 0 } /* attrPs7 */
};
static asn_SEQUENCE_specifics_t asn_SPC_TEST_PrintableString_specs_1 = {
	sizeof(struct TEST_PrintableString),
	offsetof(struct TEST_PrintableString, _asn_ctx),
	asn_MAP_TEST_PrintableString_tag2el_1,
	7,	/* Count of tags in the map */
	asn_MAP_TEST_PrintableString_oms_1,	/* Optional members */
	1, 0,	/* Root/Additions */
	-1,	/* First extension addition */
};
asn_TYPE_descriptor_t asn_DEF_TEST_PrintableString = {
	"TEST-PrintableString",
	"TEST-PrintableString",
	&asn_OP_SEQUENCE,
	asn_DEF_TEST_PrintableString_tags_1,
	sizeof(asn_DEF_TEST_PrintableString_tags_1)
		/sizeof(asn_DEF_TEST_PrintableString_tags_1[0]), /* 1 */
	asn_DEF_TEST_PrintableString_tags_1,	/* Same as above */
	sizeof(asn_DEF_TEST_PrintableString_tags_1)
		/sizeof(asn_DEF_TEST_PrintableString_tags_1[0]), /* 1 */
	{ 0, 0, SEQUENCE_constraint },
	asn_MBR_TEST_PrintableString_1,
	7,	/* Elements count */
	&asn_SPC_TEST_PrintableString_specs_1	/* Additional specs */
};

