package pkg

import (
	"bufio"
	"io"
	"os"
	"strings"

	"log"
)

// DivideLine は建物データを３つに分類する
func DivideLine(fn string) error {
	rfl, er := os.Open(fn)
	if er != nil {
		log.Fatal(er)
	}
	defer rfl.Close()

	if f, err := os.Stat("C:/data"); os.IsNotExist(err) || !f.IsDir() {
		os.MkdirAll("C:/data", os.ModePerm)
	} else {
		log.Println("存在します")
	}

	wfl1, er := os.Create("C:/data/hutsu_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl1.Close()

	wfl2, er := os.Create("C:/data/kenro_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl2.Close()

	wfl3, er := os.Create("C:/data/other_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl3.Close()

	r := bufio.NewReader(rfl)
	w1 := bufio.NewWriter(wfl1)
	w2 := bufio.NewWriter(wfl2)
	w3 := bufio.NewWriter(wfl3)

	for {
		row, er := r.ReadString('\n')
		if er != nil && er != io.EOF {
			log.Fatal(er)
		}
		if er == io.EOF && len(row) == 0 {
			break
		}
		if strings.Contains(row, "普通建物") == true {
			log.Println("普通建物")
			_, er = w1.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w1.Flush()
		}
		if strings.Contains(row, "堅ろう建物") == true {
			log.Println("堅ろう建物")
			_, er = w2.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w2.Flush()
		}
		if strings.Contains(row, "無壁舎") == true {
			log.Println("無壁舎")
			_, er = w3.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w3.Flush()
		}
	}
	if er != nil {
		log.Fatal(er)
	}
	return nil
}
