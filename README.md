## GinConf: A common configuration package for Gin project

### usage
```go
// main.go
import (
    cfg "ginconf/config"
)

type Config struct {
    cfg.Base `ini:"base"`
    ProjectName string `int:"project_name"`
}

func main() {
    r := setupRouter()
    files := []interface{}{"base.ini", "local.ini"}
    c := Config{}
    cfg.Initialize(&c, files...)
    r.Run(c.ListenAddr)
}
```

```ini
# base.ini
[base]
listen_address = :8080
```

```ini
# local.ini
project_name = myProject
```