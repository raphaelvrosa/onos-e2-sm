// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2

import (
	"encoding/hex"
	e2sm_kpm_v2_go "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/onosproject/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasCI = "00000000  20 00 7a 00 01 42 10 01  15 1f ff f0 01 02 03 40  | .z..B.........@|\n" +
	"00000010  40 01 02 03 00 17 68 18  00 1e 00 01 70 00 00 18  |@.....h.....p...|\n" +
	"00000020  00 00 00 00 00 7a 00 01  c7 00 03 14 20           |.....z...... |"

func createMeasurementCondItem() (*e2sm_kpm_v2_go.MeasurementCondItem, error) {

	measCondItem := &e2sm_kpm_v2_go.MeasurementCondItem{
		MeasType: &e2sm_kpm_v2_go.MeasurementType{
			MeasurementType: &e2sm_kpm_v2_go.MeasurementType_MeasId{
				MeasId: &e2sm_kpm_v2_go.MeasurementTypeId{
					Value: 123,
				},
			},
		},
		MatchingCond: &e2sm_kpm_v2_go.MatchingCondList{
			Value: make([]*e2sm_kpm_v2_go.MatchingCondItem, 0),
		},
	}

	mci1 := &e2sm_kpm_v2_go.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2_go.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2sm_kpm_v2_go.TestCondInfo{
				TestType: &e2sm_kpm_v2_go.TestCondType{
					TestCondType: &e2sm_kpm_v2_go.TestCondType_AMbr{
						AMbr: e2sm_kpm_v2_go.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2sm_kpm_v2_go.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2sm_kpm_v2_go.TestCondValue{
					TestCondValue: &e2sm_kpm_v2_go.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci1)

	var br int32 = 25
	var lmm int32 = 1
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	sum := e2sm_kpm_v2_go.SUM_SUM_TRUE
	plo := e2sm_kpm_v2_go.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2sm_kpm_v2_go.StartEndInd_START_END_IND_END

	mci2 := &e2sm_kpm_v2_go.MatchingCondItem{
		MatchingCondItem: &e2sm_kpm_v2_go.MatchingCondItem_MeasLabel{
			MeasLabel: &e2sm_kpm_v2_go.MeasurementLabel{
				PlmnId: &e2sm_kpm_v2_go.PlmnIdentity{
					Value: []byte{0x1, 0x2, 0x3},
				},
				SliceId: &e2sm_kpm_v2_go.Snssai{
					SD:  []byte{0x01, 0x02, 0x03},
					SSt: []byte{0x01},
				},
				FiveQi: &e2sm_kpm_v2_go.FiveQi{
					Value: 23,
				},
				QFi: &e2sm_kpm_v2_go.Qfi{
					Value: 52,
				},
				QCi: &e2sm_kpm_v2_go.Qci{
					Value: 24,
				},
				QCimax: &e2sm_kpm_v2_go.Qci{
					Value: 30,
				},
				QCimin: &e2sm_kpm_v2_go.Qci{
					Value: 1,
				},
				ARpmax: &e2sm_kpm_v2_go.Arp{
					Value: 15,
				},
				ARpmin: &e2sm_kpm_v2_go.Arp{
					Value: 1,
				},
				BitrateRange:     &br,
				LayerMuMimo:      &lmm,
				SUm:              &sum,
				DistBinX:         &dbx,
				DistBinY:         &dby,
				DistBinZ:         &dbz,
				PreLabelOverride: &plo,
				StartEndInd:      &seind,
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci2)

	//if err := measCondItem.Validate(); err != nil {
	//	return nil, err
	//}
	return measCondItem, nil
}

func Test_perEncodeMeasurementCondItem(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementCondItem PER\n%v", hex.Dump(per))

	result := e2sm_kpm_v2_go.MeasurementCondItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt")
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementCondItem PER - decoded\n%v", &result)
	assert.Equal(t, mci.GetMeasType().GetMeasId().GetValue(), result.GetMeasType().GetMeasId().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueInt(), result.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueInt())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetAMbr().Number(), result.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetAMbr().Number())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number(), result.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number())
	assert.DeepEqual(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPlmnId().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSD(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSSt(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetFiveQi().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQFi().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCi().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimax().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimin().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmax().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmin().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetBitrateRange(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetLayerMuMimo(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSUm().Number(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinX(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinX())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinY(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinY())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinZ(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPreLabelOverride().Number(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetStartEndInd().Number(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetStartEndInd().Number())
}

func Test_perMeasurementCondItemCompareBytes(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	aper.ChoiceMap = e2sm_kpm_v2_go.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt")
	assert.NilError(t, err)
	t.Logf("MeasurementCondItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasCI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}