#!/bin/bash

# Telegram bot para la descarga de videos desde la red social tiktok 
# creado por : telegram @pes528

verde="\033[1;32m"
fin="\033[0m"
OsSystem=$(uname -o)

installAndroid(){
    apt update && apt upgrade -y 
    clear
    echo -e "${verde}Instalando dependencias...${fin}\n"
    sleep 3
    apt install curl -y
    echo -e "${verde}Instalando golang...${fin}"
    apt install golang -y
    apt install git -y 
    clear
    echo -e "${verde}Instalando python...${fin}"
    apt install python -y 
    clear
    echo -e "${verde}Instalando yt-dlp...${fin}"
    pip install yt-dlp 
    
    echo -e "${verde}\n\nInstalación finalizada ${fin}"
    sleep 2
    git clone https://github.com/pes528/golangBot
    
    configLinux

}

installLinux(){
    sudo apt update -y 

    clear
    echo -e "${verde}   Instalando golang.... ${fin}"
    sleep 3
    sudo wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
    #sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    sudo rm -rf go1.21.3.linux-amd64.tar.gz
    echo -e "${verde}    Finalizado ${fin}"
    sleep 3
    echo -e "    Instalando ytdlp"
    sleep 2
    sudo wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp
    sudo chmod a+rx /usr/local/bin/yt-dlp  # Make executable

    echo -e "    Finalizado"
    git clone https://github.com/pes528/golangBot
    configLinux

}
configLinux(){
    var="golangBot"
    clear
    echo -e "${verde}\nA continuacion copia el token de tu bot y tu nombre de usuario de telegram${fin}\n"
    read -p "Token: " token 
    read -p "Tu nombre de usuario: " user
    datos=$(curl -i -s https://api.telegram.org/bot$token/getMe | grep -o 200)
    
    if [[ $datos -eq 200 ]];then
    echo -e "${verde} Datos correctos ${fin}"

    echo -e "TOKEN:$token" > $var/.env
    echo -e "USER:$user" >> $var/.env
    cd $var && go get github.com/go-telegram/bot && go get github.com/joho/godotenv 
    cd ..
    clear
    sleep 3
    echo -e "\n${verde}Configuracion exitosa....dirijase a la carpeta golangBot y escriba 'go run bot.go' para iniciar el programa ${fin}" 
    rm $var/install.sh
    rm install.sh
    echo -e "\n${verde}Iniciando ..${fin}\nPara detener el pulse las teclas ctrl+c"
    cd golangBot && go run bot.go
    else
    clear
    sleep 3
    echo "Datos incorrectos"
    rm -rf $var
    
    fi
}



if [ $OsSystem == "Android" ];then
    installAndroid
elif [ $OsSystem == "GNU/Linux" ];then
    installLinux
else
    echo "No se reconoce sistema"
fi
