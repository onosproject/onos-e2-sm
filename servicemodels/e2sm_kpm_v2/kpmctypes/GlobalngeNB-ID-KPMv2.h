/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-changed.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_GlobalngeNB_ID_KPMv2_H_
#define	_GlobalngeNB_ID_KPMv2_H_


#include "asn_application.h"

/* Including external dependencies */
#include "PLMN-Identity-KPMv2.h"
#include "ENB-ID-Choice-KPMv2.h"
#include "BIT_STRING.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* GlobalngeNB-ID-KPMv2 */
typedef struct GlobalngeNB_ID_KPMv2 {
	PLMN_Identity_KPMv2_t	 plmn_id;
	ENB_ID_Choice_KPMv2_t	 enb_id;
	BIT_STRING_t	 short_Macro_eNB_ID;
	BIT_STRING_t	 long_Macro_eNB_ID;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} GlobalngeNB_ID_KPMv2_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_GlobalngeNB_ID_KPMv2;
extern asn_SEQUENCE_specifics_t asn_SPC_GlobalngeNB_ID_KPMv2_specs_1;
extern asn_TYPE_member_t asn_MBR_GlobalngeNB_ID_KPMv2_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _GlobalngeNB_ID_KPMv2_H_ */
#include "asn_internal.h"
