// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package test_sm_ies

import "reflect"

var Choicemap = map[string]map[int]reflect.Type{
	"choice1": {
		1: reflect.TypeOf(Choice1_Choice1A{}),
	},
	"choice2": {
		1: reflect.TypeOf(Choice2_Choice2A{}),
		2: reflect.TypeOf(Choice2_Choice2B{}),
	},
	"choice3": {
		1: reflect.TypeOf(Choice3_Choice3A{}),
		2: reflect.TypeOf(Choice3_Choice3B{}),
		3: reflect.TypeOf(Choice3_Choice3C{}),
	},
	"choice4": {
		1: reflect.TypeOf(Choice4_Choice4A{}),
	},
	"constrained_choice1": {
		1: reflect.TypeOf(ConstrainedChoice1_ConstrainedChoice1A{}),
	},
	"constrained_choice2": {
		1: reflect.TypeOf(ConstrainedChoice2_ConstrainedChoice2A{}),
		2: reflect.TypeOf(ConstrainedChoice2_ConstrainedChoice2B{}),
	},
	"constrained_choice3": {
		1: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3A{}),
		2: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3B{}),
		3: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3C{}),
		4: reflect.TypeOf(ConstrainedChoice3_ConstrainedChoice3D{}),
	},
	"constrained_choice4": {
		1: reflect.TypeOf(ConstrainedChoice4_ConstrainedChoice4A{}),
	},
	"test_nested_choice": {
		1: reflect.TypeOf(TestNestedChoice_Option1{}),
		2: reflect.TypeOf(TestNestedChoice_Option2{}),
		3: reflect.TypeOf(TestNestedChoice_Option3{}),
	},
}
