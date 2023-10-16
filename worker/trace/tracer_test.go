// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package trace

import (
	"context"

	"github.com/juju/errors"
	"github.com/juju/names/v4"
	jc "github.com/juju/testing/checkers"
	"github.com/juju/worker/v3/workertest"
	gomock "go.uber.org/mock/gomock"
	gc "gopkg.in/check.v1"

	coretrace "github.com/juju/juju/core/trace"
)

type tracerSuite struct {
	baseSuite
}

var _ = gc.Suite(&tracerSuite{})

func (s *tracerSuite) TestTracer(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClient()

	tracer := s.newTracer(c)
	defer workertest.CleanKill(c, tracer)

	ctx, span := tracer.Start(context.Background(), "foo")
	c.Check(ctx, gc.NotNil)
	c.Check(span, gc.NotNil)

	defer span.End()
}

func (s *tracerSuite) TestTracerStartContext(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClient()

	tracer := s.newTracer(c)
	defer workertest.CleanKill(c, tracer)

	ctx, span := tracer.Start(context.Background(), "foo")
	defer span.End()

	select {
	case <-ctx.Done():
		c.Fatalf("context should not be done")
	default:
	}
}

func (s *tracerSuite) TestTracerStartContextShouldBeCanceled(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClient()

	tracer := s.newTracer(c)
	defer workertest.CleanKill(c, tracer)

	ctx, cancel := context.WithCancel(context.Background())

	// cancel the context straight away
	cancel()

	ctx, span := tracer.Start(ctx, "foo")
	defer span.End()

	select {
	case <-ctx.Done():
	default:
		c.Fatalf("context should be done")
	}
}

func (s *tracerSuite) TestTracerAddEvent(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClient()

	s.span.EXPECT().AddEvent(gomock.Any(), gomock.Any()).AnyTimes()

	tracer := s.newTracer(c)
	defer workertest.CleanKill(c, tracer)

	_, span := tracer.Start(context.Background(), "foo")
	defer span.End()

	span.AddEvent("bar")
	span.AddEvent("baz", coretrace.StringAttr("qux", "quux"))
}

func (s *tracerSuite) TestTracerRecordErrorWithNil(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClient()

	tracer := s.newTracer(c)
	defer workertest.CleanKill(c, tracer)

	_, span := tracer.Start(context.Background(), "foo")
	defer span.End()

	span.RecordError(nil)
}

func (s *tracerSuite) TestTracerRecordError(c *gc.C) {
	defer s.setupMocks(c).Finish()

	s.expectClient()

	s.span.EXPECT().RecordError(gomock.Any(), gomock.Any()).AnyTimes()
	s.span.EXPECT().SetStatus(gomock.Any(), gomock.Any()).AnyTimes()

	tracer := s.newTracer(c)
	defer workertest.CleanKill(c, tracer)

	_, span := tracer.Start(context.Background(), "foo")
	defer span.End()

	span.RecordError(errors.Errorf("boom"))
}

func (s *tracerSuite) newTracer(c *gc.C) TrackedTracer {
	ns := coretrace.Namespace("agent", "controller").WithTag(names.NewMachineTag("0"))
	newClient := func(context.Context, coretrace.TaggedTracerNamespace, string, bool) (Client, ClientTracerProvider, ClientTracer, error) {
		return s.client, s.clientTracerProvider, s.clientTracer, nil
	}
	tracer, err := NewTracerWorker(context.Background(), ns, "http://meshuggah.com", false, false, s.logger, newClient)
	c.Assert(err, jc.ErrorIsNil)
	return tracer
}