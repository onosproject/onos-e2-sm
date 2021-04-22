/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v2.asn1"
 */

#ifndef	_RC_PRE_Trigger_Type_H_
#define	_RC_PRE_Trigger_Type_H_


#include <asn_application.h>

/* Including external dependencies */
#include <NativeEnumerated.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum RC_PRE_Trigger_Type {
	RC_PRE_Trigger_Type_upon_change	= 0,
	RC_PRE_Trigger_Type_periodic	= 1
	/*
	 * Enumeration is extensible
	 */
} e_RC_PRE_Trigger_Type;

/* RC-PRE-Trigger-Type */
typedef long	 RC_PRE_Trigger_Type_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_RC_PRE_Trigger_Type_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_RC_PRE_Trigger_Type;
extern const asn_INTEGER_specifics_t asn_SPC_RC_PRE_Trigger_Type_specs_1;
asn_struct_free_f RC_PRE_Trigger_Type_free;
asn_struct_print_f RC_PRE_Trigger_Type_print;
asn_constr_check_f RC_PRE_Trigger_Type_constraint;
ber_type_decoder_f RC_PRE_Trigger_Type_decode_ber;
der_type_encoder_f RC_PRE_Trigger_Type_encode_der;
xer_type_decoder_f RC_PRE_Trigger_Type_decode_xer;
xer_type_encoder_f RC_PRE_Trigger_Type_encode_xer;
oer_type_decoder_f RC_PRE_Trigger_Type_decode_oer;
oer_type_encoder_f RC_PRE_Trigger_Type_encode_oer;
per_type_decoder_f RC_PRE_Trigger_Type_decode_uper;
per_type_encoder_f RC_PRE_Trigger_Type_encode_uper;
per_type_decoder_f RC_PRE_Trigger_Type_decode_aper;
per_type_encoder_f RC_PRE_Trigger_Type_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RC_PRE_Trigger_Type_H_ */
#include <asn_internal.h>
