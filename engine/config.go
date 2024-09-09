package engine

func Config(proxy string, device string, logLevel string) {
  key := new(Key)
  key.Proxy = proxy
  key.Device = device
  key.LogLevel = logLevel
  Insert(key)
}
