package handlers

import (
	"context"

	"google.golang.org/grpc"

	pb "github.com/geshtng/go-base-backend/internal/handlers/article/grpc"
	"github.com/geshtng/go-base-backend/internal/models"
	s "github.com/geshtng/go-base-backend/internal/services"
)

type articleServer struct {
	pb.UnimplementedArticleServer
	articleService s.ArticleService
}

func NewArticleServerGrpc(gserver *grpc.Server, articleService s.ArticleService) {
	server := articleServer{
		articleService: articleService,
	}

	pb.RegisterArticleServer(gserver, &server)
}

func (a *articleServer) GetAllArticles(ctx context.Context, in *pb.GetAllArticlesRequest) (*pb.GetAllArticlesResponse, error) {
	articles, err := a.articleService.GetAll()
	if err != nil {
		return nil, err
	}

	var articlesResponse []*pb.ArticleResponse

	for i := range articles {
		articleResponse := pb.ArticleResponse{
			Id:          articles[i].ID,
			Title:       articles[i].Title,
			Description: articles[i].Description,
		}

		articlesResponse = append(articlesResponse, &articleResponse)
	}

	response := &pb.GetAllArticlesResponse{
		Data: articlesResponse,
	}

	return response, nil
}

func (a *articleServer) GetArticleById(ctx context.Context, in *pb.GetArticleByIdRequest) (*pb.GetArticleByIdResponse, error) {
	article, err := a.articleService.GetByID(in.GetId())
	if err != nil {
		return nil, err
	}

	response := &pb.GetArticleByIdResponse{
		Id:          article.ID,
		Title:       article.Title,
		Description: article.Description,
	}

	return response, nil
}

func (a *articleServer) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	article := models.Article{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	newArticle, err := a.articleService.Create(article)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateArticleResponse{
		Id:          newArticle.ID,
		Title:       newArticle.Title,
		Description: newArticle.Description,
	}

	return response, nil
}

func (a *articleServer) UpdateArticle(ctx context.Context, in *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	article := models.Article{
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
	}

	updatedArticle, err := a.articleService.Update(in.GetId(), article)
	if err != nil {
		return nil, err
	}

	response := &pb.UpdateArticleResponse{
		Id:          updatedArticle.ID,
		Title:       updatedArticle.Title,
		Description: updatedArticle.Description,
	}

	return response, nil
}

func (a *articleServer) DeleteArticle(ctx context.Context, in *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	err := a.articleService.Delete(in.GetId())
	if err != nil {
		return nil, err
	}

	response := &pb.DeleteArticleResponse{
		Status: "article deleted",
	}

	return response, nil
}
