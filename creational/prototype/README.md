## 原型模式

用于克隆负载类的实例，而不是创建新的实例。这种模式有时被称为虚拟构造函数或克隆工厂。这种模式是创建型模式，它提供了一种创建对象的最佳方式。

![img.png](img.png)

### 优缺点分析
1. 优点
    - 1.1 当创建一个对象的过程很昂贵时，可以通过复制一个现有对象来避免这种开销。
    - 1.2 当一个对象的创建过程很复杂时，可以通过复制一个现有对象来避免这种复杂性。