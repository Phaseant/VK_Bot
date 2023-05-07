package telegramEvents

import (
	"log"
	"strings"

	"github.com/Phaseant/VK_Bot/internal/telegram"
)

const (
	START   = "/start"
	HELP    = "/help"
	JOKE    = "Шутка 🤡"
	POEM    = "Стих 📖"
	SONG    = "Песня 🎵"
	PICTURE = "Картинка 🖼️"
	RUSSONG = "🇷🇺"
	ENGSONG = "🇺🇸"

	JOKECOMMAND    = "/joke"
	POEMCOMMAND    = "/poem"
	SONGCOMMAND    = "/song"
	PICTURECOMMAND = "/picture"

	gopherImageURL = "https://telegra.ph/file/dd33d46494b0f3e45a997.jpg"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	JokeBtn := telegram.KeyboardButton{Text: JOKE}
	SongBtn := telegram.KeyboardButton{Text: SONG}
	PoemBtn := telegram.KeyboardButton{Text: POEM}
	PictureBtn := telegram.KeyboardButton{Text: PICTURE}

	RuSongBtn := telegram.KeyboardButton{Text: RUSSONG}
	EnSongBtn := telegram.KeyboardButton{Text: ENGSONG}

	defKeyboard := telegram.ReplyKeyboardMarkup{
		Keyboard: [][]telegram.KeyboardButton{
			[]telegram.KeyboardButton{JokeBtn, SongBtn},
			[]telegram.KeyboardButton{PoemBtn, PictureBtn}},
		IsPersistent: true}

	SongKeyboard := telegram.ReplyKeyboardMarkup{
		Keyboard: [][]telegram.KeyboardButton{
			[]telegram.KeyboardButton{RuSongBtn, EnSongBtn}},
		IsPersistent: true}

	text = strings.TrimSpace(text)

	log.Printf("got command: %s", text)

	switch text {
	case START:
		return p.tg.SendMessage(chatID, StartMessage, defKeyboard)
	case HELP:
		return p.tg.SendMessage(chatID, HelpMessage, defKeyboard)
	case JOKE, JOKECOMMAND:
		return p.tg.SendMessage(chatID, JokeMessage, defKeyboard)
	case POEM, POEMCOMMAND:
		return p.tg.SendMessage(chatID, PoemMessage, defKeyboard)
	case SONG, SONGCOMMAND:
		//keyboard with two buttons
		return p.tg.SendMessage(chatID, "Выбери язык песни", SongKeyboard)
	case PICTURE, PICTURECOMMAND:
		return p.tg.SendPicture(chatID, gopherImageURL)
	case RUSSONG:
		return p.tg.SendMessage(chatID, RusSongMessage, defKeyboard)
	case ENGSONG:
		return p.tg.SendMessage(chatID, EngSongMessage, defKeyboard)
	default:
		return p.tg.SendMessage(chatID, unknownMessage, defKeyboard)
	}
}
