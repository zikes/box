<h1 align="center">
<pre>
╔═════════╗
║   BOX   ║
╚═════════╝</pre>
</h1>

```go
b := &box.Box{
  BoxStyle: &box.DoubleStyle,
  Margin:  1,
  Padding: 3,
}

fmt.Print(box.Sprint("BOX"))
// ╔═════════╗
// ║   BOX   ║
// ╚═════════╝
```
