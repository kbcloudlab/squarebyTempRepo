package sector

import (
	"errors"

	dto "squareby.com/admin/cloudspacemanager/rest/dto/sector"
	models "squareby.com/admin/cloudspacemanager/src/models/sector"

	customError "squareby.com/admin/cloudspacemanager/src/programfiles/customerror"
	repo "squareby.com/admin/cloudspacemanager/src/repository/sector"
)

type SpaceCategoryService struct {
	spaceCategoryRepo *repo.SpaceCategoryRepo
}

func NewSpaceCategoryService() *SpaceCategoryService {
	return &SpaceCategoryService{
		spaceCategoryRepo: repo.NewSpaceCategoryRepo(),
	}
}

func (_sfs *SpaceCategoryService) CreateSpaceCategory() error {

	for _, space_category_data := range models.SpaceCategoryDataList {

		_, err := _sfs.spaceCategoryRepo.GetSpaceCategoryByCode(space_category_data.Code)

		var notFoundErr *customError.NotFoundErr = &customError.NotFoundErr{}

		if errors.As(err, &notFoundErr) {
			_, _ = _sfs.spaceCategoryRepo.CreateSpaceCategory(space_category_data)

		} else if err != nil {
			return err
		}

	}

	return nil

} //End of the method CreateSpaceCategory

func (_sfs *SpaceCategoryService) UpdateSpaceCategory(_id uint64, _spaceCategoryRequest *dto.SpaceCategoryRequest) error {

	space_category := _spaceCategoryRequest.NewSpaceCategoryFromRequest(false)

	exist_spacecategory, err := _sfs.spaceCategoryRepo.GetSpaceCategoryById(_id)

	if err != nil {
		return err
	}

	exist_spacecategory.SectorCode = space_category.SectorCode
	exist_spacecategory.Name = space_category.Name
	// exist_spacecategory.ImageURL = space_category.ImageURL

	return _sfs.spaceCategoryRepo.UpdateSpaceCategory(_id, exist_spacecategory)

} //End of the method UpdateSpaceCategory

func (_sfs *SpaceCategoryService) AddRemoveProductContainer(_id uint64, _categoryProductContainerRequest *dto.CategoryProductContainerRequest) error {

	container_id_list := _categoryProductContainerRequest.ProductContainerIdList

	exist_spacecategory, err := _sfs.spaceCategoryRepo.GetSpaceCategoryById(_id)

	if err != nil {
		return err
	}

	exist_spacecategory.ProductContainerIdList = container_id_list

	return _sfs.spaceCategoryRepo.UpdateSpaceCategory(_id, exist_spacecategory)

} //End of the method UpdateSpaceCategory

func (_sfs *SpaceCategoryService) EnableDisableSpaceCategory(_id uint64, _spaceCategoryRequest *dto.SpaceCategoryEnableRequest) error {

	space_category, err := _sfs.spaceCategoryRepo.GetSpaceCategoryById(_id)
	if err != nil {
		return err
	}

	space_category.Enabled = _spaceCategoryRequest.Enabled

	return _sfs.spaceCategoryRepo.UpdateSpaceCategory(_id, space_category)

} //End of the method EnableDisableSpaceCategory

func (_sfs *SpaceCategoryService) GetSpaceCategoryByCode(_code string) (*dto.SpaceCategoryResponse, error) {

	space_category, err := _sfs.spaceCategoryRepo.GetSpaceCategoryByCode(_code)

	if err != nil {
		return nil, err
	}

	res, err := dto.NewSpaceCategoryResponse(space_category)
	if err != nil {
		return nil, err
	}

	return res, nil

} //End of the method GetSpaceCategoryByCode

func (_sfs *SpaceCategoryService) GetSpaceCategoryById(_id uint64) (*dto.SpaceCategoryResponse, error) {

	space_category, err := _sfs.spaceCategoryRepo.GetSpaceCategoryById(_id)

	if err != nil {
		return nil, err
	}

	res, err := dto.NewSpaceCategoryResponse(space_category)
	if err != nil {
		return nil, err
	}

	return res, nil

} //End of the method GetSpaceCategoryById

func (_sfs *SpaceCategoryService) GetActiveSpaceCategoryList() (dto.SpaceCategoryResponse1List, error) {

	space_category_list, err := _sfs.spaceCategoryRepo.GetActiveSpaceCategoryList()

	if err != nil {
		return nil, err
	}

	spacecategory_res_list := dto.SpaceCategoryResponse1List{}
	for _, spacecategory := range space_category_list {
		res, err := dto.NewSpaceCategoryResponse1(spacecategory)
		if err != nil {
			return nil, err
		}
		spacecategory_res_list = append(spacecategory_res_list, res)
	}

	return spacecategory_res_list, nil

} //End of the method GetActiveSpaceCategoryList

func (_sfs *SpaceCategoryService) GetAllSpaceCategoryList() (dto.SpaceCategoryResponse1List, error) {

	space_category_list, err := _sfs.spaceCategoryRepo.GetAllSpaceCategoryList()

	if err != nil {
		return nil, err
	}

	spacecategory_res_list := dto.SpaceCategoryResponse1List{}
	for _, spacecategory := range space_category_list {
		res, err := dto.NewSpaceCategoryResponse1(spacecategory)
		if err != nil {
			return nil, err
		}
		spacecategory_res_list = append(spacecategory_res_list, res)
	}

	return spacecategory_res_list, nil

} //End of the method GetAllSpaceCategoryList
