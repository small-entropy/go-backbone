package mongo

import (
	"context"
	"log"
	"time"

	"github.com/small-entropy/go-backbone/pkg/database/abstract"
	facade "github.com/small-entropy/go-backbone/third_party/facade/mongo"
)

// Метод для установления соединения с репозиторием
func (m *MongoBD) Connect(uri string, database string) error {
	var err error
	var client *facade.Client
	// Пытаемся создать клиент для подключения к MongoDB
	m.Options = facade.GetClientOptions().ApplyURI(uri)
	if client, err = facade.NewClient(m.Options); err == nil {
		// Если удалось создать клиент, то пытаемся выполнить подключение
		// к серверу MongoDB
		m.Client = client
		m.Name = database
		// Если удалось создать клиент, то пытаемся выполнить подключение
		// к серверу MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		// Если контекст "отвалился", то выполняем очищение
		defer cancel()
		// Пытаемся подключиться к серверу MongoDB
		if err = m.Client.Connect(ctx); err != nil {
			// Если не удалось выполнить подключение к серверу MongoDB, то
			// выводим ошибку в лог
			log.Fatalln(err)
		} else {

			// Если удалось выполнить соединение с MongoDB, то пытаемся
			// выполнить пинг сервера
			if err = m.Client.Ping(ctx, nil); err == nil {
				// Если удалось произвести пинг, то выводим сообщение
				// в лог
				log.Println("Connect to MongoDB")
			} else {
				// Если не удалось произвести пинг соединения,
				// то выводим ошибку в лог
				log.Fatalln(err)
			}
		}
	} else {
		// Если не удалось создать клиент для подключения к MongoDB, то
		// выводим ошибку в лог
		log.Fatal(err)
	}
	return err
}

// Метод для получения хранилища данных
func (m *MongoBD) GetStorage(ctx *context.Context, store_name string) *facade.Collection {
	collection := m.Client.Database(m.Name).Collection(store_name)
	return collection
}

// Метод создания индексов для коллекций
func (m *MongoBD) CreateIndexes(opts []abstract.IndexOptions) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, opt := range opts {
		if err = m.CreateIndex(ctx, opt.Collection, opt.Name, opt.Unique); err != nil {
			break
		}
	}

	return err
}

// Метод создания индекса для коллекции
func (m *MongoBD) CreateIndex(ctx context.Context, collectionName string, indexName string, unique bool) error {
	var err error
	var name string
	var datastore *facade.Collection

	// Задаем модель индекса
	indexModel := facade.IndexModel{
		// Указываем, какие поля требуется индексировать
		Keys: facade.BsonD{
			{Key: indexName, Value: 1},
		},
		// Указываем опции индексации
		Options: facade.GetIndexOptions().SetUnique(unique),
	}

	// Получаем коллекцию по названию
	datastore = m.GetStorage(&ctx, collectionName)

	// Пытаемся создать индекс для записей коллекции
	if name, err = datastore.Indexes().CreateOne(ctx, indexModel); err != nil {
		// Если при попытке создать индекс возникла ошибка,
		// то выводим ее в лог
		log.Fatal(err)
	} else {
		// Если удалось создать индекс, то выводим сообщение в лог
		log.Println("Created index with name: " + name)
	}

	return err
}
