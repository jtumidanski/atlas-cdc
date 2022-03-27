package handler

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EventHandler[E any] func(logrus.FieldLogger, opentracing.Span, E)
