/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_TEST_ListExtensible2_H_
#define	_TEST_ListExtensible2_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct ItemExtensible;

/* TEST-ListExtensible2 */
typedef struct TEST_ListExtensible2 {
	A_SEQUENCE_OF(struct ItemExtensible) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TEST_ListExtensible2_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TEST_ListExtensible2;

#ifdef __cplusplus
}
#endif

#endif	/* _TEST_ListExtensible2_H_ */
#include "asn_internal.h"
