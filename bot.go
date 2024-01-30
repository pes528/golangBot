package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"net/http"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

/*
Bot de telegram para descargar videos desde tiktok creado por @pes528

*/
func main() {
	com := os.Args
	if len(com[0:]) > 1 && com[1] == "config"{
		fmt.Println("config")
		config()
	}else{
		ini()
	}

}


func ini(){
	
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	if godotenv.Load() != nil{
		log.Fatal("\n\033[1;33m Error archivo .env ejecuta 'go run golangBot config' \033[0m")
	}

	b, err := bot.New(os.Getenv("TOKEN"), opts...)
	if nil != err {
		
		fmt.Println("Error")
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, start)
	fmt.Println()
	fmt.Println("		BOT RUNNING.......")
	b.Start(ctx)
}



func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println(update.Message.Text)
	fmt.Println(ctx)
	
	usuario := os.Getenv("USERTELEGRAM")
	
	enviado := tiktok(update.Message.Text)
	//sti := update.Message.Sticker
	//s:= sti.FileID
	chat := update.Message.ID
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   enviado + "\nlink->: "+ update.Message.Text ,//update.Message.Text + " sdasd",
	})
	if enviado != "Link incorrecto"{
		file, err := os.Open("video.mp4")
		if err != nil{
			return
		}
		defer file.Close()
		b.SendChatAction(ctx, &bot.SendChatActionParams{
			ChatID: update.Message.Chat.ID,
			Action: "upload_video",
		})
		b.SendVideo(ctx, &bot.SendVideoParams{
			ChatID: update.Message.Chat.ID,
			Video: &models.InputFileUpload{Filename: "video.mp4", Data: file},
			ReplyToMessageID: chat,
			Caption: "Created by: telegram @pes528\nUser : "+usuario,
		})
		
		com := exec.Command("rm", "video.mp4")
		com.Run()
		b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID: update.Message.Chat.ID,
			MessageID: chat+1,
		})
	}else{
	b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID: update.Message.Chat.ID,
		MessageID: chat,
	})}
	
}


func start(ctx context.Context, b *bot.Bot, update *models.Update){
	//var IdSticker string = "CAACAgEAAxkBAAIBo2UqxVSSj52ACwlGBHYWDyEs7oFGAALfAAMwmxFEG-rHFOoHsaEwBA"
	b.SendSticker(ctx, &bot.SendStickerParams{
		ChatID: update.Message.Chat.ID,
		Sticker: &models.InputFileString{Data: stickerRandom()},
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "Hola "+ update.Message.Chat.FirstName + " unete al canal de telegram @tec_tricks\nEste bot puede descargar videos de tiktok solo envia el link ",
	})
}




func tiktok(link string)string{
	//var comando string = `yt-dlp -o video".%(ext)s" `+link
	com := exec.Command("yt-dlp", "-o", "video.%(ext)s",link)

	out, err := com.Output()
	if err != nil{
		fmt.Println(err)
		return string("Link incorrecto")
	}
	
	return string(out)
	
}



func stickerRandom()string{
	
	stickers := []string{
		"CAACAgEAAxkBAAIBjmUqwfR-F5fn7CMHFsFGn3hAVPH-AAI4AQACtvgQRIZn_Z9WqMrpMAQ",
		"CAACAgEAAxkBAAIBo2UqxVSSj52ACwlGBHYWDyEs7oFGAALfAAMwmxFEG-rHFOoHsaEwBA",
		"CAACAgIAAxkBAAIBvWUrDhTMdacsbYXYhj__3ICQLf_NAAKLAQACK15TC6NhvGkkNINQMAQ",
		"CAACAgIAAxkBAAIBwGUrDnlHt_9Vpitap3pBIs7T16QgAAJUHQACipHQSVRMwb4UEM47MAQ",
		"CAACAgIAAxkBAAIBw2UrDsahy72IuzxeL2uLDKt9EJ3tAAI1AQACMNSdEbS4Nf1moLZ8MAQ",
		"CAACAgIAAxkBAAIBxmUrD3D-oURz_fb6knZHTcFJ_jpmAAIYAAMNttIZfrvWCJAMlTwwBA",

	}
	num := stickers[rand.Intn(len(stickers))]
	return num
}

func config(){
	fmt.Print("\nTOKEN: ")
	var token string
	fmt.Scanln(&token)
	fmt.Print("\nUser: ")
	var use string 
	fmt.Scanln(&use)

	pet, err := http.Get("https://api.telegram.org/bot"+token+"/getMe")
	if err != nil{
		panic(err)
	}
	if pet.StatusCode != 200{
		log.Fatal("\033[1;31m Datos incorrectos \033[0m")
	}else{

		file:= os.WriteFile(".env", []byte("TOKEN:"+token+"\nUSERTELEGRAM:"+use+"\n"), 0666)
		if file != nil{
			log.Fatal(err)
		}
		fmt.Println("CONFIGURACION EXITOSA EJECUTA 'go run bot.go' para iniciar el programa")

	}




}
