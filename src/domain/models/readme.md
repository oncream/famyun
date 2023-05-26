# domain -> models
## models 职责
### 1、定义用户实体
有唯一的表示，包含各种数据验证、操作前置函数、构造函数实例化等
### 2、用户的值对象
用来描述一个事物的特征，没有唯一标识的对象 
 
#### 有两个重要原则
* 实体只能通过ID来判断两者是否相同
* 而值对象。只需根据“值”就能判断两者是否相同

#### 不可变
修改值对象，必须传入新对象。


 
#### 值对象就是struct下的struct，例子如下：
 ```go
type A struct {   
	Id int
	Name string
	Extra *B
 }
 type B struct {
 	Address string
 	Email string
 }
 ```