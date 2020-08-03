cd CanITravelTo_Backend
echo -e "\033[1;93m [Git]\e[0m Pulling Changes..."
git pull > /dev/null 2>&1

echo -e "\033[1;36m [Compose]\e[0m Docker-Compose Down"
docker-compose down > /dev/null 2>&1

echo -e "\033[1;34m [Docker]\e[0m Removing old containers and images..."
docker rm business-handler > /dev/null 2>&1
docker rm data-retriever > /dev/null 2>&1
docker rmi canitraveltobackend_business-handler > /dev/null 2>&1
docker rmi canitraveltobackend_data-retriever > /dev/null 2>&1

echo -e "\033[1;34m [Docker]\e[0m Pulling new images..."
docker pull docker.pkg.github.com/marcllort/canitravelto_backend/data-retriever:"$(git rev-parse --short HEAD)"  > /dev
docker pull docker.pkg.github.com/marcllort/canitravelto_backend/business-handler:"$(git rev-parse --short HEAD)"  > /d$

echo -e "\033[1;36m [Compose]\e[0m Docker-Compose up! Starting..."
BUSINESS_VERSION="$(git rev-parse --short HEAD)" DATA_VERSION="$(git rev-parse --short HEAD)" DB_PASSWORD="$1" docker-compose up -d  > /dev/null 2>&1
echo -e "\033[1;32m [CanITravelTo]\e[0m CanITravelTo up and running!"

cd