# Vypolnyator  

* Бот в телеге принимает запросы на создание списков и записей.  
* Есть крон, который раз в заданный срок напоминает о повторении и меняет даты в записях?  


##Todo:
1. Сервис для авторизации пользователей + DB:  
  * JWT  
  * DB user (id, telegram_id, name)
2. БД для напоминаний:  
  * note_group (id, user_id, name, description, notiy_time)  
  * note (id, group_id, name, description, notiy_time) 
 