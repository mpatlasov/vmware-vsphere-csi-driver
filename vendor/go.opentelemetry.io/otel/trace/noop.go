// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package trace // import "go.opentelemetry.io/otel/trace"

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace/embedded"
)

// NewNoopTracerProvider returns an implementation of TracerProvider that
// performs no operations. The Tracer and Spans created from the returned
// TracerProvider also perform no operations.
//
// Deprecated: Use [go.opentelemetry.io/otel/trace/noop.NewTracerProvider]
// instead.
func NewNoopTracerProvider() TracerProvider {
	return noopTracerProvider{}
}

type noopTracerProvider struct{ embedded.TracerProvider }

var _ TracerProvider = noopTracerProvider{}

// Tracer returns noop implementation of Tracer.
func (p noopTracerProvider) Tracer(string, ...TracerOption) Tracer {
	return noopTracer{}
}

// noopTracer is an implementation of Tracer that performs no operations.
type noopTracer struct{ embedded.Tracer }

var _ Tracer = noopTracer{}

// Start starts a noop span.
func (t noopTracer) Start(ctx context.Context, name string, _ ...SpanOption) (context.Context, Span) {
	span := noopSpan{}
	return ContextWithSpan(ctx, span), span
}

// noopSpan is an implementation of Span that performs no operations.
type noopSpan struct{ embedded.Span }

var _ Span = noopSpan{}

// SpanContext returns an empty span context.
func (noopSpan) SpanContext() SpanContext { return SpanContext{} }

// IsRecording always returns false.
func (noopSpan) IsRecording() bool { return false }

// SetStatus does nothing.
func (noopSpan) SetStatus(codes.Code, string) {}

// SetError does nothing.
func (noopSpan) SetError(bool) {}

// SetAttributes does nothing.
func (noopSpan) SetAttributes(...attribute.KeyValue) {}

// End does nothing.
func (noopSpan) End(...SpanOption) {}

// RecordError does nothing.
func (noopSpan) RecordError(error, ...EventOption) {}

// Tracer returns the Tracer that created this Span.
func (noopSpan) Tracer() Tracer { return noopTracer{} }

// AddEvent does nothing.
func (noopSpan) AddEvent(string, ...EventOption) {}

// SetName does nothing.
func (noopSpan) SetName(string) {}
