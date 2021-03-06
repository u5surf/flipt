// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/custom_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Custom placeholder fields.
type CustomPlaceholderFieldEnum_CustomPlaceholderField int32

const (
	// Not specified.
	CustomPlaceholderFieldEnum_UNSPECIFIED CustomPlaceholderFieldEnum_CustomPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	CustomPlaceholderFieldEnum_UNKNOWN CustomPlaceholderFieldEnum_CustomPlaceholderField = 1
	// Data Type: STRING. Required. Combination ID and ID2 must be unique per
	// offer.
	CustomPlaceholderFieldEnum_ID CustomPlaceholderFieldEnum_CustomPlaceholderField = 2
	// Data Type: STRING. Combination ID and ID2 must be unique per offer.
	CustomPlaceholderFieldEnum_ID2 CustomPlaceholderFieldEnum_CustomPlaceholderField = 3
	// Data Type: STRING. Required. Main headline with product name to be shown
	// in dynamic ad.
	CustomPlaceholderFieldEnum_ITEM_TITLE CustomPlaceholderFieldEnum_CustomPlaceholderField = 4
	// Data Type: STRING. Optional text to be shown in the image ad.
	CustomPlaceholderFieldEnum_ITEM_SUBTITLE CustomPlaceholderFieldEnum_CustomPlaceholderField = 5
	// Data Type: STRING. Optional description of the product to be shown in the
	// ad.
	CustomPlaceholderFieldEnum_ITEM_DESCRIPTION CustomPlaceholderFieldEnum_CustomPlaceholderField = 6
	// Data Type: STRING. Full address of your offer or service, including
	// postal code. This will be used to identify the closest product to the
	// user when there are multiple offers in the feed that are relevant to the
	// user.
	CustomPlaceholderFieldEnum_ITEM_ADDRESS CustomPlaceholderFieldEnum_CustomPlaceholderField = 7
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	CustomPlaceholderFieldEnum_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 8
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	CustomPlaceholderFieldEnum_FORMATTED_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 9
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	CustomPlaceholderFieldEnum_SALE_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 10
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	CustomPlaceholderFieldEnum_FORMATTED_SALE_PRICE CustomPlaceholderFieldEnum_CustomPlaceholderField = 11
	// Data Type: URL. Image to be displayed in the ad. Highly recommended for
	// image ads.
	CustomPlaceholderFieldEnum_IMAGE_URL CustomPlaceholderFieldEnum_CustomPlaceholderField = 12
	// Data Type: STRING. Used as a recommendation engine signal to serve items
	// in the same category.
	CustomPlaceholderFieldEnum_ITEM_CATEGORY CustomPlaceholderFieldEnum_CustomPlaceholderField = 13
	// Data Type: URL_LIST. Final URLs for the ad when using Upgraded
	// URLs. User will be redirected to these URLs when they click on an ad, or
	// when they click on a specific product for ads that have multiple
	// products.
	CustomPlaceholderFieldEnum_FINAL_URLS CustomPlaceholderFieldEnum_CustomPlaceholderField = 14
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	CustomPlaceholderFieldEnum_FINAL_MOBILE_URLS CustomPlaceholderFieldEnum_CustomPlaceholderField = 15
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	CustomPlaceholderFieldEnum_TRACKING_URL CustomPlaceholderFieldEnum_CustomPlaceholderField = 16
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	CustomPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS CustomPlaceholderFieldEnum_CustomPlaceholderField = 17
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	CustomPlaceholderFieldEnum_ANDROID_APP_LINK CustomPlaceholderFieldEnum_CustomPlaceholderField = 18
	// Data Type: STRING_LIST. List of recommended IDs to show together with
	// this item.
	CustomPlaceholderFieldEnum_SIMILAR_IDS CustomPlaceholderFieldEnum_CustomPlaceholderField = 19
	// Data Type: STRING. iOS app link.
	CustomPlaceholderFieldEnum_IOS_APP_LINK CustomPlaceholderFieldEnum_CustomPlaceholderField = 20
	// Data Type: INT64. iOS app store ID.
	CustomPlaceholderFieldEnum_IOS_APP_STORE_ID CustomPlaceholderFieldEnum_CustomPlaceholderField = 21
)

var CustomPlaceholderFieldEnum_CustomPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ID",
	3:  "ID2",
	4:  "ITEM_TITLE",
	5:  "ITEM_SUBTITLE",
	6:  "ITEM_DESCRIPTION",
	7:  "ITEM_ADDRESS",
	8:  "PRICE",
	9:  "FORMATTED_PRICE",
	10: "SALE_PRICE",
	11: "FORMATTED_SALE_PRICE",
	12: "IMAGE_URL",
	13: "ITEM_CATEGORY",
	14: "FINAL_URLS",
	15: "FINAL_MOBILE_URLS",
	16: "TRACKING_URL",
	17: "CONTEXTUAL_KEYWORDS",
	18: "ANDROID_APP_LINK",
	19: "SIMILAR_IDS",
	20: "IOS_APP_LINK",
	21: "IOS_APP_STORE_ID",
}
var CustomPlaceholderFieldEnum_CustomPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"ID":                   2,
	"ID2":                  3,
	"ITEM_TITLE":           4,
	"ITEM_SUBTITLE":        5,
	"ITEM_DESCRIPTION":     6,
	"ITEM_ADDRESS":         7,
	"PRICE":                8,
	"FORMATTED_PRICE":      9,
	"SALE_PRICE":           10,
	"FORMATTED_SALE_PRICE": 11,
	"IMAGE_URL":            12,
	"ITEM_CATEGORY":        13,
	"FINAL_URLS":           14,
	"FINAL_MOBILE_URLS":    15,
	"TRACKING_URL":         16,
	"CONTEXTUAL_KEYWORDS":  17,
	"ANDROID_APP_LINK":     18,
	"SIMILAR_IDS":          19,
	"IOS_APP_LINK":         20,
	"IOS_APP_STORE_ID":     21,
}

func (x CustomPlaceholderFieldEnum_CustomPlaceholderField) String() string {
	return proto.EnumName(CustomPlaceholderFieldEnum_CustomPlaceholderField_name, int32(x))
}
func (CustomPlaceholderFieldEnum_CustomPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_custom_placeholder_field_da1fc5fdf27a9f46, []int{0, 0}
}

// Values for Custom placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type CustomPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomPlaceholderFieldEnum) Reset()         { *m = CustomPlaceholderFieldEnum{} }
func (m *CustomPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*CustomPlaceholderFieldEnum) ProtoMessage()    {}
func (*CustomPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_custom_placeholder_field_da1fc5fdf27a9f46, []int{0}
}
func (m *CustomPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *CustomPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *CustomPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomPlaceholderFieldEnum.Merge(dst, src)
}
func (m *CustomPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_CustomPlaceholderFieldEnum.Size(m)
}
func (m *CustomPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.CustomPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.CustomPlaceholderFieldEnum_CustomPlaceholderField", CustomPlaceholderFieldEnum_CustomPlaceholderField_name, CustomPlaceholderFieldEnum_CustomPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/custom_placeholder_field.proto", fileDescriptor_custom_placeholder_field_da1fc5fdf27a9f46)
}

var fileDescriptor_custom_placeholder_field_da1fc5fdf27a9f46 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x8e, 0x12, 0x31,
	0x18, 0x15, 0x70, 0x41, 0x3e, 0x96, 0xa5, 0x14, 0x56, 0x8d, 0x66, 0x2f, 0xdc, 0x07, 0x18, 0x88,
	0x5e, 0xea, 0x4d, 0x99, 0x16, 0xd2, 0x30, 0x4c, 0x27, 0x6d, 0xd9, 0x75, 0x0d, 0x49, 0x83, 0xcc,
	0x38, 0x6e, 0x02, 0x0c, 0x61, 0x96, 0x7d, 0x1d, 0x13, 0x2f, 0x7d, 0x08, 0x1f, 0xc0, 0xf8, 0x50,
	0xa6, 0x1d, 0x16, 0xbc, 0x58, 0xbd, 0x69, 0xbe, 0xef, 0x9c, 0xef, 0x9c, 0xfe, 0x1d, 0xf8, 0x90,
	0x66, 0x59, 0xba, 0x4c, 0x7a, 0xf3, 0x38, 0xef, 0x15, 0xa5, 0xad, 0xee, 0xfb, 0xbd, 0x64, 0xbd,
	0x5b, 0xe5, 0xbd, 0xc5, 0x2e, 0xbf, 0xcb, 0x56, 0x66, 0xb3, 0x9c, 0x2f, 0x92, 0xaf, 0xd9, 0x32,
	0x4e, 0xb6, 0xe6, 0xcb, 0x6d, 0xb2, 0x8c, 0xbd, 0xcd, 0x36, 0xbb, 0xcb, 0xf0, 0x45, 0x21, 0xf1,
	0xe6, 0x71, 0xee, 0x1d, 0xd4, 0xde, 0x7d, 0xdf, 0x73, 0xea, 0xcb, 0x9f, 0x15, 0x78, 0xe5, 0x3b,
	0x87, 0xe8, 0x68, 0x30, 0xb4, 0x7a, 0xb6, 0xde, 0xad, 0x2e, 0xbf, 0x55, 0xe0, 0xf9, 0xe3, 0x34,
	0x6e, 0x41, 0x63, 0x1a, 0xaa, 0x88, 0xf9, 0x7c, 0xc8, 0x19, 0x45, 0x4f, 0x70, 0x03, 0x6a, 0xd3,
	0x70, 0x1c, 0x8a, 0xeb, 0x10, 0x95, 0x70, 0x15, 0xca, 0x9c, 0xa2, 0x32, 0xae, 0x41, 0x85, 0xd3,
	0xb7, 0xa8, 0x82, 0xcf, 0x00, 0xb8, 0x66, 0x13, 0xa3, 0xb9, 0x0e, 0x18, 0x7a, 0x8a, 0xdb, 0xd0,
	0x74, 0xbd, 0x9a, 0x0e, 0x0a, 0xe8, 0x04, 0x77, 0x01, 0x39, 0x88, 0x32, 0xe5, 0x4b, 0x1e, 0x69,
	0x2e, 0x42, 0x54, 0xc5, 0x08, 0x4e, 0x1d, 0x4a, 0x28, 0x95, 0x4c, 0x29, 0x54, 0xc3, 0x75, 0x38,
	0x89, 0x24, 0xf7, 0x19, 0x7a, 0x86, 0x3b, 0xd0, 0x1a, 0x0a, 0x39, 0x21, 0x5a, 0x33, 0x6a, 0x0a,
	0xb0, 0x6e, 0xb7, 0x52, 0x24, 0x60, 0xfb, 0x1e, 0xf0, 0x4b, 0xe8, 0x1e, 0x87, 0xfe, 0x62, 0x1a,
	0xb8, 0x09, 0x75, 0x3e, 0x21, 0x23, 0x66, 0xa6, 0x32, 0x40, 0xa7, 0x87, 0x33, 0xf9, 0x44, 0xb3,
	0x91, 0x90, 0x37, 0xa8, 0x69, 0xbd, 0x86, 0x3c, 0x24, 0x81, 0x9d, 0x50, 0xe8, 0x0c, 0x9f, 0x43,
	0xbb, 0xe8, 0x27, 0x62, 0xc0, 0x03, 0x56, 0xc0, 0x2d, 0x7b, 0x48, 0x2d, 0x89, 0x3f, 0xe6, 0xe1,
	0xc8, 0x79, 0x21, 0xfc, 0x02, 0x3a, 0xbe, 0x08, 0x35, 0xfb, 0xa8, 0xa7, 0x24, 0x30, 0x63, 0x76,
	0x73, 0x2d, 0x24, 0x55, 0xa8, 0x6d, 0x6f, 0x49, 0x42, 0x2a, 0x05, 0xa7, 0x86, 0x44, 0x91, 0x09,
	0x78, 0x38, 0x46, 0xd8, 0xbe, 0xa6, 0xe2, 0x13, 0x1e, 0x10, 0x69, 0x38, 0x55, 0xa8, 0xe3, 0xae,
	0x2d, 0xd4, 0x71, 0xa4, 0xeb, 0x9e, 0x67, 0x8f, 0x28, 0x2d, 0x24, 0x33, 0x9c, 0xa2, 0xf3, 0xc1,
	0xef, 0x12, 0xbc, 0x59, 0x64, 0x2b, 0xef, 0xbf, 0xdf, 0x3c, 0x78, 0xfd, 0xf8, 0x27, 0x46, 0x36,
	0x22, 0x51, 0xe9, 0xd3, 0x60, 0xaf, 0x4e, 0xb3, 0xe5, 0x7c, 0x9d, 0x7a, 0xd9, 0x36, 0xed, 0xa5,
	0xc9, 0xda, 0x05, 0xe8, 0x21, 0x72, 0x9b, 0xdb, 0xfc, 0x1f, 0x09, 0x7c, 0xef, 0xd6, 0xef, 0xe5,
	0xca, 0x88, 0x90, 0x1f, 0xe5, 0x8b, 0x51, 0x61, 0x45, 0xe2, 0xdc, 0x2b, 0x4a, 0x5b, 0x5d, 0xf5,
	0x3d, 0x9b, 0xa7, 0xfc, 0xd7, 0x03, 0x3f, 0x23, 0x71, 0x3e, 0x3b, 0xf0, 0xb3, 0xab, 0xfe, 0xcc,
	0xf1, 0x9f, 0xab, 0x6e, 0xd3, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0x53, 0x00, 0x84,
	0xf5, 0x02, 0x00, 0x00,
}
