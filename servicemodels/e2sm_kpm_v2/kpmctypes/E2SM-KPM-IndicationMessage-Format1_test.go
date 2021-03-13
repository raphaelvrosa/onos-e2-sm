// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/pdubuilder"
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createE2SMKPMIndicationMessageFormat1() *e2sm_kpm_v2.E2SmKpmIndicationMessageFormat1 {

	var integer int64 = 12345
	var rl float64 = 6789
	var cellObjID string = "onf"
	var granularity int32 = 21
	var subscriptionID int64 = 12345
	plmnID := []byte{0x21, 0x22, 0x23}
	sst := []byte{0x01}
	sd := []byte{0x01, 0x02, 0x03}
	var fiveQI int32 = 10
	var qci int32 = 15
	var qciMin int32 = 1
	var qciMax int32 = 15
	var arpMax int32 = 15
	var arpMin int32 = 1
	var bitrateRange int32 = 251
	var layerMuMimo int32 = 5
	var distX int32 = 123
	var distY int32 = 456
	var distZ int32 = 789
	startEndIndication := e2sm_kpm_v2.StartEndInd_START_END_IND_START
	var measurementName string = "trial"

	labelInfoItem, _ := pdubuilder.CreateLabelInfoItem(plmnID, sst, sd, fiveQI,
		qci, qciMax, qciMin, arpMax, arpMin, bitrateRange, layerMuMimo,
		distX, distY, distZ, startEndIndication)

	labelInfoList := e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}
	labelInfoList.Value = append(labelInfoList.Value, labelInfoItem)

	measName, _ := pdubuilder.CreateMeasurementTypeMeasName(measurementName)
	measInfoItem, _ := pdubuilder.CreateMeasurementInfoItem(measName, &labelInfoList)

	measInfoList := e2sm_kpm_v2.MeasurementInfoList{
		Value: make([]*e2sm_kpm_v2.MeasurementInfoItem, 0),
	}
	measInfoList.Value = append(measInfoList.Value, measInfoItem)

	measRecord := e2sm_kpm_v2.MeasurementRecord{
		Value: make([]*e2sm_kpm_v2.MeasurementRecordItem, 0),
	}
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemInteger(integer))
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemNoValue())
	measRecord.Value = append(measRecord.Value, pdubuilder.CreateMeasurementRecordItemReal(rl))

	measData := e2sm_kpm_v2.MeasurementData{
		Value: make([]*e2sm_kpm_v2.MeasurementRecord, 0),
	}
	measData.Value = append(measData.Value, &measRecord)

	newE2SmKpmPdu, _ := pdubuilder.CreateE2SmKpmIndicationMessage(subscriptionID, cellObjID, granularity, &measInfoList, &measData)

	return newE2SmKpmPdu.GetIndicationMessageFormat1()
}

func Test_xerEncodeE2SmKpmIndicationMessageFormat1(t *testing.T) {

	imf1 := createE2SMKPMIndicationMessageFormat1()

	xer, err := xerEncodeE2SmKpmIndicationMessageFormat1(imf1)
	assert.NilError(t, err)
	assert.Equal(t, 210, len(xer))
	t.Logf("E2SmKpmIndicationMessageFormat1 XER\n%s", string(xer))
}

func Test_xerDecodeE2SmKpmIndicationMessageFormat1(t *testing.T) {

	imf1 := createE2SMKPMIndicationMessageFormat1()

	xer, err := xerEncodeE2SmKpmIndicationMessageFormat1(imf1)
	assert.NilError(t, err)
	assert.Equal(t, 210, len(xer))
	t.Logf("E2SmKpmIndicationMessageFormat1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmKpmIndicationMessageFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmKpmIndicationMessageFormat1 XER - decoded\n%s", result)
}

func Test_perEncodeE2SmKpmIndicationMessageFormat1(t *testing.T) {

	imf1 := createE2SMKPMIndicationMessageFormat1()

	per, err := perEncodeE2SmKpmIndicationMessageFormat1(imf1)
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("E2SmKpmIndicationMessageFormat1 PER\n%s", string(per))
}

func Test_perDecodeE2SmKpmIndicationMessageFormat1(t *testing.T) {

	imf1 := createE2SMKPMIndicationMessageFormat1()

	per, err := perEncodeE2SmKpmIndicationMessageFormat1(imf1)
	assert.NilError(t, err)
	assert.Equal(t, 12, len(per))
	t.Logf("E2SmKpmIndicationMessageFormat1 PER\n%s", string(per))

	//result, err := perDecodeE2SmKpmIndicationMessageFormat1(per)
	//assert.NilError(t, err)
	//assert.Assert(t, result != nil)
	//t.Logf("E2SmKpmIndicationMessageFormat1 PER - decoded\n%s", result)
}