/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "../v2/e2sm-rc-pre_v2_rsys.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RANfunction_Name_H_
#define	_RANfunction_Name_H_


#include "asn_application.h"

/* Including external dependencies */
#include "PrintableString.h"
#include "NativeInteger.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* RANfunction-Name */
typedef struct RANfunction_Name {
	PrintableString_t	 ranFunction_ShortName;
	PrintableString_t	 ranFunction_E2SM_OID;
	PrintableString_t	 ranFunction_Description;
	long	*ranFunction_Instance;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} RANfunction_Name_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_RANfunction_Name;
extern asn_SEQUENCE_specifics_t asn_SPC_RANfunction_Name_specs_1;
extern asn_TYPE_member_t asn_MBR_RANfunction_Name_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _RANfunction_Name_H_ */
#include "asn_internal.h"
