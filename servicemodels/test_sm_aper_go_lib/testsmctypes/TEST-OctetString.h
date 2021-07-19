/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_TEST_OctetString_H_
#define	_TEST_OctetString_H_


#include "asn_application.h"

/* Including external dependencies */
#include "OCTET_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* TEST-OctetString */
typedef struct TEST_OctetString {
	OCTET_STRING_t	 attrOs1;
	OCTET_STRING_t	 attrOs2;
	OCTET_STRING_t	 attrOs3;
	OCTET_STRING_t	 attrOs4;
	OCTET_STRING_t	 attrOs5;
	OCTET_STRING_t	 attrOs6;
	OCTET_STRING_t	*attrOs7;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TEST_OctetString_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TEST_OctetString;

#ifdef __cplusplus
}
#endif

#endif	/* _TEST_OctetString_H_ */
#include "asn_internal.h"
