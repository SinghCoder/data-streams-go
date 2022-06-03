// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package kafka

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/DataDog/data-streams-go/datastreams"
)

// traceKafkaProduce appends the pathway in the context to the kafka message headers, and returns true if
// it is successful.
func traceKafkaProduce(ctx context.Context, msg kafka.Headers) bool {
	_, ctx = datastreams.SetCheckpoint(ctx, "internal")
	p, ok := datastreams.PathwayFromContext(ctx)
	if ok {
		msg.Headers = append(msg.Headers, kafka.Header{Key: datastreams.PropagationKey, Value: p.Encode()})
	}
	return ok
}