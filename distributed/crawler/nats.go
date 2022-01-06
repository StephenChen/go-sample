package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/nats-io/nats.go"
	"net/url"
	"os"
	"regexp"
	"time"
)

var domain2Collector = map[string]*colly.Collector{}
var nc *nats.Conn
var maxDepth = 10
var natsURL = "nats://localhost:4222"

func factory(urlStr string) *colly.Collector {
	u, _ := url.Parse(urlStr)
	return domain2Collector[u.Host]
}

func initABCDECollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.abcdefg.com"),
		colly.MaxDepth(maxDepth),
	)

	c.OnResponse(func(resp *colly.Response) {
		// 做些爬完后的善后
		// 比如页面已爬完的确认存进 mysql
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// 基本的反爬虫策略
		link := e.Attr("href")
		time.Sleep(time.Second * 2)

		// 认为匹配该模式的是该网站的详情页
		detailRegex, _ := regexp.Compile(`/go/go\?p=\d+$`)
		// 匹配下面模式的是该网站的详情页
		listRegex, _ := regexp.Compile(`/t/\d+#\w+`)

		// 正则 match 列表页的话，就 visit
		if listRegex.Match([]byte(link)) {
			c.Visit(e.Request.AbsoluteURL(link))
		}
		// 正则 match 落地页的话，就发消息队列
		if detailRegex.Match([]byte(link)) {
			_ = nc.Publish("tasks", []byte(link))
			nc.Flush()
		}
	})

	return c
}

func initHIJKLCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.hijklmn.com"),
		colly.MaxDepth(maxDepth),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	})

	return c
}

func init() {
	domain2Collector["www.abcdefg.com"] = initABCDECollector()
	domain2Collector["www.hijklmn.com"] = initHIJKLCollector()
	_, err := nats.Connect(natsURL)
	if err != nil {
		os.Exit(1)
	}
}

func Nats() {
	// pub
	urls := []string{"https://www.abcdefg.com", "https://www.hijklmn.com"}
	for _, url := range urls {
		instance := factory(url)
		instance.Visit(url)
	}

	//sub
	startConsumer()
}

func startConsumer() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return
	}

	sub, err := nc.QueueSubscribeSync("tasks", "workers")
	if err != nil {
		return
	}

	var msg *nats.Msg
	for {
		msg, err = sub.NextMsg(time.Hour * 10000)
		if err != nil {
			break
		}

		urlStr := string(msg.Data)
		ins := factory(urlStr)
		// 因为最下游拿到的一定是对应网站的落地页
		// 所以不用进行多余的判断了，直接爬内容即可
		ins.Visit(urlStr)
		// 防止被封杀
		time.Sleep(time.Second)
	}
}

func pub() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return
	}

	// 指定 subject 为 tasks，消息内容随意
	err = nc.Publish("tasks", []byte("your task content"))

	nc.Flush()
}

func sub() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return
	}

	// queue subscribe 相当于在消费者之间进行任务分发的分支均衡
	// 前提是所有消费者都使用 workers 这个 queue
	// nats 中的 queue 概念上类似于 Kafka 中的 consumer group
	sub, err := nc.QueueSubscribeSync("tasks", "workers")
	if err != nil {
		return
	}

	var msg *nats.Msg
	for {
		msg, err = sub.NextMsg(time.Hour * 10000)
		if err != nil {
			break
		}
		// 正确地消费到了消息
		// 可用 nats.Msg 对象处理任务
		fmt.Println(msg)
	}
}
