Для создания API для управления инвентарем на складе, вы можете следовать приведенному ниже техническому заданию (ТЗ). Этот проект будет включать в себя базовую функциональность для управления товарами на складе, включая добавление, удаление, обновление и перемещение товаров между разными локациями.

Общее описание
Цель: Разработать RESTful API для управления инвентарем, который позволит пользователям взаимодействовать с базой данных товаров на складе.

Основные функции:

Добавление товара: Пользователи могут добавлять новые товары в систему, указывая их название, количество, цену за единицу и ID склада.
Обновление товара: Пользователи могут обновлять информацию о товаре, включая название, количество, цену и местоположение.
Удаление товара: Пользователи могут удалять товары из системы.
Просмотр товара: Пользователи могут просматривать список всех товаров или получать информацию о конкретном товаре.
Перемещение товара: Пользователи могут перемещать товары между складами, указывая исходный и целевой склады и количество перемещаемых товаров.
Технические детали
Технологии:
Язык программирования: Go
База данных: PostgreSQL
Дополнительные инструменты: Docker для развертывания, Swagger для документации API
Эндпоинты API:
POST /items - добавление нового товара.
Входные данные: название, количество, цена за единицу, ID склада.
PUT /items/{id} - обновление информации о товаре.
Входные данные: название (опционально), количество (опционально), цена за единицу (опционально), ID склада (опционально).
DELETE /items/{id} - удаление товара по ID.
GET /items - получение списка всех товаров.
GET /items/{id} - получение информации о конкретном товаре.
POST /items/move - перемещение товара между складами.
Входные данные: ID товара, исходный ID склада, целевой ID склада, количество.
Безопасность:
Аутентификация через JWT для всех эндпоинтов.
Разграничение доступа: только авторизованные пользователи могут создавать, обновлять и удалять информацию.
Тестирование:
Написание юнит-тестов для каждого эндпоинта.
Интеграционные тесты для проверки взаимодействия с базой данных.
Развертывание:
Контейнеризация приложения и базы данных с использованием Docker.
Инструкция по развертыванию с использованием Docker Compose.
Исходный код:
Структурирование кода с использованием пакетов для разделения логики работы с базой данных, бизнес-логики и обработчиков API.
Использование ORM для взаимодействия с базой данных.

Документация API с использованием Swagger или Postman для облегчения тестирования и интеграции для разработчиков.

Дополнительные требования и уточнения:

Функциональные Требования
История перемещения товаров: В системе должна вестись история всех операций перемещения товаров, включая дату перемещения, исходный и целевой склады, а также количество перемещенных товаров.
Поиск и фильтрация: Добавить возможность поиска товаров по названию и фильтрацию по складу, цене и количеству через GET /items с соответствующими query параметрами.
Аудит изменений: Система должна сохранять историю изменений для каждого товара, включая изменение количества, цены и названия.
Небезопасные Операции и Ограничения
Операции добавления, обновления и удаления товаров должны быть доступны только для пользователей с ролью "Администратор".
Операция перемещения товаров должна проверять наличие достаточного количества товара на исходном складе перед выполнением перемещения.
Должна быть реализована проверка на минимально допустимое количество товара при операциях обновления и перемещения, чтобы избежать отрицательного количества товаров.
Технические Уточнения
ORM (Object-Relational Mapping): Используйте такие инструменты, как GORM или Beego ORM, для взаимодействия с PostgreSQL. Это облегчит работу с базой данных и сделает код более чистым и поддерживаемым.
Миграции Базы Данных: Реализуйте систему миграций для управления структурой базы данных. Это позволит легко обновлять схему данных при изменении требований к проекту.
Логирование: Внедрите систему логирования для отслеживания запросов к API и внутренних ошибок. Это обеспечит легкость отладки и мониторинга приложения.
Дополнительные Сценарии Использования
Отчеты: Реализуйте эндпоинты для создания отчетов о текущем состоянии склада, включая общее количество товаров, их распределение по складам и историю перемещений.
Уведомления о низком запасе: Система должна автоматически оповещать администраторов о товарах с низким запасом на складе.
Интеграция с внешними системами: Предусмотрите возможность интеграции вашего API с другими системами, например, с системами учета и ERP, через веб-хуки или API.
Документация
Swagger / OpenAPI: Используйте Swagger или OpenAPI для создания интерактивной документации API. Это позволит разработчикам легко тестировать эндпоинты и понимать структуру запросов и ответов.
Руководство пользователя: Составьте подробное руководство пользователя для вашего API, включая примеры запросов и ответов, описание аутентификации и авторизации, а также общие рекомендации по использованию API.
Версионирование API: Планируйте версионирование API с самого начала разработки, чтобы обеспечить совместимость с будущими изменениями и рас