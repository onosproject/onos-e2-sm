/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "../v2/e2sm-rc-pre_v2_rsys.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_Cell_Size_H_
#define	_Cell_Size_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum Cell_Size {
	Cell_Size_femto	= 0,
	Cell_Size_enterprise	= 1,
	Cell_Size_outdoorSmall	= 2,
	Cell_Size_macro	= 3
	/*
	 * Enumeration is extensible
	 */
} e_Cell_Size;

/* Cell-Size */
typedef long	 Cell_Size_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_Cell_Size_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_Cell_Size;
extern const asn_INTEGER_specifics_t asn_SPC_Cell_Size_specs_1;
asn_struct_free_f Cell_Size_free;
asn_struct_print_f Cell_Size_print;
asn_constr_check_f Cell_Size_constraint;
ber_type_decoder_f Cell_Size_decode_ber;
der_type_encoder_f Cell_Size_encode_der;
xer_type_decoder_f Cell_Size_decode_xer;
xer_type_encoder_f Cell_Size_encode_xer;
per_type_decoder_f Cell_Size_decode_uper;
per_type_encoder_f Cell_Size_encode_uper;
per_type_decoder_f Cell_Size_decode_aper;
per_type_encoder_f Cell_Size_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _Cell_Size_H_ */
#include "asn_internal.h"
