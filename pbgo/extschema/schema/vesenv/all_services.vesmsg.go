// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package vesenv

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *ServiceChoice) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ServiceChoice) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ServiceChoice) DeepCopy() *ServiceChoice {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ServiceChoice{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ServiceChoice) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ServiceChoice) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ServiceChoiceValidator().Validate(ctx, m, opts...)
}

type ValidateServiceChoice struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateServiceChoice) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ServiceChoice)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ServiceChoice got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetChoice().(type) {
	case *ServiceChoice_Akar:
		if fv, exists := v.FldValidators["choice.akar"]; exists {
			val := m.GetChoice().(*ServiceChoice_Akar).Akar
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("akar"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ares:
		if fv, exists := v.FldValidators["choice.ares"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ares).Ares
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ares"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Asterix:
		if fv, exists := v.FldValidators["choice.asterix"]; exists {
			val := m.GetChoice().(*ServiceChoice_Asterix).Asterix
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("asterix"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Athena:
		if fv, exists := v.FldValidators["choice.athena"]; exists {
			val := m.GetChoice().(*ServiceChoice_Athena).Athena
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("athena"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Automatix:
		if fv, exists := v.FldValidators["choice.automatix"]; exists {
			val := m.GetChoice().(*ServiceChoice_Automatix).Automatix
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("automatix"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Barracute:
		if fv, exists := v.FldValidators["choice.barracute"]; exists {
			val := m.GetChoice().(*ServiceChoice_Barracute).Barracute
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("barracute"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Blindfold:
		if fv, exists := v.FldValidators["choice.blindfold"]; exists {
			val := m.GetChoice().(*ServiceChoice_Blindfold).Blindfold
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("blindfold"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Bolt:
		if fv, exists := v.FldValidators["choice.bolt"]; exists {
			val := m.GetChoice().(*ServiceChoice_Bolt).Bolt
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("bolt"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Charmander:
		if fv, exists := v.FldValidators["choice.charmander"]; exists {
			val := m.GetChoice().(*ServiceChoice_Charmander).Charmander
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("charmander"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Charmandercar:
		if fv, exists := v.FldValidators["choice.charmandercar"]; exists {
			val := m.GetChoice().(*ServiceChoice_Charmandercar).Charmandercar
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("charmandercar"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Drogon:
		if fv, exists := v.FldValidators["choice.drogon"]; exists {
			val := m.GetChoice().(*ServiceChoice_Drogon).Drogon
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("drogon"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Eywa:
		if fv, exists := v.FldValidators["choice.eywa"]; exists {
			val := m.GetChoice().(*ServiceChoice_Eywa).Eywa
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("eywa"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Griffin:
		if fv, exists := v.FldValidators["choice.griffin"]; exists {
			val := m.GetChoice().(*ServiceChoice_Griffin).Griffin
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("griffin"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Identityauthority:
		if fv, exists := v.FldValidators["choice.identityauthority"]; exists {
			val := m.GetChoice().(*ServiceChoice_Identityauthority).Identityauthority
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("identityauthority"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Keypr:
		if fv, exists := v.FldValidators["choice.keypr"]; exists {
			val := m.GetChoice().(*ServiceChoice_Keypr).Keypr
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("keypr"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Maurice:
		if fv, exists := v.FldValidators["choice.maurice"]; exists {
			val := m.GetChoice().(*ServiceChoice_Maurice).Maurice
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("maurice"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Obelix:
		if fv, exists := v.FldValidators["choice.obelix"]; exists {
			val := m.GetChoice().(*ServiceChoice_Obelix).Obelix
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("obelix"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ocspmule:
		if fv, exists := v.FldValidators["choice.ocspmule"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ocspmule).Ocspmule
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ocspmule"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Opera:
		if fv, exists := v.FldValidators["choice.opera"]; exists {
			val := m.GetChoice().(*ServiceChoice_Opera).Opera
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("opera"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Pikachu:
		if fv, exists := v.FldValidators["choice.pikachu"]; exists {
			val := m.GetChoice().(*ServiceChoice_Pikachu).Pikachu
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("pikachu"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Pkifactory:
		if fv, exists := v.FldValidators["choice.pkifactory"]; exists {
			val := m.GetChoice().(*ServiceChoice_Pkifactory).Pkifactory
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("pkifactory"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Rakar:
		if fv, exists := v.FldValidators["choice.rakar"]; exists {
			val := m.GetChoice().(*ServiceChoice_Rakar).Rakar
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("rakar"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_SiteConsole:
		if fv, exists := v.FldValidators["choice.site_console"]; exists {
			val := m.GetChoice().(*ServiceChoice_SiteConsole).SiteConsole
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("site_console"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Vega:
		if fv, exists := v.FldValidators["choice.vega"]; exists {
			val := m.GetChoice().(*ServiceChoice_Vega).Vega
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("vega"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Voucher:
		if fv, exists := v.FldValidators["choice.voucher"]; exists {
			val := m.GetChoice().(*ServiceChoice_Voucher).Voucher
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("voucher"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Vpm:
		if fv, exists := v.FldValidators["choice.vpm"]; exists {
			val := m.GetChoice().(*ServiceChoice_Vpm).Vpm
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("vpm"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Vulpix:
		if fv, exists := v.FldValidators["choice.vulpix"]; exists {
			val := m.GetChoice().(*ServiceChoice_Vulpix).Vulpix
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("vulpix"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Warden:
		if fv, exists := v.FldValidators["choice.warden"]; exists {
			val := m.GetChoice().(*ServiceChoice_Warden).Warden
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("warden"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Wingman:
		if fv, exists := v.FldValidators["choice.wingman"]; exists {
			val := m.GetChoice().(*ServiceChoice_Wingman).Wingman
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("wingman"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Hellas:
		if fv, exists := v.FldValidators["choice.hellas"]; exists {
			val := m.GetChoice().(*ServiceChoice_Hellas).Hellas
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("hellas"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Nfvsvc:
		if fv, exists := v.FldValidators["choice.nfvsvc"]; exists {
			val := m.GetChoice().(*ServiceChoice_Nfvsvc).Nfvsvc
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("nfvsvc"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Scim:
		if fv, exists := v.FldValidators["choice.scim"]; exists {
			val := m.GetChoice().(*ServiceChoice_Scim).Scim
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("scim"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_LilacEdge:
		if fv, exists := v.FldValidators["choice.lilac_edge"]; exists {
			val := m.GetChoice().(*ServiceChoice_LilacEdge).LilacEdge
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("lilac_edge"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Aipconnector:
		if fv, exists := v.FldValidators["choice.aipconnector"]; exists {
			val := m.GetChoice().(*ServiceChoice_Aipconnector).Aipconnector
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("aipconnector"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_AkarReadonly:
		if fv, exists := v.FldValidators["choice.akar_readonly"]; exists {
			val := m.GetChoice().(*ServiceChoice_AkarReadonly).AkarReadonly
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("akar_readonly"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_AkarDnsdomain:
		if fv, exists := v.FldValidators["choice.akar_dnsdomain"]; exists {
			val := m.GetChoice().(*ServiceChoice_AkarDnsdomain).AkarDnsdomain
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("akar_dnsdomain"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Aspen:
		if fv, exists := v.FldValidators["choice.aspen"]; exists {
			val := m.GetChoice().(*ServiceChoice_Aspen).Aspen
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("aspen"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Bdbewaf:
		if fv, exists := v.FldValidators["choice.bdbewaf"]; exists {
			val := m.GetChoice().(*ServiceChoice_Bdbewaf).Bdbewaf
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("bdbewaf"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Bfdp:
		if fv, exists := v.FldValidators["choice.bfdp"]; exists {
			val := m.GetChoice().(*ServiceChoice_Bfdp).Bfdp
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("bfdp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Bifrost:
		if fv, exists := v.FldValidators["choice.bifrost"]; exists {
			val := m.GetChoice().(*ServiceChoice_Bifrost).Bifrost
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("bifrost"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Brmalerts:
		if fv, exists := v.FldValidators["choice.brmalerts"]; exists {
			val := m.GetChoice().(*ServiceChoice_Brmalerts).Brmalerts
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("brmalerts"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Cdnconnectorsvc:
		if fv, exists := v.FldValidators["choice.cdnconnectorsvc"]; exists {
			val := m.GetChoice().(*ServiceChoice_Cdnconnectorsvc).Cdnconnectorsvc
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("cdnconnectorsvc"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_ClientSideDefense:
		if fv, exists := v.FldValidators["choice.client_side_defense"]; exists {
			val := m.GetChoice().(*ServiceChoice_ClientSideDefense).ClientSideDefense
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("client_side_defense"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Doscontroller:
		if fv, exists := v.FldValidators["choice.doscontroller"]; exists {
			val := m.GetChoice().(*ServiceChoice_Doscontroller).Doscontroller
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("doscontroller"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_EywaReadonly:
		if fv, exists := v.FldValidators["choice.eywa_readonly"]; exists {
			val := m.GetChoice().(*ServiceChoice_EywaReadonly).EywaReadonly
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("eywa_readonly"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Eywaprime:
		if fv, exists := v.FldValidators["choice.eywaprime"]; exists {
			val := m.GetChoice().(*ServiceChoice_Eywaprime).Eywaprime
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("eywaprime"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_EywaprimeReadonly:
		if fv, exists := v.FldValidators["choice.eywaprime_readonly"]; exists {
			val := m.GetChoice().(*ServiceChoice_EywaprimeReadonly).EywaprimeReadonly
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("eywaprime_readonly"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ganges:
		if fv, exists := v.FldValidators["choice.ganges"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ganges).Ganges
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ganges"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Iapetus:
		if fv, exists := v.FldValidators["choice.iapetus"]; exists {
			val := m.GetChoice().(*ServiceChoice_Iapetus).Iapetus
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("iapetus"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ipp:
		if fv, exists := v.FldValidators["choice.ipp"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ipp).Ipp
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ipp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ippdev:
		if fv, exists := v.FldValidators["choice.ippdev"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ippdev).Ippdev
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ippdev"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ippprom:
		if fv, exists := v.FldValidators["choice.ippprom"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ippprom).Ippprom
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ippprom"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Josef:
		if fv, exists := v.FldValidators["choice.josef"]; exists {
			val := m.GetChoice().(*ServiceChoice_Josef).Josef
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("josef"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Kcdmux:
		if fv, exists := v.FldValidators["choice.kcdmux"]; exists {
			val := m.GetChoice().(*ServiceChoice_Kcdmux).Kcdmux
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("kcdmux"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_KcdmuxReadonly:
		if fv, exists := v.FldValidators["choice.kcdmux_readonly"]; exists {
			val := m.GetChoice().(*ServiceChoice_KcdmuxReadonly).KcdmuxReadonly
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("kcdmux_readonly"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Mars:
		if fv, exists := v.FldValidators["choice.mars"]; exists {
			val := m.GetChoice().(*ServiceChoice_Mars).Mars
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("mars"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Minerva:
		if fv, exists := v.FldValidators["choice.minerva"]; exists {
			val := m.GetChoice().(*ServiceChoice_Minerva).Minerva
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("minerva"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Nmsconnector:
		if fv, exists := v.FldValidators["choice.nmsconnector"]; exists {
			val := m.GetChoice().(*ServiceChoice_Nmsconnector).Nmsconnector
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("nmsconnector"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Nmsproxy:
		if fv, exists := v.FldValidators["choice.nmsproxy"]; exists {
			val := m.GetChoice().(*ServiceChoice_Nmsproxy).Nmsproxy
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("nmsproxy"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Olympus:
		if fv, exists := v.FldValidators["choice.olympus"]; exists {
			val := m.GetChoice().(*ServiceChoice_Olympus).Olympus
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("olympus"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Phobos:
		if fv, exists := v.FldValidators["choice.phobos"]; exists {
			val := m.GetChoice().(*ServiceChoice_Phobos).Phobos
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("phobos"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Piku:
		if fv, exists := v.FldValidators["choice.piku"]; exists {
			val := m.GetChoice().(*ServiceChoice_Piku).Piku
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("piku"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Pluto:
		if fv, exists := v.FldValidators["choice.pluto"]; exists {
			val := m.GetChoice().(*ServiceChoice_Pluto).Pluto
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("pluto"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ponos:
		if fv, exists := v.FldValidators["choice.ponos"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ponos).Ponos
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ponos"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Prism:
		if fv, exists := v.FldValidators["choice.prism"]; exists {
			val := m.GetChoice().(*ServiceChoice_Prism).Prism
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("prism"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Prismprime:
		if fv, exists := v.FldValidators["choice.prismprime"]; exists {
			val := m.GetChoice().(*ServiceChoice_Prismprime).Prismprime
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("prismprime"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Protekti:
		if fv, exists := v.FldValidators["choice.protekti"]; exists {
			val := m.GetChoice().(*ServiceChoice_Protekti).Protekti
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("protekti"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Rolex:
		if fv, exists := v.FldValidators["choice.rolex"]; exists {
			val := m.GetChoice().(*ServiceChoice_Rolex).Rolex
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("rolex"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Sed:
		if fv, exists := v.FldValidators["choice.sed"]; exists {
			val := m.GetChoice().(*ServiceChoice_Sed).Sed
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("sed"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Sentinel:
		if fv, exists := v.FldValidators["choice.sentinel"]; exists {
			val := m.GetChoice().(*ServiceChoice_Sentinel).Sentinel
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("sentinel"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Spectrum:
		if fv, exists := v.FldValidators["choice.spectrum"]; exists {
			val := m.GetChoice().(*ServiceChoice_Spectrum).Spectrum
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("spectrum"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Streak:
		if fv, exists := v.FldValidators["choice.streak"]; exists {
			val := m.GetChoice().(*ServiceChoice_Streak).Streak
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("streak"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Trafficactioner:
		if fv, exists := v.FldValidators["choice.trafficactioner"]; exists {
			val := m.GetChoice().(*ServiceChoice_Trafficactioner).Trafficactioner
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("trafficactioner"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Ver:
		if fv, exists := v.FldValidators["choice.ver"]; exists {
			val := m.GetChoice().(*ServiceChoice_Ver).Ver
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ver"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Tpmauthority:
		if fv, exists := v.FldValidators["choice.tpmauthority"]; exists {
			val := m.GetChoice().(*ServiceChoice_Tpmauthority).Tpmauthority
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("tpmauthority"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Prometheus:
		if fv, exists := v.FldValidators["choice.prometheus"]; exists {
			val := m.GetChoice().(*ServiceChoice_Prometheus).Prometheus
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("prometheus"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_ViaApiService:
		if fv, exists := v.FldValidators["choice.via_api_service"]; exists {
			val := m.GetChoice().(*ServiceChoice_ViaApiService).ViaApiService
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("via_api_service"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Viaconnector:
		if fv, exists := v.FldValidators["choice.viaconnector"]; exists {
			val := m.GetChoice().(*ServiceChoice_Viaconnector).Viaconnector
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("viaconnector"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Dymo:
		if fv, exists := v.FldValidators["choice.dymo"]; exists {
			val := m.GetChoice().(*ServiceChoice_Dymo).Dymo
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("dymo"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ServiceChoice_Conprof:
		if fv, exists := v.FldValidators["choice.conprof"]; exists {
			val := m.GetChoice().(*ServiceChoice_Conprof).Conprof
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("conprof"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultServiceChoiceValidator = func() *ValidateServiceChoice {
	v := &ValidateServiceChoice{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ServiceChoiceValidator() db.Validator {
	return DefaultServiceChoiceValidator
}
