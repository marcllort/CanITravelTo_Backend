cd CanITravelTo_Backend
echo -e "\033[1;93m [GIT]\e[0m Pulling Changes..."
git pull > /dev/null 2>&1

echo -e "\033[1;36m [Compose]\e[0m Docker-Compose Down"
docker-compose down > /dev/null 2>&1
echo -e "\033[1;34m [Docker]\e[0m Removing old containers and images..."
docker rm business-handler > /dev/null 2>&1
docker rm data-retriever > /dev/null 2>&1
docker rmi canitraveltobackend_business-handler > /dev/null 2>&1
docker rmi canitraveltobackend_data-retriever > /dev/null 2>&1

# In case the Creds folder is not there, copy
cp ../OldBinary/Creds ./DataRetriever -r -n > /dev/null 2>&1
cp ../OldBinary/Creds ./BusinessHandler -r -n > /dev/null 2>&1

echo -e "\033[1;36m [Compose]\e[0m Docker-Compose up! Starting..."
docker-compose up -d  > /dev/null 2>&1
echo -e "\033[1;32m [CanITravelTo]\e[0m CanITravelTo up and running!"

cd