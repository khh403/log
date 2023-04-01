/*
 * Copyright (c) 2023 khh403. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package log_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khh403/log"
)

func Test_Options_Validate(t *testing.T) {
	opts := &log.Options{
		Level:            "test",
		Format:           "test",
		EnableColor:      true,
		EnableCaller:     false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	errs := opts.Validate()
	expected := `[unrecognized level: "test" not a valid log format: "test"]`
	assert.Equal(t, expected, fmt.Sprintf("%s", errs))
}
