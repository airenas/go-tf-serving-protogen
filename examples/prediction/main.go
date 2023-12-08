// Sample of making a call to prediction model served by tensorflow serving on localhost
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/airenas/go-tf-serving-protogen/tensorflow/core/framework"
	"github.com/airenas/go-tf-serving-protogen/tensorflow_serving/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	grpcAddress = "localhost:8500"
	modelName   = "punctuation"
	inputName   = "word_ids"
)

func main() {
	log.Println("Prepare GRPC connection")
	conn, err := grpc.Dial(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to TensorFlow Serving: %v", err)
	}
	defer conn.Close()

	log.Println("Prepare client")
	client := apis.NewPredictionServiceClient(conn)
	log.Println("Prepare request")
	request := makeRequest()

	log.Println("Make call")
	response, err := client.Predict(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to call TensorFlow Serving Predict service: %v", err)
	}

	log.Println("Yes! Got result")
	fmt.Println("Prediction Results:", response)
}

func makeRequest() *apis.PredictRequest {
	wordIdsData := []int32{4798, 63, 234, 500, 607, 576, 62, 1250, 666, 3124,
		13062, 207, 1173, 666, 4, 2950, 28668, 8442, 1954, 0, 175, 973,
		51, 426, 1661, 429, 51, 4, 139, 11, 697, 6123, 133, 7, 2506, 34, 34, 4,
		2266, 437, 1005, 152, 109, 810, 205, 1591, 107, 259, 7, 1}

	return &apis.PredictRequest{
		ModelSpec: &apis.ModelSpec{
			Name: modelName,
		},
		Inputs: map[string]*framework.TensorProto{
			inputName: {
				Dtype: framework.DataType_DT_INT32,
				TensorShape: &framework.TensorShapeProto{
					Dim: []*framework.TensorShapeProto_Dim{
						{Size: 1},
						{Size: int64(len(wordIdsData))},
					},
				},
				IntVal: wordIdsData,
			},
		},
	}
}
