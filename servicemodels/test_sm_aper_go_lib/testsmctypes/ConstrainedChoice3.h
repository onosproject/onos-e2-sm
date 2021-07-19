/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "TEST-SM-IEs"
 * 	found in "../v1/test_sm.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_ConstrainedChoice3_H_
#define	_ConstrainedChoice3_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum ConstrainedChoice3_PR {
	ConstrainedChoice3_PR_NOTHING,	/* No components present */
	ConstrainedChoice3_PR_constrainedChoice3A,
	ConstrainedChoice3_PR_constrainedChoice3B,
	ConstrainedChoice3_PR_constrainedChoice3C,
	ConstrainedChoice3_PR_constrainedChoice3D
} ConstrainedChoice3_PR;

/* ConstrainedChoice3 */
typedef struct ConstrainedChoice3 {
	ConstrainedChoice3_PR present;
	union ConstrainedChoice3_u {
		long	 constrainedChoice3A;
		long	 constrainedChoice3B;
		unsigned long	 constrainedChoice3C;
		long	 constrainedChoice3D;
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ConstrainedChoice3_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_constrainedChoice3C_4;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_ConstrainedChoice3;
extern asn_CHOICE_specifics_t asn_SPC_ConstrainedChoice3_specs_1;
extern asn_TYPE_member_t asn_MBR_ConstrainedChoice3_1[4];
extern asn_per_constraints_t asn_PER_type_ConstrainedChoice3_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _ConstrainedChoice3_H_ */
#include "asn_internal.h"
