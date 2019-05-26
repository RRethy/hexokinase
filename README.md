# hexokinase

Fast text parser to scrape and convert colours in the form of rgb, rgb, hsl, hsla functions, three and six digit hex values, web colours names, and custom patterns into hex values.

## Installation

```
go get github.com/RRethy/hexokinase
```

## Usage

```
Usage of ./hexokinase:
  -dp string
    	disabled patterns which will not be parsed for. Comma separated list
    	with possible values of hex, rgb, rgba, hsl, hsla, names. The "names"
    	argument refers to web colour names.
  -extended
    	print results in the format "filename:lnum:colstart-colend:hex:line" (default true)
  -files string
    	files to parse (or stdout to parse stdout) (default "stdout")
  -palettes string

    	palette file names.
    	This has the ability to define additional patterns to match and how to convert
    	them into hex values. The file must be a valid json file that looks like this:

    	{
    	  "regex_pattern": "foo[0-9]bar[0-9]baz[0-9]",
    	  "colour_table": {
    	    "foo1bar1baz1": "#eb00ff",
    	    "foo2bar2baz2": "#ffeb00",
    	    "foo3bar3baz3": "#00ffeb"
    	  }
    	}

    	The "regex_pattern" key is optional. If omitted, every key in the
    	"colour_table" map will be matched instead of using the regex.
    	Any key in the "colour_table" map will be matched and have the associated hex
    	string that is provided outputted.
    	If the regex matches a string which is not a key in the "colour_table" map, it
    	will be discarded as a false positive.
    	No checking is done on the hex strings so technically they can be any string.

  -simplified
    	same as -extended but don't print the full line. Overrides -extended.
```

### Examples

```sh
# Parse foo.json and bar.json for all colours except hsl and hsl functions
# Output the results in simplified format
hexokinase -dp=hsl,hsla -files=foo.json,bar.json -s

# Parse stdin all colours
# Output the results in extended format
hexokinase

# Parse foo.json and bar.json for all colours and for the colour patterns
# specified in the palettes p1.json and p2.json.
# Output the results in extended format.
hexokinase -files=foo.json,bar.json -palettes=p1.json,p2.json
```

## Integrations

* [vim-hexokinase](https://github.com/RRethy/vim-hexokinase) - WIP

# TODO

* Cleanup public API
* Improve README

## License

`mit`
