package cli

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/chasefleming/elem-go/attrs"
	"github.com/mgred/html-filer/pkg/html"
)

type Options struct {
	Help    bool
	Out     string
	Title   string
	Base    string
	From    string
	Version bool
	Assets  []html.Asset
}

type TokenType int

const (
	Argument TokenType = iota
	Option
)

type Token struct {
	Position int
	Type     TokenType
	Value    string
	Raw      string
}

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(t []Token) *Parser {
	return &Parser{t, -1}
}

func (p *Parser) Next() (Token, bool) {
	p.pos++
	return p.current()
}

func (p *Parser) Reset() {
	p.pos--
}

func (p *Parser) current() (Token, bool) {
	if p.pos >= len(p.tokens) {
		return Token{}, false
	}
	return p.tokens[p.pos], true
}

func ParseStylesSubOptions(p *Parser, prev Token) (result []html.Asset, err error) {
	var props = attrs.Props{}
	var insert bool
LOOP:
	for arg, value := p.Next(); value != false; arg, value = p.Next() {
		if arg.Type == Option {
			switch arg.Value {
			case "insert", "i":
				insert = true
			case "media", "m":
				next, e := p.Next()
				if !e {
					return result, noArgumentForOption(arg)
				}
				props[attrs.Media] = next.Value
			default:
				p.Reset()
				break LOOP
			}
		}

		if arg.Type == Argument {
			result = append(result, html.Asset{
				Parent: "head",
				Props:  props,
				Path:   arg.Value,
				Insert: insert,
				Type:   html.Style,
			})
		}
	}

	if len(result) == 0 {
		return result, noArgumentForOption(prev)
	}
	return
}

func ParseScriptSubOptions(p *Parser, prev Token) (result []html.Asset, err error) {
	var props = attrs.Props{}
	var insert bool
	parent := "body"
LOOP:
	for arg, value := p.Next(); value != false; arg, value = p.Next() {
		if arg.Type == Option {
			switch arg.Value {
			case "module", "m":
				props[attrs.Type] = "module"
			case "head", "h":
				parent = "head"
			case "insert", "i":
				insert = true
			case "async", "a":
				props[attrs.Async] = "true"
			default:
				// If the option is none of the above
				// and we have at least on argument, we
				// know, that this is the next option and
				// return. With no arguments collected so far
				// this option must be misplaced here.
				// TODO: Do we need this?
				// if len(result) == 0 {
				// 	return result, fmt.Errorf("Option `%s` is misplaced here", arg.Value)
				// }
				p.Reset()
				break LOOP
			}
		}
		if arg.Type == Argument {
			result = append(result, html.Asset{
				Parent: parent,
				Props:  props,
				Path:   arg.Value,
				Insert: insert,
				Type:   html.Script,
			})
		}
	}

	if len(result) == 0 {
		return result, noArgumentForOption(prev)
	}
	return
}

func TokenizeArgs(args []string) []Token {
	result := make([]Token, len(args))
	for i, arg := range args {
		if strings.HasPrefix(arg, "-") {
			result[i] = Token{i, Option, strings.TrimLeft(arg, "-"), arg}
		} else {
			result[i] = Token{i, Argument, arg, arg}
		}
	}
	return result
}

func unwrapNextWithErr(p *Parser) (string, error) {
	v, e := p.Next()
	if e {
		return v.Value, nil
	}
	return "", errors.New("Could not read next Option")
}

func ProcessArgs(args []string, opts *Options) (err error) {
	tokens := TokenizeArgs(args)

	var assets []html.Asset
	parser := NewParser(tokens)
	for arg, value := parser.Next(); value != false; arg, value = parser.Next() {
		if arg.Type == Argument {
			assets = append(assets, ParseAsset(arg))
		}

		if arg.Type == Option {
			switch arg.Value {
			case "help", "h":
				opts.Help = true
			case "version", "v":
				opts.Version = true
			case "base", "b":
				opts.Base, err = unwrapNextWithErr(parser)
			case "out", "o":
				opts.Out, err = unwrapNextWithErr(parser)
			case "title", "t":
				opts.Title, err = unwrapNextWithErr(parser)
			case "assets", "a":
				var a []html.Asset
				a, err = ParseNextArgumentsAsAssets(parser, arg)
				assets = append(assets, a...)
			case "scripts", "s":
				var s []html.Asset
				s, err = ParseScriptSubOptions(parser, arg)
				assets = append(assets, s...)
			case "styles", "S":
				var s []html.Asset
				s, err = ParseStylesSubOptions(parser, arg)
				assets = append(assets, s...)
			default:
				err = fmt.Errorf("ERROR: Unknown option `%s` at position %d", arg.Raw, arg.Position)
			}

			if err != nil {
				return
			}
		}
	}

	opts.Assets = assets

	return
}

func ParseAsset(a Token) html.Asset {
	switch filepath.Ext(a.Value) {
	case ".js":
		return html.Asset{
			Type:   html.Script,
			Parent: html.BODY_TAG,
			Path:   a.Value,
			Insert: false,
		}
	case ".mjs":
		return html.Asset{
			Type:   html.Script,
			Parent: html.BODY_TAG,
			Props: attrs.Props{
				attrs.Type: "module",
			},
			Path:   a.Value,
			Insert: false,
		}
	case ".css":
		return html.Asset{
			Type:   html.Style,
			Parent: "head",
			Insert: false,
			Path:   a.Value,
		}
	}

	return html.Asset{}
}

func ParseNextArgumentsAsAssets(p *Parser, t Token) (result []html.Asset, err error) {
	for arg, value := p.Next(); value != false; arg, value = p.Next() {
		if arg.Type == Argument {
			result = append(result, ParseAsset(arg))
		} else {
			p.Reset()
			break
		}
	}

	if len(result) == 0 {
		return result, noArgumentForOption(t)
	}

	return
}

func noArgumentForOption(o Token) error {
	return fmt.Errorf("ERROR: No argument(s) for `%s` at position %d", o.Raw, o.Position)
}
