### 双色球桌面小程序

##### 程序运行步骤说明

- 采用github.com/lxn/walk设置Gui。

- 嵌入syso文件，这是微软要求的文件。

  ```
   rsrc -manifest doubleball.manifest -ico doubleball.ico -o doubleball.syso
   //rsrc 通过go get github.com/akavel/rsrc获取编译
  ```

- 生成双色球运行程序。

  ```
  go build -ldflags="-H windowsgui"
  //-ldflags="-H windowsgui" 命令是去除了命令行界面的调用，直接以gui运行
  ```

  



