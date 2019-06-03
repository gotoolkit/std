// Copyright 2019 llitfkitfk@gmail.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build jsoniter

package json

import jsoniter "github.com/json-iterator/go"

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by pkg/json package.
	NewEncoder = json.NewEncoder
)
