package goea

import (
	"fmt"
	"github.com/a2htray/goea/base"
	"os"
	"testing"
)

func TestNewHRO(t *testing.T) {
	m := 60
	n := 10
	iterNum := 300
	minimum := false
	limit := base.Limit{
		Upper: base.Boundary{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		Lower: base.Boundary{-10, -10, -10, -10, -10, -10, -10, -10, -10, -10},
	}
	logFile, _ := os.Create("D:\\workspace\\github.com\\a2htray\\goea\\log\\hro.log")
	defer logFile.Close()

	model := NewHRO(m, n, iterNum, minimum, limit, ObjectSphere)
	model.OpenLog(logFile)

	model.Run()

	for i := 0; i < iterNum; i++ {
		best := model.HistoryBestIndividuals[i]
		fmt.Printf("第%d代\n最优个体为:%v,其适应值为%f\n", i+1, best.Vector.Float64(), best.FitnessValue)
	}
}
