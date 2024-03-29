/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_TEST_Extension2_H_
#define	_TEST_Extension2_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "OCTET_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* TEST-Extension2 */
typedef struct TEST_Extension2 {
	long	 item1;
	OCTET_STRING_t	*item2;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	OCTET_STRING_t	*ext1;	/* OPTIONAL */
	OCTET_STRING_t	*ext2;	/* OPTIONAL */
	OCTET_STRING_t	*ext3;	/* OPTIONAL */
	OCTET_STRING_t	*ext4;	/* OPTIONAL */
	OCTET_STRING_t	*ext5;	/* OPTIONAL */
	OCTET_STRING_t	*ext6;	/* OPTIONAL */
	OCTET_STRING_t	*ext7;	/* OPTIONAL */
	OCTET_STRING_t	*ext8;	/* OPTIONAL */
	OCTET_STRING_t	*ext9;	/* OPTIONAL */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} TEST_Extension2_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_TEST_Extension2;

#ifdef __cplusplus
}
#endif

#endif	/* _TEST_Extension2_H_ */
#include "asn_internal.h"
