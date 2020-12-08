// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package pdudecoder

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm/kpmctypes"
	"gotest.tools/assert"
	"io/ioutil"
	"testing"
)

func Test_DecodeE2SmKpmRanfunctionDescription(t *testing.T) {
	e2smKpmRanfunctionDescription, err := ioutil.ReadFile("../test/E2SM-KPM-RANfunction-Description.xml")
	assert.NilError(t, err, "Unexpected error when loading file")
	e2SmKpmPdu, err := kpmctypes.XerDecodeE2SmKpmRanfunctionDescription(e2smKpmRanfunctionDescription)
	assert.NilError(t, err)

	ranFunctionName, ricEventList, ricReportList, err := DecodeE2SmKpmRanfunctionDescription(e2SmKpmPdu)
	assert.NilError(t, err)
	//assert.Assert(t, ranFunctionName != nil)
	assert.Equal(t, "OID123", string(ranFunctionName.RanFunctionE2SmOid))
	//assert.Assert(t, ricEventList != nil)
	assert.Equal(t, 1, len(*ricEventList))
	//assert.Assert(t, ricReportList != nil)
	assert.Equal(t, 6, len(*ricReportList))
}

func Test_DecodeE2SmKpmRanfunctionDescriptionBytes(t *testing.T) {
	rfd, err := kpmctypes.PerDecodeE2SmKpmRanfunctionDescription(ranFuncDescBytes)
	assert.NilError(t, err)
	err = rfd.Validate()
	assert.NilError(t, err)
	assert.Equal(t, "ORAN-E2SM-KPM", rfd.GetRanFunctionName().GetRanFunctionShortName())
	assert.Equal(t, "KPM monitor", rfd.GetRanFunctionName().GetRanFunctionDescription())
	assert.Equal(t, "OID123", rfd.GetRanFunctionName().GetRanFunctionE2SmOid())
	assert.Equal(t, -102650112, int(rfd.GetRanFunctionName().GetRanFunctionInstance()))

	assert.Equal(t, 6, len(rfd.GetE2SmKpmRanfunctionItem().GetRicReportStyleList()))
	assert.Equal(t, 1, len(rfd.GetE2SmKpmRanfunctionItem().GetRicEventTriggerStyleList()))
}

var ranFuncDescBytes = []byte{
	0x20, 0xC0, 0x4F, 0x52, 0x41, 0x4E, 0x2D, 0x45, 0x32, 0x53, 0x4D, 0x2D, 0x4B, 0x50, 0x4D, 0x00, 0x00, 0x05, 0x4F, 0x49,
	0x44, 0x31, 0x32, 0x33, 0x05, 0x00, 0x4B, 0x50, 0x4D, 0x20, 0x6D, 0x6F, 0x6E, 0x69, 0x74, 0x6F, 0x72, 0x08, 0x93, 0x49,
	0xF4, 0x77, 0xF9, 0xE1, 0xAF, 0x00, 0x60, 0x00, 0x01, 0x01, 0x07, 0x00, 0x50, 0x65, 0x72, 0x69, 0x6F, 0x64, 0x69, 0x63,
	0x20, 0x72, 0x65, 0x70, 0x6F, 0x72, 0x74, 0x01, 0x05, 0x14, 0x01, 0x01, 0x1D, 0x00, 0x4F, 0x2D, 0x44, 0x55, 0x20, 0x4D,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x20, 0x43, 0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x65, 0x72,
	0x20, 0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x35, 0x47, 0x43, 0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F, 0x79, 0x6D, 0x65, 0x6E, 0x74, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x02,
	0x1D, 0x00, 0x4F, 0x2D, 0x44, 0x55, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x20, 0x43,
	0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x65, 0x72, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45, 0x50, 0x43,
	0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F, 0x79, 0x6D, 0x65, 0x6E,
	0x74, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x03, 0x1E, 0x80, 0x4F, 0x2D, 0x43, 0x55, 0x2D, 0x43, 0x50, 0x20, 0x4D, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x20, 0x43, 0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x65, 0x72, 0x20,
	0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x35, 0x47, 0x43, 0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F, 0x79, 0x6D, 0x65, 0x6E, 0x74, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x04, 0x1E,
	0x80, 0x4F, 0x2D, 0x43, 0x55, 0x2D, 0x43, 0x50, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74,
	0x20, 0x43, 0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x65, 0x72, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x45,
	0x50, 0x43, 0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F, 0x79, 0x6D,
	0x65, 0x6E, 0x74, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x05, 0x1E, 0x80, 0x4F, 0x2D, 0x43, 0x55, 0x2D, 0x55, 0x50, 0x20,
	0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65, 0x6E, 0x74, 0x20, 0x43, 0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x65,
	0x72, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x35, 0x47, 0x43, 0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F, 0x79, 0x6D, 0x65, 0x6E, 0x74, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01,
	0x06, 0x1E, 0x80, 0x4F, 0x2D, 0x43, 0x55, 0x2D, 0x55, 0x50, 0x20, 0x4D, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6D, 0x65,
	0x6E, 0x74, 0x20, 0x43, 0x6F, 0x6E, 0x74, 0x61, 0x69, 0x6E, 0x65, 0x72, 0x20, 0x66, 0x6F, 0x72, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x45, 0x50, 0x43, 0x20, 0x63, 0x6F, 0x6E, 0x6E, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x64, 0x65, 0x70, 0x6C, 0x6F,
	0x79, 0x6D, 0x65, 0x6E, 0x74, 0x01, 0x01, 0x01, 0x01}
