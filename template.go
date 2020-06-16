package main

const (
	pageTemplate = `
<html>
  <head>
  <title>Cougar Angels!</title>
<style>
body {
  background-image: url("{{ .Background }}");
}
img {
  border-radius: 50%%;
  transition:all 0.3s ease;
}
</style>
<link href="data:image/x-icon;base64,AAABAAEAEBAAAAAAAABoBQAAFgAAACgAAAAQAAAAIAAAAAEACAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAAAAJ+H+ABDL/ABV5/0ADdT8AAAA3wAgvPwABgwSAIzt/AAuw/YAHd38AB3f+QAMOEQAQLrtAEC86gBu6v0APcTwAB22/QAguPoAD5r7ACHL/QAd0f0AINH9ACHZ+gAcof4AH57+ACDd/QAcp/4A5fr+ABmw/gAfp/4AHq3+AB+t/gBCx+gAILD+AMf2/AAdtv4ABwPPADTi+gCA7f4AIL/+AB/C/gAexf4AHJ3hABWd8wAoockAYun8AEKw7AAe2v4ANcPvAEC07wAhmP8Ak/D+AB7g/gC89v0AlvD+AArG+QAgpP8AKigoAB6x+QBeZOEABQMDAIbp8wArs/YAHsT8AB7T/AAhzv8AGtz8ACMBXwANPUQAHan9AB1LbQAervoAK7LZABy0+gAKJSEALCwsAItzzwBV6PsAHcT9AB/K/QAc0voAAoL2AB/T/QDc+f4Aju/9ACCX/gAfmv4Aaun+AETl/AAhmv4AHaD+AB+j/gAfpv4AIab+AAccKwActf4AH7L+ACCy/gAwAHgAIL7+ADbf/QBAquYAtvX9ALr1/QAZc54Ay+v6AB+y9gBrSe0AK7L2AKfy/gAhwPwAGtD2ADji/gBCqucAZej9AEKt5wAc3P8APrfqAEG55wAjmf0AIN78ABEyQQAMv/oAK5b9AAEBAQAl4v8AHbT9AB+6/QAdwvoAG8j6ABIFvwAgDJ4ACwAiAGTo/gAUBwoAISUkAEGv7gAglv4AHtv9AB2f/gC79PwAIKL+AAVESwA7zeUANsr3AB6r/gAiqP4An/H9AAQCBQAetP4AH7b7AB3A/gAnAI0AHBwcABeY8wASm/wAhO7+AIbt+wAczP4AIcb+AB7M/gAblPkAYej/AB/M/gBi6P8APeP6AIru/gAhz/4AIdX+AB7e/gAhmf8AId7+AB2n/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAbE0AAAAAAAAAAABqAhQWgyUFPAAAAAAAAAAdIB8il0RjmYQAAAAAAAB4M5uVaStfAIVwCgAAAAB8jHo5OVwZq1o/h59DAAAAXmtIYYl0cq2Mio49oQAAE2KAZIEODTIvZltWDAZ7AEpvoKSRRRAJdm2SjDMtQACYUFMXfZo6ITF3fxssBxIAgqmqJgBMiHmQS49JGEc7ADgaWViePnOmi6gpJF2iUgAAfqM3NmhVTqwVT2AeVwAAAAsPbhxUlKUBQU8RRpwAAAAALjSNZ6cDNRUqJJMAAAAAAAAInSeGcTBCKJYAAAAAAAAAAAAjZXVRBAAAAAAAAP+fAADwDwAA4A8AAMAHAACAAwAAgAMAAAABAAAAAQAAAAEAAAABAAAAAQAAgAMAAIADAADABwAA4A8AAPg/AAA=" rel="icon" type="image/x-icon" />
  </head>
  <body>
    <center>
    <img src="{{ .Main }}">
    </center>

  <footer><mark>This site is not affiliated with Ms.Cleckley, All data is of our users, and may not share the beliefs of the owners of this site or Ms.Cleckley. For contact, please email webmaster@cleckley.click</mark></footer>
  </body>
</html>
`
)
