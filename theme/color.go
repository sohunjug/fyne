package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
)

// Keep in mind to add new constants to the tests at theme/theme_test.go as well as test/theme_test.go.
const (
	// ColorRed is the red primary color name.
	//
	// Since: 1.4
	ColorRed = "red"
	// ColorOrange is the orange primary color name.
	//
	// Since: 1.4
	ColorOrange = "orange"
	// ColorYellow is the yellow primary color name.
	//
	// Since: 1.4
	ColorYellow = "yellow"
	// ColorGreen is the green primary color name.
	//
	// Since: 1.4
	ColorGreen = "green"
	// ColorBlue is the blue primary color name.
	//
	// Since: 1.4
	ColorBlue = "blue"
	// ColorPurple is the purple primary color name.
	//
	// Since: 1.4
	ColorPurple = "purple"
	// ColorBrown is the brown primary color name.
	//
	// Since: 1.4
	ColorBrown = "brown"
	// ColorGray is the gray primary color name.
	//
	// Since: 1.4
	ColorGray = "gray"

	// ColorNameBackground is the name of theme lookup for background color.
	//
	// Since: 2.0
	ColorNameBackground fyne.ThemeColorName = "background"

	// ColorNameButton is the name of theme lookup for button color.
	//
	// Since: 2.0
	ColorNameButton fyne.ThemeColorName = "button"

	// ColorNameDisabledButton is the name of theme lookup for disabled button color.
	//
	// Since: 2.0
	ColorNameDisabledButton fyne.ThemeColorName = "disabledButton"

	// ColorNameDisabled is the name of theme lookup for disabled foreground color.
	//
	// Since: 2.0
	ColorNameDisabled fyne.ThemeColorName = "disabled"

	// ColorNameError is the name of theme lookup for error color.
	//
	// Since: 2.0
	ColorNameError fyne.ThemeColorName = "error"

	// ColorNameErrorForeground is the name of theme lookup for a contrast color to the error color.
	//
	// Since: 2.5
	ColorNameErrorForeground fyne.ThemeColorName = "errorForeground"

	// ColorNameFocus is the name of theme lookup for focus color.
	//
	// Since: 2.0
	ColorNameFocus fyne.ThemeColorName = "focus"

	// ColorNameForeground is the name of theme lookup for foreground color.
	//
	// Since: 2.0
	ColorNameForeground fyne.ThemeColorName = "foreground"

	// ColorNameHeaderBackground is the name of theme lookup for background color of a collection header.
	//
	// Since: 2.4
	ColorNameHeaderBackground fyne.ThemeColorName = "headerBackground"

	// ColorNameHover is the name of theme lookup for hover color.
	//
	// Since: 2.0
	ColorNameHover fyne.ThemeColorName = "hover"

	// ColorNameHyperlink is the name of theme lookup for hyperlink color.
	//
	// Since: 2.4
	ColorNameHyperlink fyne.ThemeColorName = "hyperlink"

	// ColorNameInputBackground is the name of theme lookup for background color of an input field.
	//
	// Since: 2.0
	ColorNameInputBackground fyne.ThemeColorName = "inputBackground"

	// ColorNameInputBorder is the name of theme lookup for border color of an input field.
	//
	// Since: 2.3
	ColorNameInputBorder fyne.ThemeColorName = "inputBorder"

	// ColorNameMenuBackground is the name of theme lookup for background color of menus.
	//
	// Since: 2.3
	ColorNameMenuBackground fyne.ThemeColorName = "menuBackground"

	// ColorNameOverlayBackground is the name of theme lookup for background color of overlays like dialogs.
	//
	// Since: 2.3
	ColorNameOverlayBackground fyne.ThemeColorName = "overlayBackground"

	// ColorNamePlaceHolder is the name of theme lookup for placeholder text color.
	//
	// Since: 2.0
	ColorNamePlaceHolder fyne.ThemeColorName = "placeholder"

	// ColorNamePressed is the name of theme lookup for the tap overlay color.
	//
	// Since: 2.0
	ColorNamePressed fyne.ThemeColorName = "pressed"

	// ColorNamePrimary is the name of theme lookup for primary color.
	//
	// Since: 2.0
	ColorNamePrimary fyne.ThemeColorName = "primary"

	// ColorNamePrimaryForeground is the name of theme lookup for a contrast color to the primary color.
	//
	// Since: 2.5
	ColorNamePrimaryForeground fyne.ThemeColorName = "primaryForeground"

	// ColorNameScrollBar is the name of theme lookup for scrollbar color.
	//
	// Since: 2.0
	ColorNameScrollBar fyne.ThemeColorName = "scrollBar"

	// ColorNameSelection is the name of theme lookup for selection color.
	//
	// Since: 2.1
	ColorNameSelection fyne.ThemeColorName = "selection"

	// ColorNameSeparator is the name of theme lookup for separator bars.
	//
	// Since: 2.3
	ColorNameSeparator fyne.ThemeColorName = "separator"

	// ColorNameShadow is the name of theme lookup for shadow color.
	//
	// Since: 2.0
	ColorNameShadow fyne.ThemeColorName = "shadow"

	// ColorNameSuccess is the name of theme lookup for success color.
	//
	// Since: 2.3
	ColorNameSuccess fyne.ThemeColorName = "success"

	// ColorNameSuccessForeground is the name of theme lookup for a contrast color to the success color.
	//
	// Since: 2.5
	ColorNameSuccessForeground fyne.ThemeColorName = "successForeground"

	// ColorNameWarning is the name of theme lookup for warning color.
	//
	// Since: 2.3
	ColorNameWarning fyne.ThemeColorName = "warning"

	// ColorNameWarningForeground is the name of theme lookup for a contrast color to the warning color.
	//
	// Since: 2.5
	ColorNameWarningForeground fyne.ThemeColorName = "warningForeground"
)

var (
	colorDarkBackground        = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkButton            = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkDisabled          = color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	colorDarkDisabledButton    = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkError             = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorDarkErrorForeground   = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkForeground        = color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	colorDarkHeaderBackground  = color.NRGBA{R: 0x1b, G: 0x1b, B: 0x1b, A: 0xff}
	colorDarkHover             = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x0f}
	colorDarkInputBackground   = color.NRGBA{R: 0x20, G: 0x20, B: 0x23, A: 0xff}
	colorDarkInputBorder       = color.NRGBA{R: 0x39, G: 0x39, B: 0x3a, A: 0xff}
	colorDarkMenuBackground    = color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
	colorDarkOverlayBackground = color.NRGBA{R: 0x18, G: 0x1d, B: 0x25, A: 0xff}
	colorDarkPlaceholder       = color.NRGBA{R: 0xb2, G: 0xb2, B: 0xb2, A: 0xff}
	colorDarkPressed           = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	colorDarkScrollBar         = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x99}
	colorDarkSeparator         = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	colorDarkShadow            = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x66}
	colorDarkSuccess           = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	colorDarkSuccessForeground = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}
	colorDarkWarning           = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	colorDarkWarningForeground = color.NRGBA{R: 0x17, G: 0x17, B: 0x18, A: 0xff}

	colorLightBackground        = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightButton            = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightDisabled          = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightDisabledButton    = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightError             = color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	colorLightErrorForeground   = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightForeground        = color.NRGBA{R: 0x56, G: 0x56, B: 0x56, A: 0xff}
	colorLightHeaderBackground  = color.NRGBA{R: 0xf9, G: 0xf9, B: 0xf9, A: 0xff}
	colorLightHover             = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x0f}
	colorLightInputBackground   = color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}
	colorLightInputBorder       = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightMenuBackground    = color.NRGBA{R: 0xf5, G: 0xf5, B: 0xf5, A: 0xff}
	colorLightPlaceholder       = color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xff}
	colorLightPressed           = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x19}
	colorLightScrollBar         = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99}
	colorLightSeparator         = color.NRGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff}
	colorLightShadow            = color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x33}
	colorLightSuccess           = color.NRGBA{R: 0x43, G: 0xf4, B: 0x36, A: 0xff}
	colorLightSuccessForeground = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	colorLightWarning           = color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	colorLightWarningForeground = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
)

// BackgroundColor returns the theme's background color.
func BackgroundColor() color.Color {
	return safeColorLookup(ColorNameBackground, currentVariant())
}

// ButtonColor returns the theme's standard button color.
func ButtonColor() color.Color {
	return safeColorLookup(ColorNameButton, currentVariant())
}

// Color looks up the named colour for current theme and variant.
//
// Since: 2.5
func Color(name fyne.ThemeColorName) color.Color {
	return safeColorLookup(name, currentVariant())
}

// ColorForWidget looks up the named colour for the requested widget using the current theme and variant.
// If the widget theme has been overridden that theme will be used.
//
// Since: 2.5
func ColorForWidget(name fyne.ThemeColorName, w fyne.Widget) color.Color {
	return CurrentForWidget(w).Color(name, currentVariant())
}

// DisabledButtonColor returns the theme's disabled button color.
func DisabledButtonColor() color.Color {
	return safeColorLookup(ColorNameDisabledButton, currentVariant())
}

// DisabledColor returns the foreground color for a disabled UI element.
//
// Since: 2.0
func DisabledColor() color.Color {
	return safeColorLookup(ColorNameDisabled, currentVariant())
}

// DisabledTextColor returns the theme's disabled text color - this is actually the disabled color since 1.4.
//
// Deprecated: Use theme.DisabledColor() colour instead.
func DisabledTextColor() color.Color {
	return DisabledColor()
}

// ErrorColor returns the theme's error foreground color.
//
// Since: 2.0
func ErrorColor() color.Color {
	return safeColorLookup(ColorNameError, currentVariant())
}

// FocusColor returns the color used to highlight a focused widget.
func FocusColor() color.Color {
	return safeColorLookup(ColorNameFocus, currentVariant())
}

// ForegroundColor returns the theme's standard foreground color for text and icons.
//
// Since: 2.0
func ForegroundColor() color.Color {
	return safeColorLookup(ColorNameForeground, currentVariant())
}

// HeaderBackgroundColor returns the color used to draw underneath collection headers.
//
// Since: 2.4
func HeaderBackgroundColor() color.Color {
	return Current().Color(ColorNameHeaderBackground, currentVariant())
}

// HoverColor returns the color used to highlight interactive elements currently under a cursor.
func HoverColor() color.Color {
	return safeColorLookup(ColorNameHover, currentVariant())
}

// HyperlinkColor returns the color used for the Hyperlink widget and hyperlink text elements.
func HyperlinkColor() color.Color {
	return safeColorLookup(ColorNameHyperlink, currentVariant())
}

// InputBackgroundColor returns the color used to draw underneath input elements.
func InputBackgroundColor() color.Color {
	return Current().Color(ColorNameInputBackground, currentVariant())
}

// InputBorderColor returns the color used to draw underneath input elements.
//
// Since: 2.3
func InputBorderColor() color.Color {
	return Current().Color(ColorNameInputBorder, currentVariant())
}

// MenuBackgroundColor returns the theme's background color for menus.
//
// Since: 2.3
func MenuBackgroundColor() color.Color {
	return safeColorLookup(ColorNameMenuBackground, currentVariant())
}

// PrimaryForegroundColorNamed returns a theme specific color used for text and icons against the named primary color.
//
// Since: 2.5
func PrimaryForegroundColorNamed(name string) color.Color {
	switch name {
	case ColorRed:
		return colorLightBackground
	case ColorOrange:
		return colorDarkBackground
	case ColorYellow:
		return colorDarkBackground
	case ColorGreen:
		return colorDarkBackground
	case ColorPurple:
		return colorLightBackground
	case ColorBrown:
		return colorLightBackground
	case ColorGray:
		return colorDarkBackground
	}

	// We return the “on” value for ColorBlue for every other value.
	// There is no need to have it in the switch above.
	return colorLightBackground
}

// OverlayBackgroundColor returns the theme's background color for overlays like dialogs.
//
// Since: 2.3
func OverlayBackgroundColor() color.Color {
	return safeColorLookup(ColorNameOverlayBackground, currentVariant())
}

// PlaceHolderColor returns the theme's standard text color.
func PlaceHolderColor() color.Color {
	return safeColorLookup(ColorNamePlaceHolder, currentVariant())
}

// PressedColor returns the color used to overlap tapped features.
//
// Since: 2.0
func PressedColor() color.Color {
	return safeColorLookup(ColorNamePressed, currentVariant())
}

// PrimaryColor returns the color used to highlight primary features.
func PrimaryColor() color.Color {
	return safeColorLookup(ColorNamePrimary, currentVariant())
}

// PrimaryColorNamed returns a theme specific color value for a named primary color.
//
// Since: 1.4
func PrimaryColorNamed(name string) color.Color {
	switch name {
	case ColorRed:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case ColorOrange:
		return color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	case ColorYellow:
		return color.NRGBA{R: 0xff, G: 0xeb, B: 0x3b, A: 0xff}
	case ColorGreen:
		return color.NRGBA{R: 0x8b, G: 0xc3, B: 0x4a, A: 0xff}
	case ColorPurple:
		return color.NRGBA{R: 0x9c, G: 0x27, B: 0xb0, A: 0xff}
	case ColorBrown:
		return color.NRGBA{R: 0x79, G: 0x55, B: 0x48, A: 0xff}
	case ColorGray:
		return color.NRGBA{R: 0x9e, G: 0x9e, B: 0x9e, A: 0xff}
	}

	// We return the value for ColorBlue for every other value.
	// There is no need to have it in the switch above.
	return color.NRGBA{R: 0x29, G: 0x6f, B: 0xf6, A: 0xff}
}

// PrimaryColorNames returns a list of the standard primary color options.
//
// Since: 1.4
func PrimaryColorNames() []string {
	return []string{ColorRed, ColorOrange, ColorYellow, ColorGreen, ColorBlue, ColorPurple, ColorBrown, ColorGray}
}

// ScrollBarColor returns the color (and translucency) for a scrollBar.
func ScrollBarColor() color.Color {
	return safeColorLookup(ColorNameScrollBar, currentVariant())
}

// SelectionColor returns the color for a selected element.
//
// Since: 2.1
func SelectionColor() color.Color {
	return safeColorLookup(ColorNameSelection, currentVariant())
}

// SeparatorColor returns the color for the separator element.
//
// Since: 2.3
func SeparatorColor() color.Color {
	return safeColorLookup(ColorNameSeparator, currentVariant())
}

// ShadowColor returns the color (and translucency) for shadows used for indicating elevation.
func ShadowColor() color.Color {
	return safeColorLookup(ColorNameShadow, currentVariant())
}

// SuccessColor returns the theme's success foreground color.
//
// Since: 2.3
func SuccessColor() color.Color {
	return safeColorLookup(ColorNameSuccess, currentVariant())
}

// WarningColor returns the theme's warning foreground color.
//
// Since: 2.3
func WarningColor() color.Color {
	return safeColorLookup(ColorNameWarning, currentVariant())
}

func safeColorLookup(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	col := Current().Color(n, v)
	if col == nil {
		fyne.LogError("Loaded theme returned nil color", nil)
		return fallbackColor
	}
	return col
}
