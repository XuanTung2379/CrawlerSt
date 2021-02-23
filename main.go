package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// DataPost :
type DataPost struct {
	NameText    string
	NameLink    string
	TitleText   string
	TitleLink   string
	DateTime    string
	ContentMeta string
	HeadPost    string
	ContentData string
	Image       string
	Content     string
	CodeStock   []string
}

// Category : struct cac bien
type Category struct {
	linkPage    string
	rootEndPage string
	// DataPosts123 []*DataPost
	// slicelink []string
}

// DataPosts :
var DataPosts []*DataPost

// DataManyPosts :
var DataManyPosts []*DataPost

// MapStock :
var MapStock = make(map[string]bool)

// StockMarket :
const (
	StockMarket = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/36,981,35/index.htm"
	Bonds       = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/971/index.htm"
	Companies   = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/34,962,963,964,965,966,967,214/index.htm"
	RealEstate  = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/973/index.htm"
	Finance     = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/37/index.htm"
	Economy     = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/38/index.htm"
	Commodity   = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/970/index.htm"
	Industries  = "https://en.vietstock.vn/page/%d/TabId/56/Type/8/ChannelID/974,982,143/index.htm"
	rootLink    = "https://en.vietstock.vn"
	rootPage    = "https://en.vietstock.vn/page/"
)

// GetList : lay mot page bat ky
func (c *Category) GetList(Category string, page int) (error, error) {
	if Category == StockMarket {
		c.linkPage = strings.Replace(StockMarket, "%d", strconv.Itoa(page), 20)
	}
	if Category == Bonds {
		c.linkPage = strings.Replace(Bonds, "%d", strconv.Itoa(page), 20)
	}
	if Category == Companies {
		c.linkPage = strings.Replace(Companies, "%d", strconv.Itoa(page), 20)
	}
	if Category == RealEstate {
		c.linkPage = strings.Replace(RealEstate, "%d", strconv.Itoa(page), 20)
	}
	if Category == Finance {
		c.linkPage = strings.Replace(Finance, "%d", strconv.Itoa(page), 20)
	}
	if Category == Economy {
		c.linkPage = strings.Replace(Economy, "%d", strconv.Itoa(page), 20)
	}
	if Category == Commodity {
		c.linkPage = strings.Replace(Commodity, "%d", strconv.Itoa(page), 20)
	}
	if Category == Industries {
		c.linkPage = strings.Replace(Industries, "%d", strconv.Itoa(page), 20)
	}
	doc, error := goquery.NewDocument(c.linkPage)
	if error != nil {
		log.Fatal(error)
	}
	doc.Find(".News").Each(func(index int, element *goquery.Selection) {
		textOption := element.Find("a.Color1").Text()
		linkOption, _ := element.Find("h6.arial_8pt a.Color1").Attr("href")
		textTitle := strings.TrimSpace(element.Find("h3.News_title a.Color3").Text())
		linkTitle, _ := element.Find("a.Color3").Attr("href")
		dateTime := strings.TrimSpace(element.Find("p.arial_9pt").Text())
		Content := strings.TrimSpace(element.Find("p.NewsItem").Text())
		linkImg, _ := element.Find("img").Attr("src")
		fmt.Printf("%v \n %v \n %v \n %v \n %v \n %v \n %v \n\n\n", textOption, linkOption, textTitle, linkTitle, dateTime, Content, linkImg)

		// childResultLink := []string{linkTitle}
		// c.slicelink = append(c.slicelink, childResultLink...)

		// fmt.Printf("xuan tung %v \n\n", c.slicelink)
		allPosts := &DataPost{
			NameText:    textOption,
			NameLink:    linkOption,
			TitleText:   textTitle,
			TitleLink:   linkTitle,
			DateTime:    dateTime,
			Image:       linkImg,
			ContentData: Content,
		}
		DataPosts = append(DataPosts, allPosts)
	})

	return nil, nil
}

// GetManyList : Lay nhieu list
func (c *Category) GetManyList(Category string, startPage int, endPage int) (error, error) {
	if Category == StockMarket {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/36/index.htm"
	}
	if Category == Bonds {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/971,44,40/index.htm"
	}
	if Category == Companies {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/34,962,963,964,965,966,967,214/index.htm"
	}
	if Category == RealEstate {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/973/index.htm"
	}
	if Category == Finance {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/37/index.htm"
	}
	if Category == Economy {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/38,977,71/index.htm"
	}
	if Category == Commodity {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/970,41,117/index.htm"
	}
	if Category == Industries {
		c.rootEndPage = "/TabId/56/Type/8/ChannelID/143/index.htm"
	}
	for i := startPage; i <= endPage; i++ {
		Page := rootPage + strconv.Itoa(i) + c.rootEndPage

		doc, error := goquery.NewDocument(Page)
		if error != nil {
			log.Fatal(error)
		}
		doc.Find(".News").Each(func(index int, element *goquery.Selection) {
			textOption := element.Find("a.Color1").Text()
			linkOption, _ := element.Find("h6.arial_8pt a.Color1").Attr("href")
			textTitle := strings.TrimSpace(element.Find("h3.News_title a.Color3").Text())
			linkTitle, _ := element.Find("a.Color3").Attr("href")
			dateTime := strings.TrimSpace(element.Find("p.arial_9pt").Text())
			Content := strings.TrimSpace(element.Find("p.NewsItem").Text())
			linkImg, _ := element.Find("img").Attr("src")
			fmt.Printf("%v \n %v \n %v \n %v \n %v \n %v \n %v \n\n\n", textOption, linkOption, textTitle, linkTitle, dateTime, Content, linkImg)
			allManyPosts := &DataPost{
				NameText:    textOption,
				NameLink:    linkOption,
				TitleText:   textTitle,
				TitleLink:   linkTitle,
				DateTime:    dateTime,
				Image:       linkImg,
				ContentData: Content,
			}
			DataManyPosts = append(DataManyPosts, allManyPosts)
		})
		fmt.Printf("-------------------------------------------trang hoan thanh %v----------------------------------------", i)
	}
	return nil, nil
}

// GetManyDetail : lay data cua nhieu post
func (d *DataPost) GetManyDetail() error {
	showlink := &d.TitleLink
	// fmt.Println(*showlink)
	doc, err := goquery.NewDocument(rootLink + *showlink)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("tung 123", doc)
	doc.Find(".NewsDetail_Box").Each(func(index int, element *goquery.Selection) {
		textTittleChild := strings.TrimSpace(element.Find("h1.DtailTitle").Text())
		textStartChild := element.Find("h2.DtailText").Text()
		textBodyChild := strings.TrimSpace(element.Find("p.pBody").Text())
		textBodyChild2 := strings.TrimSpace(element.Find("p").Text())
		linkImgChild, _ := element.Find("img").Attr("src")
		linkAll, _ := doc.Find("a target").Attr("href")
		d.Content = textBodyChild
		fmt.Printf(" %v \n\n %v \n %v \n %v \n %v \n %v \n", textTittleChild, textStartChild, textBodyChild, textBodyChild2, linkImgChild, linkAll)
		file, err := ioutil.ReadFile("stockMai.txt")
		if err != nil {
			log.Fatal(err)
		}
		splitLine := strings.Split(string(file), "\r\n")
		for _, line := range splitLine {
			if equal := strings.Index(line, " "); equal >= 0 {
				if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
					var value bool = true
					MapStock[key] = value
				}
			}
		}
		textBodyChild = d.Content

		tc2 := strings.ReplaceAll(textBodyChild, ".", " ")
		tc3 := strings.ReplaceAll(tc2, ":", " ")
		tc4 := strings.ReplaceAll(tc3, "(", " ")
		tc5 := strings.ReplaceAll(tc4, ")", " ")
		tc6 := strings.ReplaceAll(tc5, ",", " ")
		tc7 := strings.ReplaceAll(tc6, `"`, " ")
		splitText := strings.Split(tc7, " ")
		for _, check := range splitText {
			if MapStock[check] == true {
				d.CodeStock = append(d.CodeStock, check)
			}
		}
	})
	fmt.Println(d.CodeStock)
	return nil
}

// GetDetail : lay data cua 1 post
func (d *DataPost) GetDetail() error {
	showlink := &d.TitleLink
	// fmt.Println(*showlink)
	doc, err := goquery.NewDocument(rootLink + *showlink)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("tung 123", doc)
	doc.Find(".NewsDetail_Box").Each(func(index int, element *goquery.Selection) {
		textTittleChild := strings.TrimSpace(element.Find("h1.DtailTitle").Text())
		textStartChild := element.Find("h2.DtailText").Text()
		textBodyChild := strings.TrimSpace(element.Find("p.pBody").Text())
		textBodyChild2 := strings.TrimSpace(element.Find("p").Text())
		linkImgChild, _ := element.Find("img").Attr("src")
		linkAll, _ := doc.Find("a target").Attr("href")
		d.Content = textBodyChild
		fmt.Printf(" %v \n\n %v \n %v \n %v \n %v \n %v \n", textTittleChild, textStartChild, textBodyChild, textBodyChild2, linkImgChild, linkAll)
		file, err := ioutil.ReadFile("stockMai.txt")
		if err != nil {
			log.Fatal(err)
		}
		splitLine := strings.Split(string(file), "\r\n")
		for _, line := range splitLine {
			if equal := strings.Index(line, " "); equal >= 0 {
				if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
					var value bool = true
					MapStock[key] = value
				}
			}
		}
		textBodyChild = d.Content

		tc2 := strings.ReplaceAll(textBodyChild, ".", " ")
		tc3 := strings.ReplaceAll(tc2, ":", " ")
		tc4 := strings.ReplaceAll(tc3, "(", " ")
		tc5 := strings.ReplaceAll(tc4, ")", " ")
		tc6 := strings.ReplaceAll(tc5, ",", " ")
		tc7 := strings.ReplaceAll(tc6, `"`, " ")
		splitText := strings.Split(tc7, " ")
		for _, check := range splitText {
			if MapStock[check] == true {
				d.CodeStock = append(d.CodeStock, check)
			}
		}
	})
	fmt.Println(d.CodeStock)
	return nil
}

// VTCrawler :
type VTCrawler interface {
	GetList(Category string, page int) (error, error)
	GetManyList(Category string, startPage int, endPage int) (error, error)
	// GetDetail() error
}

// VTDetailPost : lay data tung POST
type VTDetailPost interface {
	GetDetail() error
}

// var vtc VTCrawler = &DataPost{}
var vtc VTCrawler = &Category{}

var vtd VTDetailPost = &DataPost{}

func main() {
	// var c Category
	// //khai bao page muon lay
	var numberPage = 1
	//thay doi link theo page
	var chosePage = StockMarket
	// var chosePage = strings.Replace(StockMarket, "%d", numberPage, 20)
	vtc.GetList(chosePage, numberPage)
	//khai trang bat dau muon crawler
	var startPage = 2
	//khai bao tran ket thuc crawler
	var endPage = 4
	// chay func crawl nhieu page
	vtc.GetManyList(chosePage, startPage, endPage)
	// // //------------------------
	// lấy content và giá trị của 1 page
	for _, fact := range DataPosts {
		fact.GetDetail()

	}
	for _, fact := range DataManyPosts {
		fact.GetManyDetail()

	}
	fmt.Println("ket thuc")
}
