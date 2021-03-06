//
//   Copyright © 2019 Uncharted Software Inc.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package description

import (
	"io/ioutil"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/uncharted-distil/distil-compute/model"
)

func TestCreateUserDatasetPipeline(t *testing.T) {

	variables := []*model.Variable{
		{
			Name:         "test_var_0",
			OriginalType: "ordinal",
			Type:         "categorical",
			Index:        0,
		},
		{
			Name:         "test_var_1",
			OriginalType: "categorical",
			Type:         "integer",
			Index:        1,
		},
		{
			Name:         "test_var_2",
			OriginalType: "categorical",
			Type:         "integer",
			Index:        2,
		},
		{
			Name:         "test_var_3",
			OriginalType: "categorical",
			Type:         "integer",
			Index:        3,
		},
	}

	pipeline, err := CreateUserDatasetPipeline(
		"test_user_pipeline", "a test user pipeline", variables, "test_target", []string{"test_var_0", "test_var_1", "test_var_3"}, nil)

	// assert 1st is a semantic type update
	hyperParams := pipeline.GetSteps()[0].GetPrimitive().GetHyperparams()
	assert.Equal(t, []int64{1, 3}, ConvertToIntArray(hyperParams["add_columns"].GetValue().GetData().GetRaw().GetList()))
	assert.Equal(t, []string{"http://schema.org/Integer"}, ConvertToStringArray(hyperParams["add_types"].GetValue().GetData().GetRaw().GetList()))
	assert.Equal(t, []int64{}, ConvertToIntArray(hyperParams["remove_columns"].GetValue().GetData().GetRaw().GetList()))
	assert.Equal(t, []string{""}, ConvertToStringArray(hyperParams["remove_types"].GetValue().GetData().GetRaw().GetList()))

	// assert 2nd is a semantic type update
	hyperParams = pipeline.GetSteps()[1].GetPrimitive().GetHyperparams()
	assert.Equal(t, []int64{}, ConvertToIntArray(hyperParams["add_columns"].GetValue().GetData().GetRaw().GetList()))
	assert.Equal(t, []string{""}, ConvertToStringArray(hyperParams["add_types"].GetValue().GetData().GetRaw().GetList()))
	assert.Equal(t, []int64{1, 3}, ConvertToIntArray(hyperParams["remove_columns"].GetValue().GetData().GetRaw().GetList()))
	assert.Equal(t, []string{"https://metadata.datadrivendiscovery.org/types/CategoricalData"},
		ConvertToStringArray(hyperParams["remove_types"].GetValue().GetData().GetRaw().GetList()))

	// assert 3rd step is column remove and index two was remove
	hyperParams = pipeline.GetSteps()[2].GetPrimitive().GetHyperparams()
	assert.Equal(t, "0", hyperParams["resource_id"].GetValue().GetData().GetRaw().GetString_(), "0")
	assert.Equal(t, []int64{2}, ConvertToIntArray(hyperParams["columns"].GetValue().GetData().GetRaw().GetList()))

	assert.NoError(t, err)
}

func TestCreateUserDatasetPipelineMappingError(t *testing.T) {

	variables := []*model.Variable{
		{
			Name:         "test_var_0",
			OriginalType: "blordinal",
			Type:         "categorical",
			Index:        0,
		},
	}

	_, err := CreateUserDatasetPipeline(
		"test_user_pipeline", "a test user pipeline", variables, "test_target", []string{"test_var_0"}, nil)
	assert.Error(t, err)
}

func TestCreateUserDatasetEmpty(t *testing.T) {

	variables := []*model.Variable{
		{
			Name:         "test_var_0",
			OriginalType: "categorical",
			Type:         "categorical",
			Index:        0,
		},
	}

	pipeline, err := CreateUserDatasetPipeline(
		"test_user_pipeline", "a test user pipeline", variables, "test_target", []string{"test_var_0"}, nil)

	assert.Nil(t, pipeline)
	assert.Nil(t, err)
}

func TestCreatePCAFeaturesPipeline(t *testing.T) {
	pipeline, err := CreatePCAFeaturesPipeline("pca_features_test", "test pca feature ranking pipeline")
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/pca_features.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateSimonPipeline(t *testing.T) {
	pipeline, err := CreateSimonPipeline("simon_test", "test simon classification pipeline")
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/simon.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateCrocPipeline(t *testing.T) {
	pipeline, err := CreateCrocPipeline("croc_test", "test croc object detection pipeline", []string{"filename"}, []string{"objects"})
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/croc.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateDataCleaningPipeline(t *testing.T) {
	pipeline, err := CreateDataCleaningPipeline("data cleaning test", "test data cleaning pipeline")
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/datacleaning.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateUnicornPipeline(t *testing.T) {
	pipeline, err := CreateUnicornPipeline("unicorn test", "test unicorn image detection pipeline", []string{"filename"}, []string{"objects"})
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/unicorn.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateSlothPipeline(t *testing.T) {
	timeSeriesVariables := []*model.Variable{
		{Name: "time", Index: 0},
		{Name: "value", Index: 1},
	}

	pipeline, err := CreateSlothPipeline("sloth_test", "test sloth object detection pipeline", "time", "value", timeSeriesVariables)
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/sloth.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateDukePipeline(t *testing.T) {
	pipeline, err := CreateDukePipeline("duke_test", "test duke data summary pipeline")
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/duke.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateTargetRankingPipeline(t *testing.T) {
	vars := []*model.Variable{
		{
			Name:         "hall_of_fame",
			Index:        18,
			Type:         model.CategoricalType,
			OriginalType: model.CategoricalType,
		},
	}
	pipeline, err := CreateTargetRankingPipeline("target_ranking_test", "test target_ranking pipeline", "hall_of_fame", vars)
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/target_ranking.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateGoatForwardPipeline(t *testing.T) {
	pipeline, err := CreateGoatForwardPipeline("goat_forward_test", "test goat forward geocoding pipeline", "region")
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/goat_forward.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateGoatReversePipeline(t *testing.T) {
	pipeline, err := CreateGoatReversePipeline("goat_reverse_test", "test goat reverse geocoding pipeline", "lat", "lon")
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/goat_reverse.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateJoinPipeline(t *testing.T) {
	pipeline, err := CreateJoinPipeline("join_test", "test join pipeline", "Doubles", "horsepower", 0.8)
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/join.pln", data, 0644)
	assert.NoError(t, err)
}

func TestCreateTimeseriesFormatterPipeline(t *testing.T) {
	pipeline, err := CreateTimeseriesFormatterPipeline("formatter_test", "test formatter pipeline", "learningData", 1)
	assert.NoError(t, err)

	data, err := proto.Marshal(pipeline)
	assert.NoError(t, err)
	assert.NotNil(t, data)

	err = ioutil.WriteFile("/tmp/formatter.pln", data, 0644)
	assert.NoError(t, err)
}
