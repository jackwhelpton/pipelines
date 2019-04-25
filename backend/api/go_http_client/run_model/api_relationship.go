// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package run_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// APIRelationship api relationship
// swagger:model apiRelationship
type APIRelationship string

const (

	// APIRelationshipUNKNOWNRELATIONSHIP captures enum value "UNKNOWN_RELATIONSHIP"
	APIRelationshipUNKNOWNRELATIONSHIP APIRelationship = "UNKNOWN_RELATIONSHIP"

	// APIRelationshipOWNER captures enum value "OWNER"
	APIRelationshipOWNER APIRelationship = "OWNER"

	// APIRelationshipCREATOR captures enum value "CREATOR"
	APIRelationshipCREATOR APIRelationship = "CREATOR"
)

// for schema
var apiRelationshipEnum []interface{}

func init() {
	var res []APIRelationship
	if err := json.Unmarshal([]byte(`["UNKNOWN_RELATIONSHIP","OWNER","CREATOR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiRelationshipEnum = append(apiRelationshipEnum, v)
	}
}

func (m APIRelationship) validateAPIRelationshipEnum(path, location string, value APIRelationship) error {
	if err := validate.Enum(path, location, value, apiRelationshipEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this api relationship
func (m APIRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPIRelationshipEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
