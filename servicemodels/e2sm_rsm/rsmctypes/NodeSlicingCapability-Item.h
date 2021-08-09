/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RSM-IEs"
 * 	found in "../v1/e2sm-rsm.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_NodeSlicingCapability_Item_H_
#define	_NodeSlicingCapability_Item_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "SlicingType.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* NodeSlicingCapability-Item */
typedef struct NodeSlicingCapability_Item {
	unsigned long	 maxNumberOfSlicesDL;
	unsigned long	 maxNumberOfSlicesUL;
	SlicingType_t	 slicingType;
	long	 maxNumberOfUEsPerSlice;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} NodeSlicingCapability_Item_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_maxNumberOfSlicesDL_2;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_maxNumberOfSlicesUL_3;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_NodeSlicingCapability_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_NodeSlicingCapability_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_NodeSlicingCapability_Item_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _NodeSlicingCapability_Item_H_ */
#include "asn_internal.h"
