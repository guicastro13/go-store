package response

type CategoryResponse struct {
  ID string `json:"id"`
  Title string `json:"title"`
}

type ManyCategoryResponse struct {
  Categories []CategoryResponse `json:"categories"`
}
