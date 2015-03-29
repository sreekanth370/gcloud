// Copyright 2015 Google Inc. All Rights Reserved.
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

package gcs

import "time"

// Object is a record representing a particular generation of a particular
// object name in GCS.
//
// See here for more information about its fields:
//
//     https://cloud.google.com/storage/docs/json_api/v1/objects#resource
//
type Object struct {
	Name            string
	ContentType     string
	ContentLanguage string
	CacheControl    string
	Owner           string
	// TODO(jacobsa): Switch to uint64, matching underlying JSON interface.
	// Cf. https://cloud.google.com/storage/docs/json_api/v1/objects
	Size            int64
	ContentEncoding string
	// TODO(jacobsa): Switch to [16]byte like the md5 package.
	MD5            []byte
	CRC32C         uint32
	MediaLink      string
	Metadata       map[string]string
	Generation     int64
	MetaGeneration int64
	StorageClass   string
	Deleted        time.Time
	Updated        time.Time
}