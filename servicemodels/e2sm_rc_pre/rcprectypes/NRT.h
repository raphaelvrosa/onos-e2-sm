/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-RC-PRE-IEs"
 * 	found in "e2sm-rc-pre-v1.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D e2sm_rc_pre_v01_asn1/asn1c-gen`
 */

#ifndef	_NRT_H_
#define	_NRT_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "CellGlobalID.h"
#include "EARFCN.h"
#include "Cell-Size.h"
#include "PCI.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* NRT */
typedef struct NRT {
	long	 nrIndex;
	CellGlobalID_t	 cgi;
	EARFCN_t	 dl_EARFCN;
	Cell_Size_t	 cell_Size;
	PCI_t	 pci;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} NRT_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_NRT;
extern asn_SEQUENCE_specifics_t asn_SPC_NRT_specs_1;
extern asn_TYPE_member_t asn_MBR_NRT_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _NRT_H_ */
#include "asn_internal.h"
