package fake

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
)

// 用于存储已生成的 PN，确保不重复
//var PNMap = make(map[string]struct{})
//var mutex sync.Mutex

func generatePartNumber() string {
	rand.Seed(time.Now().UnixNano())
	// 定义允许的字符集，包括大写字母、数字、- 和 /
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789-/"
	// 定义长度范围
	length := rand.Intn(12) + 4 // 生成长度为4到15的随机数
	partNumber := make([]byte, length)
	for i := 0; i < length; i++ {
		// 从字符集中随机选择一个字符
		partNumber[i] = characters[rand.Intn(len(characters))]
	}
	// 使用正则表达式验证生成的航材件号
	re := regexp.MustCompile(`^[A-Z1-9]([A-Z1-9]*-\d+)?[A-Z1-9]([A-Z1-9/]*[A-Z1-9])?$`)
	pn := string(partNumber)
	// 判断键是否存在
	//mutex.Lock()
	//if _, exists := PNMap[pn]; exists {
	//	// 如果不符合要求，重新生成
	//	return generatePartNumber()
	//}
	//mutex.Unlock()
	if !re.MatchString(pn) {
		// 如果不符合要求，重新生成
		return generatePartNumber()
	}
	//mutex.Lock()
	//PNMap[pn] = struct{}{}
	//mutex.Unlock()
	return pn
}

func generateStorageNum() string {
	// 定义大写字母集合
	upperLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机仓位号
	position := string(upperLetters[rand.Intn(len(upperLetters))])
	position += fmt.Sprintf("-%d", rand.Intn(50)+1) // 生成1到50之间的随机数字

	return position
}

func generateBN() string {
	// 获取当前日期
	currentDate := time.Now().Format("20060102") // 格式为年月日，例如：20220113

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机数（这里示例生成4位随机数）
	randomNumber := rand.Intn(10000)

	// 格式化批次号
	batchNumber := fmt.Sprintf("%s-%04d", currentDate, randomNumber)

	return batchNumber
}

func FakeMaterial() {
	// 指定分批次生成的数量
	batchSize := 100 // 每批多少个
	countSize := 10  // 一共多少批

	// 分批次生成并保存数据
	for i := 0; i < countSize; i++ {
		// 生成 batchSize 条模拟数据
		var materials []models.Material
		BN := generateBN()
		for j := 0; j < batchSize; j++ {
			material := models.Material{}

			// 定义日期范围
			startDate := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
			endDate := time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC)
			diff := endDate.Sub(startDate)
			randomDuration := time.Duration(rand.Int63n(int64(diff)))
			randomDate := startDate.Add(randomDuration)

			// 定义件名
			materialNames := []string{
				"PACKING",
				"PACKING",
				"PACKING",
				"PACKING",
				"PACKING",
				"PACKING",
				"PACKING",
				"PACKING",
				"PACKING",
				"BACKUP",
				"RING",
				"O-RING",
				"SPRING",
				"NUT",
				"NUT",
				"NUT",
				"NUT",
				"NUT",
				"GEAR ASSY",
				"BEARING",
				"WASHER",
				"SEAL ASSEMBLY",
				"SEAL",
				"GUIDE",
				"ROLLER",
				"PIN",
				"MARKER",
				"RACK",
			}
			randomMaterial := materialNames[rand.Intn(len(materialNames))]
			// 最少数量
			materialMinCounts := []uint{
				10,
				20,
				30,
				40,
				50,
				50,
				50,
				50,
				50,
				50,
			}
			materialMinCount := materialMinCounts[rand.Intn(len(materialMinCounts))]

			// 生成仓位号
			material.PN = generatePartNumber()         // A/CC-12
			material.StorageNum = generateStorageNum() // A-50
			material.Count = uint(rand.Intn(100) + 1)  // 1-100随机数
			material.ExpiryAt = randomDate
			material.Name = randomMaterial
			material.BN = BN
			material.Unit = "个"
			material.Manufacturer = "霍尼韦尔"
			material.Price = uint(rand.Intn(100000) + 1) // 1-100000随机数
			material.MinCount = materialMinCount

			materials = append(materials, material)

		}

		// 创建记录
		res := global.DB.Create(&materials)
		if res.Error == nil {
			fmt.Printf("第 %d 次生成完成, 已生成 %d 条模拟数据\n", i+1, res.RowsAffected)
		}
		if res.Error != nil {
			fmt.Println(res.Error.Error())
		}
	}

	fmt.Println("生成模拟数据完成！")
}
