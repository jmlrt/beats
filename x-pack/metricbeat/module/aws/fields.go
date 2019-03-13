// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}