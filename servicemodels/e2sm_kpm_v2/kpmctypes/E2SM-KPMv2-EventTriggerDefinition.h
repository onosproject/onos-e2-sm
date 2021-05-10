/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-changed.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_E2SM_KPMv2_EventTriggerDefinition_H_
#define	_E2SM_KPMv2_EventTriggerDefinition_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_CHOICE.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR {
	E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR_NOTHING,	/* No components present */
	E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR_eventDefinition_Format1
	/* Extensions may appear below */
	
} E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR;

/* Forward declarations */
struct E2SM_KPMv2_EventTriggerDefinition_Format1;

/* E2SM-KPMv2-EventTriggerDefinition */
typedef struct E2SM_KPMv2_EventTriggerDefinition {
	struct E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats {
		E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_PR present;
		union E2SM_KPMv2_EventTriggerDefinition__eventDefinition_formats_u {
			struct E2SM_KPMv2_EventTriggerDefinition_Format1	*eventDefinition_Format1;
			/*
			 * This type is extensible,
			 * possible extensions are below.
			 */
		} choice;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} eventDefinition_formats;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_KPMv2_EventTriggerDefinition_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_KPMv2_EventTriggerDefinition;

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_KPMv2_EventTriggerDefinition_H_ */
#include "asn_internal.h"
