package main

import "fmt"

// 类的属性大写，表明当前方法可以被外部包访问
// Hero大写，说明该包的方法可以被其他包引用
type Hero struct {
	Name  string
	Age   int
	level int // 小写说明它是私有的
}

// 方法首字母大写表示对外公开（可导出）
func (this *Hero) Show() {
	fmt.Println("Name:", this.Name)
	fmt.Println("Age:", this.Age)
	fmt.Println("Level:", this.level)
}

// GetName 方法返回 string，需要加上返回值类型
func (this *Hero) GetName() string {
	return this.Name
}

// SetName 方法设置名字，没问题
func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "ZhangSan", Age: 10, level: 1}

	hero.Show()

	name := hero.GetName()
	fmt.Println("GetName:", name)

	hero.SetName("LiSi")
	hero.Show()
}
