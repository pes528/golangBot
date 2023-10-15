#!/bin/bash

# Telegram bot para la descarga de videos desde la red social tiktok 
# creado por : telegram @pes528


OsSystem=$(uname -o)
devNull = "/dev/null 2>&1"
installAndroid(){
    apt update && apt upgrade -y 
    apt install curl -y
    apt install git -y 
    apt install python -y 
    pip install yt-dlp -y 
    
    git clone https://github.com/pes528/golangBot

    configLinux

}

installLinux(){
    apt update -y 

    clear
    echo -e "   Instalando golang...."
    sleep 3
    sudo wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    rm go1.21.3.linux-amd64.tar.gz
    echo -e "    Finalizado"
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
    echo -e "\nA continuacion copia el token de tu bot y tu nombre de usuario de telegram\n"
    read -p "Token: " token 
    read -p "Tu nombre de usuario: " user
    datos=$(curl -i -s https://api.telegram.org/bot$token/getMe | grep -o 200)
    
    if [[ $datos -eq 200 ]];then
    echo " Datos correctos"

    echo -e "TOKEN:$token" > $var/.env
    echo -e "USER:$user" >> $var/.env
    cd $var && go get github.com/go-telegram/bot && go get github.com/joho/godotenv 
    cd ..
    clear
    sleep 3
    echo "Configuracion exitosa....dirijase a la carpeta golangBot y escriba 'go run bot.go' para iniciar el programa " 
    rm $var/insall.sh
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

