#!/usr/bin/env bash

# 打印带颜色的文本的函数
print_colored_text() {
  local color=$1
  local style=$2
  shift 2
  local text=$@

  case $color in
    "black")   color=0;;
    "red")     color=1;;
    "green")   color=2;;
    "yellow")  color=3;;
    "blue")    color=4;;
    "magenta") color=5;;
    "cyan")    color=6;;
    "white")   color=7;;
    *)         color=9;;
  esac

  case $style in
    "normal")    style=0;;
    "bold")      style=1;;
    "underline") style=4;;
    "reverse")   style=7;;
    *)           style=0;;
  esac

  echo -e "\033[0;3${color};4${style}m${text}\033[0m"
}

quit_with_error(){
  echo "UIF Install Error: "
  print_colored_text "red" "normal" $1
  echo ""
  exit -1
}

check_arch(){
  if [[ $arch == "x86_64" || $arch == "x64" || $arch == "amd64" ]]; then
      UIFARCH="amd64"
  elif [[ $arch == "aarch64" || $arch == "arm64" ]]; then
      UIFARCH="arm64"
  else
      UIFARCH="amd64"
  fi

  print_colored_text "Arch: ${UIFARCH}"
}

green_print(){
  print_colored_text "green" "bold" $1
}

yellow_print(){
  print_colored_text "yellow" "normal" $1
}

if [ "$(id -u)" != "0" ]; then
  quit_with_error 'Need Root User!'
fi

if ! [ -x "$(command -v systemctl)" ]; then
  quit_with_error 'Missing "systemd"'
fi

if [ -x "$(command -v yum)" ]; then
  manager="rpm"
elif [ -x "$(command -v apt)" ]; then
  manager="deb"
else
  quit_with_error "Required 'yum' or 'apt'"
fi
check_arch

FILE_NAME="uif-linux-$UIFARCH.$manager"
CACHE_PATH="./$FILE_NAME"

download() {
  DOWNLOAD_LINK="https://github.com/UIforFreedom/UIF/releases/latest/download/$FILE_NAME"
  echo "Downloading from: $DOWNLOAD_LINK"
  if ! curl -R -L -H 'Cache-Control: no-cache' -o "$1"  "$DOWNLOAD_LINK"; then
    quit_with_error 'Download failed! Please check your network or try again.'
  fi
}

sudo systemctl stop uiforfreedom
download $CACHE_PATH

if [[ $manager=="deb" ]]; then
  sudo apt remove uiforfreedom
  if ! sudo apt install $CACHE_PATH; then
    quit_with_error "Failed to install 'uif' by 'apt'."
  fi
  if ! sudo apt install ufw; then
    quit_with_error "Failed to install 'ufw' by 'apt'."
  fi
else
  sudo yum remove uiforfreedom
  if ! sudo yum localinstall $CACHE_PATH; then
    quit_with_error "Failed to install 'uif' by 'yum'."
  fi
  if ! sudo yum install ufw; then
    quit_with_error "Failed to install 'ufw' by 'yum'."
  fi
fi


if ! sudo systemctl start uiforfreedom; then
  quit_with_error "Failed to execute systemctl"
fi


chmod -R 755 /usr/bin/uif/
echo "0.0.0.0:80" >> /usr/bin/uif/uif_api_address.txt

ufw enable
ufw allow 80/tcp
ufw allow 22/tcp
systemctl restart uiforfreedom

green_print "Using port 80. 'http://{YourIPAddress}:80' to connect."
echo "Password:"
cat /usr/bin/uif/uif_key.txt
echo ""
green_print "OK!!!"
