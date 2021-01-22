/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v1.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D e2sm_rc_pre_v01_asn1/asn1c-gen`
 */

#ifndef	_CellGlobalID_H_
#define	_CellGlobalID_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum CellGlobalID_PR {
	CellGlobalID_PR_NOTHING,	/* No components present */
	CellGlobalID_PR_nr_CGI,
	CellGlobalID_PR_eUTRA_CGI
	/* Extensions may appear below */
	
} CellGlobalID_PR;

/* Forward declarations */
struct NRCGI;
struct EUTRACGI;

/* CellGlobalID */
typedef struct CellGlobalID {
	CellGlobalID_PR present;
	union CellGlobalID_u {
		struct NRCGI	*nr_CGI;
		struct EUTRACGI	*eUTRA_CGI;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} CellGlobalID_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_CellGlobalID;
extern asn_CHOICE_specifics_t asn_SPC_CellGlobalID_specs_1;
extern asn_TYPE_member_t asn_MBR_CellGlobalID_1[2];
extern asn_per_constraints_t asn_PER_type_CellGlobalID_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _CellGlobalID_H_ */
#include "asn_internal.h"
