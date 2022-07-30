package ginparam

import "github.com/gin-gonic/gin"

func QueryKeysToMap(c *gin.Context, keys []string) map[string]interface{} {
	query := make(map[string]interface{})
	for _, key := range keys {
		query[key] = c.Query(key)
	}
	return query
}

func PostKeysToMap(c *gin.Context, keys []string) map[string]interface{} {
	post := make(map[string]interface{})

	for _, key := range keys {
		post[key] = c.PostForm(key)
	}

	return post
}
