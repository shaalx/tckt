package main

import (
	"fmt"
	"net/url"
	"os"
)

type Ad struct {
	Url, From, Uid, Adid, Adgroupid, Ir string
}

func NewAd(_url, from, uid, adid, adgroupid, ir string) *Ad {
	return &Ad{
		Url:       _url,
		From:      from,
		Uid:       uid,
		Adid:      adid,
		Adgroupid: adgroupid,
		Ir:        ir,
	}
}

func init() {
	Ads = make([]*Ad, 0, 15)
	Ads = append(Ads, NewAd("http://www.ard9.com/shzs/2636003.html", "http://mdblog.daoapp.io/", "2073499", "33650", "10762", "0"))
	Ads = append(Ads, NewAd("http://kj.13322.com/?youjian$026", "http://mdblog.daoapp.io/", "2073499", "33500", "10667", "0"))
	Ads = append(Ads, NewAd("http://www.cnzjw.net/easypanel-main.html", "http://mdblog.daoapp.io/", "2073499", "33139", "10428", "0"))
	Ads = append(Ads, NewAd("http://www.nanshiw.com/rejian/1772.html", "http://mdblog.daoapp.io/", "2073499", "33641", "10752", "0"))
	Ads = append(Ads, NewAd("http://www.tootoo.cn/sale/wap.html?t=android", "http://mdblog.daoapp.io/", "2073499", "30951", "9405", "0"))
	Ads = append(Ads, NewAd("http://www.tootoo.cn/sale/wap.html?t=android", "http://mdblog.daoapp.io/", "2073499", "30952", "9405", "0"))
	Ads = append(Ads, NewAd("http://www.tootoo.cn/sale/wap.html?t=android", "http://mdblog.daoapp.io/", "2073499", "30953", "9405", "0"))
	Ads = append(Ads, NewAd("http://www.tootoo.cn/?utm_source=youjian&utm_medium=cpc&utm_campaign=yj&buyersource=roiyj&promotion_from=roiyj", "http://mdblog.daoapp.io/", "2073499", "33651", "10762", "0"))
	Ads = append(Ads, NewAd("http://www.tootoo.cn/?utm_source=youjian&utm_medium=cpc&utm_campaign=yj&buyersource=roiyj&promotion_from=roiyj", "http://mdblog.daoapp.io/", "2073499", "33652", "10762", "0"))
	Ads = append(Ads, NewAd("http://www.tootoo.cn/?utm_source=youjian&utm_medium=cpc&utm_campaign=yj&buyersource=roiyj&promotion_from=roiyj", "http://mdblog.daoapp.io/", "2073499", "33653", "10762", "0"))
	Ads = append(Ads, NewAd("http://www.cz868.com/forum.php", "http://mdblog.daoapp.io/", "2073499", "33628", "10739", "0"))
	Ads = append(Ads, NewAd("https://passfs.tmall.com/?mm_gxbid=1_6171_6aca8cbd3acb85903897daf9aebab7d0", "http://mdblog.daoapp.io/", "2073499", "32625", "10411", "0"))
	Ads = append(Ads, NewAd("http://www.117play.com/index.html", "http://mdblog.daoapp.io/", "2073499", "33652", "10763", "0"))
	Ads = append(Ads, NewAd("http://www.lovexiangsui.com", "http://mdblog.daoapp.io/", "2073499", "33592", "10720", "0"))
	Ads = append(Ads, NewAd("https://item.taobao.com/item.htm?id=524742259541", "http://mdblog.daoapp.io/", "2073499", "33654", "10764", "0"))
	Ads = append(Ads, NewAd("http://www.chenlingjian.com/55.html", "http://mdblog.daoapp.io/", "2073499", "33503", "10671", "0"))

	Sites = []string{
		"http://mdblog.daoapp.io/",
		"http://smdblog.daoapp.io/",
		"http://lmdblog.daoapp.io/",
		"http://wmdblog.daoapp.io/",
		"http://zmdblog.daoapp.io/",
		"http://topcode.daoapp.io/",
	}
}

var (
	Ads   []*Ad
	Sites []string
)

func main() {
	file, _ := os.OpenFile("target.md", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
	file.WriteString(`<style>
a:hover
{
background-color:yellow;
}
a:link    {color:blue;}
a:visited {color:blue;}
a:hover   {color:red;}
a:active  {color:yellow;}
</style>
<script src="http://lib.sinaapp.com/js/jquery/1.9.1/jquery-1.9.1.min.js"></script>
<script type="text/javascript">
$(document).ready(function () {
$("a").each(function(){
alert($(this).text())
});
});
</script>
`)
	for _, it := range Ads {
		for _, siteURI := range Sites {
			it.From = siteURI
			uri := NewQuery(it)
			file.WriteString(fmt.Sprintf("[%s](%s)\t\t", siteURI, uri.String()))
		}
	}
}

func NewQuery(ad *Ad) *url.URL {
	val := url.Values{
		"url":       []string{ad.Url},
		"from":      []string{ad.From},
		"uid":       []string{ad.Uid},
		"adid":      []string{ad.Adid},
		"adgroupid": []string{ad.Adgroupid},
		"ir":        []string{ad.Ir},
	}
	URL, _ := url.Parse("http://u.ujian.cc/")
	encd := val.Encode()
	URL.RawQuery = encd
	return URL
}
