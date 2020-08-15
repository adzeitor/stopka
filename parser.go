package stopka

import (
	"unicode"

	"github.com/adzeitor/stopka/pc"
)

func parse(line string) ([]Value, error) {
	result := words().Run(line)
	if result.Err != nil {
		return nil, result.Err
	}
	return result.Value.([]Value), nil
}

// FIXME: do this:
//   identifier = (:) <$> letter <*> many (alphaNumChar <|> char '_')
func pIdentifier() pc.Parser {
	toIdentifier := func(s []rune) Value { return Identifier(s) }
	allowed := pc.OneOf(
		pc.Satisfy(unicode.IsLetter),
		pc.Char("_+-*/?="),
	)
	name := pc.Many1(allowed, []rune{})
	return name.Map(toIdentifier)
}

func pSymbol() pc.Parser {
	toSymbol := func(s Identifier) Value { return Symbol(s) }
	quote := pc.Char(`'`)
	return quote.DiscardLeft(pIdentifier()).Map(toSymbol)
}

func pInteger() pc.Parser {
	return pc.Integer().Map(func(value int) Value {
		return Integer(value)
	})
}

func pString() pc.Parser {
	quote := pc.Str(`"`)
	content := pc.Many1(pc.NotChar(`"`), []rune{})
	toString := func(runes []rune) interface{} { return String(runes) }
	return pc.Between(quote, content, quote).Map(toString)
}

func pEmptyList() pc.Parser {
	return pc.Str(`()`).DiscardLeft(pc.Pure(List{}))
}

func pList() pc.Parser {
	// return function is needed because of recursion between pList and words
	return func(state pc.State) pc.State {
		openParens := pc.Char(`(`)
		closeParens := pc.Char(`)`)
		toList := func(values []Value) List { return values }
		return pc.Between(openParens, words(), closeParens).Map(toList)(state)
	}
}

func word() pc.Parser {
	return pc.OneOf(
		pString(),
		pInteger(),
		pIdentifier(),
		pSymbol(),
		pEmptyList(),
		pList(),
	)
}

func words() pc.Parser {
	// FIXME: pc.Str -> pc.Many1(pc.Char)
	space := pc.Many(pc.Char(" "), []rune{})
	return pc.Many(pc.Between(space, word(), space), []Value{})
}
