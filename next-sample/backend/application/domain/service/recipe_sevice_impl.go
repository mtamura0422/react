package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/react/next-sample/backend/domain/entity"
	"github.com/react/next-sample/backend/domain/repositories"
	"github.com/react/next-sample/backend/domain/service/port"
	"github.com/react/next-sample/backend/infrastructure/db"
)

var _ port.RecipeService = (*RecipeServiceImpl)(nil)

type RecipeServiceImpl struct {
	txRepo *db.TxRepository
	rRepo  repositories.RecipeRepository
	rmRepo repositories.RecipeMaterialRepository
}

func NewRecipeService(
	txRepo *db.TxRepository,
	rRepo repositories.RecipeRepository,
	rmRepo repositories.RecipeMaterialRepository,
) port.RecipeService {
	return &RecipeServiceImpl{txRepo: txRepo, rRepo: rRepo, rmRepo: rmRepo}
}

func (s *RecipeServiceImpl) createRecipe(input *entity.Recipe) func(ctx context.Context) (interface{}, error) {

	return func(ctx context.Context) (interface{}, error) {
		// 画像アップロード
		filename, err := uploadImg(input)
		if err != nil {
			return nil, err
		}

		input.Filename = filename
		eRecipe, err := s.rRepo.Create(ctx, input)
		if err != nil {
			return nil, err
		}
		log.Printf("createRecipe tuuka2")
		for _, recipeMaterial := range input.RecipeMaterials {
			log.Printf("createRecipe tuuka3")
			recipeMaterial.RecipeId = eRecipe.Id
			_, err := s.rmRepo.Create(ctx, &recipeMaterial)
			if err != nil {
				return nil, err
			}

		}

		return eRecipe, nil
	}
}

func (s *RecipeServiceImpl) CreateRecipeTx(ctx context.Context, input *entity.Recipe) (interface{}, error) {

	v, err := s.txRepo.RunInTx(ctx, s.createRecipe(input))
	if err != nil {
		return v, err
	}

	return v, nil
}

func uploadImg(input *entity.Recipe) (string, error) {

	if input.RecipeImage.File == nil {
		return "", nil
	}

	defer input.RecipeImage.File.Close()

	// 存在していなければ、保存用のディレクトリを作成します。
	err := os.MkdirAll("./images", os.ModePerm)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(input.RecipeImage.FileHeader.Filename))

	// 保存用ディレクトリ内に新しいファイルを作成します。
	dst, err := os.Create(fmt.Sprintf("./images/%s", filename))
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// アップロードされたファイルを先程作ったファイルにコピーします。
	_, err = io.Copy(dst, input.RecipeImage.File)
	if err != nil {
		return "", err
	}

	return filename, nil
}
