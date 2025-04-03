---
author: Lux Barker
paging: '%d / %d'
---

# dwnld: a cross-platform desktop app made with go & svelte

agenda:

- the what
- the why
- the how

---

## the what

we're going to talk about:

- what it is
- what it is not
- quick demo

---

### what it is

- minimal and simple
- a gui wrapper around [yt-dlp](https://github.com/yt-dlp/yt-dlp)

---

### what it is not

- an implementation of a wrapper around yt-dlp
  - someone else already did the wrapping of yt-dlp - [go-ytdlp](github.com/lrstanley/go-ytdlp)!
- a full replacement for everything yt-dlp can do

---

### ok cool but can i see it?

- `<insert demo here>`

---

## the why

2 main reasons:

- i wanted to learn new things
  - [wails](https://wails.io/)
  - [svelte](https://svelte.dev/)
  - [tailwind](https://v3.tailwindcss.com/)
- i wanted to make something my partner could (and would) use
  - he is not a techy person
  - he is fairly resistant to new software - once he find something that works, he sticks with it
    - as someone who loves trying out new tools, his commitment to sticking with things impresses and confuses me
  - but he also likes simplicity & minimalism
  - the web-based tools he was using before were:
    - often visually loud & obnoxious to use
    - offered features he didn't care about
    - but didn't offer features he actually wanted

---

## the how (slash the things i learned (slash what i'd still like to implement))

- what even is wails
- what wails is not
- frontend (front-end? front end??) is hard
- what's next

---

### wails is

- a go framework for building cross-platform gui applications with a webview front-end
- a lightweight electron alternative
- automatic generation of typescript models from go structs
- an easy way to call go methods from typescript
- seriously, all you have to do is add your structs to a config option when calling `wails.Run`:

```go
type app struct {}
type returned struct { val int }
func (a *app) CallMe() returnedStruct { return returnedStruct{ val: 3 } }
func main() {
 err := wails.Run() &options.App{
  Bind: []interface{}{
   app,
  },
  // ...
 }
}
```

this results in a typescript model for `returnedStruct` and a function called `CallMe`!

- there's also support for some things like native dialogs & menus
- and many more nice things!

---

### wails is not

- a full-on electron replacement
  - it uses native webview libraries instead of embedding a full chromium instance
  - this makes it lighter than electron, but a little more limited
    - [can i webview](https://caniwebview.com/)
    - example: had to build my own directory input component
      - there is a `webkitdirectory` attribute that you can supposedly add to `<input type=file>` elements for this
        - [can i use](https://caniuse.com/input-file-directory) reports it works in safari/chrome/firefox
        - it did not work for me in webview :(

---

### frontend is hard

- but also kind of fun?
- svelte is relatively simple and i probably would have had a less hard time if i didn't try using sveltekit
- i like sveltekit (a lot!) but it seems to not really want to be used in spa-mode, which wails requires
- you can't use server load functions with wails
  - many resources assume you are using server load functions
- [shadcn-svelte](https://next.shadcn-svelte.com/) is wonderful
- i'm still figuring out tailwind and maybe should have just used vanilla css to re-learn how that even really works

---

### what's next

i'm putting dwnld down for now BUT future improvements could include:

- letting the user select preferred codecs, file formats, all the things the yt-dlp [format option allows](https://github.com/yt-dlp/yt-dlp?tab=readme-ov-file#format-selection)
  - including allowing for only audio
- saving multiple output template strings
- more granular subtitle options
- github action to build & release binaries
- test on linux

---

## that's all, thank you

i had a really fun time with this!
please reach out if you're interested in learning more or pairing on wails, svelte, or anything else go-related!
