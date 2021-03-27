# HTML-Parser
HTML link Parser. Go exercise

The goal of this project is to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link it should return a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Input: 

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

Output:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```
