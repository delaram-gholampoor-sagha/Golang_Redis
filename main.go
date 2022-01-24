package main

import (
	"os"

	"github.com/Delaram-Gholampoor-Sagha/Golang_Redis/cache"
	"github.com/Delaram-Gholampoor-Sagha/Golang_Redis/repository"
	"github.com/Delaram-Gholampoor-Sagha/Golang_Redis/service"
	router "github.com/Delaram-Gholampoor-Sagha/Golang_Redis/http"
	"github.com/Delaram-Gholampoor-Sagha/Golang_Redis/controller"

)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
 
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
