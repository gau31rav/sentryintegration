package stream_integrations

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

var (
	stream = flag.String("stream", "tet", "tet")
	region = flag.String("region", "us-west-1", "us-west-1")
)

func PutRecord() {
	flag.Parse()

	s := session.New(&aws.Config{Region: aws.String(*region), Credentials: credentials.NewStaticCredentials("", "", "")})
	kc := kinesis.New(s)

	streamName := aws.String(*stream)

	putOutput, err := kc.PutRecord(&kinesis.PutRecordInput{
		Data:         []byte("hoge"),
		StreamName:   streamName,
		PartitionKey: aws.String("key1"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", putOutput)

	// put 10 records using PutRecords API
	entries := make([]*kinesis.PutRecordsRequestEntry, 10)
	for i := 0; i < len(entries); i++ {
		entries[i] = &kinesis.PutRecordsRequestEntry{
			Data:         []byte(fmt.Sprintf("hoge%d", i)),
			PartitionKey: aws.String("key2"),
		}
	}
	fmt.Printf("%v\n", entries)
	putsOutput, err := kc.PutRecords(&kinesis.PutRecordsInput{
		Records:    entries,
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	// putsOutput has Records, and its shard id and sequece enumber.
	fmt.Printf("%v\n", putsOutput)

	// retrieve iterator
	iteratorOutput, err := kc.GetShardIterator(&kinesis.GetShardIteratorInput{
		// Shard Id is provided when making put record(s) request.
		ShardId:           putOutput.ShardId,
		ShardIteratorType: aws.String("TRIM_HORIZON"),
		// ShardIteratorType: aws.String("AT_SEQUENCE_NUMBER"),
		// ShardIteratorType: aws.String("LATEST"),
		StreamName: streamName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", iteratorOutput)

	// get records use shard iterator for making request
	records, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: iteratorOutput.ShardIterator,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", records)

	// and, you can iteratively make GetRecords request using records.NextShardIterator
	recordsSecond, err := kc.GetRecords(&kinesis.GetRecordsInput{
		ShardIterator: records.NextShardIterator,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", recordsSecond)

}
