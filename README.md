# egg-service-api

<p align="justify">
	
Это сервис для передачи в базу данных информации от разработанной мной яцеварки. Визуализация работы устройства:
</justify> <br /> 
 <img  src="./assets/scene.gif" width="100%">
 <br /> 

<p align="justify">
	
Для симуляции работы устройства спроектирована схема в Proteus с использованием микроконтроллера Ардуино Uno, датчиков веса HX-711, ТЭНа. Устройство работает в трех режимах (вкрутую всмятку и в мешочек). Программа микроконтроллера написана на C++ и позволяет проверять наличие яиц и воды в отсеках, готовить яйца с помощью трех программ и посылать информацию о приготовленных порциях на сервер через подключенный по COM-порту сетевой адаптер. 
</justify> <br /> 
 <img  src="./assets/SCHEME.png" width="100%">
 <br /> 

<p align="justify">
	
В программе Компас 3D был создан чертеж устройства.
</justify> <br /> 
 <img  src="./assets/Куратор2000.png" width="100%">
 <br /> 

 <p align="justify">
	
Работу серверной части (сервер БД и сервер API полняты в контейнере) можно проверить с помощью Postman/Insomnia
</justify> <br /> 
 <img  src="./assets/SampleRequest.jpg" width="100%">
 <br /> 

 <p align="justify">
 Информация отразилась  базе данных
</justify> <br /> 
 <img  src="./assets/Results.jpg" width="100%">
 <br /> 

 <p align="justify">
 Клиент был написан на Java, он связывается с API сервером с помощью HTTP 1.1 и парсит ответ в объект класса Meal, после чего выводит историю приготовленных блюд.
</justify> <br /> 
 <img  src="./assets/test db.gif" width="100%">
 <br /> 
