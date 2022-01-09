# wsession

![GitHub Repo stars](https://img.shields.io/github/stars/wyy-go/wsession?style=social)
![GitHub](https://img.shields.io/github/license/wyy-go/wsession)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/wyy-go/wsession)
![GitHub CI Status](https://img.shields.io/github/workflow/status/wyy-go/wsession/ci?label=CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/wyy-go/wsession)](https://goreportcard.com/report/github.com/wyy-go/wsession)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/wyy-go/wsession?tab=doc)
[![codecov](https://codecov.io/gh/wyy-go/wsession/branch/main/graph/badge.svg)](https://codecov.io/gh/wyy-go/wsession)

Gin middleware for session management with multi-backend support:

- [cookie-based](#cookie-based)
- [Redis](#redis)
- [memstore](#memstore)


## Usage

### Start using it

Download and install it:

```bash
go get github.com/wyy-go/wsession
```

Import it in your code:

```go
import "github.com/wyy-go/wsession"
```

## Basic Examples

### single session

```go
package main

import (
  "github.com/wyy-go/wsession"
  "github.com/wyy-go/wsession/cookie"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store := cookie.NewStore([]byte("secret"))
  r.Use(wsession.New("mysession", store))

  r.GET("/hello", func(c *gin.Context) {
    session := wsession.Default(c)

    if session.Get("hello") != "world" {
      session.Set("hello", "world")
      session.Save()
    }

    c.JSON(200, gin.H{"hello": session.Get("hello")})
  })
  r.Run(":8000")
}
```

### multiple sessions

```go
package main

import (
  "github.com/wyy-go/wsession"
  "github.com/wyy-go/wsession/cookie"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store := cookie.NewStore([]byte("secret"))
  sessionNames := []string{"a", "b"}
  r.Use(wsession.NewMany(sessionNames, store))

  r.GET("/hello", func(c *gin.Context) {
    sessionA := wsession.DefaultMany(c, "a")
    sessionB := wsession.DefaultMany(c, "b")

    if sessionA.Get("hello") != "world!" {
      sessionA.Set("hello", "world!")
      sessionA.Save()
    }

    if sessionB.Get("hello") != "world?" {
      sessionB.Set("hello", "world?")
      sessionB.Save()
    }

    c.JSON(200, gin.H{
      "a": sessionA.Get("hello"),
      "b": sessionB.Get("hello"),
    })
  })
  r.Run(":8000")
}
```

## Backend Examples

### cookie-based

```go
package main

import (
  "github.com/wyy-go/wsession"
  "github.com/wyy-go/wsession/cookie"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store := cookie.NewStore([]byte("secret"))
  r.Use(wsession.New("mysession", store))

  r.GET("/incr", func(c *gin.Context) {
    session := wsession.Default(c)
    var count int
    v := session.Get("count")
    if v == nil {
      count = 0
    } else {
      count = v.(int)
      count++
    }
    session.Set("count", count)
    session.Save()
    c.JSON(200, gin.H{"count": count})
  })
  r.Run(":8000")
}
```

### Redis

```go
package main

import (
  "github.com/wyy-go/wsession"
  "github.com/wyy-go/wsession/redis"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
  r.Use(wsession.New("mysession", store))

  r.GET("/incr", func(c *gin.Context) {
    session := wsession.Default(c)
    var count int
    v := session.Get("count")
    if v == nil {
      count = 0
    } else {
      count = v.(int)
      count++
    }
    session.Set("count", count)
    session.Save()
    c.JSON(200, gin.H{"count": count})
  })
  r.Run(":8000")
}
```

### memstore

```go
package main

import (
  "github.com/wyy-go/wsession"
  "github.com/wyy-go/wsession/memstore"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  store := memstore.NewStore([]byte("secret"))
  r.Use(wsession.New("mysession", store))

  r.GET("/incr", func(c *gin.Context) {
    session := wsession.Default(c)
    var count int
    v := session.Get("count")
    if v == nil {
      count = 0
    } else {
      count = v.(int)
      count++
    }
    session.Set("count", count)
    session.Save()
    c.JSON(200, gin.H{"count": count})
  })
  r.Run(":8000")
}
```