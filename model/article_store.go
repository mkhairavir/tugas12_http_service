package model

type ArticleStoreInMemory struct {
	ArticleMap []Article
}

func NewArticleStoreInMemory() *ArticleStoreInMemory {
	return &ArticleStoreInMemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat Web", Body: "Hallo world!"},
		},
	}
}

func (store *ArticleStoreInMemory) Save(article *Article) error {
	//1. menghitung length
	lastID := len(store.ArticleMap)

	//set article id
	article.ID = lastID + 1

	// push to article map slice
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}

func (store *ArticleStoreInMemory) Edit(id int, title, body string) error {
	store.ArticleMap[id-1] = Article{ID: id, Title: title, Body: body}

	return nil
}

func (store *ArticleStoreInMemory) Del(id int) error {
	store.ArticleMap = append(store.ArticleMap[:id-1], store.ArticleMap[id:]...)

	return nil
}
