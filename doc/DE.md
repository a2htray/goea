差分算法
===============

Price, K. V. (1996, June). Differential evolution: a fast and simple numerical optimizer. In Proceedings of North American Fuzzy Information Processing (pp. 524-527). IEEE.

### 介绍

在 DE 算法中，函数参数编码成浮点型变量，并对变量应用简单的数据变换，以求得函数的较优解。
其中，算法的交叉操作采用单向的方式，并以当前最优解作为参照。
同时，选择操作可加快种群的收敛，采取“择优录取”的策略。
值得注意的是，文章中引入一种新颖的采样技术，以增加种群间的扰动性。
