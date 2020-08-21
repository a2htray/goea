## 人工蜂群算法

Karaboga, D., & Basturk, B. (2007). A powerful and efficient algorithm for numerical function optimization: artificial bee colony (ABC) algorithm. Journal of global optimization, 39(3), 459-471.

### 介绍

现阶段已经开发出多个启发式随机算法，用于解决组合问题和数值优化问题 [1]。
ABC 算法属于群体智能算法，个体间相互交流并自适应于整个种群。
ABC 算法模拟蜂群寻找食物的动物本能，其目标是寻找食物量最多的食物来源。

在 ABC 算法中，蜂群中的有三种蜂，分别为：雇佣蜂(employed bee)、观察蜂(onlooker bee)和侦查蜂(scout bee)。
起初，雇佣蜂与观察蜂数量相同，各为蜂群数量的一半，当食物来源耗尽，观察蜂会变成侦查蜂，以寻找新的食物来源。
对于每一个食物来源，有且仅有一个雇佣蜂，即雇佣蜂的数量与蜂房周围的食物来源的数量相同。

* 雇佣蜂：利用先前的蜜源信息寻找新的蜜源并与观察蜂分享蜜源信息
* 观察蜂：在蜂房中等待并依据雇佣蜂分享的信息寻找新的蜜源
* 侦查蜂：寻找一个新的有价值的蜜源，它们在蜂房附近随机的寻找蜜源 [2]

雇佣蜂和观察蜂负责模型的开发性(exploitation)，侦查蜂负责模型的探索性(exploration)。

一个食物来源代表一个可行解，而雇佣蜂的数量等于食物来源的数量，又等于观察蜂的数量，则整个蜂群数量为可行解数量的两倍。

### 过程

每一次的迭代过程包含 3 个步骤：

1. 派遣雇佣蜂到指定食物来源，并计算食物量
2. 观察蜂接收到派遣雇的消息，根据一定策略选择食物来源
3. 确定侦查蜂，并寻找可能的食物来源

在初始阶段，蜂群随机选择食物来源，并确定食物数量，侦察蜂接收到食物来源信息
第二个阶段，雇佣蜂根据之前的食物量信息判断是否选择新的食物来源
第三个阶段，观察蜂根据雇佣蜂所给信息依概率选择合适的食物来源，若为非较优食物来源，则由侦查蜂寻找新的食物来源

包含一个预设的限制次数，当一个食物来源未改变的次数大于等于该值时，侦查蜂负责开辟新的食物来源

### 发现

在编码的过程中发现，ABC 算法对解空间各维区间敏感。当各维区间以 0 对称中心时，各维越界的概率最低。
原因在于其内部算法中的随机数取值区间 [-1, 0] 以 0 为对称中心，

### 参考

1. Pham, D., & Karaboga, D. (2012). Intelligent optimisation techniques: genetic algorithms, tabu search, simulated annealing and neural networks. Springer Science & Business Media.
2. [人工蜂群算法（Artifical Bee Colony）](https://blog.csdn.net/u013927464/article/details/82722471)