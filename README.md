# health-check
A small Go package for quickly adding a `/health` endpoint to RESTful APIs

### Dependencies
- [Echo](https://labstack.com/echo)

### Usage
```
import "github.com/jessemillar/health"
```

```
func main() {
  port := ":8000"
  e := echo.New()

  e.Get("/health", health.Check)

  e.Run(fasthttp.New(port))
}
```
