package demoxml2

type Alignment struct {
	Horizontal string `xml:"horizontal,attr"`
	Vertical   string `xml:"vertical,attr"`
}
type B struct {
	Val string `xml:"val,attr"`
}
type BgColor struct {
	Auto string `xml:"auto,attr"`
}
type Border struct {
	Bottom Bottom `xml:"bottom,attr"`
	Left   Left   `xml:"left,attr"`
	Right  Right  `xml:"right,attr"`
	Top    Top    `xml:"top,attr"`
}
type Borders struct {
	Border []Border `xml:"border,attr"`
	Count  string   `xml:"count,attr"`
}
type Bottom struct {
	Color Color  `xml:"color,attr"`
	Style string `xml:"style,attr"`
}
type CellStyle struct {
	BuiltinID string `xml:"builtinId,attr"`
	Name      string `xml:"name,attr"`
	XfID      string `xml:"xfId,attr"`
}
type CellStyleXfs struct {
	Count string `xml:"count,attr"`
	Xf    Xf     `xml:"xf,attr"`
}
type CellStyles struct {
	CellStyle CellStyle `xml:"cellStyle,attr"`
	Count     string    `xml:"count,attr"`
}
type CellXfs struct {
	Count string `xml:"count,attr"`
	Xf    []Xf   `xml:"xf,attr"`
}
type Color struct {
	Indexed string `xml:"indexed,attr"`
}
type Colors struct {
	IndexedColors IndexedColors `xml:"indexedColors,attr"`
}
type Dxfs struct {
	Count string `xml:"count,attr"`
}
type FgColor struct {
	Indexed string `xml:"indexed,attr"`
}
type Fill struct {
	PatternFill PatternFill `xml:"patternFill,attr"`
}
type Fills struct {
	Count string `xml:"count,attr"`
	Fill  []Fill `xml:"fill,attr"`
}
type Font struct {
	B     B     `xml:"b,attr"`
	Color Color `xml:"color,attr"`
	Name  Name  `xml:"name,attr"`
	Sz    Sz    `xml:"sz,attr"`
}
type Fonts struct {
	Count string `xml:"count,attr"`
	Font  []Font `xml:"font,attr"`
}
type IndexedColors struct {
	RgbColor []RgbColor `xml:"rgbColor,attr"`
}
type Left struct {
	Color Color  `xml:"color,attr"`
	Style string `xml:"style,attr"`
}
type Name struct {
	Val string `xml:"val,attr"`
}
type NumFmt struct {
	FormatCode string `xml:"formatCode,attr"`
	NumFmtID   string `xml:"numFmtId,attr"`
}
type NumFmts struct {
	Count  string `xml:"count,attr"`
	NumFmt NumFmt `xml:"numFmt,attr"`
}
type PatternFill struct {
	BgColor     BgColor `xml:"bgColor,attr"`
	FgColor     FgColor `xml:"fgColor,attr"`
	PatternType string  `xml:"patternType,attr"`
}
type RgbColor struct {
	Rgb string `xml:"rgb,attr"`
}
type Right struct {
	Color Color  `xml:"color,attr"`
	Style string `xml:"style,attr"`
}
type StyleSheet struct {
	Borders      Borders      `xml:"borders,attr"`
	CellStyleXfs CellStyleXfs `xml:"cellStyleXfs,attr"`
	CellStyles   CellStyles   `xml:"cellStyles,attr"`
	CellXfs      CellXfs      `xml:"cellXfs,attr"`
	Colors       Colors       `xml:"colors,attr"`
	Dxfs         Dxfs         `xml:"dxfs,attr"`
	Fills        Fills        `xml:"fills,attr"`
	Fonts        Fonts        `xml:"fonts,attr"`
	NumFmts      NumFmts      `xml:"numFmts,attr"`
	TableStyles  TableStyles  `xml:"tableStyles,attr"`
	Xmlns        string       `xml:"xmlns,attr"`
}
type Sz struct {
	Val string `xml:"val,attr"`
}
type TableStyles struct {
	Count string `xml:"count,attr"`
}
type Top struct {
	Color Color  `xml:"color,attr"`
	Style string `xml:"style,attr"`
}
type Xf struct {
	Alignment         Alignment `xml:"alignment,attr"`
	ApplyAlignment    string    `xml:"applyAlignment,attr"`
	ApplyBorder       string    `xml:"applyBorder,attr"`
	ApplyFill         string    `xml:"applyFill,attr"`
	ApplyFont         string    `xml:"applyFont,attr"`
	ApplyNumberFormat string    `xml:"applyNumberFormat,attr"`
	ApplyProtection   string    `xml:"applyProtection,attr"`
	BorderID          string    `xml:"borderId,attr"`
	FillID            string    `xml:"fillId,attr"`
	FontID            string    `xml:"fontId,attr"`
	NumFmtID          string    `xml:"numFmtId,attr"`
}
