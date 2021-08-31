// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by x-pack/elastic-agent/dev-tools/cmd/buildspec/buildspec.go - DO NOT EDIT.

package program

import (
	"strings"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
)

var Supported []Spec
var SupportedMap map[string]Spec

func init() {
	// Packed Files
	// spec/apm-server.yml
	// spec/endpoint.yml
	// spec/filebeat.yml
	// spec/fleet-server.yml
	// spec/heartbeat.yml
	// spec/metricbeat.yml
	// spec/osquerybeat.yml
	// spec/packetbeat.yml
	unpacked := packer.MustUnpack("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")
	SupportedMap = make(map[string]Spec)

	for f, v := range unpacked {
		s, err := NewSpecFromBytes(v)
		if err != nil {
			panic("Cannot read spec from " + f)
		}
		Supported = append(Supported, s)
		SupportedMap[strings.ToLower(s.Cmd)] = s
	}
}
