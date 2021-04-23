/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-MHO-IEs"
 * 	found in "e2sm-mho.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2SM_MHO_ControlMessage_Format1_H_
#define	_E2SM_MHO_ControlMessage_Format1_H_


#include "asn_application.h"

/* Including external dependencies */
#include "CellGlobalID.h"
#include "UE-Identity.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* E2SM-MHO-ControlMessage-Format1 */
typedef struct E2SM_MHO_ControlMessage_Format1 {
	CellGlobalID_t	 serving_cgi;
	UE_Identity_t	 uedID;
	CellGlobalID_t	 target_cgi;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_MHO_ControlMessage_Format1_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_MHO_ControlMessage_Format1;
extern asn_SEQUENCE_specifics_t asn_SPC_E2SM_MHO_ControlMessage_Format1_specs_1;
extern asn_TYPE_member_t asn_MBR_E2SM_MHO_ControlMessage_Format1_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_MHO_ControlMessage_Format1_H_ */
#include "asn_internal.h"
