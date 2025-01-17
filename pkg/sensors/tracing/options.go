// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

package tracing

import (
	"fmt"
	"strconv"

	"github.com/cilium/tetragon/pkg/k8s/apis/cilium.io/v1alpha1"
	"github.com/cilium/tetragon/pkg/logger"
	"github.com/cilium/tetragon/pkg/option"
)

type kprobeOptions struct {
	DisableKprobeMulti bool
}

type opt struct {
	set func(val string, options *kprobeOptions) error
}

// Allowed kprobe options
var opts = map[string]opt{
	option.KeyDisableKprobeMulti: opt{
		set: func(str string, options *kprobeOptions) (err error) {
			options.DisableKprobeMulti, err = strconv.ParseBool(str)
			return err
		},
	},
}

func getKprobeOptions(specs []v1alpha1.OptionSpec) (*kprobeOptions, error) {
	options := &kprobeOptions{}

	for _, spec := range specs {
		opt, ok := opts[spec.Name]
		if ok {
			if err := opt.set(spec.Value, options); err != nil {
				return nil, fmt.Errorf("failed to set option %s: %s", spec.Name, err)
			}
			logger.GetLogger().Infof("Set option %s = %s", spec.Name, spec.Value)
		}
	}

	return options, nil
}
