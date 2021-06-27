package dao

import (
	"blog-service/pkg/tracer"
	"context"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type Dao struct {
	engine *gorm.DB
}

func New(ctx context.Context, engine *gorm.DB) *Dao {
	//从context获取span
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return &Dao{engine: engine}
	}
	return &Dao{engine: engine.Set(tracer.ParentSpanGormKey, parentSpan)}
}

