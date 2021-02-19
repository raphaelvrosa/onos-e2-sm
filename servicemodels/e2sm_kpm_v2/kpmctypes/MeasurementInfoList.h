/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-KPM-IEs"
 * 	found in "../v2/e2sm_kpm_v2.asn"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_MeasurementInfoList_H_
#define	_MeasurementInfoList_H_


#include "asn_application.h"

/* Including external dependencies */
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct MeasurementInfoItem;

/* MeasurementInfoList */
typedef struct MeasurementInfoList {
	A_SEQUENCE_OF(struct MeasurementInfoItem) list;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} MeasurementInfoList_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_MeasurementInfoList;
extern asn_SET_OF_specifics_t asn_SPC_MeasurementInfoList_specs_1;
extern asn_TYPE_member_t asn_MBR_MeasurementInfoList_1[1];
extern asn_per_constraints_t asn_PER_type_MeasurementInfoList_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _MeasurementInfoList_H_ */
#include "asn_internal.h"
