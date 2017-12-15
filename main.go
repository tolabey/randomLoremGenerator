package main

import (
	"strings"
	"fmt"
	"math/rand"

)

func main() {

	fmt.Println("word number")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("paragraf count")
	var j int
	_, err2 := fmt.Scanf("%d", &j)
	if err2 != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(randomLorem(start(), i, j))

}


func start() []string {
	trimList := []string{".", ",", "/", "!"}

	lorem := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent non nulla et nulla malesuada ullamcorper at eget libero. Suspendisse potenti. Sed vel dui interdum, aliquet nunc et, tincidunt ante. Sed ultricies in enim quis auctor. Donec aliquet velit vel mauris elementum varius. Proin lacinia sodales sem, quis molestie erat facilisis eu. Nullam vitae condimentum lectus. Maecenas id laoreet dolor. Cras sit amet bibendum tortor, porttitor consectetur nisl. Nullam mollis sed sem in porta. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Proin commodo volutpat volutpat. Sed scelerisque porta tortor, quis interdum odio."
	lorem2 := "Etiam viverra felis erat, sit amet congue ex aliquam ut. Sed pulvinar, massa sed lacinia laoreet, enim ex scelerisque tortor, nec dapibus ligula tellus eget lectus. Phasellus at varius lorem, et facilisis massa. Sed vitae purus tristique, sollicitudin felis tristique, rutrum diam. Sed ac leo vitae magna commodo dignissim et sed ante. Vivamus non eros libero. Aliquam accumsan tristique arcu ut feugiat. Sed lorem mi, mattis in elit et, tincidunt pellentesque quam. Suspendisse ipsum metus, venenatis nec venenatis sit amet, sollicitudin vel turpis. Etiam pellentesque venenatis eros. Aenean quis nibh et nisl feugiat dapibus. Ut sodales orci arcu, sit amet convallis dolor rhoncus non."
	lorem3 := "Fusce in mollis arcu, quis fringilla quam. In mauris nibh, dapibus ac urna eget, dictum blandit sem. Aliquam egestas posuere odio, vitae hendrerit est ultrices nec. Nullam at feugiat libero. Duis sem magna, accumsan in tellus ac, suscipit cursus urna. Curabitur bibendum dignissim lectus, in aliquam nunc fringilla id. Nam in erat et enim eleifend cursus. Phasellus pharetra nisi non magna porttitor scelerisque."
	lorem4 := "Fusce tempor varius sagittis. In ultricies, purus et posuere cursus, massa lectus placerat nunc, id finibus ex justo ut odio. Mauris et gravida turpis, quis hendrerit arcu. Quisque quis scelerisque augue. Sed cursus, massa a condimentum sagittis, dolor orci laoreet ipsum, et sagittis sem massa in nibh. Proin et ligula a augue eleifend porta eu sit amet tellus. Donec tincidunt elementum scelerisque. Aenean faucibus dui at purus tempor, a dapibus dui consectetur. Duis luctus felis tellus, id posuere nibh tincidunt vitae. Vestibulum eu augue tristique, ullamcorper libero vitae, venenatis libero."
	lorem5 := "Mauris maximus leo et interdum commodo. Duis sit amet vulputate tortor. Nulla urna mi, dignissim ut odio et, mollis pulvinar elit. Maecenas at leo et est hendrerit vulputate. Quisque aliquam consequat enim, placerat varius massa volutpat ut. Nullam vitae nunc condimentum, sollicitudin justo sed, posuere lacus. Ut condimentum odio interdum ex placerat molestie. Vestibulum bibendum arcu ut quam finibus pulvinar. Nam vel quam felis."
	lorem6 := "Morbi euismod odio volutpat lobortis tempor. Integer elit diam, facilisis sit amet tincidunt sit amet, suscipit quis ex. Duis sit amet suscipit ex, at vehicula sapien. Etiam a velit ac lorem laoreet porttitor. Curabitur egestas porta elementum. Sed euismod sapien at dapibus pharetra. Pellentesque porttitor quis tortor et auctor. Curabitur neque leo, pretium eu neque ac, volutpat pulvinar odio. Donec ac finibus leo. Ut ultricies lacus sit amet ipsum vulputate, id imperdiet ipsum aliquet"

	loremList := []string{lorem, lorem2, lorem3, lorem4, lorem5, lorem6}
	mapping := map[string]string{}

	for _, value := range trimList {
		for i, str := range loremList {
			loremList[i] = strings.Replace(str, value, "", -1)

		}
	}

	for _, value := range loremList {
		ipsum := strings.Split(value, " ")
		for _, vl := range ipsum {
			mapping[vl] = vl
		}
	}

	var loremUniqueWordList []string

	for _, value := range mapping{
		loremUniqueWordList = append(loremUniqueWordList, value)
	}

	return loremUniqueWordList
}

func randomLorem (loremList []string, wordCount int, paragCount int) string{

	var returnVal string
    var k int
    for i := 0; i < paragCount; i++ {
		for j := 0; j < wordCount; j++ {
			k  = rand.Intn(len(loremList))
			if j == 0 {
				returnVal = returnVal + loremList[k]
				continue
			}
			returnVal = returnVal + " " + loremList[k]
		}

		returnVal = returnVal + "."
		returnVal = returnVal + "\n"
	}

	return returnVal

}