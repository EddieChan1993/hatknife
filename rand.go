package hatknife

import (
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

const RandTotal = 1000             //默认权重总和
type RandPoolTyp = map[int32]int32 //随机池类型k-poolId（需要随机的id） v-weight（权重）
const stdStrDigit = "0123456789"

// RandInt32 返回一个(0,total]的随机数
func RandInt32(total int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(total) + 1
}

func RandNumSlice(pool []int32, n int32) []int32 {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })
	return pool[:n]
}

// RandSlice 随机切片值
func RandSlice(pool []int32) (sliceVal, sliceIndex int32) {
	total := len(pool)
	rand.Seed(time.Now().UnixNano())
	index := rand.Int31n(int32(total)) //[0,total)
	return pool[index], index
}

// RandBetween 两个数之间随机[start,end]
func RandBetween(start, end int32) int32 {
	ca := end - start
	rand.Seed(time.Now().UnixNano())
	randCa := rand.Int31n(ca + 1) //[0,ca]
	return randCa + start
}

// RandWeight 计算权重
func RandWeight(pool RandPoolTyp) (weightList, poolIds []int32, weightTotal int32) {
	poolIds = make([]int32, 0, len(pool))
	weightList = make([]int32, 0, len(pool))
	weightTotal = int32(0)
	for id, w := range pool {
		poolIds = append(poolIds, id)
		weightTotal += w
		weightList = append(weightList, weightTotal)
	}
	return
}

// RandOneOnceWeight 指定权重随机
func RandOneOnceWeight(weightList, poolIds []int32, weightTotal int32) (poolId int32) {
	randInt := RandInt32(weightTotal)
	index := 0
	for i, w := range weightList {
		index = i
		if w >= randInt {
			break
		}
	}
	return poolIds[index]
}

// RandOne 随机产出一个
// pool 奖池；k-奖品id v-奖品权重
func RandOne(pool RandPoolTyp) (poolId int32) {
	poolIds := make([]int32, 0, len(pool))
	weightList := make([]int32, 0, len(pool))
	weightTotal := int32(0)
	for id, w := range pool {
		poolIds = append(poolIds, id)
		weightTotal += w
		weightList = append(weightList, weightTotal)
	}
	randInt := RandInt32(weightTotal)
	index := 0
	for i, w := range weightList {
		index = i
		if w >= randInt {
			break
		}
	}
	return poolIds[index]
}

// RandMulti 随机产出多个
// pool 奖池；k-奖品id v-奖品权重
func RandMulti(pool RandPoolTyp, randTimes int32) (poolRandIds []int32) {
	if len(pool) == 0 {
		//权重和奖池不等
		return nil
	}
	poolIds := make([]int32, 0, len(pool))
	weightList := make([]int32, 0, len(pool))
	weightTotal := int32(0)
	for id, w := range pool {
		poolIds = append(poolIds, id)
		weightTotal += w
		weightList = append(weightList, weightTotal)
	}
	poolRandIds = make([]int32, 0, randTimes)
	for i := int32(0); i < randTimes; i++ {
		randInt := RandInt32(weightTotal)
		index := 0
		for j, w := range weightList {
			index = j
			if w >= randInt {
				break
			}
		}
		poolRandIds = append(poolRandIds, poolIds[index])
	}
	return poolRandIds
}

// RandMultiNoRepeat  随机产出N个不重复奖品
// pool 奖池；k-奖品id v-奖品权重
func RandMultiNoRepeat(pool RandPoolTyp, randTimes int32) (poolRandIds []int32) {
	if int32(len(pool)) < randTimes {
		//权重和奖池不等，奖池少于需要随机的个数
		return nil
	} else if int32(len(pool)) == randTimes {
		poolIds := make([]int32, 0, len(pool))
		for id := range pool {
			poolIds = append(poolIds, id)
		}
		return poolIds
	}
	poolRandIds = make([]int32, 0, randTimes)
	for i := int32(0); i < randTimes; i++ {
		poolIds := make([]int32, 0, len(pool))
		weightList := make([]int32, 0, len(pool))
		weightTotal := int32(0)
		for id, w := range pool {
			poolIds = append(poolIds, id)
			weightTotal += w
			weightList = append(weightList, weightTotal)
		}
		randInt := RandInt32(weightTotal)
		index := 0
		for j, w := range weightList {
			index = j
			if w >= randInt {
				break
			}
		}
		delete(pool, poolIds[index])
		poolRandIds = append(poolRandIds, poolIds[index])
	}
	return poolRandIds
}

// RandNumByLen 生成纯数字字符串
func RandNumByLen(length int) string {
	return randStr(stdStrDigit, length)
}

func randStr(std string, length int) string {
	if length <= 0 {
		return ""
	}

	bytes := []byte(std)
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}

	return bytesToString(result)
}

func bytesToString(b []byte) (s string) {
	_bptr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	_sptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	_sptr.Data = _bptr.Data
	_sptr.Len = _bptr.Len
	return s
}
