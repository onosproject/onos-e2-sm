/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.3-changed.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RIC_KPMNode_Item_H_
#define	_RIC_KPMNode_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "GlobalKPMv2node-ID.h"
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct Cell_Measurement_Object_Item;

/* RIC-KPMNode-Item */
typedef struct RIC_KPMNode_Item {
	GlobalKPMv2node_ID_t	 ric_KPMNode_Type;
	struct RIC_KPMNode_Item__cell_Measurement_Object_List {
		A_SEQUENCE_OF(struct Cell_Measurement_Object_Item) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} *cell_Measurement_Object_List;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RIC_KPMNode_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RIC_KPMNode_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_RIC_KPMNode_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_RIC_KPMNode_Item_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _RIC_KPMNode_Item_H_ */
#include "asn_internal.h"
