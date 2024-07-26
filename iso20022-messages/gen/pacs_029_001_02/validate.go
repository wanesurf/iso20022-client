// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:pacs.029.001.02
package pacs_029_001_02

import (
	"fmt"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
	"github.com/moov-io/base"
)

// XSD Element validations

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	iso.AddError(&errs, baseName+".MulSttlmReq", v.MulSttlmReq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations

func (v AccountIdentification4Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountIdentification4Choice"
	if v.IBAN != nil {
		iso.AddError(&errs, baseName+".IBAN", v.IBAN.Validate())
	}
	if v.Othr != nil {
		iso.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AccountSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountSchemeName1Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ActiveCurrencyAndAmount) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveCurrencyAndAmount"

	iso.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AddressType3Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AddressType3Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AmountAndDirection5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AmountAndDirection5"
	iso.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	if v.CdtDbt != nil {
		iso.AddError(&errs, baseName+".CdtDbt", v.CdtDbt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccount40) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccount40"
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.Tp != nil {
		iso.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	}
	if v.Ccy != nil {
		iso.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())
	}
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.Prxy != nil {
		iso.AddError(&errs, baseName+".Prxy", v.Prxy.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccountType2Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccountType2Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemIdentification3Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemIdentification3Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Contact13) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Contact13"
	if v.NmPrfx != nil {
		iso.AddError(&errs, baseName+".NmPrfx", v.NmPrfx.Validate())
	}
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PhneNb != nil {
		iso.AddError(&errs, baseName+".PhneNb", v.PhneNb.Validate())
	}
	if v.MobNb != nil {
		iso.AddError(&errs, baseName+".MobNb", v.MobNb.Validate())
	}
	if v.FaxNb != nil {
		iso.AddError(&errs, baseName+".FaxNb", v.FaxNb.Validate())
	}
	if v.URLAdr != nil {
		iso.AddError(&errs, baseName+".URLAdr", v.URLAdr.Validate())
	}
	if v.EmailAdr != nil {
		iso.AddError(&errs, baseName+".EmailAdr", v.EmailAdr.Validate())
	}
	if v.EmailPurp != nil {
		iso.AddError(&errs, baseName+".EmailPurp", v.EmailPurp.Validate())
	}
	if v.JobTitl != nil {
		iso.AddError(&errs, baseName+".JobTitl", v.JobTitl.Validate())
	}
	if v.Rspnsblty != nil {
		iso.AddError(&errs, baseName+".Rspnsblty", v.Rspnsblty.Validate())
	}
	if v.Dept != nil {
		iso.AddError(&errs, baseName+".Dept", v.Dept.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if v.PrefrdMtd != nil {
		iso.AddError(&errs, baseName+".PrefrdMtd", v.PrefrdMtd.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndPlaceOfBirth1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndPlaceOfBirth1"
	iso.AddError(&errs, baseName+".BirthDt", v.BirthDt.Validate())
	if v.PrvcOfBirth != nil {
		iso.AddError(&errs, baseName+".PrvcOfBirth", v.PrvcOfBirth.Validate())
	}
	iso.AddError(&errs, baseName+".CityOfBirth", v.CityOfBirth.Validate())
	iso.AddError(&errs, baseName+".CtryOfBirth", v.CtryOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericAccountIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericAccountIdentification1"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		iso.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericIdentification30) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericIdentification30"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	iso.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification3"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		iso.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericPersonIdentification2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericPersonIdentification2"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		iso.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader104) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader104"
	iso.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	iso.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	iso.AddError(&errs, baseName+".NbOfSttlmReqs", v.NbOfSttlmReqs.Validate())
	if v.CtrlSum != nil {
		iso.AddError(&errs, baseName+".CtrlSum", v.CtrlSum.Validate())
	}
	if v.SttlmInf != nil {
		iso.AddError(&errs, baseName+".SttlmInf", v.SttlmInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MovementRecord2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MovementRecord2"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SeqNb != nil {
		iso.AddError(&errs, baseName+".SeqNb", v.SeqNb.Validate())
	}
	iso.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	if v.SttlmAgt != nil {
		iso.AddError(&errs, baseName+".SttlmAgt", v.SttlmAgt.Validate())
	}
	if v.SttlmAgtAcct != nil {
		iso.AddError(&errs, baseName+".SttlmAgtAcct", v.SttlmAgtAcct.Validate())
	}
	if v.Ptcpt != nil {
		iso.AddError(&errs, baseName+".Ptcpt", v.Ptcpt.Validate())
	}
	if v.PtcptAcct != nil {
		iso.AddError(&errs, baseName+".PtcptAcct", v.PtcptAcct.Validate())
	}
	if v.Ref != nil {
		iso.AddError(&errs, baseName+".Ref", v.Ref.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MultilateralSettlementRequest3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MultilateralSettlementRequest3"
	iso.AddError(&errs, baseName+".InstrId", v.InstrId.Validate())
	if v.InstrPrty != nil {
		iso.AddError(&errs, baseName+".InstrPrty", v.InstrPrty.Validate())
	}
	if v.SttlmTmReq != nil {
		iso.AddError(&errs, baseName+".SttlmTmReq", v.SttlmTmReq.Validate())
	}
	if v.SttlmPrty != nil {
		iso.AddError(&errs, baseName+".SttlmPrty", v.SttlmPrty.Validate())
	}
	if v.SttlmCycl != nil {
		iso.AddError(&errs, baseName+".SttlmCycl", v.SttlmCycl.Validate())
	}
	if v.NbOfMvmntRcrds != nil {
		iso.AddError(&errs, baseName+".NbOfMvmntRcrds", v.NbOfMvmntRcrds.Validate())
	}
	for indx := range v.MvmntRcrd {
		iso.AddError(&errs, baseName+".MvmntRcrd", v.MvmntRcrd[indx].Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MultilateralSettlementRequestV02) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MultilateralSettlementRequestV02"
	iso.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	for indx := range v.SttlmReq {
		iso.AddError(&errs, baseName+".SttlmReq", v.SttlmReq[indx].Validate())
	}
	if v.SplmtryData != nil {
		iso.AddError(&errs, baseName+".SplmtryData", v.SplmtryData.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification39) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification39"
	if v.AnyBIC != nil {
		iso.AddError(&errs, baseName+".AnyBIC", v.AnyBIC.Validate())
	}
	if v.LEI != nil {
		iso.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentificationSchemeName1Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OtherContact1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OtherContact1"
	iso.AddError(&errs, baseName+".ChanlTp", v.ChanlTp.Validate())
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party52Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party52Choice"
	if v.OrgId != nil {
		iso.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		iso.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification272) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification272"
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		iso.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtryOfRes != nil {
		iso.AddError(&errs, baseName+".CtryOfRes", v.CtryOfRes.Validate())
	}
	if v.CtctDtls != nil {
		iso.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification18) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification18"
	if v.DtAndPlcOfBirth != nil {
		iso.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentificationSchemeName1Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress27) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress27"
	if v.AdrTp != nil {
		iso.AddError(&errs, baseName+".AdrTp", v.AdrTp.Validate())
	}
	if v.CareOf != nil {
		iso.AddError(&errs, baseName+".CareOf", v.CareOf.Validate())
	}
	if v.Dept != nil {
		iso.AddError(&errs, baseName+".Dept", v.Dept.Validate())
	}
	if v.SubDept != nil {
		iso.AddError(&errs, baseName+".SubDept", v.SubDept.Validate())
	}
	if v.StrtNm != nil {
		iso.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	}
	if v.BldgNb != nil {
		iso.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	if v.BldgNm != nil {
		iso.AddError(&errs, baseName+".BldgNm", v.BldgNm.Validate())
	}
	if v.Flr != nil {
		iso.AddError(&errs, baseName+".Flr", v.Flr.Validate())
	}
	if v.UnitNb != nil {
		iso.AddError(&errs, baseName+".UnitNb", v.UnitNb.Validate())
	}
	if v.PstBx != nil {
		iso.AddError(&errs, baseName+".PstBx", v.PstBx.Validate())
	}
	if v.Room != nil {
		iso.AddError(&errs, baseName+".Room", v.Room.Validate())
	}
	if v.PstCd != nil {
		iso.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	}
	if v.TwnNm != nil {
		iso.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	}
	if v.TwnLctnNm != nil {
		iso.AddError(&errs, baseName+".TwnLctnNm", v.TwnLctnNm.Validate())
	}
	if v.DstrctNm != nil {
		iso.AddError(&errs, baseName+".DstrctNm", v.DstrctNm.Validate())
	}
	if v.CtrySubDvsn != nil {
		iso.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	}
	if v.Ctry != nil {
		iso.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	}
	if v.AdrLine != nil {
		for indx := range v.AdrLine {
			iso.AddError(&errs, baseName+".AdrLine", v.AdrLine[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ProxyAccountIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProxyAccountIdentification1"
	if v.Tp != nil {
		iso.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	}
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ProxyAccountType1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ProxyAccountType1Choice"
	if v.Cd != nil {
		iso.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SettlementInstruction14) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SettlementInstruction14"
	iso.AddError(&errs, baseName+".SttlmMtd", v.SttlmMtd.Validate())
	if v.SttlmAcct != nil {
		iso.AddError(&errs, baseName+".SttlmAcct", v.SttlmAcct.Validate())
	}
	if v.ClrSys != nil {
		iso.AddError(&errs, baseName+".ClrSys", v.ClrSys.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SettlementTimeRequest2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SettlementTimeRequest2"
	if v.CLSTm != nil {
		iso.AddError(&errs, baseName+".CLSTm", v.CLSTm.Validate())
	}
	if v.TillTm != nil {
		iso.AddError(&errs, baseName+".TillTm", v.TillTm.Validate())
	}
	if v.FrTm != nil {
		iso.AddError(&errs, baseName+".FrTm", v.FrTm.Validate())
	}
	if v.RjctTm != nil {
		iso.AddError(&errs, baseName+".RjctTm", v.RjctTm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SupplementaryData1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SupplementaryData1"
	if v.PlcAndNm != nil {
		iso.AddError(&errs, baseName+".PlcAndNm", v.PlcAndNm.Validate())
	}
	iso.AddError(&errs, baseName+".Envlp", v.Envlp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SupplementaryDataEnvelope1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v ActiveCurrencyCode) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z]{3,3}`); err != nil {
		return err
	}
	return nil
}

func (v ActiveOrHistoricCurrencyCode) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z]{3,3}`); err != nil {
		return err
	}
	return nil
}

func (v AddressType2Code) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY"); err != nil {
		return err
	}
	return nil
}

func (v AnyBICDec2014Identifier) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}`); err != nil {
		return err
	}
	return nil
}

func (v CountryCode) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z]{2,2}`); err != nil {
		return err
	}
	return nil
}

func (v CreditDebitCode) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "CRDT", "DBIT"); err != nil {
		return err
	}
	return nil
}

func (v DecimalNumber) Validate() error {
	if err := iso.ValidateFractionDigits(fmt.Sprintf("%v", v), 17); err != nil {
		return err
	}
	if err := iso.ValidateTotalDigits(fmt.Sprintf("%v", v), 18); err != nil {
		return err
	}
	return nil
}

func (v Exact4AlphaNumericText) Validate() error {
	if err := iso.ValidatePattern(string(v), `[a-zA-Z0-9]{4}`); err != nil {
		return err
	}
	return nil
}

func (v ExternalAccountIdentification1Code) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalCashAccountType1Code) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalCashClearingSystem1Code) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 3); err != nil {
		return err
	}
	return nil
}

func (v ExternalOrganisationIdentification1Code) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalPersonIdentification1Code) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalProxyAccountType1Code) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v IBAN2007Identifier) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v ISOTime) Validate() error {
	return nil
}

func (v LEIIdentifier) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z0-9]{18,18}[0-9]{2,2}`); err != nil {
		return err
	}
	return nil
}

func (v Max128Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 128); err != nil {
		return err
	}
	return nil
}

func (v Max140Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 140); err != nil {
		return err
	}
	return nil
}

func (v Max15NumericText) Validate() error {
	if err := iso.ValidatePattern(string(v), `[0-9]{1,15}`); err != nil {
		return err
	}
	return nil
}

func (v Max16Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 16); err != nil {
		return err
	}
	return nil
}

func (v Max2048Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 2048); err != nil {
		return err
	}
	return nil
}

func (v Max256Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 256); err != nil {
		return err
	}
	return nil
}

func (v Max34Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 34); err != nil {
		return err
	}
	return nil
}

func (v Max350Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 350); err != nil {
		return err
	}
	return nil
}

func (v Max35Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max4Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v Max70Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 70); err != nil {
		return err
	}
	return nil
}

func (v NamePrefix2Code) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "DOCT", "MADM", "MISS", "MIST", "MIKS"); err != nil {
		return err
	}
	return nil
}

func (v Number) Validate() error {
	if err := iso.ValidateFractionDigits(fmt.Sprintf("%v", v), 0); err != nil {
		return err
	}
	if err := iso.ValidateTotalDigits(fmt.Sprintf("%v", v), 18); err != nil {
		return err
	}
	return nil
}

func (v PhoneNumber) Validate() error {
	if err := iso.ValidatePattern(string(v), `\+[0-9]{1,3}-[0-9()+\-]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v PreferredContactMethod2Code) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "MAIL", "FAXX", "LETT", "CELL", "ONLI", "PHON"); err != nil {
		return err
	}
	return nil
}

func (v Priority3Code) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "URGT", "HIGH", "NORM"); err != nil {
		return err
	}
	return nil
}

func (v SettlementMethod2Code) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "INDA", "INGA", "CLRG"); err != nil {
		return err
	}
	return nil
}