// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	{{.Imports}}
	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	{{if .Logger}}"github.com/onosproject/onos-lib-go/pkg/logging"{{end}}
)
{{if .Logger}}var log = logging.GetLogger(){{end}}
func PerEncode{{.MessageName}}(msg *{{.ProtoName}}.{{.MessageName}}) ([]byte, error) {

	log.Debugf("Obtained {{.MessageNameInLogging}} message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating {{.MessageNameInLogging}} PDU %s", err.Error())
	}

	{{if .CanonicalChoicePresence}}per, err := aper.MarshalWithParams(msg, "{{.Parameters}}", choiceOptions.{{.ChoiceMapName}}, choiceOptions.{{.CanonicalChoiceMapName}}){{else}}per, err := aper.MarshalWithParams(msg, "{{.Parameters}}", choiceOptions.{{.ChoiceMapName}}, nil){{end}}
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded {{.MessageNameInLogging}} PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecode{{.MessageName}}(per []byte) (*{{.ProtoName}}.{{.MessageName}}, error) {

	log.Debugf("Obtained {{.MessageNameInLogging}} PER bytes are\n%v", hex.Dump(per))

	result := {{.ProtoName}}.{{.MessageName}}{}
	{{if .CanonicalChoicePresence}}err := aper.UnmarshalWithParams(per, &result, "{{.Parameters}}", choiceOptions.{{.ChoiceMapName}}, choiceOptions.{{.CanonicalChoiceMapName}}){{else}}err := aper.UnmarshalWithParams(per, &result, "{{.Parameters}}", choiceOptions.{{.ChoiceMapName}}, nil){{end}}
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded {{.MessageNameInLogging}} from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating {{.MessageNameInLogging}} PDU %s", err.Error())
	}

	return &result, nil
}