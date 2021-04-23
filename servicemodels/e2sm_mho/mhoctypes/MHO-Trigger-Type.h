/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-MHO-IEs"
 * 	found in "e2sm-mho.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_MHO_Trigger_Type_H_
#define	_MHO_Trigger_Type_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum MHO_Trigger_Type {
	MHO_Trigger_Type_periodic	= 0,
	MHO_Trigger_Type_upon_rcv_meas_report	= 1,
	MHO_Trigger_Type_upon_change_rrc_status	= 2
	/*
	 * Enumeration is extensible
	 */
} e_MHO_Trigger_Type;

/* MHO-Trigger-Type */
typedef long	 MHO_Trigger_Type_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_MHO_Trigger_Type_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_MHO_Trigger_Type;
extern const asn_INTEGER_specifics_t asn_SPC_MHO_Trigger_Type_specs_1;
asn_struct_free_f MHO_Trigger_Type_free;
asn_struct_print_f MHO_Trigger_Type_print;
asn_constr_check_f MHO_Trigger_Type_constraint;
ber_type_decoder_f MHO_Trigger_Type_decode_ber;
der_type_encoder_f MHO_Trigger_Type_encode_der;
xer_type_decoder_f MHO_Trigger_Type_decode_xer;
xer_type_encoder_f MHO_Trigger_Type_encode_xer;
per_type_decoder_f MHO_Trigger_Type_decode_uper;
per_type_encoder_f MHO_Trigger_Type_encode_uper;
per_type_decoder_f MHO_Trigger_Type_decode_aper;
per_type_encoder_f MHO_Trigger_Type_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _MHO_Trigger_Type_H_ */
#include "asn_internal.h"
