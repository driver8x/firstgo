// go test -bench=. github.com/driver8x/firstgo/searchtree
package searchtree_test

import (
	"bufio"
	"fmt"
	"github.com/driver8x/firstgo/searchtree"
	"os"
	"testing"
)

var tree1 searchtree.Tree

func TestMain(m *testing.M) {
	infile, err := os.Open("../dict.txt")
	if err != nil {
		panic("Failed to open file")
	}

	i := 0
	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {

		line := scanner.Text()
		if e := scanner.Err(); e != nil {
			panic("Scanner error")
		}
		tree1.Add(line, []int{8})
		i++
	}
	infile.Close()
	fmt.Println("Tree built\n")
	os.Exit(m.Run())
}

func BenchmarkAdd0(b *testing.B) {
	tree1.Add("sdkaflhgllkhwei", []int{8, 888, 88})
	tree1.Add("aghfashjjk", []int{8, 888, 88})
	tree1.Add("fouhdxzbxzxbs", []int{8, 888, 88})
	tree1.Add("weatygacvzxcv", []int{8, 888, 88})
	tree1.Add("zbncxfgaertf", []int{8, 888, 88})
	tree1.Add("jytdfgzxbawe", []int{8, 888, 88})
	tree1.Add("shjtdjdnszerthnbszbaewragasdfgargrafcvavaertqah", []int{8, 888, 88})
	tree1.Add("pqowernflakvn;alsiehoufhlaksdfnalnvausjehdfiouqbcvkjaBnv", []int{8, 888, 88})
	tree1.Add("Nalgkahnlnlkjlakvglkagh;lkLKBVH;ALKHEERTGOINlgkaalk", []int{8, 888, 88})
	tree1.Add("88888888", []int{8, 888, 88})
	tree1.Add("dalkvnpoaweirvnlzvoabfgoiqnvlknas192gh9(&*T)*&T9faudwbfkqjwh9&T*&^!ojkblkjbakvjbdiaua098vhjn239uvnIV)87rhqpq9n2bp4nbua0v89&BVC9drfawpnhe1!fjhh9VC", []int{8, 888, 88})
}

func BenchmarkSearch0(b *testing.B) {
	tree1.Find("alba")
	tree1.Find("ceru")
	tree1.Find("adglkahl")
	tree1.Find("zi")
	tree1.Find("bromomethylbenz")
	tree1.Find("asdkljgha09bvha09dusifhfg")
	tree1.Find("Nalgkahnlnlkjlakvglkagh;lkLKBVH;ALKHEERTGOINlgkaalk")
	tree1.Find("qwoertuipqhv09ahn")
	tree1.Find("c")
	tree1.Find("qiwoethlknn,.xzcvn")
	tree1.Find("230-yhnbvglka")
}
