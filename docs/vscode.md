# VsCode 插件

## Go Interface Annotations
- interface 接口提示插件

## koroFileHeader
- 文档注释插件
- 插件配置
```
// 自动添加文件头部注释
"fileheader.configObj": {
  "autoAdd": true, // 自动添加头部注释开启才能自动添加`
  "autoAlready": true, // 默认开启`
},
// 文件头部注释
"fileheader.customMade": {
  "version": "V1.0.0", //版本号
  "Author": "lilun", //作者
  "Date": "Do not edit", //文件创建创建时间
  "LastEditors": "lilun", //最后编辑作者
  "LastEditTime": "Do not Edit", //最后编辑时间
  "company": "轩田科技", //公司名称
  "Mailbox": "lilun@sharetek.com.cn",//邮箱
  "FilePath": "Do not edit", // 文件在项目中的相对路径 自动更新
  "Descripttion": "", //文本描述
}, 
//函数注释
"fileheader.cursorMode": {
  "description": "", //方法描述
  "param ": "", //参数
  "return": ""
},
```
- 使用快捷键
  - 文件头部注释： 在当前编辑文件中使用快捷键即可生成文件头部注释。
    - ctrl+alt+i
  - 函数注释： 将光标放在函数行或者将光标放在函数上方的空白行 使用快捷键即可生成函数注释。
    - ctrl+alt+t
