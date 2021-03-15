/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.0.2-rm.asn"
 * 	`asn1c -pdu=all -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_SubscriptionID_H_
#define	_SubscriptionID_H_


#include "asn_application.h"

/* Including external dependencies */
#include "INTEGER.h"

#ifdef __cplusplus
extern "C" {
#endif

/* SubscriptionID */
typedef INTEGER_t	 SubscriptionID_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_SubscriptionID_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_SubscriptionID;
asn_struct_free_f SubscriptionID_free;
asn_struct_print_f SubscriptionID_print;
asn_constr_check_f SubscriptionID_constraint;
ber_type_decoder_f SubscriptionID_decode_ber;
der_type_encoder_f SubscriptionID_encode_der;
xer_type_decoder_f SubscriptionID_decode_xer;
xer_type_encoder_f SubscriptionID_encode_xer;
per_type_decoder_f SubscriptionID_decode_uper;
per_type_encoder_f SubscriptionID_encode_uper;
per_type_decoder_f SubscriptionID_decode_aper;
per_type_encoder_f SubscriptionID_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _SubscriptionID_H_ */
#include "asn_internal.h"
