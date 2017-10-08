# cli.crawler

[![Build Status](https://travis-ci.org/cameronnewman/cli.crawler.svg?branch=master)](https://travis-ci.org/cameronnewman/cli.crawler) [![GoDoc](https://godoc.org/github.com/cameronnewman/cli.crawler?status.svg)](http://godoc.org/github.com/cameronnewman/cli.crawler) [![Report card](https://goreportcard.com/badge/github.com/cameronnewman/cli.crawler)](https://goreportcard.com/report/github.com/cameronnewman/cli.crawler)

a simple cli web crawler


Usage

```
lappy:~ root$ ./cli.crawler --url https://facebook.com
{
	"links": [
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://www.facebook.com/",
			"original_url_hash": "108ec1825ce89b9f50be3f8e197ee4252f011be4",
			"redirects": false,
			"destination_url": "https://www.facebook.com/",
			"destination_url_hash": "108ec1825ce89b9f50be3f8e197ee4252f011be4"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://www.facebook.com/recover/initiate?lwv=110",
			"original_url_hash": "1593d5a072a7c153a7f48a75d87e1c721ec4d2bb",
			"redirects": true,
			"destination_url": "https://www.facebook.com/login/identify?ctx=recover\u0026lwv=110",
			"destination_url_hash": "f7985456126a8607db62df0c9d064e1d8a1a9685"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://zh-cn.facebook.com/",
			"original_url_hash": "5ebd2d09c4394c0b5fb859c1c66d045aa4dad544",
			"redirects": false,
			"destination_url": "https://zh-cn.facebook.com/",
			"destination_url_hash": "5ebd2d09c4394c0b5fb859c1c66d045aa4dad544"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://ko-kr.facebook.com/",
			"original_url_hash": "a1e9199d28108eb2f9f04ea959ecbffbf1c87410",
			"redirects": false,
			"destination_url": "https://ko-kr.facebook.com/",
			"destination_url_hash": "a1e9199d28108eb2f9f04ea959ecbffbf1c87410"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://ja-jp.facebook.com/",
			"original_url_hash": "9607cc21fc2e388a464e8ee2d741b9bd3db391d1",
			"redirects": false,
			"destination_url": "https://ja-jp.facebook.com/",
			"destination_url_hash": "9607cc21fc2e388a464e8ee2d741b9bd3db391d1"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://fr-fr.facebook.com/",
			"original_url_hash": "1d7a5661979ffea224bb3ba7d300b1e4c03b3689",
			"redirects": false,
			"destination_url": "https://fr-fr.facebook.com/",
			"destination_url_hash": "1d7a5661979ffea224bb3ba7d300b1e4c03b3689"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://es-la.facebook.com/",
			"original_url_hash": "6c991391dc185dd11a8e91bc8349f299bfc81585",
			"redirects": false,
			"destination_url": "https://es-la.facebook.com/",
			"destination_url_hash": "6c991391dc185dd11a8e91bc8349f299bfc81585"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://de-de.facebook.com/",
			"original_url_hash": "81c033379d9b1981d0f430b7cab2920de576e8f0",
			"redirects": false,
			"destination_url": "https://de-de.facebook.com/",
			"destination_url_hash": "81c033379d9b1981d0f430b7cab2920de576e8f0"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://it-it.facebook.com/",
			"original_url_hash": "e62de1ca552223c4756dbd1736f09fe5e773b9fe",
			"redirects": false,
			"destination_url": "https://it-it.facebook.com/",
			"destination_url_hash": "e62de1ca552223c4756dbd1736f09fe5e773b9fe"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://pt-br.facebook.com/",
			"original_url_hash": "d4b40e6f6595c0563188addd8ece6dc876b2a369",
			"redirects": false,
			"destination_url": "https://pt-br.facebook.com/",
			"destination_url_hash": "d4b40e6f6595c0563188addd8ece6dc876b2a369"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://ar-ar.facebook.com/",
			"original_url_hash": "f66653fec2800a347e27886065dd365886b726ab",
			"redirects": false,
			"destination_url": "https://ar-ar.facebook.com/",
			"destination_url_hash": "f66653fec2800a347e27886065dd365886b726ab"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://hi-in.facebook.com/",
			"original_url_hash": "60b29e06f26fd7989064502d6d55e543f401951f",
			"redirects": false,
			"destination_url": "https://hi-in.facebook.com/",
			"destination_url_hash": "60b29e06f26fd7989064502d6d55e543f401951f"
		},
		{
			"domain": "messenger.com",
			"tld": "com",
			"original_url": "https://messenger.com/",
			"original_url_hash": "0cd2a4846cc056b7ce31bd9cf23fb56bbb7cb21b",
			"redirects": true,
			"destination_url": "https://www.messenger.com/",
			"destination_url_hash": "cdc912efc117aa3d2232ecbe38e6d62916aadc16"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "http://l.facebook.com/l.php?h=ATP0JqcY62kf-j4LrZ28FD94-0tHO5w-vs0mkta3Lvz-28l46CvwHcgzzjhyQRdgUrLu8hUuxK3MrOvt3uC70TorbI0774D9VZHvC2DY7ulygfFqF-Q-bnsTWei2Rg\u0026u=http%3A//momentsapp.com/",
			"original_url_hash": "29fa0d15941cdffd4718637e877c6f4fe10a4add",
			"redirects": false,
			"destination_url": "http://l.facebook.com/l.php?h=ATP0JqcY62kf-j4LrZ28FD94-0tHO5w-vs0mkta3Lvz-28l46CvwHcgzzjhyQRdgUrLu8hUuxK3MrOvt3uC70TorbI0774D9VZHvC2DY7ulygfFqF-Q-bnsTWei2Rg\u0026u=http%3A//momentsapp.com/",
			"destination_url_hash": "29fa0d15941cdffd4718637e877c6f4fe10a4add"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://l.facebook.com/l.php?h=ATNOtGYv8Pqzm7boTd1_PEk414vef1_QVqgbCwhdYmG9riIDpooT2g5wjjpKGI5fXT41ViU9jGylVKdblrOToP6Kc0Jv4VC1l1sf9gxNLE9j5hhxVi6uqwsVphzQmg\u0026u=https%3A//instagram.com/",
			"original_url_hash": "8d44968997ab722f524420deeb25465f9ec5af9c",
			"redirects": false,
			"destination_url": "https://l.facebook.com/l.php?h=ATNOtGYv8Pqzm7boTd1_PEk414vef1_QVqgbCwhdYmG9riIDpooT2g5wjjpKGI5fXT41ViU9jGylVKdblrOToP6Kc0Jv4VC1l1sf9gxNLE9j5hhxVi6uqwsVphzQmg\u0026u=https%3A//instagram.com/",
			"destination_url_hash": "8d44968997ab722f524420deeb25465f9ec5af9c"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://developers.facebook.com/?ref=pf",
			"original_url_hash": "e874fb5e01201333a023e62c824858d5bb0e5858",
			"redirects": false,
			"destination_url": "https://developers.facebook.com/?ref=pf",
			"destination_url_hash": "e874fb5e01201333a023e62c824858d5bb0e5858"
		},
		{
			"domain": "facebook.com",
			"tld": "com",
			"original_url": "https://www.facebook.com/help/568137493302217",
			"original_url_hash": "a8d47b592d730c50200e8d2a07224203af182975",
			"redirects": false,
			"destination_url": "https://www.facebook.com/help/568137493302217",
			"destination_url_hash": "a8d47b592d730c50200e8d2a07224203af182975"
		}
	],
	"count": 17,
	"page": "https://www.facebook.com"
}
```
