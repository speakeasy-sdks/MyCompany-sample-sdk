// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// UnitOfMeasureType - Enum representing various unit of measures.<p>Members:</p><ul><li><i>BX</i> - Box (Norwegian: Eske)</li><li><i>MTR</i> - Meter (m) (Norwegian: Meter)</li><li><i>KMT</i> - Kilometer (km) (Norwegian: Kilometer)</li><li><i>KGM</i> - Kilogram (kg) (Norwegian: Kilogram)</li><li><i>EA</i> - Each (Norwegian: Stykke (stk))</li><li><i>LTR</i> - Liter (L) (Norwegian: Liter)</li><li><i>HUR</i> - Hour (Norwegian: Time)</li><li><i>DAY</i> - Day (Norwegian: Dag)</li><li><i>MTK</i> - Square meter (m^2) (Norwegian: Kvadratmeter)</li><li><i>MTQ</i> - Cubic meter (m^3) (Norwegian: Kubikkmeter)</li><li><i>TNE</i> - Metric ton (t) (Norwegian: Metrisk tonn)</li><li><i>MON</i> - Month (Norwegian: Måned)</li><li><i>ANN</i> - Year (Norwegian: År)</li><li><i>QAN</i> - Quarter of a year (Norwegian: Kvartal)</li><li><i>NL</i> - Load (Norwegian: Lass)</li><li><i>XRO</i> - Roll (Norwegian: Rull)</li><li><i>XBA</i> - Barrel (Norwegian: Tønne)</li><li><i>XPL</i> - Pail (Norwegian: Spann)</li><li><i>XPC</i> - Parcel (Norwegian: Kolli)</li><li><i>PR</i> - Pair (Norwegian: Par)</li><li><i>XCI</i> - Canister (Norwegian: Boks)</li><li><i>XBG</i> - Bag (Norwegian: Pose)</li><li><i>SET</i> - Set (Norwegian: Sett)</li><li><i>XTY</i> - Tank (Norwegian: Tank)</li><li><i>XOF</i> - Pallet (Norwegian: Pall)</li><li><i>FTK</i> - Square foot (ft^2) (Norwegian: Kvadratfot)</li><li><i>KWH</i> - Kilowatt hour (kWh) (Norwegian: Kilowattime)</li><li><i>MWH</i> - Megawatt hour (mWh) (Norwegian: Megawattime)</li><li><i>LBR</i> - Pound (lb) (Norwegian: Pund)</li><li><i>CMT</i> - Centimeter (cm) (Norwegian: Centimeter)</li><li><i>DMT</i> - Decimeter (dm) (Norwegian: Desimeter)</li><li><i>LM</i> - Linear meter (Norwegian: Lineær meter)</li><li><i>XPK</i> - Package (Norwegian: Pakke (pk))</li><li><i>GRM</i> - Gram (g) (Norwegian: Gram)</li><li><i>HGM</i> - Hectogram (hg) (Norwegian: Hektogram)</li><li><i>XFL</i> - Flask (Norwegian: Flaske)</li><li><i>XBE</i> - Bundle (Norwegian: Bunt)</li><li><i>E54</i> - Trip (Norwegian: Tur)</li><li><i>MMT</i> - Millimeter (mm) (Norwegian: Millimeter)</li><li><i>DAA</i> - Decare (Norwegian: Dekar)</li><li><i>H18</i> - Hectare (Norwegian: Hektar)</li><li><i>MLT</i> - Milliliter (mL) (Norwegian: Milliliter)</li><li><i>HLT</i> - Hectoliter (hL) (Norwegian: Hektoliter)</li><li><i>DLT</i> - Deciliter (dL) (Norwegian: Desiliter)</li><li><i>AK</i> - Fathom (Norwegian: Favn)</li><li><i>XCR</i> - Crate (Norwegian: Kasse)</li><li><i>E14</i> - Kilocalorie (kcal) (Norwegian: Kilokalori)</li><li><i>MJ</i> - Megajoule (MJ) (Norwegian: Megajoule)</li><li><i>J57</i> - Barrel (petroleum) (Norwegian: Fat)</li><li><i>XJG</i> - Jug (Norwegian: Kanne)</li><li><i>XCT</i> - Carton (Norwegian: Kartong)</li><li><i>XSA</i> - Sack (Norwegian: Sekk)</li><li><i>XTU</i> - Tube (Norwegian: Tube)</li><li><i>WEE</i> - Week (Norwegian: Uke)</li><li><i>XCA</i> - Can (Rectangular) (Norwegian: Boks (Rektangulær))</li><li><i>XCN</i> - Container (Norwegian: Konteiner)</li><li><i>NAR</i> - Number of articles (Norwegian: Antall artikler)</li><li><i>M4</i> - Monetary value (Norwegian: Pengeverdi)</li><li><i>XVQ</i> - Bulk (Norwegian: Bulk)</li><li><i>P1</i> - Percent (%) (Norwegian: Prosent)</li><li><i>MFU</i> - Milk Forage Unit (Norwegian: Forenhet melk (FEm))</li><li><i>KMK</i> - Square kilometer (km^2) (Norwegian: Kvadratkilometer)</li><li><i>LM3</i> - Loose cubic meter (Norwegian: Løskubikkmeter)</li><li><i>FOT</i> - Foot (ft) (Norwegian: Fot)</li><li><i>FM3</i> - Solid cubic meter (Norwegian: Fastkubikkmeter)</li></ul>
type UnitOfMeasureType string

const (
	UnitOfMeasureTypeBx  UnitOfMeasureType = "BX"
	UnitOfMeasureTypeMtr UnitOfMeasureType = "MTR"
	UnitOfMeasureTypeKmt UnitOfMeasureType = "KMT"
	UnitOfMeasureTypeKgm UnitOfMeasureType = "KGM"
	UnitOfMeasureTypeEa  UnitOfMeasureType = "EA"
	UnitOfMeasureTypeLtr UnitOfMeasureType = "LTR"
	UnitOfMeasureTypeHur UnitOfMeasureType = "HUR"
	UnitOfMeasureTypeDay UnitOfMeasureType = "DAY"
	UnitOfMeasureTypeMtk UnitOfMeasureType = "MTK"
	UnitOfMeasureTypeMtq UnitOfMeasureType = "MTQ"
	UnitOfMeasureTypeTne UnitOfMeasureType = "TNE"
	UnitOfMeasureTypeMon UnitOfMeasureType = "MON"
	UnitOfMeasureTypeAnn UnitOfMeasureType = "ANN"
	UnitOfMeasureTypeQan UnitOfMeasureType = "QAN"
	UnitOfMeasureTypeNl  UnitOfMeasureType = "NL"
	UnitOfMeasureTypeXro UnitOfMeasureType = "XRO"
	UnitOfMeasureTypeXba UnitOfMeasureType = "XBA"
	UnitOfMeasureTypeXpl UnitOfMeasureType = "XPL"
	UnitOfMeasureTypeXpc UnitOfMeasureType = "XPC"
	UnitOfMeasureTypePr  UnitOfMeasureType = "PR"
	UnitOfMeasureTypeXci UnitOfMeasureType = "XCI"
	UnitOfMeasureTypeXbg UnitOfMeasureType = "XBG"
	UnitOfMeasureTypeSet UnitOfMeasureType = "SET"
	UnitOfMeasureTypeXty UnitOfMeasureType = "XTY"
	UnitOfMeasureTypeXof UnitOfMeasureType = "XOF"
	UnitOfMeasureTypeFtk UnitOfMeasureType = "FTK"
	UnitOfMeasureTypeKwh UnitOfMeasureType = "KWH"
	UnitOfMeasureTypeMwh UnitOfMeasureType = "MWH"
	UnitOfMeasureTypeLbr UnitOfMeasureType = "LBR"
	UnitOfMeasureTypeCmt UnitOfMeasureType = "CMT"
	UnitOfMeasureTypeDmt UnitOfMeasureType = "DMT"
	UnitOfMeasureTypeLm  UnitOfMeasureType = "LM"
	UnitOfMeasureTypeXpk UnitOfMeasureType = "XPK"
	UnitOfMeasureTypeGrm UnitOfMeasureType = "GRM"
	UnitOfMeasureTypeHgm UnitOfMeasureType = "HGM"
	UnitOfMeasureTypeXfl UnitOfMeasureType = "XFL"
	UnitOfMeasureTypeXbe UnitOfMeasureType = "XBE"
	UnitOfMeasureTypeE54 UnitOfMeasureType = "E54"
	UnitOfMeasureTypeMmt UnitOfMeasureType = "MMT"
	UnitOfMeasureTypeDaa UnitOfMeasureType = "DAA"
	UnitOfMeasureTypeH18 UnitOfMeasureType = "H18"
	UnitOfMeasureTypeMlt UnitOfMeasureType = "MLT"
	UnitOfMeasureTypeHlt UnitOfMeasureType = "HLT"
	UnitOfMeasureTypeDlt UnitOfMeasureType = "DLT"
	UnitOfMeasureTypeAk  UnitOfMeasureType = "AK"
	UnitOfMeasureTypeXcr UnitOfMeasureType = "XCR"
	UnitOfMeasureTypeE14 UnitOfMeasureType = "E14"
	UnitOfMeasureTypeMj  UnitOfMeasureType = "MJ"
	UnitOfMeasureTypeJ57 UnitOfMeasureType = "J57"
	UnitOfMeasureTypeXjg UnitOfMeasureType = "XJG"
	UnitOfMeasureTypeXct UnitOfMeasureType = "XCT"
	UnitOfMeasureTypeXsa UnitOfMeasureType = "XSA"
	UnitOfMeasureTypeXtu UnitOfMeasureType = "XTU"
	UnitOfMeasureTypeWee UnitOfMeasureType = "WEE"
	UnitOfMeasureTypeXca UnitOfMeasureType = "XCA"
	UnitOfMeasureTypeXcn UnitOfMeasureType = "XCN"
	UnitOfMeasureTypeNar UnitOfMeasureType = "NAR"
	UnitOfMeasureTypeM4  UnitOfMeasureType = "M4"
	UnitOfMeasureTypeXvq UnitOfMeasureType = "XVQ"
	UnitOfMeasureTypeP1  UnitOfMeasureType = "P1"
	UnitOfMeasureTypeMfu UnitOfMeasureType = "MFU"
	UnitOfMeasureTypeKmk UnitOfMeasureType = "KMK"
	UnitOfMeasureTypeLm3 UnitOfMeasureType = "LM3"
	UnitOfMeasureTypeFot UnitOfMeasureType = "FOT"
	UnitOfMeasureTypeFm3 UnitOfMeasureType = "FM3"
)

func (e UnitOfMeasureType) ToPointer() *UnitOfMeasureType {
	return &e
}

func (e *UnitOfMeasureType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BX":
		fallthrough
	case "MTR":
		fallthrough
	case "KMT":
		fallthrough
	case "KGM":
		fallthrough
	case "EA":
		fallthrough
	case "LTR":
		fallthrough
	case "HUR":
		fallthrough
	case "DAY":
		fallthrough
	case "MTK":
		fallthrough
	case "MTQ":
		fallthrough
	case "TNE":
		fallthrough
	case "MON":
		fallthrough
	case "ANN":
		fallthrough
	case "QAN":
		fallthrough
	case "NL":
		fallthrough
	case "XRO":
		fallthrough
	case "XBA":
		fallthrough
	case "XPL":
		fallthrough
	case "XPC":
		fallthrough
	case "PR":
		fallthrough
	case "XCI":
		fallthrough
	case "XBG":
		fallthrough
	case "SET":
		fallthrough
	case "XTY":
		fallthrough
	case "XOF":
		fallthrough
	case "FTK":
		fallthrough
	case "KWH":
		fallthrough
	case "MWH":
		fallthrough
	case "LBR":
		fallthrough
	case "CMT":
		fallthrough
	case "DMT":
		fallthrough
	case "LM":
		fallthrough
	case "XPK":
		fallthrough
	case "GRM":
		fallthrough
	case "HGM":
		fallthrough
	case "XFL":
		fallthrough
	case "XBE":
		fallthrough
	case "E54":
		fallthrough
	case "MMT":
		fallthrough
	case "DAA":
		fallthrough
	case "H18":
		fallthrough
	case "MLT":
		fallthrough
	case "HLT":
		fallthrough
	case "DLT":
		fallthrough
	case "AK":
		fallthrough
	case "XCR":
		fallthrough
	case "E14":
		fallthrough
	case "MJ":
		fallthrough
	case "J57":
		fallthrough
	case "XJG":
		fallthrough
	case "XCT":
		fallthrough
	case "XSA":
		fallthrough
	case "XTU":
		fallthrough
	case "WEE":
		fallthrough
	case "XCA":
		fallthrough
	case "XCN":
		fallthrough
	case "NAR":
		fallthrough
	case "M4":
		fallthrough
	case "XVQ":
		fallthrough
	case "P1":
		fallthrough
	case "MFU":
		fallthrough
	case "KMK":
		fallthrough
	case "LM3":
		fallthrough
	case "FOT":
		fallthrough
	case "FM3":
		*e = UnitOfMeasureType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UnitOfMeasureType: %v", v)
	}
}