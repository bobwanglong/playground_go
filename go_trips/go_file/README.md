### go 操作文件可以使用的库

```shell
Go语言官方库：os、io/ioutil、bufio涵盖了文件操作的所有场景，

os提供了对文件IO直接调用的方法，bufio提供缓冲区操作文件的方法，io/ioutil也提供对文件IO直接调用的方法，不过Go语言在Go1.16版本已经弃用了io/ioutil库，这个io/ioutil包是一个定义不明确且难以理解的东西集合。该软件包提供的所有功能都已移至其他软件包，所以io/ioutil中操作文件的方法都在io库有相同含义的方法，大家以后在使用到ioutil中的方法是可以通过注释在其他包找到对应的方法。
```
### 写文件
os/ioutil也提供WriteFile方法可以快速处理创建/打开文件/写数据/关闭文件，使用示例如下
```go
func writeAll(filename string) error {
 err := os.WriteFile("asong.txt", []byte("Hi asong\n"), 0666)
 if err != nil {
  return err
 }
 return nil
}
```
#### 按行写文件
os、buffo写数据都没有提供按行写入的方法，所以我们可以在调用os.WriteString、bufio.WriteString方法是在数据中加入换行符即可，来看示例
```go
import (
 "bufio"
 "log"
 "os"
)
// 直接操作IO
func writeLine(filename string) error {
 data := []string{
  "asong",
  "test",
  "123",
 }
 f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
 if err != nil{
  return err
 }

 for _, line := range data{
  _,err := f.WriteString(line + "\n")
  if err != nil{
   return err
  }
 }
 f.Close()
 return nil
}
// 使用缓存区写入
func writeLine2(filename string) error {
 file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
 if err != nil {
  return err
 }

 // 为这个文件创建buffered writer
 bufferedWriter := bufio.NewWriter(file)
 
 for i:=0; i < 2; i++{
  // 写字符串到buffer
  bytesWritten, err := bufferedWriter.WriteString(
   "asong真帅\n",
  )
  if err != nil {
   return err
  }
  log.Printf("Bytes written: %d\n", bytesWritten)
 }
 // 写内存buffer到硬盘
 err = bufferedWriter.Flush()
 if err != nil{
  return err
 }

 file.Close()
 return nil
}
```
#### 偏移量写入
```go
import "os"

func writeAt(filename string) error {
 data := []byte{
  0x41, // A
  0x73, // s
  0x20, // space
  0x20, // space
  0x67, // g
 }
 f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
 if err != nil{
  return err
 }
 _, err = f.Write(data)
 if err != nil{
  return err
 }

 replaceSplace := []byte{
  0x6F, // o
  0x6E, // n
 }
 _, err = f.WriteAt(replaceSplace, 2)
 if err != nil{
  return err
 }
 f.Close()
 return nil
}
```
### 缓存区写入
os库中的方法对文件都是直接的IO操作，频繁的IO操作会增加CPU的中断频率，所以我们可以使用内存缓存区来减少IO操作，在写字节到硬盘前使用内存缓存，当内存缓存区的容量到达一定数值时在写内存数据buffer到硬盘，bufio就是这样示一个库，来个例子我们看一下怎么使用：
```go
import (
 "bufio"
 "log"
 "os"
)

func writeBuffer(filename string) error {
 file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
 if err != nil {
  return err
 }

 // 为这个文件创建buffered writer
 bufferedWriter := bufio.NewWriter(file)

 // 写字符串到buffer
 bytesWritten, err := bufferedWriter.WriteString(
  "asong真帅\n",
 )
 if err != nil {
  return err
 }
 log.Printf("Bytes written: %d\n", bytesWritten)

 // 检查缓存中的字节数
 unflushedBufferSize := bufferedWriter.Buffered()
 log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

 // 还有多少字节可用（未使用的缓存大小）
 bytesAvailable := bufferedWriter.Available()
 if err != nil {
  return err
 }
 log.Printf("Available buffer: %d\n", bytesAvailable)
 // 写内存buffer到硬盘
 err = bufferedWriter.Flush()
 if err != nil{
  return err
 }

 file.Close()
 return nil
}

```

### 读文件
有两种方式我们可以读取全文件：

os、io/ioutil中提供了readFile方法可以快速读取全文
io/ioutil中提供了ReadAll方法在打开文件句柄后可以读取全文；
```go
import (
 "io/ioutil"
 "log"
 "os"
)

func readAll(filename string) error {
 data, err := os.ReadFile(filename)
 if err != nil {
  return err
 }
 log.Printf("read %s content is %s", filename, data)
 return nil
}

func ReadAll2(filename string) error {
 file, err := os.Open("asong.txt")
 if err != nil {
  return err
 }

 content, err := ioutil.ReadAll(file)
 log.Printf("read %s content is %s\n", filename, content)

 file.Close()
 return nil
}
```
#### 逐行读取

os库中提供了Read方法是按照字节长度读取，如果我们想要按行读取文件需要配合bufio一起使用，bufio中提供了三种方法ReadLine、ReadBytes("\n")、ReadString("\n")可以按行读取数据，下面我使用ReadBytes("\n")来写个例子
```go
func readLine(filename string) error {
 file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
 if err != nil {
  return err
 }
 bufferedReader := bufio.NewReader(file)
 for {
  // ReadLine is a low-level line-reading primitive. Most callers should use
  // ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
  lineBytes, err := bufferedReader.ReadBytes('\n')
  bufferedReader.ReadLine()
  line := strings.TrimSpace(string(lineBytes))
  if err != nil && err != io.EOF {
   return err
  }
  if err == io.EOF {
   break
  }
  log.Printf("readline %s every line data is %s\n", filename, line)
 }
 file.Close()
 return nil
}
```

#### 按块读取文件
os库的Read方法
os库配合bufio.NewReader调用Read方法
os库配合io库的ReadFull、ReadAtLeast方法
```go

// use bufio.NewReader
func readByte(filename string) error {
 file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
 if err != nil {
  return err
 }
 // 创建 Reader
 r := bufio.NewReader(file)

 // 每次读取 2 个字节
 buf := make([]byte, 2)
 for {
  n, err := r.Read(buf)
  if err != nil && err != io.EOF {
   return err
  }

  if n == 0 {
   break
  }
  log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
 }
 file.Close()
 return nil
}

// use os
func readByte2(filename string) error{
 file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
 if err != nil {
  return err
 }

 // 每次读取 2 个字节
 buf := make([]byte, 2)
 for {
  n, err := file.Read(buf)
  if err != nil && err != io.EOF {
   return err
  }

  if n == 0 {
   break
  }
  log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
 }
 file.Close()
 return nil
}


// use os and io.ReadAtLeast
func readByte3(filename string) error{
 file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
 if err != nil {
  return err
 }

 // 每次读取 2 个字节
 buf := make([]byte, 2)
 for {
  n, err := io.ReadAtLeast(file, buf, 0)
  if err != nil && err != io.EOF {
   return err
  }

  if n == 0 {
   break
  }
  log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
 }
 file.Close()
 return nil
}
```
#### 分隔符读取

bufio包中提供了Scanner扫描器模块，它的主要作用是把数据流分割成一个个标记并除去它们之间的空格，他支持我们定制Split函数做为分隔函数，分隔符可以不是一个简单的字节或者字符，我们可以自定义分隔函数，在分隔函数实现分隔规则以及指针移动多少，返回什么数据，如果没有定制Split函数，那么就会使用默认ScanLines作为分隔函数，也就是使用换行作为分隔符，bufio中还提供了默认方法ScanRunes、ScanWrods，下面我们用SacnWrods方法写个例子，获取用空格分隔的文本

```go
func readScanner(filename string) error {
 file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
 if err != nil {
  return err
 }

 scanner := bufio.NewScanner(file)
 // 可以定制Split函数做分隔函数
 // ScanWords 是scanner自带的分隔函数用来找空格分隔的文本字
 scanner.Split(bufio.ScanWords)
 for {
  success := scanner.Scan()
  if success == false {
   // 出现错误或者EOF是返回Error
   err = scanner.Err()
   if err == nil {
    log.Println("Scan completed and reached EOF")
    break
   } else {
    return err
   }
  }
  // 得到数据，Bytes() 或者 Text()
  log.Printf("readScanner get data is %s", scanner.Text())
 }
 file.Close()
 return nil
}
```

### 打包/解包
Go语言的archive包中提供了tar、zip两种打包/解包方法，这里以zip的打包/解包为例子：

zip解包示例：
```go
import (
 "archive/zip"
 "fmt"
 "io"
 "log"
 "os"
)

func main()  {
 // Open a zip archive for reading.
 r, err := zip.OpenReader("asong.zip")
 if err != nil {
  log.Fatal(err)
 }
 defer r.Close()
 // Iterate through the files in the archive,
 // printing some of their contents.
 for _, f := range r.File {
  fmt.Printf("Contents of %s:\n", f.Name)
  rc, err := f.Open()
  if err != nil {
   log.Fatal(err)
  }
  _, err = io.CopyN(os.Stdout, rc, 68)
  if err != nil {
   log.Fatal(err)
  }
  rc.Close()
 }
}
```
zip打包示例：
```go
func writerZip()  {
 // Create archive
 zipPath := "out.zip"
 zipFile, err := os.Create(zipPath)
 if err != nil {
  log.Fatal(err)
 }

 // Create a new zip archive.
 w := zip.NewWriter(zipFile)
 // Add some files to the archive.
 var files = []struct {
  Name, Body string
 }{
  {"asong.txt", "This archive contains some text files."},
  {"todo.txt", "Get animal handling licence.\nWrite more examples."},
 }
 for _, file := range files {
  f, err := w.Create(file.Name)
  if err != nil {
   log.Fatal(err)
  }
  _, err = f.Write([]byte(file.Body))
  if err != nil {
   log.Fatal(err)
  }
 }
 // Make sure to check the error on Close.
 err = w.Close()
 if err != nil {
  log.Fatal(err)
 }
}
```