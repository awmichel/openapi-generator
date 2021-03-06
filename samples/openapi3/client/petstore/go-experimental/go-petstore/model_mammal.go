/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
	"fmt"
)

// Mammal struct for Mammal
type Mammal struct {
	MammalInterface interface { GetClassName() string }
}

func (s Mammal) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.MammalInterface)
}

func (s *Mammal) UnmarshalJSON(src []byte) error {
	var err error
	var unmarshaled map[string]interface{}
	err = json.Unmarshal(src, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["className"]; ok {
		switch v {
			case "whale":
				var result *Whale = &Whale{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.MammalInterface = result
				return nil
			case "zebra":
				var result *Zebra = &Zebra{}
				err = json.Unmarshal(src, result)
				if err != nil {
					return err
				}
				s.MammalInterface = result
				return nil
			default:
				return fmt.Errorf("No oneOf model has 'className' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'className' not found in unmarshaled payload: %+v", unmarshaled)
	}
}
type NullableMammal struct {
	value *Mammal
	isSet bool
}

func (v NullableMammal) Get() *Mammal {
	return v.value
}

func (v *NullableMammal) Set(val *Mammal) {
	v.value = val
	v.isSet = true
}

func (v NullableMammal) IsSet() bool {
	return v.isSet
}

func (v *NullableMammal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMammal(val *Mammal) *NullableMammal {
	return &NullableMammal{value: val, isSet: true}
}

func (v NullableMammal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMammal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
