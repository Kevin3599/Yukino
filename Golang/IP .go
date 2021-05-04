fmt.Printf("Your name is %s", input)
var name string
fmt.Println("请输入姓名:")
// 当程序执行到 fmt.Scanl(&name), 程序会停止这里, 等待用户输入, 并回车
fmt.Scanln(&name)