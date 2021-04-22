/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v2.asn1"
 */

#ifndef	_RANparameterDef_Item_H_
#define	_RANparameterDef_Item_H_


#include <asn_application.h>

/* Including external dependencies */
#include "RANparameter-ID.h"
#include "RANparameter-Name.h"
#include "RANparameter-Type.h"
#include <constr_SEQUENCE.h>

#ifdef __cplusplus
extern "C" {
#endif

/* RANparameterDef-Item */
typedef struct RANparameterDef_Item {
	RANparameter_ID_t	 ranParameter_ID;
	RANparameter_Name_t	 ranParameter_Name;
	RANparameter_Type_t	 ranParameter_Type;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RANparameterDef_Item_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RANparameterDef_Item;
extern asn_SEQUENCE_specifics_t asn_SPC_RANparameterDef_Item_specs_1;
extern asn_TYPE_member_t asn_MBR_RANparameterDef_Item_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _RANparameterDef_Item_H_ */
#include <asn_internal.h>
