sudo systemctl stop canitravelto.service
rm CanITravelTo
mv CanITravelToUpdated CanITravelTo
chmod 700 CanITravelTo
sudo systemctl enable canitravelto.service
sudo systemctl start canitravelto.service
