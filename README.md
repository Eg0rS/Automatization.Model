# 1)    docker-compose up --build 
# 2)    и идешь за пивом, образы весят 3.3gb
# 3)    в KafDrop (http://localhost:9090/) нужно создать 3 топика 
    storage_topic - эвенты деталей ожидающих отправку на склад
    verification_topic - эвенты деталей ожидающих процедуру проверки
    processing_topic - эвенты деталей ожидающих каких то манипуляций 
# 4)    в базе пострегреса можно посмотреть какие детали лежат и в каком состоянии (localhost:5431 postgres:postgres)
# 5)    крад апи, поднятое на порту 80 содерждит endpoint'ы:
    /ping - get - проверка активности сервиса, если нет докера

    /add - post - добавление детали в очередь на обрабтку, в теле массив моделей 
    {
        "long": "длинна детатали, float"
	    "width": "ширина детатали, float"
	    "height": "высота детатали, float"
	    "color": "цвет детатали, text"
    }
    возващает массив, id добавленных деталей в том же порядке что и добавление, 0 - означает ошибку добавления 

    /get/{id} - get - возвращает модель детали по id 
    {
        "id": "id детатали, int"
        "long": "длинна детатали, float"
	    "width": "ширина детатали, float"
	    "height": "высота детатали, float"
	    "color": "цвет детатали, text"
        "event_data": "время добавления делали а очередь детатали, datatimestamp"
        "is_deleted": "флаг удаления детали, bool"
    }
    /get/all - get - возвращает массив всех деталей
    
    /delete/{id} - delete - обновляет флаг is_deleted у детали по id, удаленные детали по конвейру идти перестают
    /update - patch - обновляет все поля в модели детали кроме id 

# 6*)	Clickhouse связывание кликахуса и кафки 
	CREATE TABLE storage_queue
(
    Id        UInt64,
    Long      Float64,
    Width     Float64,
    Height    Float64,
    Color     String,
    EventDate Date,
    IsDeleted UInt8
) ENGINE = Kafka('kafka:29092', 'storage_topic', 'group', 'JSONEachRow');

CREATE TABLE stats_storage
(
    Id        UInt64,
    Long      Float64,
    Width     Float64,
    Height    Float64,
    Color     String,
    EventDate Date,
    IsDeleted UInt8
) ENGINE = MergeTree()
      ORDER BY (Id);


CREATE MATERIALIZED VIEW storage TO stats_storage
AS SELECT Id, Long, Width, Height, Color, EventDate, IsDeleted
FROM storage_queue;

CREATE TABLE verification_queue
(
    Id        UInt64,
    Long      Float64,
    Width     Float64,
    Height    Float64,
    Color     String,
    EventDate Date,
    IsDeleted UInt8
) ENGINE = Kafka('kafka:29092', 'verification_topic', 'group', 'JSONEachRow');

CREATE TABLE verification_storage
(
    Id        UInt64,
    Long      Float64,
    Width     Float64,
    Height    Float64,
    Color     String,
    EventDate Date,
    IsDeleted UInt8
) ENGINE = MergeTree()
      ORDER BY (Id);


CREATE MATERIALIZED VIEW verification TO verification_storage
AS SELECT Id, Long, Width, Height, Color, EventDate, IsDeleted
FROM storage_queue;

CREATE TABLE processing_queue
(
    Id        UInt64,
    Long      Float64,
    Width     Float64,
    Height    Float64,
    Color     String,
    EventDate Date,
    IsDeleted UInt8
) ENGINE = Kafka('kafka:29092', 'processing_topic', 'group', 'JSONEachRow');

CREATE TABLE processing_storage
(
    Id        UInt64,
    Long      Float64,
    Width     Float64,
    Height    Float64,
    Color     String,
    EventDate Date,
    IsDeleted UInt8
) ENGINE = MergeTree()
      ORDER BY (Id);


CREATE MATERIALIZED VIEW processing TO processing_storage
AS SELECT Id, Long, Width, Height, Color, EventDate, IsDeleted
FROM storage_queue; 
