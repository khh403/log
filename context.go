/*
 * Copyright (c) 2023 khh403. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package log

import (
	"context"
)

type key int

const (
	logContextKey key = iota
)

// WithContext returns a copy of context in which the log value is set.
func (l *zapLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

// FromContext returns the value of the log key on the ctx.
func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		logger := ctx.Value(logContextKey)
		if logger != nil {
			return logger.(Logger)
		}
	}
	return WithName("Unknown-Context")
}
