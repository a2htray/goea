进化算法 Golang 实现
==========================

### 算法

1. [三系杂交水稻优化算法(HRO)](./doc/HRO.md) [1]
2. [人工蜂群算法](./doc/ABC.md) [2]
3. [简单差分算法 DE/rand/1/bin](./doc/DE-rand-1-bin.md) [3]
4. [差分算法 DE/best/2/bin](./doc/DE-best-2-bin.md) [4]

### 用法

`DE`

```go
import "github.com/a2htray/goea"

// 简单差分算法 DE/rand/1/bin
func main() {
    iterNum := 10
    de := goea.NewDE(4, 2, goea.Boundary{
        Lower: []float64{-10, -10, -10},
        Upper: []float64{10, 10, 10},
    }, iterNum, goea.ObjectSphere, goea.DefaultDEConfig())
    
    de.Run()

    fmt.Println(de.HistoryBestFNC)
}
```

```go
import "github.com/a2htray/goea"

// 差分算法 DE/best/2/bin
func main() {
    config := goea.DefaultDEConfig()
    config.Mode = goea.DEModeBest2bin
    
    iterNum := 20
    de := goea.NewDE(10, 3, goea.Boundary{
        Lower: []float64{-10, -10, -10},
        Upper: []float64{10, 10, 10},
    }, iterNum, goea.ObjectSphere, goea.config)
    
    de.Run()
    
    fmt.Println(de.HistoryBestFNC)
}
```

**目标函数**

[目标函数说明](./doc/FUNCTIONS.md)

### 参考文献

参考文献格式均使用 APA 格式

#### 算法

1. Ye, Z., Ma, L., & Chen, H. (2016, August). A hybrid rice optimization algorithm. In 2016 11th International Conference on Computer Science & Education (ICCSE) (pp. 169-174). IEEE.
2. Karaboga, D., & Basturk, B. (2007). A powerful and efficient algorithm for numerical function optimization: artificial bee colony (ABC) algorithm. Journal of global optimization, 39(3), 459-471.
3. Storn, R., & Price, K. (1997). Differential evolution–a simple and efficient heuristic for global optimization over continuous spaces. Journal of global optimization, 11(4), 341-359.
4. Price, K. V. (1996, June). Differential evolution: a fast and simple numerical optimizer. In Proceedings of North American Fuzzy Information Processing (pp. 524-527). IEEE.

#### 目标函数

1. Tang, K., Yao, X., Suganthan, P. N., MacNish, C., Chen, Y. P., Chen, C. M., & Yang, Z. (2010). Benchmark Functions for the CEC 2010 Special Session and Competition on Large Scale Global Optimization. University of Science and Technology of China (USTC), School of Computer Science and Technology, Nature Inspired Computation and Applications Laboratory (NICAL), Hefei, Anhui. China. Tech. Rep, Tech. Rep.