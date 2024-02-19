# Go domain crawler

A simple tool to collect all document type files/hrefs from a domain.

## Getting Started

### Installation
```
go get -u github.com/kojoluh/go-domain-crawler
```

### Usage

```
go-domain-crawler -domain https://test.com -docType js
```

Pipe the response into other tools
```
go-domain-crawler -domain https://bbc.com -docType a > links.txt
go-domain-crawler -domain https://bbc.com -docType js | wc -l
go-domain-crawler -domain https://bbc.com -docType css | httprobe
```

### Reference

- https://www.kelche.co/blog/go/flag/
- https://github.com/PuerkitoBio/goquery
