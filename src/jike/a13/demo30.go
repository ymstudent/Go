package main

import "fmt"

type Cat struct {
	name string
	scientificName string
	category string
}

func New(name, scientificName, category string) Cat {
	return Cat{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

//要想改变数据，就需要使用指针方法，否则只是在原数据的副本上做修改（除非接受者本身就是引用类型，如切片，字典等）
func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)", cat.scientificName, cat.category, cat.name)
}

func main()  {
	cat := New("little pig", "American ShortHair", "cat")
	cat.SetName("monster") //go在调用方法时，自动做了处理，如果值类型接受者找不到对应的方法，就会去指针类型的方法中寻找
	fmt.Printf("the cat : %s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("the cat : %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	//一个自定义数据的方法集合中，仅包含它所有的值方法，但这个类型的指针类型的方法集合中却包含了所有的值方法和指针方法
	_, ok := interface{}(cat).(Pet)
	fmt.Printf("cat implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*cat implements interface Pet: %v\n", ok)
}