Два варианта выполнения домашнего задания:

1. Добавить эндпоинты для админа
    * для блокировки/разблокировки пользователя;
    * назначение/удаление роли пользователя;
    * сброс пароля пользователя с установкой нового сгенерированного пароля
      (подразумевается, что новый пароль будет передан пользователю через другие каналы связи, например, почту)

2. Вариант со звездочкой *
   Добавить телеграмм бота для входа в сервис через ввод кода, отправленного сервисом авторизации в телеграмм-бота.
   Для этого вам понадобятся эндпоинты:
    * привязки телеграмм бота к авторизованному аккаунту;
    * входа с логином и отправкой кода в телеграмм-чат;
    * подтверждения кода из телеграмм чата. После успешной авторизации, код должен быть сброшен.


Критерии приемки:

1. Если вы выбрали простой первый вариант, то критерии приемки ДЗ остаются такими же как и раньше.
   Максимальная оценка минус коэффициент за критичные замечания в код ревью.

2. Если вы выбрали вариант со звездочкой:
    * при корректной работе приложения, будет поставлена наивысшая оценка вне зависимости от замечаний на код ревью;
    * при некорректной работе приложения в любом из эндпоинтов работы с телеграмм ботом, оценка будет занижена до
      минимальной (0) вне зависимости от количества ошибок и верно выполненных шагов.

upd:
для первого варианта домашки нужно сделать "назначение/удаление ролИ пользователя", а не ролей.
Имеющаяся структура данных пользователя не подразумевает, что пользователь может иметь несколько ролей. Поэтому при удалении роли пользователь должен получить роль по умолчанию, те user