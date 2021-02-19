/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../e2sm_kpm_ies_v2.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2SM_KPM_ActionDefinition_H_
#define	_E2SM_KPM_ActionDefinition_H_


#include "asn_application.h"

/* Including external dependencies */
#include "RIC-Style-Type.h"
#include "constr_CHOICE.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum E2SM_KPM_ActionDefinition__actionDefinition_formats_PR {
	E2SM_KPM_ActionDefinition__actionDefinition_formats_PR_NOTHING,	/* No components present */
	E2SM_KPM_ActionDefinition__actionDefinition_formats_PR_actionDefinition_Format1
	/* Extensions may appear below */
	
} E2SM_KPM_ActionDefinition__actionDefinition_formats_PR;

/* Forward declarations */
struct E2SM_KPM_ActionDefinition_Format1;

/* E2SM-KPM-ActionDefinition */
typedef struct E2SM_KPM_ActionDefinition {
	RIC_Style_Type_t	 ric_Style_Type;
	struct E2SM_KPM_ActionDefinition__actionDefinition_formats {
		E2SM_KPM_ActionDefinition__actionDefinition_formats_PR present;
		union E2SM_KPM_ActionDefinition__actionDefinition_formats_u {
			struct E2SM_KPM_ActionDefinition_Format1	*actionDefinition_Format1;
			/*
			 * This type is extensible,
			 * possible extensions are below.
			 */
		} choice;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} actionDefinition_formats;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_KPM_ActionDefinition_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_KPM_ActionDefinition;

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_KPM_ActionDefinition_H_ */
#include "asn_internal.h"
