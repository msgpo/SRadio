package appindicator

//#cgo CFLAGS: -I/usr/include/libappindicator3-0.1
//#cgo LDFLAGS: -lappindicator3
//#cgo pkg-config: gtk+-3.0
//#include <stdlib.h>
//#include <gtk/gtk.h>
//#include <libappindicator/app-indicator.h>
import "C"
import "unsafe"
import "github.com/gotk3/gotk3/gtk"

// AppIndicatorCatogory
type Category int

const (
	CategoryApplicationStatus Category = C.APP_INDICATOR_CATEGORY_APPLICATION_STATUS
	CategoryCommunications             = C.APP_INDICATOR_CATEGORY_COMMUNICATIONS
	CategorySystemServices             = C.APP_INDICATOR_CATEGORY_SYSTEM_SERVICES
	CategoryHardware                   = C.APP_INDICATOR_CATEGORY_HARDWARE
	CategoryOther                      = C.APP_INDICATOR_CATEGORY_OTHER
)

// AppIndicatorStatus
type Status int

const (
	StatusPassive   Status = C.APP_INDICATOR_STATUS_PASSIVE
	StatusActive           = C.APP_INDICATOR_STATUS_ACTIVE
	StatusAttention        = C.APP_INDICATOR_STATUS_ATTENTION
)

type AppIndicator struct {
	IndicatorPtr unsafe.Pointer
}

// Creates a new AppIndicator.
func NewAppIndicator(id, iconName string, category Category) *AppIndicator {
	idString := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(idString))
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))

	indicator := unsafe.Pointer(C.app_indicator_new(idString, iconNameString, C.AppIndicatorCategory(category)))

	return &AppIndicator{
		IndicatorPtr: indicator,
	}
}

// Creates a new AppIndicator using a specific icon path.
func NewAppIndicatorWithPath(id, iconName, iconPath string, category int) *AppIndicator {
	idString := (*C.gchar)(unsafe.Pointer(C.CString(id)))
	defer C.free(unsafe.Pointer(idString))
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))
	iconPathString := (*C.gchar)(unsafe.Pointer(C.CString(iconPath)))
	defer C.free(unsafe.Pointer(iconPathString))

	indicator := unsafe.Pointer(C.app_indicator_new_with_path(idString, iconNameString, C.AppIndicatorCategory(category), iconPathString))

	return &AppIndicator{
		IndicatorPtr: indicator,
	}
}

// Sets the status of the indicator.
func (indicator *AppIndicator) SetStatus(status Status) {
	C.app_indicator_set_status((*C.AppIndicator)(indicator.IndicatorPtr), C.AppIndicatorStatus(status))
}

// Sets the attention icon of the indicator.
func (indicator *AppIndicator) SetAttentionIcon(iconName, iconDescription string) {
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))
	iconDescriptionString := (*C.gchar)(unsafe.Pointer(C.CString(iconDescription)))
	defer C.free(unsafe.Pointer(iconDescriptionString))

	C.app_indicator_set_attention_icon_full((*C.AppIndicator)(indicator.IndicatorPtr), iconNameString, iconDescriptionString)
}

func (indicator *AppIndicator) SetMenu(menu *gtk.Menu) {
	menu.ShowAll()
	indicator.C_SetMenu(unsafe.Pointer(menu.Native()))
}

// Sets the menu that should be shown the indicator is clicked on in the panel. An application indicator will not be rendered unless it has a menu.
// This is the C version of the function and should only be used it not using GoGtk.
func (indicator *AppIndicator) C_SetMenu(menu unsafe.Pointer) {
	C.app_indicator_set_menu((*C.AppIndicator)(indicator.IndicatorPtr), (*C.GtkMenu)(menu))
}

// Sets the default icon to use when the status is active but not set to attention. In most cases this should be the application icon for the program.
func (indicator *AppIndicator) SetIcon(iconName, iconDescription string) {
	iconNameString := (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(iconNameString))
	iconDescriptionString := (*C.gchar)(unsafe.Pointer(C.CString(iconDescription)))
	defer C.free(unsafe.Pointer(iconDescriptionString))

	C.app_indicator_set_icon_full((*C.AppIndicator)(indicator.IndicatorPtr), iconNameString, iconDescriptionString)
}

// Sets the path to use when searching for icons.
func (indicator *AppIndicator) SetIconThemePath(iconPath string) {
	iconPathString := (*C.gchar)(unsafe.Pointer(C.CString(iconPath)))
	defer C.free(unsafe.Pointer(iconPathString))

	C.app_indicator_set_icon_theme_path((*C.AppIndicator)(indicator.IndicatorPtr), iconPathString)
}

// Sets the label and guide of the indicator.
func (indicator *AppIndicator) SetLabel(label, guide string) {
	labelString := (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(labelString))
	guideString := (*C.gchar)(unsafe.Pointer(C.CString(guide)))
	defer C.free(unsafe.Pointer(guideString))

	C.app_indicator_set_label((*C.AppIndicator)(indicator.IndicatorPtr), labelString, guideString)
}

// Sets the ordering index for the indicator which affects the placement of it on the panel. For almost all app indicators this is not he function you're looking for.
func (indicator *AppIndicator) SetOrderingIndex(index uint32) {
	C.app_indicator_set_ordering_index((*C.AppIndicator)(indicator.IndicatorPtr), C.guint32(index))
}

// Sets the menu to be activated when a secondary activation event (i.e middle-click) is emitted over the indicator icon/label.
// This is the C version of the function and should only be used it not using GoGtk.
func (indicator *AppIndicator) C_SetSecondaryActivateTarget(menu unsafe.Pointer) {
	C.app_indicator_set_secondary_activate_target((*C.AppIndicator)(indicator.IndicatorPtr), (*C.GtkWidget)(menu))
}

// Sets the title of the indicator, or how it should be referred in a human readable form.
func (indicator *AppIndicator) SetTitle(title string) {
	titleString := (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(titleString))

	C.app_indicator_set_title((*C.AppIndicator)(indicator.IndicatorPtr), titleString)
}

func (indicator AppIndicator) GetId() string {
	retVal := C.app_indicator_get_id((*C.AppIndicator)(indicator.IndicatorPtr))
	return C.GoString((*C.char)(unsafe.Pointer(retVal)))
}

func (indicator AppIndicator) GetCategory() Category {
	return Category(C.app_indicator_get_category((*C.AppIndicator)(indicator.IndicatorPtr)))
}

func (indicator AppIndicator) GetStatus() Status {
	return Status(C.app_indicator_get_status((*C.AppIndicator)(indicator.IndicatorPtr)))
}

// Gets the current icon and description that is associated with the indicator.
func (indicator AppIndicator) GetIcon() (string, string) {
	retVal := C.app_indicator_get_icon((*C.AppIndicator)(indicator.IndicatorPtr))
	iconName := C.GoString((*C.char)(unsafe.Pointer(retVal)))

	retVal = C.app_indicator_get_icon_desc((*C.AppIndicator)(indicator.IndicatorPtr))
	iconDesc := C.GoString((*C.char)(unsafe.Pointer(retVal)))

	return iconName, iconDesc
}

func (indicator AppIndicator) GetIconThemePath() string {
	retVal := C.app_indicator_get_icon_theme_path((*C.AppIndicator)(indicator.IndicatorPtr))
	return C.GoString((*C.char)(unsafe.Pointer(retVal)))
}

// Gets the current attention icon and description that is associated with the indicator.
func (indicator AppIndicator) GetAttentionIcon() (string, string) {
	retVal := C.app_indicator_get_attention_icon((*C.AppIndicator)(indicator.IndicatorPtr))
	iconName := C.GoString((*C.char)(unsafe.Pointer(retVal)))

	retVal = C.app_indicator_get_attention_icon_desc((*C.AppIndicator)(indicator.IndicatorPtr))
	iconDesc := C.GoString((*C.char)(unsafe.Pointer(retVal)))

	return iconName, iconDesc
}

// This is the C version of the function and should only be used it not using GoGtk.
func (indicator AppIndicator) C_GetMenu() unsafe.Pointer {
	return unsafe.Pointer(C.app_indicator_get_menu((*C.AppIndicator)(indicator.IndicatorPtr)))
}

// Gets the current label and guide that is associated with the indicator.
func (indicator AppIndicator) GetLabel() (string, string) {
	retVal := C.app_indicator_get_label((*C.AppIndicator)(indicator.IndicatorPtr))
	label := C.GoString((*C.char)(unsafe.Pointer(retVal)))

	retVal = C.app_indicator_get_label_guide((*C.AppIndicator)(indicator.IndicatorPtr))
	guide := C.GoString((*C.char)(unsafe.Pointer(retVal)))

	return label, guide
}

func (indicator AppIndicator) GetOrderingIndex() uint32 {
	return uint32(C.app_indicator_get_ordering_index((*C.AppIndicator)(indicator.IndicatorPtr)))
}

// This is the C version of the function and should only be used it not using GoGtk.
func (indicator AppIndicator) C_GetSecondaryActivateTarget() unsafe.Pointer {
	return unsafe.Pointer(C.app_indicator_get_secondary_activate_target((*C.AppIndicator)(indicator.IndicatorPtr)))
}

func (indicator AppIndicator) GetTitle() string {
	retVal := C.app_indicator_get_title((*C.AppIndicator)(indicator.IndicatorPtr))
	return C.GoString((*C.char)(unsafe.Pointer(retVal)))
}

func (indicator *AppIndicator) BuildMenuFromDesktop(filePath, profile string) {
	filePathString := (*C.gchar)(unsafe.Pointer(C.CString(filePath)))
	defer C.free(unsafe.Pointer(filePathString))
	profileString := (*C.gchar)(unsafe.Pointer(C.CString(profile)))
	defer C.free(unsafe.Pointer(profileString))

	C.app_indicator_build_menu_from_desktop((*C.AppIndicator)(indicator.IndicatorPtr), filePathString, profileString)
}
