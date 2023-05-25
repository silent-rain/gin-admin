/*文件上传*/
package dto

// Image 返回的图片信息
type Image struct {
	Name string `json:"name"` // 上传图片名称
	Url  string `json:"url"`  // URL 地址
}
