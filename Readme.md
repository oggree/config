# Config

It's just a simple tool for load yaml files inside ./configs folder to viper.

Then you can access them by viper

```go
Get(key string) : any
GetBool(key string) : bool
GetFloat64(key string) : float64
GetInt(key string) : int
GetIntSlice(key string) : []int
GetString(key string) : string
GetStringMap(key string) : map[string]any
GetStringMapString(key string) : map[string]string
GetStringSlice(key string) : []string
GetTime(key string) : time.Time
GetDuration(key string) : time.Duration
IsSet(key string) : bool
AllSettings() : map[string]any
```