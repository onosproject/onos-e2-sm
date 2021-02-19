/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../e2sm_kpm_ies_v2.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_MeasurementRecordItem_H_
#define	_MeasurementRecordItem_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeReal.h"
#include "NULL.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum MeasurementRecordItem_PR {
	MeasurementRecordItem_PR_NOTHING,	/* No components present */
	MeasurementRecordItem_PR_integer,
	MeasurementRecordItem_PR_real,
	MeasurementRecordItem_PR_noValue
	/* Extensions may appear below */
	
} MeasurementRecordItem_PR;

/* MeasurementRecordItem */
typedef struct MeasurementRecordItem {
	MeasurementRecordItem_PR present;
	union MeasurementRecordItem_u {
		long	 integer;
		double	 real;
		NULL_t	 noValue;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} MeasurementRecordItem_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_MeasurementRecordItem;
extern asn_CHOICE_specifics_t asn_SPC_MeasurementRecordItem_specs_1;
extern asn_TYPE_member_t asn_MBR_MeasurementRecordItem_1[3];
extern asn_per_constraints_t asn_PER_type_MeasurementRecordItem_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _MeasurementRecordItem_H_ */
#include "asn_internal.h"
